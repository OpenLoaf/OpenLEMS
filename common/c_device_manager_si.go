package common

import (
	"common/c_base"
)

type IDeviceManager interface {
	Start()                      // 启动服务
	Shutdown()                   // 停止管理器（释放资源、退出 goroutine）
	Cleanup() error              // 清理过期/无效资源（定时调用）
	Status() c_base.EServerState // 运行状态

	GetDeviceById(deviceId string) c_base.IDevice              // 通过设备ID获取设备
	GetDeviceConfigById(deviceId string) *c_base.SDeviceConfig //获取指定驱动的详细信息

	IteratorAssAllDevicesWrapper(deviceWrapper func(config *c_base.SDeviceConfig, device c_base.IDevice) bool)

	GetTopDeviceConfigs() []*c_base.SDeviceConfig // 获取顶层的设备列表
	GetAllDriversInfo() []*c_base.SDriverInfo     // 获取所有驱动的详细信息
	GetDriverInfo(driverName string) (*c_base.SDriverInfo, error)

	IsProtocolActive(protocolId string) bool // 协议是否激活
	//IsDriverAvailable(ctx context.Context, driverName string) bool // 检查驱动是否可用
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
