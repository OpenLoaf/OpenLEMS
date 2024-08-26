package c_device

import "ems-plan/c_base"

type EEnergyStoreStatus int // 储能状态（PCS状态）

type EGridMode int // 电网状态

const (
	EPcsStatusUnknown   EEnergyStoreStatus = iota // 未知状态（也可能是离线）
	EPcsStatusOff                                 // 关机
	EPcsStatusStandby                             // 待机,只有待机状态下，才算是准备运行
	EPcsStatusCharge                              // 充电
	EPcsStatusDischarge                           // 放电
	EPcsStatusFault                               // 故障
	EPcsStatusSync                                // 同步中（多个设备的状态不一致）
)

const (
	EGridUnknown EGridMode = iota // 未知状态
	EGridOn                       // 并网
	EGridOff                      // 离网
	EGridSync                     // 同步中
)

type IPcsBasic interface {
	SetReset() error                           // 复位
	SetStatus(status EEnergyStoreStatus) error // 设置状态(开机、关机、待机、充电、放电、故障、同步中
	SetGridMode(mode EGridMode) error          // 设置电网状态
	GetStatus() (EEnergyStoreStatus, error)    // 状态
	GetGridMode() (EGridMode, error)           // 电网状态

	SetPower(power float64) error         // 设置有功功率
	SetReactivePower(power float64) error // 设置无功功率
	SetPowerFactor(factor float32) error  // 设置功率因数
	GetTargetPower() float64              // 获取目标有功功率
	GetTargetReactivePower() float64      // 获取目标无功功率
	GetTargetPowerFactor() float32        // 获取目标功率因数
	GetPower() (float64, error)           // 有功功率
	GetApparentPower() (float64, error)   // 视在功率
	GetReactivePower() (float64, error)   // 无功功率

	GetRatedPower() (float64, error)     // 额定功率， -1代表未知
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
