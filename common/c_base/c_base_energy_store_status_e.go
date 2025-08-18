//go:generate stringer -type=EEnergyStoreStatus -output=c_base_energy_store_status_e_string.go
package c_base

type EEnergyStoreStatus int // 储能状态（PCS状态）

const (
	EPcsStatusUnknown   EEnergyStoreStatus = iota // 未知状态（也可能是离线）
	EPcsStatusOff                                 // 关机
	EPcsStatusStandby                             // 待机,只有待机状态下，才算是准备运行
	EPcsStatusCharge                              // 充电
	EPcsStatusDischarge                           // 放电
	EPcsStatusFault                               // 故障
	EPcsStatusSync                                // 同步中（多个设备的状态不一致）
	EPcsBooting                                   // 启动中
)
