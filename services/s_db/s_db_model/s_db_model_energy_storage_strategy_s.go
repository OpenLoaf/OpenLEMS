package s_db_model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 数据库相关常量
const (
	// 表名
	TableEnergyStorageStrategy = "energy_storage_strategy"

	// 表特有字段
	FieldEssName        = "name"
	FieldEssDescription = "description"
	FieldEssPriority    = "priority"
	FieldEssStatus      = "status"
	FieldEssIsDefault   = "is_default"
	FieldEssDateRange   = "date_range"
	FieldEssTimeRange   = "time_range"
	FieldEssConfig      = "config"
	FieldEssDeviceIds   = "ess_device_ids"
	FieldEssCreatedBy   = "created_by"
)

// SEnergyStorageStrategyModel 储能策略表结构
type SEnergyStorageStrategyModel struct {
	g.Meta       `orm:"table:energy_storage_strategy"`
	Id           string      `json:"id" orm:"id"`
	Name         string      `json:"name" orm:"name"`
	Description  string      `json:"description" orm:"description"`
	Priority     int         `json:"priority" orm:"priority"`
	Status       string      `json:"status" orm:"status"`
	IsDefault    bool        `json:"isDefault" orm:"is_default"`
	DateRange    string      `json:"dateRange" orm:"date_range"`        // JSON
	TimeRange    string      `json:"timeRange" orm:"time_range"`        // JSON
	Config       string      `json:"config" orm:"config"`               // JSON
	EssDeviceIds string      `json:"essDeviceIds" orm:"ess_device_ids"` // JSON array
	CreatedBy    string      `json:"createdBy" orm:"created_by"`
	CreatedAt    *gtime.Time `json:"createdAt" orm:"created_at"`
	UpdatedAt    *gtime.Time `json:"updatedAt" orm:"updated_at"`
}

// Create 创建记录
func (m *SEnergyStorageStrategyModel) Create(ctx context.Context) error {
	_, err := g.Model(TableEnergyStorageStrategy).Ctx(ctx).Insert(m)
	return err
}

// GetById 根据ID获取记录
func (m *SEnergyStorageStrategyModel) GetById(ctx context.Context, id string) error {
	return g.Model(TableEnergyStorageStrategy).Ctx(ctx).Where(FieldId, id).Scan(m)
}

// Update 更新整行
func (m *SEnergyStorageStrategyModel) Update(ctx context.Context) error {
	_, err := g.Model(TableEnergyStorageStrategy).Ctx(ctx).Where(FieldId, m.Id).Update(m)
	return err
}

// UpdateFields 更新指定字段
func (m *SEnergyStorageStrategyModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableEnergyStorageStrategy).Ctx(ctx).Where(FieldId, m.Id).Update(data)
	return err
}

// Delete 删除记录
func (m *SEnergyStorageStrategyModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableEnergyStorageStrategy).Ctx(ctx).Where(FieldId, m.Id).Delete()
	return err
}
