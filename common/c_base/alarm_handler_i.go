package c_base

type IAlarmHandler interface {
	ClearAlarm()                              // 清除告警
	RegisterAlarmNotify(chan<- *SAlarmDetail) // 注册告警通知
	GetAlarmLevel() EAlarmLevel               // 获取告警等级
	GetAlarmDetails() []*SAlarmDetail         // 获取告警详情列表
	HandlerAlarmDetail(alarm *SAlarmDetail)   // 处理告警
	GetMonitorChan() chan<- *SAlarmDetail     // 获取告警监听通道,用于给下级设备注册告警监听
}
