package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetEnergyStorageStrategyListReq struct {
	g.Meta   `path:"/strategy/energy-storage" method:"get" tags:"策略相关" summary:"查询储能策略列表"`
	Page     int    `json:"page" dc:"页码"`
	PageSize int    `json:"pageSize" dc:"每页数量"`
	Status   string `json:"status" dc:"all|active|inactive|expired|conflict"`
	Priority string `json:"priority" dc:"1|2|3|4|5|all"`
	Keyword  string `json:"keyword" dc:"关键词"`
}

type GetEnergyStorageStrategyListRes struct {
	List  []*EnergyStorageStrategy `json:"list"`
	Total int                      `json:"total"`
}
