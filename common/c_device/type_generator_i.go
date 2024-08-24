package c_device

import "ems-plan/c_telemetry"

type IGeneratorBasic interface {
	c_telemetry.IAcGrid
	c_telemetry.IAcControl
	c_telemetry.IAcTelemetry
	c_telemetry.IStatisticsQuantity
	c_telemetry.IPowerLimit
}

type IGenerator interface {
	IInfo
	IGeneratorBasic
}
