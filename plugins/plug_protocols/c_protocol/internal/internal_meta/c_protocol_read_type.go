package internal_meta

import (
	"bytes"
	"common/c_base"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

func ReadTypeReadValue(d c_base.EReadType, bytes []byte, bitLength uint8, endianness c_base.ECharSequence) (any, error) {
	dataLength := len(bytes)
	switch d {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3, c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7, c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11, c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		// Bit类型的需要先取出对应的位
		if bitLength == 0 {
			bitLength = 1
		}
		// 先找到d的位置，再找到bit end的位置
		start := d / 16
		end := (uint8(d)+bitLength)/16 + 2
		// 0b00000101
		bits := gbinary.DecodeBytesToBits(bytes[start:end])
		_bitsLength := len(bits)
		endBit := _bitsLength - int(d)
		bits = bits[endBit-int(bitLength) : endBit] // 数据倒置
		if bitLength == 1 {
			return bits[0] == 1, nil // 如果只有一个bit，直接返回bool类型
		}
		if bitLength > 0 {
			if bitLength <= 8 {
				//bits = endianness.FillUpSizeBit(bits, 8)
				bits = ECharSequenceFillUpSizeBit(endianness, bits, 8)
				toBytes := gbinary.EncodeBitsToBytes(bits)
				//return endianness.DecodeToUint8(toBytes), nil
				return ECharSequenceDecodeToUint8(endianness, toBytes), nil
			} else if bitLength <= 16 {
				//return endianness.DecodeToUint16(gbinary.EncodeBitsToBytes(bits)), nil
				return ECharSequenceDecodeToUint16(endianness, gbinary.EncodeBitsToBytes(bits)), nil
			} else if bitLength <= 32 {
				//return endianness.DecodeToUint32(gbinary.EncodeBitsToBytes(bits)), nil
				return ECharSequenceDecodeToUint32(endianness, gbinary.EncodeBitsToBytes(bits)), nil
			} else if bitLength <= 64 {
				//return endianness.DecodeToUint64(gbinary.EncodeBitsToBytes(bits)), nil
				return ECharSequenceDecodeToUint64(endianness, gbinary.EncodeBitsToBytes(bits)), nil
			}
		}
		return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
	case c_base.RInt8:
		if dataLength < 1 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToInt8(bytes[:1]), nil
		return ECharSequenceDecodeToInt8(endianness, bytes[:1]), nil
	case c_base.RUint8:
		if dataLength < 1 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToUint8(bytes[:1]), nil
		return ECharSequenceDecodeToUint8(endianness, bytes[:1]), nil
	case c_base.RInt16:
		if dataLength < 2 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToInt16(bytes[:2]), nil
		return ECharSequenceDecodeToInt16(endianness, bytes[:2]), nil
	case c_base.RUint16:
		if dataLength < 2 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToUint16(bytes[:2]), nil
		return ECharSequenceDecodeToUint16(endianness, bytes[:2]), nil
	case c_base.RInt32:
		if dataLength < 4 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToInt32(bytes[:4]), nil
		return ECharSequenceDecodeToInt32(endianness, bytes[:4]), nil
	case c_base.RUint32:
		if dataLength < 4 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToUint32(bytes[:4]), nil
		return ECharSequenceDecodeToUint32(endianness, bytes[:4]), nil
	case c_base.RInt64:
		if dataLength < 8 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToInt64(bytes[:8]), nil
		return ECharSequenceDecodeToInt64(endianness, bytes[:8]), nil
	case c_base.RUint64:
		if dataLength < 8 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		//return endianness.DecodeToUint64(bytes[:8]), nil
		return ECharSequenceDecodeToUint64(endianness, bytes[:8]), nil
	case c_base.RFloat32:
		if dataLength < 4 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return ECharSequenceDecodeToFloat32(endianness, bytes[:4]), nil
	case c_base.RFloat64:
		if dataLength < 8 {
			return nil, fmt.Errorf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return ECharSequenceDecodeToFloat64(endianness, bytes[:8]), nil
	case c_base.RBcd16:
		return BcdToDecimalMulti(bytes[:2]), nil
	}
	panic(`unknown data type`)
}

func ReadTypeTransform(d c_base.EReadType, value any, bitLength uint8, factor float32, offset int) any {
	switch d {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3, c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7, c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11, c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		if bitLength == 1 || bitLength == 0 {
			v := gconv.Int(value)
			return v == 1
		} else if bitLength <= 8 {
			return Calc(gconv.Uint8(value), factor, offset)
		} else if bitLength <= 16 {
			return Calc(gconv.Uint16(value), factor, offset)
		} else if bitLength <= 32 {
			return Calc(gconv.Uint32(value), factor, offset)
		} else if bitLength <= 64 {
			return Calc(gconv.Uint64(value), factor, offset)
		}
	case c_base.RInt8:
		return Calc[int8](gconv.Int8(value), factor, offset)
	case c_base.RUint8:
		return Calc[uint8](gconv.Uint8(value), factor, offset)
	case c_base.RInt16:
		return Calc[int16](gconv.Int16(value), factor, offset)
	case c_base.RUint16:
		return Calc[uint16](gconv.Uint16(value), factor, offset)
	case c_base.RBcd16:
		return Calc[uint16](gconv.Uint16(value), factor, offset)
	case c_base.RInt32:
		return Calc[int32](gconv.Int32(value), factor, offset)
	case c_base.RUint32:
		return Calc[uint32](gconv.Uint32(value), factor, offset)
	case c_base.RInt64:
		return Calc[int64](gconv.Int64(value), factor, offset)
	case c_base.RUint64:
		return Calc[uint64](gconv.Uint64(value), factor, offset)
	case c_base.RFloat32:
		return Calc[float32](gconv.Float32(value), factor, offset)
	case c_base.RFloat64:
		return Calc[float64](gconv.Float64(value), factor, offset)
	default:
		return value
	}
	panic(`unknown data type`)
}

// ReadTypeRegisterSize 寄存器大小
func ReadTypeRegisterSize(d c_base.EReadType) uint16 {
	// 一个寄存器是16位，所以一个寄存器是2个字节
	switch d {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3, c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7, c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11, c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		return 1
	case c_base.RInt8, c_base.RUint8:
		return 1
	case c_base.RInt16, c_base.RUint16:
		return 1
	case c_base.RInt32, c_base.RUint32, c_base.RFloat32:
		return 2
	case c_base.RInt64, c_base.RUint64, c_base.RFloat64:
		return 4
	case c_base.RBcd16:
		return 2
	}
	panic(`unknown data type`)
}

func ReadTypeEncoder(d c_base.EReadType, value int64, factor float32, offset int, endianness c_base.ECharSequence) []byte {
	if factor != 0 {
		value = int64(float32(value)/factor) - int64(offset)
	} else {
		value = value - int64(offset)
	}

	switch d {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3, c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7, c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11, c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		panic(`bit type not support encoder`)
	case c_base.RInt8:
		return ECharSequenceEncodeInt8(endianness, int8(value))
	case c_base.RUint8:
		return ECharSequenceEncodeUint8(endianness, uint8(value))
	case c_base.RInt16:
		return ECharSequenceEncodeInt16(endianness, int16(value))
	case c_base.RUint16:
		return ECharSequenceEncodeUint16(endianness, uint16(value))
	case c_base.RInt32:
		return ECharSequenceEncodeInt32(endianness, int32(value))
	case c_base.RUint32:
		return ECharSequenceEncodeUint32(endianness, uint32(value))
	case c_base.RInt64:
		return ECharSequenceEncodeInt64(endianness, value)
	case c_base.RUint64:
		return ECharSequenceEncodeUint64(endianness, uint64(value))
	case c_base.RFloat32:
		return ECharSequenceEncodeFloat32(endianness, float32(value))
	case c_base.RFloat64:
		return ECharSequenceEncodeFloat64(endianness, float64(value))
	case c_base.RBcd16:
		return DecimalToBCD16Bytes(int(value))
	}
	panic(`unknown data type`)
}

func ReadTypeGetReflectKind(d c_base.EReadType, bitLength uint8) reflect.Kind {
	switch d {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3, c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7, c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11, c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		if bitLength == 1 || bitLength == 0 {
			return reflect.Bool
		} else if bitLength <= 8 {
			return reflect.Uint8
		} else if bitLength <= 16 {
			return reflect.Uint16
		} else if bitLength <= 32 {
			return reflect.Uint32
		} else if bitLength <= 64 {
			return reflect.Uint64
		}
	case c_base.RInt8:
		return reflect.Int8
	case c_base.RUint8:
		return reflect.Uint8
	case c_base.RInt16:
		return reflect.Int16
	case c_base.RUint16:
		return reflect.Uint16
	case c_base.RInt32:
		return reflect.Int32
	case c_base.RUint32:
		return reflect.Uint32
	case c_base.RInt64:
		return reflect.Int64
	case c_base.RUint64:
		return reflect.Uint64
	case c_base.RFloat32:
		return reflect.Float32
	case c_base.RFloat64:
		return reflect.Float64
	case c_base.RBcd16:
		return reflect.Int16
	}
	panic(`unknown data type`)
}

func formatHex(b []byte) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < len(b); i++ {
		buffer.WriteString(fmt.Sprintf("%X", b[i]))
		if i != len(b)-1 {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("]\n")

	return buffer.String()
}

// DecimalToBCD16 将一个最多四位的十进制数转换为16位BCD码
func DecimalToBCD16(n int) uint16 {
	var bcd uint16 = 0
	bcd |= uint16((n / 1000) << 12)      // 最高4位
	bcd |= uint16(((n / 100) % 10) << 8) // 次高4位
	bcd |= uint16(((n / 10) % 10) << 4)  // 次低4位
	bcd |= uint16(n % 10)                // 最低4位
	return bcd
}

// Bcd16ToDecimal 将16位BCD码转换为十进制数
func Bcd16ToDecimal(bcd uint16) int {
	// 通过移位和按位操作提取每个4位数
	thousands := int((bcd >> 12) & 0xF) // 最高4位
	hundreds := int((bcd >> 8) & 0xF)   // 次高4位
	tens := int((bcd >> 4) & 0xF)       // 次低4位
	ones := int(bcd & 0xF)              // 最低4位

	// 将每个提取的十进制数还原成完整的十进制数
	return thousands*1000 + hundreds*100 + tens*10 + ones
}

// DecimalToBCD16Bytes 将一个最多四位的十进制数转换为16位BCD码，并返回 []byte
func DecimalToBCD16Bytes(n int) []byte {
	var bcd uint16 = 0
	bcd |= uint16((n / 1000) << 12)      // 最高4位
	bcd |= uint16(((n / 100) % 10) << 8) // 次高4位
	bcd |= uint16(((n / 10) % 10) << 4)  // 次低4位
	bcd |= uint16(n % 10)                // 最低4位

	// 将16位的BCD拆分为2个字节
	highByte := byte(bcd >> 8)  // 高8位
	lowByte := byte(bcd & 0xFF) // 低8位

	return []byte{highByte, lowByte}
}

// BcdToDecimalMulti 处理多个字节的 BCD 编码
func BcdToDecimalMulti(bcd []byte) int {
	result := 0
	for _, byteVal := range bcd {
		result = result*100 + BcdToDecimal(byteVal)
	}
	return result
}

// BcdToDecimal 将 BCD 格式转换为十进制整数
func BcdToDecimal(bcd byte) int {
	high := int(bcd >> 4)  // 取出高 4 位
	low := int(bcd & 0x0F) // 取出低 4 位
	return high*10 + low
}
