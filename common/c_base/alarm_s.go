package c_base

import (
	"sync"
	"time"
)

type SAlarm struct {
	rwMutex       sync.RWMutex
	*SAlarmDetail // 最高等级告警
	details       []*SAlarmDetail
}

type SAlarmDetail struct {
	DeviceId   string      `json:"deviceId" dc:"设备ID"`
	DeviceType EDeviceType `json:"deviceType" dc:"设备类型"`
	Level      AlarmLevel  `json:"level" dc:"告警级别"`
	Meta       *Meta       `json:"meta" dc:"告警元数据"`
	HappenTime time.Time   `json:"happenTime" dc:"发生时间"`
	Value      any         `json:"value" dc:"数值"`
}

func (s *SAlarm) Clear() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.details = []*SAlarmDetail{}
	s.SAlarmDetail = nil
}

func (s *SAlarm) Add(deviceId string, deviceType EDeviceType, meta *Meta, value any) {
	if meta.Level == ENone {
		return
	}
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	alarm := &SAlarmDetail{
		DeviceId:   deviceId,
		DeviceType: deviceType,
		Level:      meta.Level,
		Meta:       meta,
		HappenTime: time.Now(),
		Value:      value,
	}
	s.details = append(s.details, alarm)

	if s.SAlarmDetail == nil {
		s.SAlarmDetail = alarm
	}
	// 更新最高等级告警
	if s.SAlarmDetail.Level < meta.Level {
		s.SAlarmDetail = alarm
	}

}

func (s *SAlarm) Remove(deviceId string, meta *Meta) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	for i, detail := range s.details {
		if detail.Meta == meta && detail.DeviceId == deviceId {
			s.details = append(s.details[:i], s.details[i+1:]...)
			break
		}
	}
}

func (s *SAlarm) List() []*SAlarmDetail {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.details
}
