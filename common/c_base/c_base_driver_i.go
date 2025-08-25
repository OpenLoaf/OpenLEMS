package c_base

type IDriver interface {
	IDevice

	Init() error               // 初始化
	Shutdown()                 // 关闭
	GetConfig() *SDeviceConfig // 获取配置
}
