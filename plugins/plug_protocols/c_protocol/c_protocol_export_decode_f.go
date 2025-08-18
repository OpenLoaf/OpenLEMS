package c_protocol

import (
	"c_protocol/internal/internal_meta"
	"common/c_base"
)

func ECharSequenceDecodeToUint16(e c_base.ECharSequence, b []byte) uint16 {
	return internal_meta.ECharSequenceDecodeToUint16(e, b)
}
