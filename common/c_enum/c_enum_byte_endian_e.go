//go:generate stringer -type=EByteEndian -trimprefix=ByteEndian -output=c_enum_byte_endian_e_string.go
package c_enum

// EByteEndian 字节序 (针对一个16位寄存器内部的字节顺序)
type EByteEndian int

const (
	ByteEndianBig    EByteEndian = iota // 大端字节序 (标准Modbus) [AB] -> 高字节A在前
	ByteEndianLittle                    // 小端字节序 [BA] -> 低字节B在前
)
