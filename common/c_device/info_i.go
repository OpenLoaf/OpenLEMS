package c_device

import "ems-plan/c_telemetry"

type IInfo interface {
	IConfig
	c_telemetry.ICollection // device肯定有采集的信息

	GetDescription() SDescription
}
