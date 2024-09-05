package c_device

import (
	"ems-plan/c_base"
)

type IStationEntrance interface {
	c_base.IDriver
	GetGridFrequency() (float32, error) // 电网频率
	GetVa() (float32, error)            // A相电压
	GetVb() (float32, error)            // B相电压
	GetVc() (float32, error)            // C相电压
	GetIa() (float32, error)            // A相电流
	GetIb() (float32, error)            // B相电流
	GetIc() (float32, error)            // C相电流
	GetPa() (float32, error)            // A相有功功率
	GetPb() (float32, error)            // B相有功功率
	GetPc() (float32, error)            // C相有功功率

	SetPower(power float64) error         // 设置有功功率
	SetReactivePower(power float64) error // 设置无功功率
	SetPowerFactor(factor float32) error  // 设置功率因数

	GetTargetPower() float64         // 获取目标有功功率
	GetTargetReactivePower() float64 // 获取目标无功功率
	GetTargetPowerFactor() float32   // 获取目标功率因数

	GetPower() (float64, error)         // 有功功率
	GetApparentPower() (float64, error) // 视在功率
	GetReactivePower() (float64, error) // 无功功率

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量

	GetPtCt() (float32, float32, error)             // PT CT
	GetFrequency() (float32, error)                 // 频率
	GetPowerFactor() (float32, error)               // 功率因数
	GetVoltage() (float32, float32, float32, error) // 三相电压
	GetCurrent() (float32, float32, float32, error) // 三相电流

	GetAllowControl() bool // 是否允许控制
	GetChildren() []c_base.IDriver
}
