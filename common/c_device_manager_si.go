package common

import (
	"common/c_base"
	"context"
)

type IDeviceManager interface {
	IManager

	GetAllDriverNames() []string                                                       // 获取所有驱动名称
	GetAllDriversInfo(ctx context.Context) []c_base.SDriverInfo                        // 获取所有驱动的详细信息
	GetDriverInfo(ctx context.Context, driverName string) (*c_base.SDriverInfo, error) //获取指定驱动的详细信息

	IsProtocolActive(protocolId string) bool                       // 协议是否激活
	IsDriverAvailable(ctx context.Context, driverName string) bool // 检查驱动是否可用
	GetDeviceById(deviceId string) c_base.IDevice
}

var deviceManager IDeviceManager

func RegisterDeviceManager(cmd IDeviceManager) {
	deviceManager = cmd
}

func GetDeviceManager() IDeviceManager {
	if deviceManager == nil {
		panic("device manager is nil !")
	}
	return deviceManager
}
