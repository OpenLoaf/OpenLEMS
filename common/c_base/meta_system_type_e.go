package c_base

import (
	"common/c_util"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

type ESystemType int // 读取到数据后，转换为到系统类型

const (
	SUseReadType ESystemType = iota // 自动使用ReadType的类型。
	SBool
	SInt8
	SUint8
	SInt16
	SUint16
	SInt32
	SUint32
	SInt64
	SUint64
	SFloat32
	SFloat64
	SString
)

func (s ESystemType) Transform(value any, readType EReadType, bitLength uint8, factor float32, offset int) any {

	switch s {
	case SUseReadType:
		return readType.Transform(value, bitLength, factor, offset) // 使用读取的类型
	case SInt32:
		return calc(gconv.Int32(value), factor, offset)
	case SBool:
		return calc(gconv.Int8(value), factor, offset) != 0
	case SInt8:
		return calc(gconv.Int8(value), factor, offset)
	case SUint8:
		return calc(gconv.Uint8(value), factor, offset)
	case SInt16:
		return calc(gconv.Int16(value), factor, offset)
	case SUint16:
		return calc(gconv.Uint16(value), factor, offset)
	case SUint32:
		return calc(gconv.Uint32(value), factor, offset)
	case SInt64:
		return calc(gconv.Int64(value), factor, offset)
	case SUint64:
		return calc(gconv.Uint64(value), factor, offset)
	case SFloat32:
		return calc(gconv.Float32(value), factor, offset)
	case SFloat64:
		return calc(gconv.Float64(value), factor, offset)
	case SString:
		return gconv.String(value)

	default:
		return value
	}
}

func (s ESystemType) GetReflectKind(readType EReadType, bitLength uint8) reflect.Kind {
	switch s {
	case SUseReadType:
		return readType.GetReflectKind(bitLength)
	case SBool:
		return reflect.Bool
	case SInt8:
		return reflect.Int8
	case SUint8:
		return reflect.Uint8
	case SInt16:
		return reflect.Int16
	case SUint16:
		return reflect.Uint16
	case SInt32:
		return reflect.Int32
	case SUint32:
		return reflect.Uint32
	case SInt64:
		return reflect.Int64
	case SUint64:
		return reflect.Uint64
	case SFloat32:
		return reflect.Float32
	case SFloat64:
		return reflect.Float64
	case SString:
		return reflect.String
	}
	panic(gerror.New("未知的SystemType类型！"))
}

func calc[T c_util.Number](result T, factor float32, offset int) T {
	// 先乘
	if factor != 1 && factor != 0 {
		result = T(factor * float32(result))
	}
	// 偏移值
	if offset != 0 {
		result = T(offset) + result
	}
	return result
}
