package internal_meta

import (
	"common/c_base"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/pkg/errors"
	"p_base/internal/util_binary"
)

func ECharSequenceEncodeString(e c_base.ECharSequence, s string) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeString(s)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeString(s)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeString(s)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeString(s)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeInt(e c_base.ECharSequence, i int) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeInt(i)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeInt(i)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeInt(i)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeInt(i)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeUint(e c_base.ECharSequence, u uint) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeUint(u)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeUint(u)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeUint(u)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeUint(u)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeBool(e c_base.ECharSequence, b bool) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeBool(b)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeBool(b)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeBool(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeBool(b)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeInt8(e c_base.ECharSequence, i int8) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeInt8(i)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeInt8(i)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeInt8(i)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeInt8(i)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeUint8(e c_base.ECharSequence, u uint8) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeUint8(u)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeUint8(u)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeUint8(u)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeUint8(u)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeInt16(e c_base.ECharSequence, i int16) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeInt16(i)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeInt16(i)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeInt16(i)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeInt16(i)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeUint16(e c_base.ECharSequence, u uint16) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeUint16(u)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeUint16(u)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeUint16(u)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeUint16(u)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeInt32(e c_base.ECharSequence, i int32) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeInt32(i)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeInt32(i)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeInt32(i)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeInt32(i)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeUint32(e c_base.ECharSequence, u uint32) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeUint32(u)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeUint32(u)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeUint32(u)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeUint32(u)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeInt64(e c_base.ECharSequence, i int64) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeInt64(i)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeInt64(i)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeInt64(i)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeInt64(i)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeUint64(e c_base.ECharSequence, u uint64) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeUint64(u)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeUint64(u)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeUint64(u)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeUint64(u)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeFloat32(e c_base.ECharSequence, f float32) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeFloat32(f)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeFloat32(f)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeFloat32(f)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeFloat32(f)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceEncodeFloat64(e c_base.ECharSequence, f float64) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeEncodeFloat64(f)
	case c_base.ELittleEndian:
		return gbinary.LeEncodeFloat64(f)
	case c_base.EMiddleEndian:
		return util_binary.MeEncodeFloat64(f)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeEncodeFloat64(f)
	default:
		panic(errors.New("不支持的Endianness类型！"))
	}
}
