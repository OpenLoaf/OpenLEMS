package c_device

import "common/c_base"

type IDriverInstances interface {
	RegisterInstance(info c_base.IDevice)

	// FindAll 获取所有设备实例, 参数为空时获取所有设备实例, 参数为true时获取虚拟设备实例, 参数为false时获取实体设备实例
	FindAll(isVirtual ...bool) []c_base.IDevice

	FindById(id string) c_base.IDevice

	FindByType(t c_base.EDeviceType) []c_base.IDevice

	RemoveById(id string)

	GetStationEnergyStore() IStationEnergyStore
}
