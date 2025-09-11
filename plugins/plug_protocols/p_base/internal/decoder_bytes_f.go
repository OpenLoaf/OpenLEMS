package internal

import (
	"common/c_enum"
	"encoding/binary"
	"math"
	"strings"
	"unicode/utf16"

	"github.com/pkg/errors"
)

// DecoderBytes 将字节数组解析为各种数据格式
// 此函数是协议解析的核心，支持以下协议：
// - Modbus TCP/RTU: 支持所有标准数据类型和字节序
// - CANbus: 支持位级数据、BCD码、多字节序
// - IEC 61850: 支持IEEE 754浮点数、字符串
// - S7: 支持西门子PLC的数据格式
// - 其他工业协议: 通过自定义格式扩展
//
// 参数说明：
// - bytes: 原始字节数据
// - byteIndex: 字节起始索引
// - byteLength: 字节长度（0表示纯位模式）
// - bitIndex: 位起始索引
// - bitLength: 位长度（0表示纯字节模式）
// - byteEndian: 字节序（大端/小端）
// - wordOrder: 字序（高字在前/低字在前）
// - dataFormat: 数据格式（整数、浮点数、BCD、字符串等）
// - returnFormat: 返回格式类型
// - offset: 偏移量
// - factor: 系数
//
// 三种读取模式：
//  1. 纯字节模式 (bitLength=0): 读取 byteIndex 开始的 byteLength 字节
//  2. 纯位模式 (byteLength=0): 读取 bitIndex 开始的 bitLength 位
//  3. 混合模式 (两者都不为0): 先读取字节，再从中提取指定位
func DecoderBytes(bytes []byte, byteIndex uint16, byteLength uint16, bitIndex uint16, bitLength uint16, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder, dataFormat c_enum.EDataFormat, returnFormat c_enum.EValueType, offset int, factor float32) (any, error) {
	// 参数验证
	if len(bytes) == 0 {
		return nil, errors.New("empty input data")
	}

	// 确保系数不为0
	factor = EnsureFactorNotZero(factor)

	// 获取有效的字节和位长度
	effectiveByteLength := GetEffectiveByteLength(byteLength, dataFormat)
	effectiveBitLength := GetEffectiveBitLength(bitLength, dataFormat)

	// 验证参数组合的有效性
	if effectiveByteLength == 0 && effectiveBitLength == 0 {
		return nil, errors.New("both byteLength and bitLength cannot be zero")
	}

	// 根据参数组合确定读取模式
	var data []byte
	var bitStart, actualBitLength uint8

	if IsByteMode(byteLength, dataFormat) && !IsBitMode(bitLength, dataFormat) {
		// 模式1：纯字节模式 - 读取 byteIndex 开始的 byteLength 字节
		if int(byteIndex)+int(effectiveByteLength) > len(bytes) {
			return nil, errors.Errorf("insufficient data: requested byte range [%d:%d] exceeds data length %d",
				byteIndex, byteIndex+effectiveByteLength-1, len(bytes))
		}
		data = bytes[byteIndex : byteIndex+effectiveByteLength]
		bitStart = 0
		actualBitLength = 0
	} else if IsBitMode(bitLength, dataFormat) && !IsByteMode(byteLength, dataFormat) {
		// 模式2：纯位模式 - 读取 bitIndex 开始的 bitLength 位
		// 计算需要的字节数
		endBitIndex := uint32(bitIndex) + uint32(effectiveBitLength)
		if endBitIndex > 65535 {
			return nil, errors.New("bit range too large: exceeds uint16 limit")
		}
		requiredBytes := uint16((endBitIndex + 7) / 8)

		if uint16(len(bytes)) < requiredBytes {
			return nil, errors.Errorf("insufficient data: need %d bytes for bit range, got %d", requiredBytes, len(bytes))
		}

		data = bytes[0:requiredBytes]
		bitStart = uint8(bitIndex % 8) // 在字节内的位偏移
		actualBitLength = uint8(effectiveBitLength)
	} else {
		// 模式3：混合模式 - 先读取字节，再从中提取位
		if int(byteIndex)+int(effectiveByteLength) > len(bytes) {
			return nil, errors.Errorf("insufficient data: requested byte range [%d:%d] exceeds data length %d",
				byteIndex, byteIndex+effectiveByteLength-1, len(bytes))
		}

		// 检查位索引是否在字节范围内
		if bitIndex >= uint16(effectiveByteLength*8) {
			return nil, errors.Errorf("bitIndex %d exceeds byte range [0:%d]", bitIndex, effectiveByteLength*8-1)
		}

		// 检查位长度是否超出字节范围
		if bitIndex+effectiveBitLength > uint16(effectiveByteLength*8) {
			return nil, errors.Errorf("bit range [%d:%d] exceeds byte range [0:%d]",
				bitIndex, bitIndex+effectiveBitLength-1, effectiveByteLength*8-1)
		}

		data = bytes[byteIndex : byteIndex+effectiveByteLength]
		bitStart = uint8(bitIndex % 8) // 在字节内的位偏移
		actualBitLength = uint8(effectiveBitLength)
	}

	// 根据数据格式进行解析
	var rawValue any
	var err error

	switch dataFormat {
	case c_enum.DataFormatUInt8:
		rawValue, err = decodeUInt8(data, byteEndian, wordOrder)
	case c_enum.DataFormatInt8:
		rawValue, err = decodeInt8(data, byteEndian, wordOrder)
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
		rawValue, err = decodeBitRange(data, bitStart, actualBitLength)
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
	finalValue := ApplyFactorAndOffset(rawValue, factor, offset)

	// 转换为目标格式
	result := ConvertToReturnFormat(finalValue, returnFormat)

	return result, nil
}

// 数值类型解析函数

func decodeUInt8(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 1 {
		return nil, errors.New("insufficient data for uint8")
	}

	// uint8 只有一个字节，不需要字节序和字序处理
	return uint8(data[0]), nil
}

func decodeInt8(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 1 {
		return nil, errors.New("insufficient data for int8")
	}

	// int8 只有一个字节，不需要字节序和字序处理
	return int8(data[0]), nil
}

func decodeUInt16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for uint16")
	}

	reordered := ReorderBytes(data[:2], byteEndian, wordOrder)
	return binary.BigEndian.Uint16(reordered), nil
}

