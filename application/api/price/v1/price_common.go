package v1

import (
	"reflect"
	"s_db/s_db_model"
	"s_price"
	"time"

	"common/c_enum"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/pkg/errors"
)

// Price DTO
type Price struct {
	Id            int                      `json:"id" dc:"电价ID"`
	Description   string                   `json:"description" dc:"电价描述"`
	Priority      int                      `json:"priority" dc:"优先级，数值越小优先级越高"`
	Status        c_enum.EStatus           `json:"status" dc:"启用状态"`
	IsActive      bool                     `json:"isActive" dc:"当前是否在时间/日期范围内生效（不含启用状态）"`
	DateRange     *s_price.SDateRange      `json:"dateRange" dc:"日期范围配置"`
	TimeRange     *s_price.STimeRange      `json:"timeRange" dc:"时间范围配置（weekday/custom/monthly）"`
	PriceSegments []*s_price.SPriceSegment `json:"priceSegments" dc:"电价时段配置"`
	RemoteId      *string                  `json:"remoteId" dc:"远程电价ID"`
	CreatedAt     *time.Time               `json:"createdAt" dc:"创建时间"`
	UpdatedAt     *time.Time               `json:"updatedAt" dc:"更新时间"`
	CreatedBy     string                   `json:"createdBy" dc:"创建人"`
}

// UnmarshalValue 将数据库Model转换为DTO
func (p *Price) UnmarshalValue(value interface{}) error {
	if model, ok := value.(*s_db_model.SPriceModel); ok {
		// 基础字段映射
		p.Id = model.Id
		p.Description = model.Description
		p.Priority = model.Priority
		p.Status = c_enum.ParseStatus(model.Status)
		p.CreatedBy = model.CreatedBy

		// 时间字段转换
		if model.CreatedAt != nil {
			p.CreatedAt = &model.CreatedAt.Time
		}
		if model.UpdatedAt != nil {
			p.UpdatedAt = &model.UpdatedAt.Time
		}

		// 远程ID处理
		if model.RemoteId != "" {
			p.RemoteId = &model.RemoteId
		}

		// JSON字段反序列化
		if model.DateRange != "" {
			var dateRange s_price.SDateRange
			if err := gjson.DecodeTo(model.DateRange, &dateRange); err == nil {
				p.DateRange = &dateRange
			}
		}

		if model.TimeRange != "" {
			var timeRange s_price.STimeRange
			if err := gjson.DecodeTo(model.TimeRange, &timeRange); err == nil {
				p.TimeRange = &timeRange
			}
		}

		if model.PriceSegments != "" {
			var segments []*s_price.SPriceSegment
			if err := gjson.DecodeTo(model.PriceSegments, &segments); err == nil {
				p.PriceSegments = segments
			}
		}

		// 计算当前是否生效（基于日期/时间范围，不包含启用状态）
		p.IsActive = IsActive(time.Now(), p.DateRange, p.TimeRange)

		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}

// IsActive 判断在给定时间点是否命中日期/时间范围
func IsActive(now time.Time, dr *s_price.SDateRange, tr *s_price.STimeRange) bool {
	// 简化的时间范围判断逻辑
	if dr == nil && tr == nil {
		return true
	}

	// 日期范围判断
	if dr != nil {
		if !dr.IsLongTerm && dr.StartDate != "" {
			if startDate, err := time.Parse("2006-01-02", dr.StartDate); err == nil && now.Before(startDate) {
				return false
			}
		}
		if !dr.IsLongTerm && dr.EndDate != "" {
			if endDate, err := time.Parse("2006-01-02", dr.EndDate); err == nil && now.After(endDate) {
				return false
			}
		}
	}

	// 时间范围判断（简化版）
	if tr != nil {
		switch tr.Type {
		case "weekday":
			// 工作日判断
			if tr.WeekdayType == "workday" && (now.Weekday() == time.Saturday || now.Weekday() == time.Sunday) {
				return false
			}
			if tr.WeekdayType == "weekend" && now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
				return false
			}
		}
	}

	return true
}
