package c_base

type IService interface {
	// Start 启动EMS 服务
	Start(activeDeviceRootId string)

	// Stop 停止EMS服务
	Stop()

	// InitDriver 初始化驱动
	InitDriver(clientCache map[string]any, config *SDriverConfig, protocolConfigList []*SProtocolConfig) IDriver
}
