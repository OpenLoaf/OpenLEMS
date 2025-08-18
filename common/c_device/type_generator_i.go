package c_device

import "common/c_base"

type IGeneratorBasic interface {
	GetGridFrequency() (float32, error) // 电网频率
	GetUa() (float32, error)            // A相电压
	GetUb() (float32, error)            // B相电压
	GetUc() (float32, error)            // C相电压
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

	GetRatedPower() uint32               // 额定功率， -1代表未知
	GetMaxInputPower() (float64, error)  // 最大充电功率、最大输入功率限制
	GetMaxOutputPower() (float64, error) // 最大放电功率、最大输出功率限制
}

type IGenerator interface {
	c_base.IDevice
	IGeneratorBasic
}
