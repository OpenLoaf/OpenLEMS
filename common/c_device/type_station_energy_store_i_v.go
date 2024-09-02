package c_device

import (
	"ems-plan/c_base"
)

type IStationEnergyStore interface {
	c_base.IDriver
	IEnergyStoreBasic

	GetAllowControl() bool // 是否允许控制
	GetChildren() []c_base.IDriver
}
