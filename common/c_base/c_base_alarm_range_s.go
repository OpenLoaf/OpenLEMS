package c_base

import "common/c_enum"

type SAlarmRangeTrigger struct {
	Error *SAlarmOvertop //  超出范围时触发故障
	Alert *SAlarmOvertop //  超出范围时触发警报
	Warn  *SAlarmOvertop // 超出范围时触发告警
}

type SAlarmOvertop struct {
	Before float64 //  超出范围前
	After  float64 //  超出范围后
}

// CheckAlarm 检查数值是否超出告警范围，返回是否触发告警和告警级别
func (s *SAlarmOvertop) CheckAlarm(value float64, level c_enum.EAlarmLevel) (trigger bool, alarmLevel c_enum.EAlarmLevel) {
	if s == nil {
		return false, c_enum.EAlarmLevelNone
	}

	// 检查是否超出范围
	if (s.Before != 0 && value < s.Before) || (s.After != 0 && value > s.After) {
		return true, level
	}

	return false, c_enum.EAlarmLevelNone
}
