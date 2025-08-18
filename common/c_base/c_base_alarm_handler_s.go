package c_base

import (
	"context"
	"fmt"
	"sync"
)

/*
SAlarmHandler
告警实现结构体，使用时只需要 *c_base.SAlarmHandler 嵌套进目标结构体，然后初始化目标结构体时：

	SAlarmHandler: &c_base.SAlarmHandler{
		ctx: ctx,
	},
*/
type SAlarmHandler struct {
	AlarmHappened  func(alarm *SAlarmDetail) // 告警发生
	AlarmDisappear func(alarm *SAlarmDetail) // 告警消失
	Fc             func()
	Ctx            context.Context
	monitorOnce    sync.Once          // 只监听一次
	monitorChan    chan *SAlarmDetail // 监听

	rwMutex       sync.RWMutex
	maxLevelAlarm *SAlarmDetail // 最高等级告警
	details       []*SAlarmDetail

	notifyChanList []chan<- *SAlarmDetail
}

func (s *SAlarmHandler) ResetAlarm() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.details = []*SAlarmDetail{}
	s.maxLevelAlarm = nil
}

func (s *SAlarmHandler) GetAlarmLevel() EAlarmLevel {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.maxLevelAlarm == nil {
		return ENone
	}
	return s.maxLevelAlarm.Level
}

func (s *SAlarmHandler) RegisterMonitorChan(details chan<- *SAlarmDetail) {
	s.notifyChanList = append(s.notifyChanList, details)
}

func (s *SAlarmHandler) GetAlarmDetails() []*SAlarmDetail {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.details
}

func (s *SAlarmHandler) TriggerAlarm(alarm *SAlarmDetail) {
	if alarm.Level == ENone {
		return
	}

	needNotify := false
	if alarm.IsTrigger {
		needNotify = s.addDetail(alarm)
	} else {
		needNotify = s.remove(alarm.DeviceId, alarm.Meta)
	}

	if needNotify {
		//g.Log().Noticef(s.ctx, "告警处理：%s", maxLevelAlarm.ToString())
		for _, notifyChan := range s.notifyChanList {
			notifyChan <- alarm
		}
	}
}

func (s *SAlarmHandler) GetMonitorChan() chan<- *SAlarmDetail {
	s.monitorOnce.Do(func() {
		s.monitorChan = make(chan *SAlarmDetail)
		go func() {
			for {
				select {
				case detail, ok := <-s.monitorChan:
					if !ok {
						fmt.Println("关闭告警监听Goroutine")
						return
					}

					s.TriggerAlarm(detail)
				}
			}
		}()
		fmt.Println("启动告警监听Goroutine")
	})
	return s.monitorChan
}

func (s *SAlarmHandler) addDetail(detail *SAlarmDetail) bool {
	if s.isExist(detail.DeviceId, detail.Meta, detail.Value) {
		return false
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	s.details = append(s.details, detail)

	if s.maxLevelAlarm == nil {
		//g.Log().Warning(s.Ctx, detail.ToString())
		s.maxLevelAlarm = detail
		return true
	}
	// 更新最高等级告警
	if s.maxLevelAlarm.Level < detail.Level {
		//fmt.Printf("%s 比原来的告警等级[%s]大！", detail.ToString(), s.maxLevelAlarm.Level.String())
		s.maxLevelAlarm = detail
	}

	if s.AlarmHappened != nil {
		s.AlarmHappened(detail)
	}

	return true
}

func (s *SAlarmHandler) isExist(deviceId string, meta *Meta, value any) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	for _, detail := range s.details {
		if detail.Meta == meta && detail.DeviceId == deviceId {
			// 如果值不相等,先清楚原来的告警，再从新添加
			if detail.Value != value {
				for i, _detail := range s.details {
					if _detail.Meta == meta && _detail.DeviceId == deviceId {
						//g.Log().Noticef(s.Ctx, _detail.ToString())
						s.details = append(s.details[:i], s.details[i+1:]...)
						break
					}
				}
				if len(s.details) == 0 {
					s.maxLevelAlarm = nil
				}
				return false
			}
			return true
		}
	}
	return false
}

func (s *SAlarmHandler) remove(deviceId string, meta *Meta) bool {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	isRemove := false
	for i, detail := range s.details {
		if detail.Meta == meta && detail.DeviceId == deviceId {
			//g.Log().Noticef(s.Ctx, "-- 清除 Id:%s 的告警：%s(%s)告警！数值:%v", detail.DeviceId, detail.Meta.Name, detail.Meta.Cn, detail.Value)
			s.details = append(s.details[:i], s.details[i+1:]...)
			isRemove = true

			if s.AlarmDisappear != nil {
				s.AlarmDisappear(detail)
			}
			break
		}
	}
	if len(s.details) == 0 {
		s.maxLevelAlarm = nil
	}

	return isRemove
}
