package c_base

type IAlarm interface {
	ResetAlarm()
	GetLevel() EAlarmLevel
	GetAlarmDetails() []*SAlarmDetail // 获取告警详情列表
}
