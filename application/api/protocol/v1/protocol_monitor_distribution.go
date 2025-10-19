package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PostProtocolMonitorDistributionReq struct {
	g.Meta        `path:"/protocol/monitor/distribution" method:"post" tags:"协议相关" summary:"获取协议状态分布数据" role:"user"`
	ProtocolTypes []string `json:"protocolTypes" dc:"协议类型列表"`
}

type PostProtocolMonitorDistributionRes struct {
	TypeDistribution   map[string]int             `json:"typeDistribution" dc:"协议类型分布"`
	StatusDistribution ProtocolStatusDistribution `json:"statusDistribution" dc:"协议状态分布"`
}

type ProtocolStatusDistribution struct {
	Active   int `json:"active" dc:"激活状态协议数量"`
	Inactive int `json:"inactive" dc:"非激活状态协议数量"`
}
