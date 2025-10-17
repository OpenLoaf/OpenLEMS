package v1

import (
	"common/c_enum"
	"p_energy_manage"

	"github.com/gogf/gf/v2/frame/g"
)

type UpdateEnergyStorageStrategyReq struct {
	g.Meta       `path:"/strategy/energy-storage/{id}" method:"put" tags:"策略相关" summary:"更新储能策略" role:"admin"`
	Id           string                           `json:"id" in:"path"`
	Name         string                           `json:"name"`
	Description  string                           `json:"description"`
	Priority     int                              `json:"priority"`
	Status       c_enum.EStrategyStatus           `json:"status"`
	DateRange    *p_energy_manage.SDateRange      `json:"dateRange"`
	TimeRange    *p_energy_manage.STimeRange      `json:"timeRange"`
	Config       *p_energy_manage.SStrategyConfig `json:"config"`
	EssDeviceIds []string                         `json:"essDeviceIds"`
	IsDefault    bool                             `json:"isDefault"`
}

type UpdateEnergyStorageStrategyRes struct {
	// 空响应结构体（操作类接口）
}
