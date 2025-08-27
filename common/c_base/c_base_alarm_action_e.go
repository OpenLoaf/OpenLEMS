//go:generate stringer -type=EAlarmAction -trimprefix=EAlarmAction -output=c_base_alarm_action_e_string.go
package c_base

type EAlarmAction int

const (
	EAlarmActionLevelUp         EAlarmAction = iota // 告警等级上升了
	EAlarmActionLevelDown                           // 告警等级下降了
	EAlarmActionFirstTrigger                        // 首次触发告警
	EAlarmActionFirstClear                          // 首次触发告警消除
	EAlarmActionNotFirstTrigger                     // 非首次触发告警
	EAlarmActionNotFirstClear                       // 非首次触发告警消除
	EAlarmActionReset                               // 告警重置
)

// IsLevelChange 判断是否是告警等级发生变化
func (s EAlarmAction) IsLevelChange() bool {
	return s == EAlarmActionLevelUp || s == EAlarmActionLevelDown
}

// IsTrigger 判断是否是告警触发
func (s EAlarmAction) IsTrigger() bool {
	return s == EAlarmActionFirstTrigger || s == EAlarmActionNotFirstTrigger
}

// IsClear 判断是否是告警清除
func (s EAlarmAction) IsClear() bool {
	return s == EAlarmActionFirstClear || s == EAlarmActionNotFirstClear
}
