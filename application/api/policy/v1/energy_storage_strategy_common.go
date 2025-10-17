package v1

import (
	"time"

	"common/c_enum"
	"p_energy_manage"
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
	CreatedAt    *time.Time                       `json:"createdAt"`
	UpdatedAt    *time.Time                       `json:"updatedAt"`
	CreatedBy    string                           `json:"createdBy,omitempty"`
	IsDefault    bool                             `json:"isDefault,omitempty"`
}
