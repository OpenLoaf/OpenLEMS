package v1

import (
	"p_energy_storage"
	"reflect"
	"time"

	"common/c_enum"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/pkg/errors"
)

// EnergyStorage DTO
type EnergyStorage struct {
	Id          int                               `json:"id" dc:"策略ID"`
	Name        string                            `json:"name" dc:"策略名称"`
	Description string                            `json:"description,omitempty" dc:"策略描述"`
	Priority    int                               `json:"priority" dc:"优先级，数值越小优先级越高"`
	Status      c_enum.EStatus                    `json:"status" dc:"启用状态"`
	DateRange   *p_energy_storage.SDateRange      `json:"dateRange" dc:"日期范围配置"`
	TimeRange   *p_energy_storage.STimeRange      `json:"timeRange" dc:"时间范围配置（weekday/custom/monthly）"`
	Config      *p_energy_storage.SStrategyConfig `json:"config" dc:"策略执行配置"`
	IsActive    bool                              `json:"isActive" dc:"当前是否在时间/日期范围内生效（不含启用状态）"`
	CreatedAt   *time.Time                        `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *time.Time                        `json:"updatedAt" dc:"更新时间"`
	CreatedBy   string                            `json:"createdBy,omitempty" dc:"创建人"`
}

// UnmarshalValue 将数据库Model转换为DTO
func (s *EnergyStorage) UnmarshalValue(value interface{}) error {
	if model, ok := value.(*s_db_model.SEnergyStorageModel); ok {
		// 基础字段映射
		s.Id = model.Id
		s.Name = model.Name
		s.Description = model.Description
		s.Priority = model.Priority
		s.Status = c_enum.ParseStatus(model.Status)
		s.CreatedBy = model.CreatedBy

		// 时间字段转换
		if model.CreatedAt != nil {
			s.CreatedAt = &model.CreatedAt.Time
		}
		if model.UpdatedAt != nil {
			s.UpdatedAt = &model.UpdatedAt.Time
		}

		// JSON字段反序列化
		if model.DateRange != "" {
			var dateRange p_energy_storage.SDateRange
			if err := gjson.DecodeTo(model.DateRange, &dateRange); err == nil {
				s.DateRange = &dateRange
			}
		}

		if model.TimeRange != "" {
			var timeRange p_energy_storage.STimeRange
			if err := gjson.DecodeTo(model.TimeRange, &timeRange); err == nil {
				s.TimeRange = &timeRange
			}
		}

		if model.Config != "" {
			var config p_energy_storage.SStrategyConfig
			if err := gjson.DecodeTo(model.Config, &config); err == nil {
				s.Config = &config
			}
		}

		// 计算当前是否生效（基于日期/时间范围，不包含启用状态）
		s.IsActive = p_energy_storage.IsActive(time.Now(), s.DateRange, s.TimeRange)

		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}
