package c_device

import "ems-plan/c_base"

type IEnergyStoreBasic interface {
	GetCellMinTemp() (float32, error)    // 电芯最低温度
	GetCellMaxTemp() (float32, error)    // 电芯最高温度
	GetCellAvgTemp() (float32, error)    // 电芯平均温度
	GetCellMinVoltage() (float32, error) // 电芯最低电压
	GetCellMaxVoltage() (float32, error) // 电芯最高电压
	GetCellAvgVoltage() (float32, error) // 电芯平均电压

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
