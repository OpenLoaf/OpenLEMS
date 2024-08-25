package c_device

import "ems-plan/c_telemetry"

type ILoadBasic interface {
	c_telemetry.IAcControl
	c_telemetry.IAcTelemetry
	c_telemetry.IStatisticsIncomingQuantity
	c_telemetry.IPowerLimit
}

type ILoad interface {
	IInfo
	ILoadBasic

	c_telemetry.IAlarmHandler
}
