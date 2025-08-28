package c_alarm

import "sync"

var (
	alarmManagerInstance IAlarmManager
	alarmManagerOnce     sync.Once
)

// RegisterAlarmManager 注册告警管理器
func RegisterAlarmManager(manager IAlarmManager) {
	alarmManagerOnce.Do(func() {
		alarmManagerInstance = manager
	})
}

// GetAlarmManager 获取告警管理器
func GetAlarmManager() IAlarmManager {
	return alarmManagerInstance
}
