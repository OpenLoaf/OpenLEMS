package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PostProtocolMonitorOverviewReq struct {
	g.Meta        `path:"/protocol/monitor/overview" method:"post" tags:"协议相关" summary:"获取协议监控概览数据" role:"user"`
	ProtocolIds   []string `json:"protocolIds" dc:"协议ID列表"`
	ProtocolTypes []string `json:"protocolTypes" dc:"协议类型列表"`
}

type PostProtocolMonitorOverviewRes struct {
	TotalProtocols    int     `json:"totalProtocols" dc:"协议总数"`
	ActiveProtocols   int     `json:"activeProtocols" dc:"活跃协议数"`
	AvgSuccessRate    float64 `json:"avgSuccessRate" dc:"平均成功率"`
	AvgResponseMs     float64 `json:"avgResponseMs" dc:"平均响应时间"`
	AbnormalProtocols int     `json:"abnormalProtocols" dc:"异常协议数"`
}
