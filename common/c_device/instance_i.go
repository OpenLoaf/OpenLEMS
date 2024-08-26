package c_device

import "ems-plan/c_base"

type IDriverInstances interface {
	RegisterInstance(info c_base.IDriver)

	FindById(id string) c_base.IDriver

	FindAll() []c_base.IDriver

	FindByType(t c_base.EDeviceType) []c_base.IDriver

	RemoveById(id string)
}
