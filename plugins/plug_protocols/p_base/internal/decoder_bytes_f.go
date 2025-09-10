package internal

import (
	"common/c_enum"
	"encoding/binary"
	"math"
	"strings"
	"unicode/utf16"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

// DecoderBytes 通用字节解析函数，支持多种协议的数据解析
//
// 此函数是协议解析的核心，支持以下协议：
// - Modbus TCP/RTU: 支持所有标准数据类型和字节序
// - CANbus: 支持位级数据、BCD码、多字节序
// - IEC 61850: 支持IEEE 754浮点数、字符串
// - S7: 支持西门子PLC的数据格式
// - 其他工业协议: 通过自定义格式扩展
//
// 参数说明：
//   - bytes: 原始字节数据
//   - index: 起始字节索引
//   - length: 数据长度（字节数）
//   - byteEndian: 字节序（大端/小端）
//   - wordOrder: 字序（高字在前/低字在前）
//   - dataFormat: 数据格式（整数、浮点数、BCD、字符串等）
//   - returnFormat: 返回格式类型
//   - offset: 偏移量
//   - factor: 系数
//   - min: 最小值验证（数值类型）或最小长度验证（字符串类型，0表示不验证）
//   - max: 最大值验证（数值类型）或最大长度验证（字符串类型，0表示不验证）
//
// 使用示例：
//
//	// Modbus 16位整数解析（带范围验证）
//	result, err := DecoderBytes(data, 0, 2, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, SUseReadType, 0, 1.0, 0, 1000)
//	if err != nil {
//		// 处理错误
//	}
//
//	// CANbus BCD码解析（无范围验证）
//	result, err := DecoderBytes(data, 2, 2, ByteEndianLittle, WordOrderHighLow, DataFormatBCD, SInt32, 0, 0.1, 0, 0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// IEEE 754浮点数解析（带范围验证）
//	result, err := DecoderBytes(data, 0, 4, ByteEndianBig, WordOrderHighLow, DataFormatFloat32, SFloat64, 0, 1.0, -100, 100)
//	if err != nil {
//		// 处理错误
//	}
//
//	// ASCII字符串解析（带长度验证）
//	result, err := DecoderBytes(data, 0, 10, ByteEndianBig, WordOrderHighLow, DataFormatStringASCII, SString, 0, 1.0, 3, 20)
//	if err != nil {
//		// 处理错误
//	}
func DecoderBytes(bytes []byte, index uint16, length uint16, isBit bool, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder, dataFormat c_enum.EDataFormat, returnFormat c_enum.ESystemType, offset int, factor float32, min, max int64) (any, error) {
	// 参数验证
	if len(bytes) == 0 {
		return nil, errors.New("empty input data")
	}

	// 确保系数不为0
	if factor == 0 {
		factor = 1.0
	}

	// 根据 isBit 参数计算实际需要的数据
	var data []byte
	var bitStart, bitLength uint8

	if isBit {
		// 位索引模式：index 和 length 都是位级别的
		bitStart = uint8(index % 8) // 在字节内的位偏移
		bitLength = uint8(length)   // 位长度

		// 计算需要的字节数，防止uint16溢出
		byteIndex := index / 8
		// 使用uint32进行计算以防止溢出
		endBitIndex := uint32(index) + uint32(length)
		if endBitIndex > 65535 {
			return nil, errors.New("bit range too large: exceeds uint16 limit")
		}
		requiredBytes := uint16((endBitIndex + 7) / 8)

		if uint16(len(bytes)) < requiredBytes {
			return nil, errors.Errorf("insufficient data: need %d bytes for bit range, got %d", requiredBytes, len(bytes))
		}

		data = bytes[byteIndex:requiredBytes]
	} else {
		// 字节索引模式：index 和 length 都是字节级别的
		// 检查数据长度是否足够，防止uint16溢出
		indexInt := int(index)
		lengthInt := int(length)
		endIndex := indexInt + lengthInt

		if endIndex > len(bytes) {
			return nil, errors.Errorf("insufficient data: requested range [%d:%d] exceeds data length %d",
				index, endIndex-1, len(bytes))
		}

		data = bytes[indexInt:endIndex]
		bitStart = 0
		bitLength = 0
	}

	// 根据数据格式进行解析
	var rawValue any
	var err error

	switch dataFormat {
	case c_enum.DataFormatUInt16:
		rawValue, err = decodeUInt16(data, byteEndian, wordOrder)
	case c_enum.DataFormatInt16:
		rawValue, err = decodeInt16(data, byteEndian, wordOrder)
	case c_enum.DataFormatUInt32:
		rawValue, err = decodeUInt32(data, byteEndian, wordOrder)
	case c_enum.DataFormatInt32:
		rawValue, err = decodeInt32(data, byteEndian, wordOrder)
	case c_enum.DataFormatFloat32:
		rawValue, err = decodeFloat32(data, byteEndian, wordOrder)
	case c_enum.DataFormatFloat64:
		rawValue, err = decodeFloat64(data, byteEndian, wordOrder)
	case c_enum.DataFormatBCD:
		rawValue, err = decodeBCD16(data, byteEndian, wordOrder)
	case c_enum.DataFormatBCD32:
		rawValue, err = decodeBCD32(data, byteEndian, wordOrder)
	case c_enum.DataFormatStringASCII:
		rawValue, err = decodeASCIIString(data)
	case c_enum.DataFormatStringUTF16:
		rawValue, err = decodeUTF16String(data, byteEndian, wordOrder)
	case c_enum.DataFormatBits:
		rawValue, err = decodeBits(data, byteEndian, wordOrder)
	case c_enum.DataFormatBitRange:
		rawValue, err = decodeBitRange(data, bitStart, bitLength)
	case c_enum.DataFormatCustom:
		// 自定义格式需要外部提供解析函数
		return nil, errors.New("custom data format not supported")
	default:
		return nil, errors.Errorf("unsupported data format: %v", dataFormat)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to decode data")
	}

	// 应用系数和偏移量
	finalValue := applyFactorAndOffset(rawValue, factor, offset)

	// 转换为目标格式
	result := convertToReturnFormat(finalValue, returnFormat)

	// 数据范围验证
	if err := validateValueRange(result, min, max); err != nil {
		return nil, errors.Wrap(err, "value validation failed")
	}

	return result, nil
}

// 数据范围验证函数
func validateValueRange(value any, min, max int64) error {
	// 如果min和max都为0，表示不进行范围验证
	if min == 0 && max == 0 {
		return nil
	}

	// 字符串类型进行长度验证
	if str, ok := value.(string); ok {
		length := int64(len(str))
		// 对于字符串长度，min和max应该是非负数，但我们仍然正确处理
		if min != 0 && length < min {
			return errors.Errorf("string length %d is below minimum %d", length, min)
		}
		if max != 0 && length > max {
			return errors.Errorf("string length %d is above maximum %d", length, max)
		}
		return nil
	}

	// 布尔类型不进行数值范围验证
	if _, ok := value.(bool); ok {
		return nil
	}

	// 处理uint64类型，避免溢出风险
	if uintVal, ok := value.(uint64); ok {
		// 检查uint64值是否在int64范围内
		if uintVal > math.MaxInt64 {
			return errors.Errorf("uint64 value %d exceeds int64 maximum %d", uintVal, math.MaxInt64)
		}
		val := int64(uintVal)
		// 修复：正确处理负数范围验证
		if min != 0 && val < min {
			return errors.Errorf("value %d is below minimum %d", val, min)
		}
		if max != 0 && val > max {
			return errors.Errorf("value %d is above maximum %d", val, max)
		}
		return nil
	}

	// 处理其他数值类型
	val := cvt.Int64(value)

	// 修复：正确处理负数范围验证
	// 只有当min或max不为0时才进行相应的检查
	if min != 0 && val < min {
		return errors.Errorf("value %d is below minimum %d", val, min)
	}

	if max != 0 && val > max {
		return errors.Errorf("value %d is above maximum %d", val, max)
	}

	return nil
}

// 字节序和字序处理辅助函数

// reorderBytes 根据字节序和字序重新排列字节
// 遵循Modbus协议标准：先处理字序（Word Order），再处理字节序（Byte Endian）
// 性能优化：预分配内存，减少内存分配次数
func reorderBytes(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) []byte {
	if len(data) < 2 {
		return data
	}

	// 预分配结果切片，避免多次内存分配
	result := make([]byte, len(data))
	copy(result, data)

	// 第一步：处理字序（Word Order）- 交换16位字的顺序
	// 这在Modbus协议中首先发生
	if wordOrder == c_enum.WordOrderLowHigh {
		switch len(result) {
		case 4:
			// 32位数据：[AB CD EF GH] -> [EF GH AB CD]
			result[0], result[2] = result[2], result[0]
			result[1], result[3] = result[3], result[1]
		case 8:
			// 64位数据：[AB CD EF GH IJ KL MN OP] -> [IJ KL MN OP AB CD EF GH]
			// 交换前4字节和后4字节
			for i := 0; i < 4; i++ {
				result[i], result[i+4] = result[i+4], result[i]
			}
		default:
			// 对于其他长度，按4字节分组交换（保持向后兼容）
			if len(result) >= 4 {
				for i := 0; i < len(result)-3; i += 4 {
					result[i], result[i+2] = result[i+2], result[i]
					result[i+1], result[i+3] = result[i+3], result[i+1]
				}
			}
		}
	}

	// 第二步：处理字节序（Byte Endian）- 在每个16位字内交换字节
	// 这在Modbus协议中第二步发生
	if byteEndian == c_enum.ByteEndianLittle {
		// 小端字节序：反转每个16位字内的字节
		// 例如：[AB CD] -> [BA DC]
		for i := 0; i < len(result); i += 2 {
			if i+1 < len(result) {
				result[i], result[i+1] = result[i+1], result[i]
			}
		}
	}

	return result
}

// 数值类型解析函数

func decodeUInt16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for uint16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)
	return binary.BigEndian.Uint16(reordered), nil
}

