package c_meta

import (
	"ems-plan/util/util_binary"
	"github.com/gogf/gf/v2/encoding/gbinary"
)

func (e ECharSequence) EncodeString(s string) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeString(s)
	case ELittleEndian:
		return gbinary.LeEncodeString(s)
	case EMiddleEndian:
		return util_binary.MeEncodeString(s)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeString(s)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeInt(i int) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeInt(i)
	case ELittleEndian:
		return gbinary.LeEncodeInt(i)
	case EMiddleEndian:
		return util_binary.MeEncodeInt(i)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeInt(i)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeUint(u uint) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeUint(u)
	case ELittleEndian:
		return gbinary.LeEncodeUint(u)
	case EMiddleEndian:
		return util_binary.MeEncodeUint(u)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeUint(u)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeBool(b bool) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeBool(b)
	case ELittleEndian:
		return gbinary.LeEncodeBool(b)
	case EMiddleEndian:
		return util_binary.MeEncodeBool(b)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeBool(b)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeInt8(i int8) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeInt8(i)
	case ELittleEndian:
		return gbinary.LeEncodeInt8(i)
	case EMiddleEndian:
		return util_binary.MeEncodeInt8(i)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeInt8(i)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeUint8(u uint8) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeUint8(u)
	case ELittleEndian:
		return gbinary.LeEncodeUint8(u)
	case EMiddleEndian:
		return util_binary.MeEncodeUint8(u)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeUint8(u)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeInt16(i int16) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeInt16(i)
	case ELittleEndian:
		return gbinary.LeEncodeInt16(i)
	case EMiddleEndian:
		return util_binary.MeEncodeInt16(i)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeInt16(i)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeUint16(u uint16) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeUint16(u)
	case ELittleEndian:
		return gbinary.LeEncodeUint16(u)
	case EMiddleEndian:
		return util_binary.MeEncodeUint16(u)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeUint16(u)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeInt32(i int32) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeInt32(i)
	case ELittleEndian:
		return gbinary.LeEncodeInt32(i)
	case EMiddleEndian:
		return util_binary.MeEncodeInt32(i)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeInt32(i)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeUint32(u uint32) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeUint32(u)
	case ELittleEndian:
		return gbinary.LeEncodeUint32(u)
	case EMiddleEndian:
		return util_binary.MeEncodeUint32(u)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeUint32(u)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeInt64(i int64) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeInt64(i)
	case ELittleEndian:
		return gbinary.LeEncodeInt64(i)
	case EMiddleEndian:
		return util_binary.MeEncodeInt64(i)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeInt64(i)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeUint64(u uint64) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeUint64(u)
	case ELittleEndian:
		return gbinary.LeEncodeUint64(u)
	case EMiddleEndian:
		return util_binary.MeEncodeUint64(u)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeUint64(u)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeFloat32(f float32) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeFloat32(f)
	case ELittleEndian:
		return gbinary.LeEncodeFloat32(f)
	case EMiddleEndian:
		return util_binary.MeEncodeFloat32(f)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeFloat32(f)
	default:
		panic("不支持的Endianness类型！")
	}
}

func (e ECharSequence) EncodeFloat64(f float64) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeEncodeFloat64(f)
	case ELittleEndian:
		return gbinary.LeEncodeFloat64(f)
	case EMiddleEndian:
		return util_binary.MeEncodeFloat64(f)
	case EReverseMiddleEndian:
		return util_binary.RmeEncodeFloat64(f)
	default:
		panic("不支持的Endianness类型！")
	}
}
