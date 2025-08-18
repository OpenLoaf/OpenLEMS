package internal_meta

import (
	"common/c_base"
	"common/c_util"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

func SystemTypeTransform(s c_base.ESystemType, value any, readType c_base.EReadType, bitLength uint8, factor float32, offset int) any {

	switch s {
	case c_base.SUseReadType:
		return ReadTypeTransform(readType, value, bitLength, factor, offset) // 使用读取的类型
	case c_base.SInt32:
		return Calc(gconv.Int32(value), factor, offset)
	case c_base.SBool:
		return Calc(gconv.Int8(value), factor, offset) != 0
	case c_base.SInt8:
		return Calc(gconv.Int8(value), factor, offset)
	case c_base.SUint8:
		return Calc(gconv.Uint8(value), factor, offset)
	case c_base.SInt16:
		return Calc(gconv.Int16(value), factor, offset)
	case c_base.SUint16:
		return Calc(gconv.Uint16(value), factor, offset)
	case c_base.SUint32:
		return Calc(gconv.Uint32(value), factor, offset)
	case c_base.SInt64:
		return Calc(gconv.Int64(value), factor, offset)
	case c_base.SUint64:
		return Calc(gconv.Uint64(value), factor, offset)
	case c_base.SFloat32:
		return Calc(gconv.Float32(value), factor, offset)
	case c_base.SFloat64:
		return Calc(gconv.Float64(value), factor, offset)
	case c_base.SString:
		return gconv.String(value)

	default:
		return value
	}
}

func SystemTypeGetReflectKind(s c_base.ESystemType, readType c_base.EReadType, bitLength uint8) reflect.Kind {
	switch s {
	case c_base.SUseReadType:

		return ReadTypeGetReflectKind(readType, bitLength)
	case c_base.SBool:
		return reflect.Bool
	case c_base.SInt8:
		return reflect.Int8
	case c_base.SUint8:
		return reflect.Uint8
	case c_base.SInt16:
		return reflect.Int16
	case c_base.SUint16:
		return reflect.Uint16
	case c_base.SInt32:
		return reflect.Int32
	case c_base.SUint32:
		return reflect.Uint32
	case c_base.SInt64:
		return reflect.Int64
	case c_base.SUint64:
		return reflect.Uint64
	case c_base.SFloat32:
		return reflect.Float32
	case c_base.SFloat64:
		return reflect.Float64
	case c_base.SString:
		return reflect.String
	}
	panic(gerror.New("未知的SystemType类型！"))
}

func Calc[T c_util.Number](result T, factor float32, offset int) T {
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
