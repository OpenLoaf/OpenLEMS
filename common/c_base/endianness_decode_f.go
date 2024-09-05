package c_base

import (
	"ems-plan/util/util_binary"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (e ECharSequence) DecodeToString(b []byte) string {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToString(b)
	case ELittleEndian:
		return gbinary.LeDecodeToString(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToString(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToString(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToInt(b []byte) int {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToInt(b)
	case ELittleEndian:
		return gbinary.LeDecodeToInt(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToInt(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToUint(b []byte) uint {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToUint(b)
	case ELittleEndian:
		return gbinary.LeDecodeToUint(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToUint(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToBool(b []byte) bool {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToBool(b)
	case ELittleEndian:
		return gbinary.LeDecodeToBool(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToBool(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToBool(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToInt8(b []byte) int8 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToInt8(b)
	case ELittleEndian:
		return gbinary.LeDecodeToInt8(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToInt8(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt8(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToUint8(b []byte) uint8 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToUint8(b)
	case ELittleEndian:
		return gbinary.LeDecodeToUint8(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToUint8(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint8(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToInt16(b []byte) int16 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToInt16(b)
	case ELittleEndian:
		return gbinary.LeDecodeToInt16(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToInt16(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt16(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToUint16(b []byte) uint16 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToUint16(b)
	case ELittleEndian:
		return gbinary.LeDecodeToUint16(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToUint16(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint16(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToInt32(b []byte) int32 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToInt32(b)
	case ELittleEndian:
		return gbinary.LeDecodeToInt32(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToInt32(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToUint32(b []byte) uint32 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToUint32(b)
	case ELittleEndian:
		return gbinary.LeDecodeToUint32(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToUint32(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToInt64(b []byte) int64 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToInt64(b)
	case ELittleEndian:
		return gbinary.LeDecodeToInt64(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToInt64(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToUint64(b []byte) uint64 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToUint64(b)
	case ELittleEndian:
		return gbinary.LeDecodeToUint64(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToUint64(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToFloat32(b []byte) float32 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToFloat32(b)
	case ELittleEndian:
		return gbinary.LeDecodeToFloat32(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToFloat32(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToFloat32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) DecodeToFloat64(b []byte) float64 {
	switch e {
	case EBigEndian:
		return gbinary.BeDecodeToFloat64(b)
	case ELittleEndian:
		return gbinary.LeDecodeToFloat64(b)
	case EMiddleEndian:
		return util_binary.MeDecodeToFloat64(b)
	case EReverseMiddleEndian:
		return util_binary.RmeDecodeToFloat64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}
