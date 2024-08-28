package c_base

import "context"

type IAlarmHandler interface {
	ClearAlarm()                                                 // 清除告警
	RegisterAlarmNotify(chan<- *SAlarmDetail)                    // 注册告警通知
	GetAlarmLevel() EAlarmLevel                                  // 获取告警等级
	GetAlarmDetails() []*SAlarmDetail                            // 获取告警详情列表
	HandlerAlarmDetail(ctx context.Context, alarm *SAlarmDetail) // 处理告警
}
