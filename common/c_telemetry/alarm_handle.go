package c_telemetry

import "ems-plan/c_meta"

type IAlarmHandler interface {
	HandleAlarm(self c_meta.SAlarmDetail, global c_meta.SAlarmDetail) error
}
