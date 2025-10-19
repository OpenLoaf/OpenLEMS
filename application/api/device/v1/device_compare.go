package v1

import (
	"common/c_chart"

	"github.com/gogf/gf/v2/frame/g"
)

// PostDeviceCompareReq 设备数据对比请求
type PostDeviceCompareReq struct {
	g.Meta    `path:"/device/compare" method:"post" tags:"设备相关" summary:"获取多设备数据对比" role:"user"`
	DeviceIds []string `json:"deviceIds" v:"required" dc:"设备ID列表"`
	LineKeys  []string `json:"lineKeys" dc:"线图点位名称列表"`
	BarKeys   []string `json:"barKeys" dc:"柱状图点位名称列表"`
	StartTime *int64   `json:"startTime" dc:"开始时间（毫秒时间戳）"`
	EndTime   *int64   `json:"endTime" dc:"结束时间（毫秒时间戳）"`
	Step      int      `json:"step" dc:"步长（毫秒）"`
}

// PostDeviceCompareRes 设备数据对比响应
type PostDeviceCompareRes struct {
	*c_chart.ChartData
}
