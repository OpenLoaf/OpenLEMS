package c_station

import "ems-plan/c_device"

type IStationEnergyStore interface {
	IStation
	c_device.IEnergyStoreBasic

	GetChildren() []c_device.IEnergyStore
}
