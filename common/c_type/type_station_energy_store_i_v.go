package c_type

import (
	"common/c_base"
)

type IStationEnergyStore interface {
	c_base.IDevice
	IEnergyStoreBasic

	GetAllowControl() bool // 是否允许控制
	GetChildren() []c_base.IDevice
}
