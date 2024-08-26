package c_base

type IAlarmHandler interface {
	HandleAlarm(self SAlarmDetail, global SAlarmDetail) error
}
