package c_device

import "ems-plan/c_telemetry"

type IAmmeterBasic interface {
	c_telemetry.IAcGrid
	c_telemetry.IAcStatisticsQuantity

	GetPowerFactor() (float32, error) // 功率因数
}

type IAmmeter interface {
	IInfo
	IAmmeterBasic
	GetPtCt() (float32, float32, error) // PT CT
}
