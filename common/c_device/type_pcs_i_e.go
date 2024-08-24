package c_device

import "ems-plan/c_telemetry"

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
	c_telemetry.IAcStatisticsQuantity
	c_telemetry.IPowerLimit

	SetReset() error                           // 复位
	SetStatus(status EEnergyStoreStatus) error // 设置状态(开机、关机、待机、充电、放电、故障、同步中
	SetGridMode(mode EGridMode) error          // 设置电网状态

	GetStatus() (EEnergyStoreStatus, error) // 状态
	GetGridMode() (EGridMode, error)        // 电网状态
}

type IPcs interface {
	IInfo
	IPcsBasic
}
