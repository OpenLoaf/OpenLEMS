package common

import (
	"common/c_base"
)

type IDeviceManager interface {
	Start()                      // 启动服务
	Shutdown()                   // 停止管理器（释放资源、退出 goroutine）
	Cleanup() error              // 清理过期/无效资源（定时调用）
	Status() c_enum.EServerState // 运行状态

	GetDeviceById(deviceId string) c_base.IDevice              // 通过设备ID获取设备
	GetDeviceNameById(deviceId string) string                  // 获取设备名称
	GetDeviceConfigById(deviceId string) *c_base.SDeviceConfig //获取指定驱动的详细信息

	IteratorAllDevices(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool)
	IteratorChildDevicesById(deviceId string, iterator func(config *c_base.SDeviceConfig, device c_base.IDevice) bool)  //  按设备ID遍历该设备及其所有子设备
	IteratorParentDevicesById(deviceId string, iterator func(config *c_base.SDeviceConfig, device c_base.IDevice) bool) // 递归当前设备和所有父设备

	GetDeviceConfigTree() []*c_base.SDeviceConfig // 获取顶层的设备列表
	GetAllDriversInfo() []*c_base.SDriverInfo     // 获取所有驱动的详细信息
	GetDriverInfo(driverName string) (*c_base.SDriverInfo, error)

	IsProtocolActive(protocolId string) bool // 协议是否激活
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
