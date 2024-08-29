package c_device

import "ems-plan/c_base"

type IPcsBasic interface {
	SetReset() error                                  // 复位
	SetStatus(status c_base.EEnergyStoreStatus) error // 设置状态(开机、关机、待机、充电、放电、故障、同步中
	SetGridMode(mode c_base.EGridMode) error          // 设置电网状态
	GetStatus() (c_base.EEnergyStoreStatus, error)    // 状态
	GetGridMode() (c_base.EGridMode, error)           // 电网状态

	SetPower(power int32) error          // 设置有功功率
	SetReactivePower(power int32) error  // 设置无功功率
	SetPowerFactor(factor float32) error // 设置功率因数
	GetTargetPower() int32               // 获取目标有功功率
	GetTargetReactivePower() int32       // 获取目标无功功率
	GetTargetPowerFactor() float32       // 获取目标功率因数
	GetPower() (float64, error)          // 有功功率
	GetApparentPower() (float64, error)  // 视在功率
	GetReactivePower() (float64, error)  // 无功功率

	GetRatedPower() (int32, error)       // 额定功率， -1代表未知
	GetMaxInputPower() (float64, error)  // 最大充电功率、最大输入功率限制
	GetMaxOutputPower() (float64, error) // 最大放电功率、最大输出功率限制

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量
}

type IPcs interface {
	c_base.IDriver
	IPcsBasic
}
