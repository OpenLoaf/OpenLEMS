package c_base

type EDeviceStatus int // 设备状态

const (
	EStatusUninitialized EDeviceStatus = iota // 未初始化
	EStatusInitSuccess                        // 成功
	EStatusInitFail                           // 失败
)
