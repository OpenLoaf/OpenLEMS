package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceTelemetryReq struct {
	g.Meta   `path:"/device/telemetry" method:"get" tags:"设备相关" summary:"获取所有设备的遥测信息"`
	DeviceId string `json:"deviceId" dc:"设备ID"`
}

type DeviceTelemetryData struct {
	LastUpdateTime  *time.Time     `json:"lastUpdateTime" dc:"最后更新时间"`
	TelemetryValues map[string]any `json:"telemetryValues" dc:"遥测值"`
}

type GetDeviceTelemetryRes struct {
	ProtocolStatus map[string]string               `json:"protocolStatus" dc:"协议状态"`
	AlarmLevelMap  map[string]string               `json:"AlarmLevel,omitempty"`
	Telemetry      map[string]*DeviceTelemetryData `json:"telemetry" dc:"设备遥测信息"`
}

type GetDeviceStatusReq struct {
	g.Meta   `path:"/device/status" method:"get" tags:"设备相关" summary:"获取设备状态信息"`
	DeviceId string `json:"deviceId" dc:"设备ID，为空时返回所有设备状态"`
}

type DeviceStatusData struct {
	ProtocolStatus string     `json:"protocolStatus" dc:"协议状态"`
	AlarmLevel     string     `json:"alarmLevel" dc:"告警级别"`
	LastUpdateTime *time.Time `json:"lastUpdateTime" dc:"最后更新时间"`
}

type GetDeviceStatusRes struct {
	DeviceStatus map[string]*DeviceStatusData `json:"deviceStatus" dc:"设备状态信息"`
}
