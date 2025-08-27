package c_base

type IAlarm interface {
	GetAlarmLevel() EAlarmLevel        // 获取当前的告警级别
	GetAlarmList() []*MetaValueWrapper // 获取告警详情列表

	UpdateAlarm(deviceId string, deviceType EDeviceType, meta *Meta, value any) // 接收告警，根据Trigger函数返回值决定触发或清除

	ResetAlarm()                                                                                                                                                       // 重置清除所有告警
	RegisterAlarmHandlerFunc(alarmAction EAlarmAction, handler func(alarm *MetaValueWrapper, currentMaxAlarmLevel EAlarmLevel, isFirstHandler bool), sortValue ...int) // 注册告警处理函数,  isHandler代表是否处理过
}
