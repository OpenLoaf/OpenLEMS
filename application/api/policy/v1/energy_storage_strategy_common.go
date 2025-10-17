package v1

import (
	"reflect"
	"time"

	"common/c_enum"
	"p_energy_manage"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/pkg/errors"
)

// EnergyStorageStrategy DTO
type EnergyStorageStrategy struct {
	Id           string                           `json:"id"`
	Name         string                           `json:"name"`
	Description  string                           `json:"description,omitempty"`
	Priority     int                              `json:"priority"`
	Status       c_enum.EStatus                   `json:"status"`
	DateRange    *p_energy_manage.SDateRange      `json:"dateRange"`
	TimeRange    *p_energy_manage.STimeRange      `json:"timeRange"`
	Config       *p_energy_manage.SStrategyConfig `json:"config"`
	EssDeviceIds []string                         `json:"essDeviceIds"`
	IsActive     bool                             `json:"isActive"`
	CreatedAt    *time.Time                       `json:"createdAt"`
	UpdatedAt    *time.Time                       `json:"updatedAt"`
	CreatedBy    string                           `json:"createdBy,omitempty"`
	IsDefault    bool                             `json:"isDefault,omitempty"`
}

// UnmarshalValue 将数据库Model转换为DTO
func (s *EnergyStorageStrategy) UnmarshalValue(value interface{}) error {
	if model, ok := value.(*s_db_model.SEnergyStorageStrategyModel); ok {
		// 基础字段映射
		s.Id = model.Id
		s.Name = model.Name
		s.Description = model.Description
		s.Priority = model.Priority
		s.Status = c_enum.ParseStatus(model.Status)
		s.IsDefault = model.IsDefault
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
			var dateRange p_energy_manage.SDateRange
			if err := gjson.DecodeTo(model.DateRange, &dateRange); err == nil {
				s.DateRange = &dateRange
			}
		}

		if model.TimeRange != "" {
			var timeRange p_energy_manage.STimeRange
			if err := gjson.DecodeTo(model.TimeRange, &timeRange); err == nil {
				s.TimeRange = &timeRange
			}
		}

		if model.Config != "" {
			var config p_energy_manage.SStrategyConfig
			if err := gjson.DecodeTo(model.Config, &config); err == nil {
				s.Config = &config
			}
		}

		if model.EssDeviceIds != "" {
			var ids []string
			if err := gjson.DecodeTo(model.EssDeviceIds, &ids); err == nil {
				s.EssDeviceIds = ids
			}
		}

		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}
