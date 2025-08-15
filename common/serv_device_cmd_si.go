package common

import (
	"common/c_base"
	"common/c_error"
)

type IDeviceCmd interface {
	// Start 启动EMS 服务
	Start(activeDeviceRootId string)

	// Stop 停止EMS服务
	Stop()

	// IsProtocolActive 协议是否激活
	IsProtocolActive(protocolId string) bool

	// InitDriver 初始化驱动
	InitDriver(clientCache map[string]any, config *c_base.SDriverConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDriver
}

var deviceCmd IDeviceCmd

func RegisterDeviceCmd(cmd IDeviceCmd) {
	deviceCmd = cmd
}

func GetDeviceCmd() (IDeviceCmd, error) {
	if deviceCmd == nil {
		return nil, c_error.NonSupportError
	}
	return deviceCmd, nil
}
