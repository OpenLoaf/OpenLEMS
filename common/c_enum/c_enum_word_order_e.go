package c_base

// EWordOrder 字序 (针对多个16位寄存器之间的顺序)
type EWordOrder int

const (
	WordOrderHighLow EWordOrder = iota // 高字在前，低字在后 (标准Modbus) [AB CD]
	WordOrderLowHigh                   // 低字在前，高字在后 (字交换) [CD AB]
)
