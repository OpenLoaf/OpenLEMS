package c_device

import "ems-plan/c_base"

type IAmmeterBasic interface {
	GetVa() (float32, error) // A相电压
	GetVb() (float32, error) // B相电压
	GetVc() (float32, error) // C相电压
	GetIa() (float32, error) // A相电流
	GetIb() (float32, error) // B相电流
	GetIc() (float32, error) // C相电流
	GetPa() (float32, error) // A相有功功率
	GetPb() (float32, error) // B相有功功率
	GetPc() (float32, error) // C相有功功率

	GetPtCt() (float32, float32, error) // PT CT
	GetFrequency() (float32, error)     // 频率
	GetPowerFactor() (float32, error)   // 功率因数

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量

}

type IAmmeter interface {
	c_base.IDriver
	IAmmeterBasic
}
