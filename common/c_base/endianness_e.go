//go:generate stringer -type=ECharSequence -output=endianness_e_string.go
package c_base

import (
	"common/util/util_binary"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ECharSequence uint8

const (
	EBigEndian           ECharSequence = iota // AB CD 大端是
	ELittleEndian                             // DC BA 小端是
	EMiddleEndian                             // CD AB 中端序是
	EReverseMiddleEndian                      // BA DC 反中端序是
)

func (e ECharSequence) FillUpSize(b []byte, l int) []byte {
	switch e {
	case EBigEndian:
		return gbinary.BeFillUpSize(b, l)
	case ELittleEndian:
		return gbinary.LeFillUpSize(b, l)
	case EMiddleEndian:
		return util_binary.MeFillUpSize(b, l)
	case EReverseMiddleEndian:
		return util_binary.RmeFillUpSize(b, l)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}

func (e ECharSequence) FillUpSizeBit(b []gbinary.Bit, l int) []gbinary.Bit {
	switch e {
	case EBigEndian, EReverseMiddleEndian:
		return util_binary.FillUpBitSizeLeft(b, l)
	case ELittleEndian, EMiddleEndian:
		return util_binary.FillUpBitSizeRight(b, l)
	default:
		panic(gerror.New("不支持的Endianness类型！"))
	}
}
