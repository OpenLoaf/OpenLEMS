package v1

import (
	"common/c_chart"

	"github.com/gogf/gf/v2/frame/g"
)

type PostProtocolMonitorTrendReq struct {
	g.Meta        `path:"/protocol/monitor/trend" method:"post" tags:"协议相关" summary:"获取协议性能趋势数据" role:"user"`
	ProtocolIds   []string `json:"protocolIds" dc:"协议ID列表"`
	ProtocolTypes []string `json:"protocolTypes" dc:"协议类型列表"`
	MetricKeys    []string `json:"metricKeys" v:"required" dc:"指标key列表"`
	StartTime     *int64   `json:"startTime" dc:"开始时间"`
	EndTime       *int64   `json:"endTime" dc:"结束时间"`
	Step          int      `json:"step" dc:"步长"`
}

type PostProtocolMonitorTrendRes struct {
	*c_chart.ChartData
}
