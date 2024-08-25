package c_group

type IInfo interface {
	// IConfig 设备配置信息
	IConfig

	AllowControl() bool // 是否允许控制

	// GetFunctionList 获取设备功能列表
	GetFunctionList() []*SFunction
}
