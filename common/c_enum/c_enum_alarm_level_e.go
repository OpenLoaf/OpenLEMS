//go:generate stringer -type=EAlarmLevel -trimprefix=EAlarmLevel -output=c_enum_alarm_level_e_string.go
package c_enum

type EAlarmLevel int

const (
	EAlarmLevelNone  EAlarmLevel = iota // 默认非告警
	EAlarmLevelWarn                     // 警告，不影响系统
	EAlarmLevelAlert                    // 警报，系统会限制功能
	EAlarmLevelError                    // 故障，系统会使得设备停机
)
