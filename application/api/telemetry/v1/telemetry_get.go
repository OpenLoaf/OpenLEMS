package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetTelemetryGetReq struct {
	g.Meta       `path:"/telemetry/get" method:"get" tags:"遥测相关" summary:"获取某个遥测设备的当前值（用于测试）"`
	DeviceId     string `json:"deviceId,omitempty" v:"required|length:1,32#请输入设备Key|设备Key长度为:min到:max位" dc:"设备Key"`
	TelemetryKey string `json:"telemetryKey,omitempty" v:"required|length:1,32#请输入遥测Key|遥测Key长度为:min到:max位" dc:"遥测Key"`
}

type GetTelemetryGetRes struct {
	TestJoinKey      string `json:"testJoinKey" dc:"测试联合Key"`
	DeviceId         string `json:"deviceId" dc:"设备Key"`
	TelemetryKey     string `json:"telemetryKey" dc:"遥测Key"`
	TelemetryKeyName string `json:"telemetryKeyName" dc:"遥测名称"`
	Value            any    `json:"value" dc:"遥测值"`
}
