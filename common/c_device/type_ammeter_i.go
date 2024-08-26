package c_device

import "ems-plan/c_base"

type IAmmeterBasic interface {
	GetPtCt() (float32, float32, error)             // PT CT
	GetFrequency() (float32, error)                 // 频率
	GetPowerFactor() (float32, error)               // 功率因数
	GetVoltage() (float32, float32, float32, error) // 三相电压
	GetCurrent() (float32, float32, float32, error) // 三相电流

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量

}

type IAmmeter interface {
	c_base.IDriver
	IAmmeterBasic
}
