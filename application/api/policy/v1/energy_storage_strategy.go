package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// EnergyStorageStrategy DTO
type EnergyStorageStrategy struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description,omitempty"`
	Priority     int         `json:"priority"`
	Status       string      `json:"status"`
	DateRange    interface{} `json:"dateRange"`
	TimeRange    interface{} `json:"timeRange"`
	Config       interface{} `json:"config"`
	EssDeviceIds []string    `json:"essDeviceIds"`
	CreatedAt    *time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time  `json:"updatedAt"`
	CreatedBy    string      `json:"createdBy,omitempty"`
	IsDefault    bool        `json:"isDefault,omitempty"`
}

type CreateEnergyStorageStrategyReq struct {
	g.Meta       `path:"/strategy/energy-storage" method:"post" tags:"策略相关" summary:"创建储能策略"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Priority     int         `json:"priority"`
	Status       string      `json:"status"`
	DateRange    interface{} `json:"dateRange"`
	TimeRange    interface{} `json:"timeRange"`
	Config       interface{} `json:"config"`
	EssDeviceIds []string    `json:"essDeviceIds"`
	IsDefault    bool        `json:"isDefault"`
}

type CreateEnergyStorageStrategyRes struct {
	// 空响应结构体（操作类接口）
}

type UpdateEnergyStorageStrategyReq struct {
	g.Meta       `path:"/strategy/energy-storage/{id}" method:"put" tags:"策略相关" summary:"更新储能策略"`
	Id           string      `json:"id" in:"path"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Priority     int         `json:"priority"`
	Status       string      `json:"status"`
	DateRange    interface{} `json:"dateRange"`
	TimeRange    interface{} `json:"timeRange"`
	Config       interface{} `json:"config"`
	EssDeviceIds []string    `json:"essDeviceIds"`
	IsDefault    bool        `json:"isDefault"`
}

type UpdateEnergyStorageStrategyRes struct {
	// 空响应结构体（操作类接口）
}

type DeleteEnergyStorageStrategyReq struct {
	g.Meta `path:"/strategy/energy-storage/{id}" method:"delete" tags:"策略相关" summary:"删除储能策略"`
	Id     string `json:"id" in:"path"`
}

type DeleteEnergyStorageStrategyRes struct{}

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

type GetEnergyStorageStrategyDetailReq struct {
	g.Meta `path:"/strategy/energy-storage/{id}" method:"get" tags:"策略相关" summary:"获取储能策略详情"`
	Id     string `json:"id" in:"path"`
}

type GetEnergyStorageStrategyDetailRes = EnergyStorageStrategy

type DetectEnergyStorageStrategyConflictsReq struct {
	g.Meta      `path:"/strategy/energy-storage/conflicts" method:"post" tags:"策略相关" summary:"储能策略冲突检测"`
	StrategyIds []string                 `json:"strategyIds"`
	Candidates  []*EnergyStorageStrategy `json:"candidates"`
}

type DetectEnergyStorageStrategyConflictsRes struct {
	Conflicts []struct {
		StrategyId    string   `json:"strategyId"`
		ConflictWith  []string `json:"conflictWith"`
		ConflictDates []string `json:"conflictDates"`
	} `json:"conflicts"`
}

type ActivateEnergyStorageStrategyReq struct {
	g.Meta `path:"/strategy/energy-storage/activate" method:"post" tags:"策略相关" summary:"激活或停用储能策略"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

type ActivateEnergyStorageStrategyRes struct{}
