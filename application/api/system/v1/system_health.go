package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetHealthReq struct {
	g.Meta `path:"/health" method:"get" tags:"系统相关" summary:"系统健康检查" noAuth:"true"`
}

type GetHealthRes struct {
	Status    string            `json:"status" dc:"系统状态: healthy, unhealthy"`
	Timestamp string            `json:"timestamp" dc:"检查时间戳"`
	Services  map[string]string `json:"services" dc:"关键服务状态"`
	Version   string            `json:"version" dc:"系统版本"`
}
