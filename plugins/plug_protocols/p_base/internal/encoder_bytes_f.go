package internal

import (
	"common/c_enum"
	"encoding/binary"
	"math"
	"unicode/utf16"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

// EncoderBytes 将各种数据格式编码为字节数组
// 此函数用于生成Modbus、CANbus、IEC 61850、S7等协议的发送指令
// 参数说明：
// - value: 要编码的值
// - dataFormat: 数据格式
// - byteEndian: 字节序
// - wordOrder: 字序
// - offset: 偏移量（编码前应用）
// - factor: 系数（编码前应用）
func EncoderBytes(value any, dataFormat c_enum.EDataFormat, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder, offset int, factor float32) ([]byte, error) {
	// 参数验证
	if value == nil {
		return nil, errors.New("input value cannot be nil")
	}

	// 确保系数不为0
	factor = EnsureFactorNotZero(factor)

	// 应用偏移量和系数（编码前需要反向应用）
	processedValue := ApplyFactorAndOffsetForEncoding(value, factor, offset)

	// 根据数据格式进行编码
	var result []byte
	var err error

	switch dataFormat {
	case c_enum.DataFormatUInt8:
		result, err = encodeUInt8(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatInt8:
		result, err = encodeInt8(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatUInt16:
		result, err = encodeUInt16(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatInt16:
		result, err = encodeInt16(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatUInt32:
		result, err = encodeUInt32(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatInt32:
		result, err = encodeInt32(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatFloat32:
		result, err = encodeFloat32(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatFloat64:
		result, err = encodeFloat64(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatBCD:
		result, err = encodeBCD16(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatBCD32:
		result, err = encodeBCD32(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatStringASCII:
		result, err = encodeASCIIString(processedValue)
	case c_enum.DataFormatStringUTF16:
		result, err = encodeUTF16String(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatBits:
		result, err = encodeBits(processedValue, byteEndian, wordOrder)
	case c_enum.DataFormatBitRange:
		result, err = encodeBitRange(processedValue)
	case c_enum.DataFormatCustom:
		// 自定义格式需要外部提供编码函数
		return nil, errors.New("custom data format not supported")
	default:
		return nil, errors.Errorf("unsupported data format: %v", dataFormat)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to encode data")
	}

	return result, nil
}

// 数值类型编码函数

func encodeUInt8(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Uint8(value)

	// uint8 只有一个字节，不需要字节序和字序处理
	return []byte{byte(val)}, nil
}

func encodeInt8(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Int8(value)

	// int8 只有一个字节，不需要字节序和字序处理
	return []byte{byte(val)}, nil
}

func encodeUInt16(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Uint16(value)

	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, val)

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeInt16(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Int16(value)

	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(val))

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeUInt32(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Uint32(value)

	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, val)

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeInt32(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Int32(value)

	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(val))

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeFloat32(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Float32(value)

	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, math.Float32bits(val))

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeFloat64(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Float64(value)

	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, math.Float64bits(val))

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

// BCD码编码函数

func encodeBCD16(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Int(value)

	// 确保值在BCD16范围内 (0-9999)
	if val < 0 || val > 9999 {
		return nil, errors.Errorf("value %d is out of BCD16 range (0-9999)", val)
	}

	data := make([]byte, 2)

	// BCD编码：将十进制数字转换为BCD格式
	// 高字节：千位和百位
	data[0] = byte((val/1000)<<4 | (val/100)%10)
	// 低字节：十位和个位
	data[1] = byte(((val/10)%10)<<4 | val%10)

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

func encodeBCD32(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	val := cvt.Int(value)

	// 确保值在BCD32范围内 (0-99999999)
	if val < 0 || val > 99999999 {
		return nil, errors.Errorf("value %d is out of BCD32 range (0-99999999)", val)
	}

	data := make([]byte, 4)

	// BCD编码：将十进制数字转换为BCD格式
	for i := 0; i < 4; i++ {
		digit := (val / int(math.Pow10(7-2*i))) % 100
		data[i] = byte((digit/10)<<4 | digit%10)
	}

	return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
}

// 字符串编码函数

func encodeASCIIString(value any) ([]byte, error) {
	str := cvt.String(value)
	return []byte(str), nil
}

func encodeUTF16String(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	str := cvt.String(value)

	// 转换为UTF-16代码点
	runes := []rune(str)
	codePoints := utf16.Encode(runes)

	// 转换为字节数组
	result := make([]byte, len(codePoints)*2)
	for i, cp := range codePoints {
		binary.BigEndian.PutUint16(result[i*2:], cp)
	}

	// UTF-16字符串只需要处理字节序，不需要字序交换
	// 字序交换会破坏UTF-16编码的正确性
	if byteEndian == c_enum.ByteEndianLittle {
		// 小端字节序：在每个16位字内交换字节
		for i := 0; i < len(result); i += 2 {
			if i+1 < len(result) {
				result[i], result[i+1] = result[i+1], result[i]
			}
		}
	}

	return result, nil
}

// 位图编码函数

func encodeBits(value any, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) ([]byte, error) {
	// 根据值的类型和大小确定编码方式
	switch v := value.(type) {
	case uint8:
		return []byte{byte(v)}, nil
	case uint16:
		data := make([]byte, 2)
		binary.BigEndian.PutUint16(data, v)
		return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
	case uint32:
		data := make([]byte, 4)
		binary.BigEndian.PutUint32(data, v)
		return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
	case []byte:
		// 对于字节数组，直接返回
		return v, nil
	default:
		// 尝试转换为uint32
		val := cvt.Uint32(value)
		data := make([]byte, 4)
		binary.BigEndian.PutUint32(data, val)
		return ReorderBytesForEncoding(data, byteEndian, wordOrder), nil
	}
}

// 位范围编码函数
// 将值编码为指定位长度的位数据
func encodeBitRange(value any) ([]byte, error) {
	val := cvt.Uint64(value)

	// 根据值的大小确定需要的字节数
	var result []byte

	if val <= 0xFF {
		result = []byte{byte(val)}
	} else if val <= 0xFFFF {
		result = make([]byte, 2)
		binary.BigEndian.PutUint16(result, uint16(val))

	} else if val <= 0xFFFFFFFF {
		result = make([]byte, 4)
		binary.BigEndian.PutUint32(result, uint32(val))
	} else {
		result = make([]byte, 8)
		binary.BigEndian.PutUint64(result, val)
	}

	return result, nil
}
