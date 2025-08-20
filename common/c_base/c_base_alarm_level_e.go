//go:generate stringer -type=EAlarmLevel -trimprefix=E -output=c_base_alarm_level_e_string.go
package c_base

type EAlarmLevel int

const (
	ENone  EAlarmLevel = iota // 默认非告警
	EWarn                     // 警告，系统正常工作
	EAlarm                    // 警报，系统降低功率
	EError                    // 故障，一旦有一个系统全部停机
)
