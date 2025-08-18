package internal_meta

import (
	"c_protocol/internal/util_binary"
	"common/c_base"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
)

func ECharSequenceFillUpSize(e c_base.ECharSequence, b []byte, l int) []byte {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeFillUpSize(b, l)
	case c_base.ELittleEndian:
		return gbinary.LeFillUpSize(b, l)
	case c_base.EMiddleEndian:
		return util_binary.MeFillUpSize(b, l)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeFillUpSize(b, l)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceFillUpSizeBit(e c_base.ECharSequence, b []gbinary.Bit, l int) []gbinary.Bit {
	switch e {
	case c_base.EBigEndian, c_base.EReverseMiddleEndian:
		return util_binary.FillUpBitSizeLeft(b, l)
	case c_base.ELittleEndian, c_base.EMiddleEndian:
		return util_binary.FillUpBitSizeRight(b, l)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToString(e c_base.ECharSequence, b []byte) string {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToString(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToString(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToString(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToString(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToInt(e c_base.ECharSequence, b []byte) int {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToInt(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToInt(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToInt(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToUint(e c_base.ECharSequence, b []byte) uint {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToUint(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToUint(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToUint(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToBool(e c_base.ECharSequence, b []byte) bool {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToBool(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToBool(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToBool(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToBool(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToInt8(e c_base.ECharSequence, b []byte) int8 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToInt8(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToInt8(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToInt8(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt8(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToUint8(e c_base.ECharSequence, b []byte) uint8 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToUint8(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToUint8(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToUint8(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint8(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToInt16(e c_base.ECharSequence, b []byte) int16 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToInt16(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToInt16(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToInt16(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt16(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToUint16(e c_base.ECharSequence, b []byte) uint16 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToUint16(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToUint16(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToUint16(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint16(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToInt32(e c_base.ECharSequence, b []byte) int32 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToInt32(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToInt32(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToInt32(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToUint32(e c_base.ECharSequence, b []byte) uint32 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToUint32(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToUint32(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToUint32(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToInt64(e c_base.ECharSequence, b []byte) int64 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToInt64(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToInt64(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToInt64(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToInt64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToUint64(e c_base.ECharSequence, b []byte) uint64 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToUint64(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToUint64(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToUint64(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToUint64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToFloat32(e c_base.ECharSequence, b []byte) float32 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToFloat32(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToFloat32(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToFloat32(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToFloat32(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func ECharSequenceDecodeToFloat64(e c_base.ECharSequence, b []byte) float64 {
	switch e {
	case c_base.EBigEndian:
		return gbinary.BeDecodeToFloat64(b)
	case c_base.ELittleEndian:
		return gbinary.LeDecodeToFloat64(b)
	case c_base.EMiddleEndian:
		return util_binary.MeDecodeToFloat64(b)
	case c_base.EReverseMiddleEndian:
		return util_binary.RmeDecodeToFloat64(b)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}
