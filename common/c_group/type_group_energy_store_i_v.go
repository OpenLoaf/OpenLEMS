package c_group

import "ems-plan/c_device"

type IGroupEnergyStore interface {
	IInfo
	c_device.IEnergyStoreBasic

	GetChildren() []c_device.IEnergyStore
}
