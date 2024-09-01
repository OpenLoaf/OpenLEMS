package c_station

import (
	"ems-plan/c_base"
	"ems-plan/c_device"
)

type IStationEnergyStore interface {
	c_base.IDriver
	c_device.IEnergyStoreBasic

	GetAllowControl() bool // 是否允许控制
	GetChildren() []c_base.IDriver
}
