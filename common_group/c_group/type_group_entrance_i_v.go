package c_group

import (
	"ems-plan/c_device"
	"ems-plan/c_telemetry"
)

type IGroupEntrance interface {
	IInfo
	c_telemetry.IAcGrid
	c_telemetry.IAcControl
	c_telemetry.IAcStatisticsQuantity
	c_device.IAmmeterBasic

	GetChildren() []c_device.IAmmeter
}
