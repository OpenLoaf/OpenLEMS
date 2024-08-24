package c_device

import "ems-plan/c_telemetry"

type IChargeBasic interface {
	c_telemetry.IAcControl
	c_telemetry.IAcTelemetry
	c_telemetry.IStatisticsIncomingQuantity
	c_telemetry.IPowerLimit

	GetCarSoc() (float64, error) // 负值代表无数据
}

type ICharge interface {
	IInfo
	ILoadBasic
}
