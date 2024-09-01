package entity

import "ems-plan/c_base"

type DeviceTelemetry struct {
	DeviceId      string                       `json:"deviceId" dc:"设备Id"`
	I8nName       string                       `json:"name" dc:"名称"`
	TelemetryKeys map[string]*c_base.SFunction `json:"telemetryKeys" dc:"点位"`
}
