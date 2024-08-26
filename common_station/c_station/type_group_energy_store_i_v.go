package c_station

import "ems-plan/c_device"

type IGroupEnergyStore interface {
	IGroup
	c_device.IEnergyStoreBasic

	GetChildren() []c_device.IEnergyStore
}
