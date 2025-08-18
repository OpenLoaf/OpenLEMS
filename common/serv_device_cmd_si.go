package common

import (
	"common/c_base"
	"context"
)

type IDeviceManager interface {
	IManager

	//InitDriver(clientCache map[string]any, config *c_base.SDeviceConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDevice // 初始化驱动

	GetAllDriverNames() []string                                                       // 获取所有驱动名称
	GetAllDriversInfo(ctx context.Context) []c_base.SDriverInfo                        // 获取所有驱动的详细信息
	GetDriverInfo(ctx context.Context, driverName string) (*c_base.SDriverInfo, error) //获取指定驱动的详细信息

	IsProtocolActive(protocolId string) bool                       // 协议是否激活
	IsDriverAvailable(ctx context.Context, driverName string) bool // 检查驱动是否可用

	//GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.SDriverInfo        //根据设备类型获取驱动信息
	//GetDriverDescription(ctx context.Context, driverName string) (*c_base.SDriverDescription, error) // 获取驱动描述信息
	//GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType                                // 获取支持的设备类型列表

	//CreateDriver(ctx context.Context, deviceConfig *c_base.SDeviceConfig) (c_base.IDevice, error) // 创建驱动实例

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
