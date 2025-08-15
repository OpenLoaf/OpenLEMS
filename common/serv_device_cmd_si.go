package common

import "common/c_base"

type IDeviceCmd interface {
	// Start 启动EMS 服务
	Start(activeDeviceRootId string)

	// Stop 停止EMS服务
	Stop()

	// InitDriver 初始化驱动
	InitDriver(clientCache map[string]any, config *c_base.SDriverConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDriver
}
