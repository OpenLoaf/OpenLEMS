package c_none

import (
	"common/c_base"
)

type sNoneAlarm struct {
}

func (s *sNoneAlarm) ResetAlarm() {

}

func (s *sNoneAlarm) TriggerAlarm(alarm *c_base.SAlarmDetail) {

}

func (s *sNoneAlarm) RegisterMonitorChan(details chan<- *c_base.SAlarmDetail) {

}

func (s *sNoneAlarm) GetAlarmLevel() c_base.EAlarmLevel {
	return c_base.ENone
}

func (s *sNoneAlarm) GetAlarmDetails() []*c_base.SAlarmDetail {
	return nil
}

func (s *sNoneAlarm) GetMonitorChan() chan<- *c_base.SAlarmDetail {
	return nil
}
