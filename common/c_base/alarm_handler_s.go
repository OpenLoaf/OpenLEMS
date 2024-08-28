package c_base

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

type SAlarmHandler struct {
	rwMutex        sync.RWMutex
	*SAlarmDetail  // 最高等级告警
	details        []*SAlarmDetail
	notifyChanList []chan<- *SAlarmDetail
}

func (s *SAlarmHandler) ClearAlarm() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.details = []*SAlarmDetail{}
	s.SAlarmDetail = nil
}

func (s *SAlarmHandler) GetAlarmLevel() EAlarmLevel {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.SAlarmDetail == nil {
		return ENone
	}
	return s.SAlarmDetail.Level
}

func (s *SAlarmHandler) RegisterAlarmNotify(details chan<- *SAlarmDetail) {
	s.notifyChanList = append(s.notifyChanList, details)
}

func (s *SAlarmHandler) GetAlarmDetails() []*SAlarmDetail {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.details
}

func (s *SAlarmHandler) HandlerAlarmDetail(ctx context.Context, alarm *SAlarmDetail) {
	if alarm.Level == ENone {
		return
	}

	needNotify := false
	if alarm.IsTrigger {
		needNotify = s.addDetail(ctx, alarm)

	} else {
		needNotify = s.remove(ctx, alarm.DeviceId, alarm.Meta)
	}

	if needNotify {
		for _, notifyChan := range s.notifyChanList {
			notifyChan <- alarm
		}
	}
}

func (s *SAlarmHandler) addDetail(ctx context.Context, detail *SAlarmDetail) bool {
	if s.isExist(ctx, detail.DeviceId, detail.Meta, detail.Value) {
		return false
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	s.details = append(s.details, detail)

	if s.SAlarmDetail == nil {
		g.Log().Warning(ctx, detail.ToString())
		s.SAlarmDetail = detail
		return true
	}
	// 更新最高等级告警
	if s.SAlarmDetail.Level < detail.Level {
		g.Log().Warningf(ctx, "%s 比原来的告警等级[%s]大！", detail.ToString(), s.SAlarmDetail.Level.Name())

		s.SAlarmDetail = detail
	}
	return true
}

func (s *SAlarmHandler) isExist(ctx context.Context, deviceId string, meta *Meta, value any) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	for _, detail := range s.details {
		if detail.Meta == meta && detail.DeviceId == deviceId {
			// 如果值不相等,先清楚原来的告警，再从新添加
			if detail.Value != value {
				for i, _detail := range s.details {
					if _detail.Meta == meta && _detail.DeviceId == deviceId {
						g.Log().Noticef(ctx, _detail.ToString())
						s.details = append(s.details[:i], s.details[i+1:]...)
						break
					}
				}
				if len(s.details) == 0 {
					s.SAlarmDetail = nil
				}
				return false
			}
			return true
		}
	}
	return false
}

func (s *SAlarmHandler) remove(ctx context.Context, deviceId string, meta *Meta) bool {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	isRemove := false
	for i, detail := range s.details {
		if detail.Meta == meta && detail.DeviceId == deviceId {
			g.Log().Noticef(ctx, "-- 清除 Id:%s 的告警：%s(%s)告警！数值:%v", detail.DeviceId, detail.Meta.Name, detail.Meta.Cn, detail.Value)
			s.details = append(s.details[:i], s.details[i+1:]...)
			isRemove = true
			break
		}
	}
	if len(s.details) == 0 {
		s.SAlarmDetail = nil
	}

	return isRemove
}
