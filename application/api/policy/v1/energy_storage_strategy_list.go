package v1

import (
	"common/c_enum"

	"github.com/gogf/gf/v2/frame/g"
)

type GetEnergyStorageStrategyListReq struct {
	g.Meta   `path:"/strategy/energy-storage" method:"get" tags:"策略相关" summary:"查询储能策略列表"`
	Page     int             `json:"page" dc:"页码" v:"required|min:1"`
	PageSize int             `json:"pageSize" dc:"每页数量" v:"required|min:1|max:100"`
	Status   *c_enum.EStatus `json:"status" dc:"状态筛选"`
	Priority *int            `json:"priority" dc:"优先级" v:"between:1,5"`
	Keyword  *string         `json:"keyword" dc:"关键词" v:"length:0,50"`
}

type GetEnergyStorageStrategyListRes struct {
	List  []*EnergyStorageStrategy `json:"list"`
	Total int                      `json:"total"`
}
