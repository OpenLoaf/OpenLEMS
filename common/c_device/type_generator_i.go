package c_device

import "ems-plan/c_base"

type IGeneratorBasic interface {
	GetGridFrequency() (float32, error)                 // 电网频率
	GetGridVoltage() (float32, float32, float32, error) // 电网电压 A、B、C
	GetGridCurrent() (float32, float32, float32, error) // 电网电流 A、B、C
	GetGridPower() (float32, float32, float32, error)   // 电网功率 A、B、C

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

	GetRatedPower() (float64, error)     // 额定功率， -1代表未知
	GetMaxInputPower() (float64, error)  // 最大充电功率、最大输入功率限制
	GetMaxOutputPower() (float64, error) // 最大放电功率、最大输出功率限制
}

type IGenerator interface {
	c_base.IDriver
	IGeneratorBasic
}
