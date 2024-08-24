package c_device

import "ems-plan/c_telemetry"

type IEnergyStoreBasic interface {
	IBmsBasic
	IPcsBasic
	IFireBasic
	GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity
}

type IEnergyStore interface {
	IInfo
	IEnergyStoreBasic
}
