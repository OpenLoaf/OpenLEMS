package c_device

import "ems-plan/c_telemetry"

type IPvBase interface {
	c_telemetry.IAcGrid
	c_telemetry.IAcControl
	c_telemetry.IAcTelemetry
	c_telemetry.IDcTelemetry
	c_telemetry.IStatisticsQuantity

	GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity
}

type IPv interface {
	IInfo
	IPvBase
}
