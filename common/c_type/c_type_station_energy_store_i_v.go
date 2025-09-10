package c_type

import (
	"common/c_base"
)

type IStationEnergyStore interface {
	c_base.IDriver
	IEnergyStoreBasic

	GetAllowControl() *bool // 是否允许控制
}
