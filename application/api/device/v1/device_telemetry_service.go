package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceTelemetryServiceReq struct {
	g.Meta   `path:"/device/telemetry-service/{deviceId}" method:"get" tags:"设备相关" summary:"获取指定设备的所有 Telemetry 和 Service"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type GetDeviceTelemetryServiceRes struct {
	DeviceId   string                   `json:"deviceId" dc:"设备ID"`
	DeviceName string                   `json:"deviceName" dc:"设备名称"`
	Telemetry  []*c_base.STelemetry     `json:"telemetry" dc:"遥测信息列表"`
	Service    []*c_base.SDriverService `json:"service" dc:"自定义服务列表"`
}
