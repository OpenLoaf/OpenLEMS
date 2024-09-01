package c_station

import (
	"ems-plan/c_device"
	"time"
)

type IStationEnergyStore interface {
	IStation
	c_device.IEnergyStoreBasic

	GetLastUpdateTime() *time.Time
	GetChildren() []c_device.IEnergyStore
}
