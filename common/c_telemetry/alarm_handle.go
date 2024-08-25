package c_telemetry

type IAlarmHandler interface {
	HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error
}
