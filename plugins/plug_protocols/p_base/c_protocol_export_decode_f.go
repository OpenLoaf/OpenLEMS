package p_base

import (
	"common/c_base"
	"p_base/internal/internal_meta"
)

func ECharSequenceDecodeToUint16(e c_base.ECharSequence, b []byte) uint16 {
	return internal_meta.ECharSequenceDecodeToUint16(e, b)
}
