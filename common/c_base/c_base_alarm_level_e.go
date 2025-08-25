//go:generate stringer -type=EAlarmLevel -trimprefix=E -output=c_base_alarm_level_e_string.go
package c_base

type EAlarmLevel int

const (
	ENone  EAlarmLevel = iota // 默认非告警
	EWarn                     // 警告，不影响系统
	EAlarm                    // 警报，系统会限制功能
	EError                    // 故障，系统会使得设备停机
)
