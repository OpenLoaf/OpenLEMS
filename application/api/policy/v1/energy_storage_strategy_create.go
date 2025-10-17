package v1

import (
	"common/c_enum"
	"p_energy_manage"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateEnergyStorageStrategyReq struct {
	g.Meta      `path:"/strategy/energy-storage" method:"post" tags:"策略相关" summary:"创建储能策略" role:"admin"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Priority    int                              `json:"priority"`
	Status      c_enum.EStatus                   `json:"status"`
	DateRange   *p_energy_manage.SDateRange      `json:"dateRange"`
	TimeRange   *p_energy_manage.STimeRange      `json:"timeRange"`
	Config      *p_energy_manage.SStrategyConfig `json:"config"`
	IsDefault   bool                             `json:"isDefault"`
}

type CreateEnergyStorageStrategyRes struct {
	// 空响应结构体（操作类接口）
}
