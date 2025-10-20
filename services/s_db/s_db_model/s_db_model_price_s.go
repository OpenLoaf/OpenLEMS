package s_db_model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 数据库相关常量
const (
	// 表名
	TablePrice = "price"

	// 电价表特有字段
	FieldPriceDescription = "description"
	FieldPricePriority    = "priority"
	FieldPriceStatus      = "status"
	FieldPriceDateRange   = "date_range"
	FieldPriceTimeRange   = "time_range"
	FieldPriceSegments    = "price_segments"
	FieldPriceRemoteId    = "remote_id"
	FieldPriceCreatedBy   = "created_by"
)

// SPriceModel 电价表结构
type SPriceModel struct {
	g.Meta        `orm:"table:price"`
	Id            int         `json:"id" orm:"id,primary,auto_increment"`                                                 // 主键ID
	Description   string      `json:"description" orm:"description"`                                                      // 电价描述
	Priority      int         `json:"priority" orm:"priority" v:"required|between:1,5"`                                   // 优先级 (1-5)
	Status        string      `json:"status" orm:"status" v:"required|in:Enable,Enabled,Disable,Disabled,Deleted,Delete"` // 状态
	DateRange     string      `json:"dateRange" orm:"date_range" v:"required"`                                            // 日期范围 (JSON)
	TimeRange     string      `json:"timeRange" orm:"time_range" v:"required"`                                            // 时间范围 (JSON)
	PriceSegments string      `json:"priceSegments" orm:"price_segments" v:"required"`                                    // 电价时段 (JSON)
	RemoteId      string      `json:"remoteId" orm:"remote_id"`                                                           // 远程电价ID
	CreatedBy     string      `json:"createdBy" orm:"created_by"`                                                         // 创建人
	CreatedAt     *gtime.Time `json:"createdAt" orm:"created_at,created_at"`                                              // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt" orm:"updated_at,updated_at"`                                              // 更新时间
}

// Create 创建记录
func (m *SPriceModel) Create(ctx context.Context) error {
	// 排除ID字段，让数据库自动生成
	_, err := g.Model(TablePrice).Ctx(ctx).FieldsEx(FieldId).Insert(m)
	return err
}

// GetById 根据ID获取记录
func (m *SPriceModel) GetById(ctx context.Context, id int) error {
	return g.Model(TablePrice).Ctx(ctx).Where(FieldId, id).Scan(m)
}

// Update 更新记录
func (m *SPriceModel) Update(ctx context.Context) error {
	_, err := g.Model(TablePrice).Ctx(ctx).
		Where(FieldId, m.Id).
		FieldsEx(FieldId, FieldCreatedAt).
		Update(m)
	return err
}

// UpdateFields 更新指定字段
func (m *SPriceModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TablePrice).Ctx(ctx).Where(FieldId, m.Id).Update(data)
	return err
}

// Delete 删除记录
func (m *SPriceModel) Delete(ctx context.Context) error {
	_, err := g.Model(TablePrice).Ctx(ctx).Where(FieldId, m.Id).Delete()
	return err
}
