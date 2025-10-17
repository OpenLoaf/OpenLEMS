//go:generate stringer -type=EValueType -trimprefix=E -output=c_enum_point_system_type_e_string.go
package c_enum

type EValueType int // 读取到数据后，转换为到系统类型

const (
	EInt16 EValueType = iota // 默认为int16
	EBool
	EInt8
	EUint8
	EUint16
	EInt32
	EUint32
	EInt64
	EUint64
	EFloat32
	EFloat64
	EString
)
