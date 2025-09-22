package entity

import (
	"errors"
	"s_db/s_db_model"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

// SAutomation 自动化任务实体
type SAutomation struct {
	Id             int        `json:"id" dc:"任务ID"`
	Name           string     `json:"name" dc:"任务名称"`
	StartTime      *time.Time `json:"startTime,omitempty" dc:"开始时间"`
	EndTime        *time.Time `json:"endTime,omitempty" dc:"结束时间"`
	TimeRangeType  string     `json:"timeRangeType" dc:"时间范围类型"`
	TimeRangeValue string     `json:"timeRangeValue" dc:"时间范围值"`
	TriggerRule    string     `json:"triggerRule" dc:"触发规则（JSON格式）"`
	ExecuteRule    string     `json:"executeRule" dc:"执行规则（JSON格式）"`
	Enabled        bool       `json:"enabled" dc:"是否启用"`
	CreatedAt      *time.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt      *time.Time `json:"updatedAt" dc:"更新时间"`
}

// UnmarshalValue 实现数据转换方法
func (a *SAutomation) UnmarshalValue(value interface{}) error {
	if model, ok := value.(*s_db_model.SAutomationModel); ok {
		*a = SAutomation{
			Id:             model.Id,
			Name:           model.Name,
			TimeRangeType:  model.TimeRangeType,
			TimeRangeValue: model.TimeRangeValue,
			TriggerRule:    model.TriggerRule,
			ExecuteRule:    model.ExecuteRule,
			Enabled:        model.Enabled,
		}

		// 时间字段转换
		if model.StartTime != nil {
			a.StartTime = &model.StartTime.Time
		}
		if model.EndTime != nil {
			a.EndTime = &model.EndTime.Time
		}
		if model.CreatedAt != nil {
			a.CreatedAt = &model.CreatedAt.Time
		}
		if model.UpdatedAt != nil {
			a.UpdatedAt = &model.UpdatedAt.Time
		}

		return nil
	}

	// 尝试使用 gconv.Scan 进行简单转换
	if err := gconv.Scan(value, a); err != nil {
		return errors.New("unsupported value type for SAutomation")
	}

	return nil
}
