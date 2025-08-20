//go:generate stringer -type=EReadType -trimprefix=R -output=c_base_meta_read_type_e_string.go
package c_base

type (
	EReadType int // 读取数据类型
)

const (
	RBit0 EReadType = iota // Bit类型的读取该地址的第x位，比如0x1000读取到数值为 1001 0010，R_Bit_1的值为1
	RBit1                  // Bit类型的读取到的值，如果BitLength为1，类型就是Bool，否则就是根据长度：Uint8 、Uint16、Uint32、Uint64自动扩展
	RBit2
	RBit3
	RBit4
	RBit5
	RBit6
	RBit7
	RBit8
	RBit9
	RBit10
	RBit11
	RBit12
	RBit13
	RBit14
	RBit15

	RBcd16 // 16位BCD码

	RInt8
	RUint8
	RInt16
	RUint16
	RInt32
	RUint32
	RInt64
	RUint64
	RFloat32
	RFloat64
)
