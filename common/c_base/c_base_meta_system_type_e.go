package c_base

type ESystemType int // 读取到数据后，转换为到系统类型

const (
	SUseReadType ESystemType = iota // 自动使用ReadType的类型。
	SBool
	SInt8
	SUint8
	SInt16
	SUint16
	SInt32
	SUint32
	SInt64
	SUint64
	SFloat32
	SFloat64
	SString
)
