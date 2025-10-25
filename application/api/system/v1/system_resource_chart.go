package v1

import (
	"common/c_chart"

	"github.com/gogf/gf/v2/frame/g"
)

type PostSystemResourceChartReq struct {
	g.Meta    `path:"/system/resource/chart" method:"post" tags:"系统相关" summary:"获取系统资源使用图表" role:"user"`
	Category  string `json:"category" v:"required|in:process,network,service,storage" dc:"资源类别: process-进程资源, network-网络, service-服务, storage-存储"`
	StartTime *int64 `json:"startTime" dc:"开始时间戳(毫秒)"`
	EndTime   *int64 `json:"endTime" dc:"结束时间戳(毫秒)"`
	Step      int    `json:"step" dc:"数据步长(毫秒)，0表示不过滤"`
}

type PostSystemResourceChartRes struct {
	ChartData *c_chart.ChartData `json:"chartData" dc:"图表数据"`
}