func decodeInt16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for int16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)
	val := binary.BigEndian.Uint16(reordered)
	return int16(val), nil
}

func decodeUInt32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for uint32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	return binary.BigEndian.Uint32(reordered), nil
}

func decodeInt32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for int32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	val := binary.BigEndian.Uint32(reordered)
	return int32(val), nil
}

func decodeFloat32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for float32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint32(reordered)
	return math.Float32frombits(bits), nil
}

func decodeFloat64(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 8 {
		return nil, errors.New("insufficient data for float64")
	}

	reordered := reorderBytes(data[:8], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint64(reordered)
	return math.Float64frombits(bits), nil
}

// BCD码解析函数

func decodeBCD16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for BCD16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)

	// BCD解码：每个字节包含两个十进制数字
	high := int(reordered[0]>>4)*1000 + int(reordered[0]&0x0F)*100
	low := int(reordered[1]>>4)*10 + int(reordered[1]&0x0F)

	return high + low, nil
}

func decodeBCD32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for BCD32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)

	result := 0
	for i := 0; i < 4; i++ {
		// 每个字节包含两个BCD数字
		high := int(reordered[i]>>4) * int(math.Pow10(7-2*i))
		low := int(reordered[i]&0x0F) * int(math.Pow10(6-2*i))
		result += high + low
	}

	return result, nil
}

