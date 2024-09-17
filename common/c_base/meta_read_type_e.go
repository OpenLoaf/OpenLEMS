//go:generate stringer -type=EReadType -output=meta_read_type_e_string.go
package c_base

import (
	"bytes"
	"common/util"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

type (
	EReadType int // 读取数据类型
)

const (
	RBit0 EReadType = iota // Bit类型的读取该地址的第x位，比如0x1000读取到数值为 1001 0010，R_Bit_1的值为1
	RBit1                  // Bit类型的读取到的值，如果BitLength为1，类型就是Bool，否则就是根据长度：Uint8 、Uint16、Uint32、Uint64自动扩展
	RBit2
	RBit3
	RBit4
	RBit5
	RBit6
	RBit7
	RBit8
	RBit9
	RBit10
	RBit11
	RBit12
	RBit13
	RBit14
	RBit15

	RBcd16 // 16位BCD码

	RInt8
	RUint8
	RInt16
	RUint16
	RInt32
	RUint32
	RInt64
	RUint64
	RFloat32
	RFloat64
)

func (d EReadType) ReadValue(bytes []byte, bitLength uint8, endianness ECharSequence) (any, error) {
	dataLength := len(bytes)
	switch d {
	case RBit0, RBit1, RBit2, RBit3, RBit4, RBit5, RBit6, RBit7, RBit8, RBit9, RBit10, RBit11, RBit12, RBit13, RBit14, RBit15:
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
		bits = bits[endBit-int(bitLength) : endBit] // 倒过来取
		if bitLength == 1 {
			return bits[0] == 1, nil // 如果只有一个bit，直接返回bool类型
		}

		if bitLength > 0 {
			if bitLength <= 8 {
				bits = endianness.FillUpSizeBit(bits, 8)
				toBytes := gbinary.EncodeBitsToBytes(bits)
				return endianness.DecodeToUint8(toBytes), nil
			} else if bitLength <= 16 {
				return endianness.DecodeToUint16(gbinary.EncodeBitsToBytes(bits)), nil
			} else if bitLength <= 32 {
				return endianness.DecodeToUint32(gbinary.EncodeBitsToBytes(bits)), nil
			} else if bitLength <= 64 {
				return endianness.DecodeToUint64(gbinary.EncodeBitsToBytes(bits)), nil
			}
		}

		return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
	case RInt8:
		if dataLength < 1 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToInt8(bytes[:1]), nil
	case RUint8:
		if dataLength < 1 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToUint8(bytes[:1]), nil
	case RInt16:
		if dataLength < 2 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToInt16(bytes[:2]), nil
	case RUint16:
		if dataLength < 2 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToUint16(bytes[:2]), nil
	case RInt32:
		if dataLength < 4 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToInt32(bytes[:4]), nil
	case RUint32:
		if dataLength < 4 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToUint32(bytes[:4]), nil
	case RInt64:
		if dataLength < 8 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToInt64(bytes[:8]), nil
	case RUint64:
		if dataLength < 8 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToUint64(bytes[:8]), nil
	case RFloat32:
		if dataLength < 4 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToFloat32(bytes[:4]), nil
	case RFloat64:
		if dataLength < 8 {
			return nil, gerror.Newf("ReadValue失败！读取类型:%s 字符序列:%s 获取到的数据为：%s", d, endianness, formatHex(bytes))
		}
		return endianness.DecodeToFloat64(bytes[:8]), nil
	case RBcd16:
		return util.BcdToDecimalMulti(bytes[:2]), nil
	}
	panic(`unknown data type`)
}

func (d EReadType) Transform(value any, bitLength uint8, factor float32, offset int) any {
	switch d {
	case RBit0, RBit1, RBit2, RBit3, RBit4, RBit5, RBit6, RBit7, RBit8, RBit9, RBit10, RBit11, RBit12, RBit13, RBit14, RBit15:
		if bitLength == 1 || bitLength == 0 {
			v := gconv.Int(value)
			return v == 1
		} else if bitLength <= 8 {
			return calc(gconv.Uint8(value), factor, offset)
		} else if bitLength <= 16 {
			return calc(gconv.Uint16(value), factor, offset)
		} else if bitLength <= 32 {
			return calc(gconv.Uint32(value), factor, offset)
		} else if bitLength <= 64 {
			return calc(gconv.Uint64(value), factor, offset)
		}
	case RInt8:
		return calc[int8](gconv.Int8(value), factor, offset)
	case RUint8:
		return calc[uint8](gconv.Uint8(value), factor, offset)
	case RInt16:
		return calc[int16](gconv.Int16(value), factor, offset)
	case RUint16:
		return calc[uint16](gconv.Uint16(value), factor, offset)
	case RBcd16:
		return calc[uint16](gconv.Uint16(value), factor, offset)
	case RInt32:
		return calc[int32](gconv.Int32(value), factor, offset)
	case RUint32:
		return calc[uint32](gconv.Uint32(value), factor, offset)
	case RInt64:
		return calc[int64](gconv.Int64(value), factor, offset)
	case RUint64:
		return calc[uint64](gconv.Uint64(value), factor, offset)
	case RFloat32:
		return calc[float32](gconv.Float32(value), factor, offset)
	case RFloat64:
		return calc[float64](gconv.Float64(value), factor, offset)
	default:
		return value
	}
	panic(`unknown data type`)
}

// RegisterSize 寄存器大小
func (d EReadType) RegisterSize() uint16 {
	// 一个寄存器是16位，所以一个寄存器是2个字节
	switch d {
	case RBit0, RBit1, RBit2, RBit3, RBit4, RBit5, RBit6, RBit7, RBit8, RBit9, RBit10, RBit11, RBit12, RBit13, RBit14, RBit15:
		return 1
	case RInt8, RUint8:
		return 1
	case RInt16, RUint16:
		return 1
	case RInt32, RUint32, RFloat32:
		return 2
	case RInt64, RUint64, RFloat64:
		return 4
	case RBcd16:
		return 2
	}
	panic(`unknown data type`)
}

func (d EReadType) Encoder(value int64, factor float32, offset int, endianness ECharSequence) []byte {
	if factor != 0 {
		value = int64(float32(value)/factor) - int64(offset)
	} else {
		value = value - int64(offset)
	}

	switch d {
	case RBit0, RBit1, RBit2, RBit3, RBit4, RBit5, RBit6, RBit7, RBit8, RBit9, RBit10, RBit11, RBit12, RBit13, RBit14, RBit15:
		panic(`bit type not support encoder`)
	case RInt8:
		return endianness.EncodeInt8(int8(value))
	case RUint8:
		return endianness.EncodeUint8(uint8(value))
	case RInt16:
		return endianness.EncodeInt16(int16(value))
	case RUint16:
		return endianness.EncodeUint16(uint16(value))
	case RInt32:
		return endianness.EncodeInt32(int32(value))
	case RUint32:
		return endianness.EncodeUint32(uint32(value))
	case RInt64:
		return endianness.EncodeInt64(value)
	case RUint64:
		return endianness.EncodeUint64(uint64(value))
	case RFloat32:
		return endianness.EncodeFloat32(float32(value))
	case RFloat64:
		return endianness.EncodeFloat64(float64(value))
	case RBcd16:
		return util.DecimalToBCD16Bytes(int(value))
	}
	panic(`unknown data type`)
}

func (d EReadType) GetReflectKind(bitLength uint8) reflect.Kind {
	switch d {
	case RBit0, RBit1, RBit2, RBit3, RBit4, RBit5, RBit6, RBit7, RBit8, RBit9, RBit10, RBit11, RBit12, RBit13, RBit14, RBit15:
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
	case RInt8:
		return reflect.Int8
	case RUint8:
		return reflect.Uint8
	case RInt16:
		return reflect.Int16
	case RUint16:
		return reflect.Uint16
	case RInt32:
		return reflect.Int32
	case RUint32:
		return reflect.Uint32
	case RInt64:
		return reflect.Int64
	case RUint64:
		return reflect.Uint64
	case RFloat32:
		return reflect.Float32
	case RFloat64:
		return reflect.Float64
	case RBcd16:
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
