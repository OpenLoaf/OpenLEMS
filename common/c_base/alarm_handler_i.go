package c_base

type IAlarm interface {
	ClearAlarm()                              // 清除告警
	TriggerAlarm(alarm *SAlarmDetail)         // 触发告警
	RegisterMonitorChan(chan<- *SAlarmDetail) // 注册告警通知

	GetAlarmLevel() EAlarmLevel           // 获取告警等级
	GetAlarmDetails() []*SAlarmDetail     // 获取告警详情列表
	GetMonitorChan() chan<- *SAlarmDetail // 获取告警监听通道,用于给下级设备注册告警监听
}
