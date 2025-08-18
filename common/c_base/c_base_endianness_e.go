//go:generate stringer -type=ECharSequence -output=c_base_endianness_e_string.go
package c_base

type ECharSequence uint8

const (
	EBigEndian           ECharSequence = iota // AB CD 大端是
	ELittleEndian                             // DC BA 小端是
	EMiddleEndian                             // CD AB 中端序是
	EReverseMiddleEndian                      // BA DC 反中端序是
)
