package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceTelemetryReq struct {
	g.Meta `path:"/device/telemetry" method:"get" tags:"设备相关" summary:"获取所有设备的遥测信息"`
}

type DeviceTelemetryData struct {
	LastUpdateTime  *time.Time     `json:"lastUpdateTime" dc:"最后更新时间"`
	TelemetryValues map[string]any `json:"telemetryValues" dc:"遥测值"`
}

type GetDeviceTelemetryRes struct {
	Telemetry map[string]*DeviceTelemetryData `json:"telemetry" dc:"设备遥测信息"`
}
