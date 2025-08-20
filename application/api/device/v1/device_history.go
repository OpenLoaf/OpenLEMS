package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceHistoryReq struct {
	g.Meta        `path:"/device/history" method:"get" tags:"设备相关" summary:"获取设备历史数据"`
	DeviceId      string   `json:"deviceId" v:"required" dc:"设备ID"`
	TelemetryKeys []string `json:"telemetryKeys" v:"required" dc:"遥测点位名称列表"`
	Date          string   `json:"date" v:"required|regex:^\\d{4}-\\d{2}-\\d{2}$" dc:"日期格式：yyyy-MM-dd"`
}

type TelemetryHistoryData struct {
	TelemetryKey string                 `json:"telemetryKey" dc:"遥测点位名称"`
	Timestamp    string                 `json:"timestamp" dc:"时间戳"`
	Value        interface{}            `json:"value" dc:"数值"`
	Quality      int                    `json:"quality" dc:"数据质量"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" dc:"元数据"`
}

type GetDeviceHistoryRes struct {
	DeviceId string                  `json:"deviceId" dc:"设备ID"`
	Date     string                  `json:"date" dc:"查询日期"`
	Data     []*TelemetryHistoryData `json:"data" dc:"历史数据列表"`
}
