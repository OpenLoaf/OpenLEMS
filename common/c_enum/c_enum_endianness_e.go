//go:generate stringer -type=ECharSequence -trimprefix=E -output=c_enum_endianness_e_string.go
package c_enum

type ECharSequence uint8 // 字节编码

const (
	EBigEndian           ECharSequence = iota // AB CD 大端 (标准）
	ELittleEndian                             // DC BA 小端是
	EMiddleEndian                             // CD AB 中端序是
	EReverseMiddleEndian                      // BA DC 反中端序是

	EBcd
	EIeee754 // float
)
