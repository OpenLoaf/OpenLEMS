package c_device

import "ems-plan/c_base"

type IDriverInstances interface {
	RegisterInstance(info c_base.IDriver)

	// FindAll 获取所有设备实例, 参数为空时获取所有设备实例, 参数为true时获取虚拟设备实例, 参数为false时获取实体设备实例
	FindAll(isVirtual ...bool) []c_base.IDriver

	FindById(id string) c_base.IDriver

	FindByType(t c_base.EDeviceType) []c_base.IDriver

	RemoveById(id string)
}
