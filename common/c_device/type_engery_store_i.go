package c_device

import "ems-plan/c_base"

type IEnergyStoreBasic interface {
	GetSoc() (float32, error) // 电池当前容量 %
	GetSoh() (float32, error) // 电池健康 %

	GetCapacity() (uint32, error) // 电池容量kWh
	GetCycleCount() (uint, error) // 循环次数

	GetDcPower() (float64, error) // 直流功率

	IPcsBasic
	IFireBasic
}

type IEnergyStore interface {
	c_base.IDriver
	IEnergyStoreBasic
}
