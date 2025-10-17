package v1

import (
	"common/c_enum"
	"s_policy"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateEnergyStorageStrategyReq struct {
	g.Meta      `path:"/strategy/energy-storage" method:"post" tags:"策略相关" summary:"创建储能策略" role:"admin"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Priority    int                       `json:"priority"`
	Status      c_enum.EStatus            `json:"status"`
	DateRange   *s_policy.SDateRange      `json:"dateRange"`
	TimeRange   *s_policy.STimeRange      `json:"timeRange"`
	Config      *s_policy.SStrategyConfig `json:"config"`
}

type CreateEnergyStorageStrategyRes struct {
	// 空响应结构体（操作类接口）
}
