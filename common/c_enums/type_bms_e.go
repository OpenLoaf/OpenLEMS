//go:generate stringer -type=EBmsStatus -output=type_bms_e_string.go
package c_type

type EBmsStatus int

const (
	EBmsStatusUnknown   EBmsStatus = iota // 未知
	EBmsStatusOff                         // 关机
	EBmsStatusStandby                     // 待机
	EBmsStatusCharge                      // 充电
	EBmsStatusDischarge                   // 放电
	EBmsStatusFault                       // 故障
)
