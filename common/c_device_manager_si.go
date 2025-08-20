package common

import (
	"common/c_base"
	"context"
)

type IDeviceManager interface {
	IManager

	//\(deviceId string) bool
	GetDeviceById(deviceId string) c_base.IDeviceWrapper // 通过设备ID获取设备

	IteratorAssAllDevicesWrapper(deviceWrapper func(device c_base.IDeviceWrapper))

	GetAllDriverNames() []string                                // 获取所有驱动名称
	GetAllDriversInfo(ctx context.Context) []c_base.SDriverInfo // 获取所有驱动的详细信息

	GetDriverInfo(ctx context.Context, driverName string) (*c_base.SDriverInfo, error) //获取指定驱动的详细信息

	IsProtocolActive(protocolId string) bool                       // 协议是否激活
	IsDriverAvailable(ctx context.Context, driverName string) bool // 检查驱动是否可用
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