// 字符串解析函数

func decodeASCIIString(data []byte) (any, error) {
	// 移除尾部的null字符
	str := strings.TrimRight(string(data), "\x00")
	return str, nil
}

func decodeUTF16String(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return "", errors.New("insufficient data for UTF16 string")
	}

	// 确保数据长度是偶数
	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	}

	// UTF-16字符串只需要处理字节序，不需要字序交换
	// 字序交换会破坏UTF-16编码的正确性
	result := make([]byte, len(data))
	if byteEndian == c_enum.ByteEndianLittle {
		// 小端字节序：在每个16位字内交换字节
		for i := 0; i < len(data); i += 2 {
			if i+1 < len(data) {
				result[i] = data[i+1]
				result[i+1] = data[i]
			} else {
				result[i] = data[i]
			}
		}
	} else {
		// 大端字节序：直接复制
		copy(result, data)
	}

	// 转换为UTF-16代码点
	codePoints := make([]uint16, len(result)/2)
	for i := 0; i < len(result); i += 2 {
		codePoints[i/2] = binary.BigEndian.Uint16(result[i : i+2])
	}

	// 转换为UTF-8字符串
	runes := utf16.Decode(codePoints)
	return string(runes), nil
}

// 位图解析函数

func decodeBits(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("insufficient data for bits")
	}

	// 对于位图，通常返回原始字节数组或转换为整数
	if len(data) == 1 {
		return uint8(data[0]), nil
	} else if len(data) == 2 {
		reordered := reorderBytes(data, byteEndian, wordOrder)
		return binary.BigEndian.Uint16(reordered), nil
	} else if len(data) == 4 {
		reordered := reorderBytes(data, byteEndian, wordOrder)
		return binary.BigEndian.Uint32(reordered), nil
	} else {
		// 对于更长的位图，返回字节数组
		return data, nil
	}
}

