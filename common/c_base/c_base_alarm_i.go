package c_base

import "common/c_enum"

type IAlarm interface {
	GetAlarmLevel() c_enum.EAlarmLevel // 获取当前的告警级别
	GetAlarmList() []*SPointValue      // 获取告警详情列表

	UpdateAlarm(deviceId string, point IPoint, value any) // 接收告警，根据Trigger函数返回值决定触发或清除

	ResetAlarm()                                                                                                                                                                // 重置清除所有告警
	ClearAlarm(deviceId string, point string)                                                                                                                                   // 仅清除某个告警（不屏蔽）
	IgnoreClearAlarm(deviceId string, point string)                                                                                                                             // 忽略清除某个告警
	RegisterAlarmHandlerFunc(alarmAction c_enum.EAlarmAction, handler func(alarm *SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool), sortValue ...int) // 注册告警处理函数,  isHandler代表是否处理过
}
