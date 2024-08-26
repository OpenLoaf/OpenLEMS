package c_device

import "ems-plan/c_base"

type IEnergyStoreBasic interface {
	IBmsBasic
	IPcsBasic
	IFireBasic
}

type IEnergyStore interface {
	c_base.IDriver
	IEnergyStoreBasic
}