// 位范围解析函数
// 从字节数组中提取特定范围的位，支持跨字节操作
func decodeBitRange(data []byte, bitStart, bitLength uint8) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("insufficient data for bit range")
	}

	if bitLength == 0 || bitLength > 64 {
		return nil, errors.New("bit length must be 1-64")
	}

	maxBitIndex := bitStart + bitLength - 1
	maxByteIndex := maxBitIndex / 8
	if maxByteIndex >= uint8(len(data)) {
		return nil, errors.Errorf("bit range exceeds data length")
	}

	// 使用位操作组合字节，防止溢出
	var result uint64
	for i := uint8(0); i <= maxByteIndex; i++ {
		shiftAmount := 8 * i
		// 防止左移超过64位导致溢出
		if shiftAmount >= 64 {
			break
		}
		result |= uint64(data[i]) << shiftAmount
	}

	// 提取所需位范围
	result = (result >> bitStart) & ((1 << bitLength) - 1)

	// 返回适当类型
	switch {
	case bitLength <= 8:
		return uint8(result), nil
	case bitLength <= 16:
		return uint16(result), nil
	case bitLength <= 32:
		return uint32(result), nil
	default:
		return result, nil
	}
}

// 系数和偏移量处理
// 使用cvt库简化类型转换

func applyFactorAndOffset(value any, factor float32, offset int) any {
	// 如果系数为1且偏移为0，直接返回原值
	if factor == 1.0 && offset == 0 {
		return value
	}

	// 字符串类型不应用系数和偏移量
	if _, ok := value.(string); ok {
		return value
	}

	// 使用cvt库转换为float64，然后应用系数和偏移量
	val := cvt.Float64(value)
	return val*float64(factor) + float64(offset)
}

// 转换为目标格式

func convertToReturnFormat(value any, returnFormat c_enum.ESystemType) any {
	if value == nil {
		return nil
	}

	switch returnFormat {
	case c_enum.SBool:
		return cvt.Bool(value)
	case c_enum.SInt8:
		return cvt.Int8(value)
	case c_enum.SUint8:
		return cvt.Uint8(value)
	case c_enum.SInt16:
		return cvt.Int16(value)
	case c_enum.SUint16:
		return cvt.Uint16(value)
	case c_enum.SInt32:
		return cvt.Int32(value)
	case c_enum.SUint32:
		return cvt.Uint32(value)
	case c_enum.SInt64:
		return cvt.Int64(value)
	case c_enum.SUint64:
		return cvt.Uint64(value)
	case c_enum.SFloat32:
		return cvt.Float32(value)
	case c_enum.SFloat64:
		return cvt.Float64(value)
	case c_enum.SString:
		return cvt.String(value)
	default:
		return value
	}
}