func decodeInt16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for int16")
	}

	reordered := ReorderBytes(data[:2], byteEndian, wordOrder)
	val := binary.BigEndian.Uint16(reordered)
	return int16(val), nil
}

func decodeUInt32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for uint32")
	}

	reordered := ReorderBytes(data[:4], byteEndian, wordOrder)
	return binary.BigEndian.Uint32(reordered), nil
}

func decodeInt32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for int32")
	}

	reordered := ReorderBytes(data[:4], byteEndian, wordOrder)
	val := binary.BigEndian.Uint32(reordered)
	return int32(val), nil
}

func decodeFloat32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for float32")
	}

	reordered := ReorderBytes(data[:4], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint32(reordered)
	return math.Float32frombits(bits), nil
}

func decodeFloat64(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 8 {
		return nil, errors.New("insufficient data for float64")
	}

	reordered := ReorderBytes(data[:8], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint64(reordered)
	return math.Float64frombits(bits), nil
}

// BCD码解析函数

func decodeBCD16(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for BCD16")
	}

	reordered := ReorderBytes(data[:2], byteEndian, wordOrder)

	// BCD解码：每个字节包含两个十进制数字
	high := int(reordered[0]>>4)*1000 + int(reordered[0]&0x0F)*100
	low := int(reordered[1]>>4)*10 + int(reordered[1]&0x0F)

	return high + low, nil
}

func decodeBCD32(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for BCD32")
	}

	reordered := ReorderBytes(data[:4], byteEndian, wordOrder)

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
		reordered := ReorderBytes(data, byteEndian, wordOrder)
		return binary.BigEndian.Uint16(reordered), nil
	} else if len(data) == 4 {
		reordered := ReorderBytes(data, byteEndian, wordOrder)
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
