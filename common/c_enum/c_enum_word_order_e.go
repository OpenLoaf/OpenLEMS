//go:generate stringer -type=EWordOrder -trimprefix=EWordOrder -output=c_enum_word_order_e_string.go
package c_enum

// EWordOrder 字序 (针对多个16位寄存器之间的顺序)
type EWordOrder int

const (
	WordOrderHighLow EWordOrder = iota // 高字在前，低字在后 (标准Modbus) [AB CD]
	WordOrderLowHigh                   // 低字在前，高字在后 (字交换) [CD AB]
)
