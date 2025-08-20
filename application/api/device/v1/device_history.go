package v1

import (
	"common/c_chart"
	"github.com/gogf/gf/v2/frame/g"
)

type PostDeviceHistoryReq struct {
	g.Meta        `path:"/device/history" method:"post" tags:"设备相关" summary:"获取设备历史数据"`
	DeviceId      string   `json:"deviceId" v:"required" dc:"设备ID"`
	TelemetryKeys []string `json:"telemetryKeys" v:"required" dc:"遥测点位名称列表"`
	StartTime     *int     `json:"startTime" dc:"开始时间"`
	EndTime       *int     `json:"endTime" dc:"结束时间"`
	Step          int      `json:"step" dc:"步长"`
}

type PostDeviceHistoryRes struct {
	*c_chart.ChartData
}
