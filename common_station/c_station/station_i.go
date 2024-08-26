package c_station

type IGroup interface {
	// IStationConfig 设备配置信息
	IStationConfig

	AllowControl() bool // 是否允许控制

	// GetFunctionList 获取设备功能列表
	GetFunctionList() []*SFunction
}
