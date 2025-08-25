package c_base

import (
	"time"
)

type SAlarmDetail struct {
	DeviceId   string      `json:"deviceId" dc:"设备ID"`
	DeviceType EDeviceType `json:"deviceType" dc:"设备类型"`
	Level      EAlarmLevel `json:"level" dc:"告警级别"`
	Meta       *Meta       `json:"meta" dc:"告警元数据"`
	HappenTime *time.Time  `json:"happenTime" dc:"发生时间"`
	IsTrigger  bool        `json:"isTrigger" dc:"是否触发"`
	Value      any         `json:"value" dc:"数值"`
}

type sAlarm struct {
	notifyChan chan<- *SAlarmDetail
	cache      []*SAlarmDetail
}

func NewAlarm(notifyChan chan<- *SAlarmDetail) IAlarm {
	return &sAlarm{
		notifyChan: notifyChan,
		cache:      make([]*SAlarmDetail, 0),
	}
}

func (s *sAlarm) ResetAlarm() {
	s.cache = make([]*SAlarmDetail, 0)
}

func (s *sAlarm) GetLevel() EAlarmLevel {
	return EAlarm
}

func (s *sAlarm) GetAlarmDetails() []*SAlarmDetail {
	return s.cache
}
