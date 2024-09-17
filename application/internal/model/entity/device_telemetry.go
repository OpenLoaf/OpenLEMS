package entity

import "common/c_base"

type DeviceTelemetry struct {
	DeviceId      string               `json:"deviceId" dc:"设备Id"`
	I8nName       string               `json:"name" dc:"名称"`
	TelemetryKeys []*c_base.STelemetry `json:"telemetryKeys" dc:"点位"`
}
