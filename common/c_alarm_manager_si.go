package common

import (
	"common/c_alarm"
)

// RegisterAlarmManager 注册告警管理器
func RegisterAlarmManager(manager c_alarm.IAlarmManager) {
	c_alarm.RegisterAlarmManager(manager)
}

// GetAlarmManager 获取告警管理器
func GetAlarmManager() c_alarm.IAlarmManager {
	return c_alarm.GetAlarmManager()
}
