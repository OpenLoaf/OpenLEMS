package entity

import (
	"encoding/json"
	"errors"
	"s_db/s_db_model"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

// SAutomation 自动化任务实体
type SAutomation struct {
	Id             int                       `json:"id" dc:"任务ID"`
	Name           string                    `json:"name" dc:"任务名称"`
	StartTime      *time.Time                `json:"startTime,omitempty" dc:"开始时间"`
	EndTime        *time.Time                `json:"endTime,omitempty" dc:"结束时间"`
	TimeRangeType  *string                   `json:"timeRangeType" dc:"时间范围类型"`
	TimeRangeValue *string                   `json:"timeRangeValue" dc:"时间范围值"`
	TriggerRule    *EAutomationTriggerConfig `json:"triggerRule" dc:"触发规则实体"`
	ExecuteRule    *EAutomationExecuteConfig `json:"executeRule" dc:"执行规则实体"`
	Enabled        bool                      `json:"enabled" dc:"是否启用"`
	CreatedAt      *time.Time                `json:"createdAt" dc:"创建时间"`
	UpdatedAt      *time.Time                `json:"updatedAt" dc:"更新时间"`
}

// EAutomationDeviceCondition 设备触发条件结构体
type EAutomationDeviceCondition struct {
	DeviceId string `json:"deviceId"` // 设备ID
	From     string `json:"from"`     // 是从哪里去取值
	Rule     string `json:"rule"`     // 规则表达式，如 "P>30", "Ia<100"
}

// EAutomationTimeCondition 时间触发条件结构体
type EAutomationTimeCondition struct {
	// 基础时间条件
	Hour   *int `json:"hour,omitempty"`   // 小时 (0-23)，nil 表示不限制
	Minute *int `json:"minute,omitempty"` // 分钟 (0-59)，nil 表示不限制

	// 日期条件
	DayOfWeek  *int `json:"dayOfWeek,omitempty"`  // 星期几 (0-6，0=周日)，nil 表示不限制
	DayOfMonth *int `json:"dayOfMonth,omitempty"` // 月的第几天 (1-31)，nil 表示不限制

	// 月份条件
	Month *int `json:"month,omitempty"` // 月份 (1-12)，nil 表示不限制

	// 时间范围条件
	StartTime string `json:"startTime,omitempty"` // 开始时间 (HH:MM 格式)
	EndTime   string `json:"endTime,omitempty"`   // 结束时间 (HH:MM 格式)
}

// EAutomationTriggerCondition 自动化触发条件结构体
type EAutomationTriggerCondition struct {
	// 设备值触发条件
	DeviceCondition *EAutomationDeviceCondition `json:"deviceCondition,omitempty"`

	// 时间触发条件
	TimeCondition *EAutomationTimeCondition `json:"timeCondition,omitempty"`
}

// EAutomationTriggerConfig 自动化触发配置结构体
type EAutomationTriggerConfig struct {
	AnyMatch          []*EAutomationTriggerCondition `json:"anyMatch"`          // 任意匹配条件（OR 逻辑）
	SubMatch          []*EAutomationTriggerCondition `json:"subMatch"`          // 子匹配条件
	SubMatchAll       *bool                          `json:"subMatchAll"`       // 子匹配是否全部满足
	ExecutionInterval int                            `json:"executionInterval"` // 执行间隔（秒），0表示实时执行
}

// EAutomationExecuteRule 自动化执行规则结构体
type EAutomationExecuteRule struct {
	DeviceId string `json:"deviceId"` // 设备ID
	Service  string `json:"service"`  // 服务名称
	Params   []any  `json:"params"`   // 参数列表
}

// EAutomationExecuteConfig 自动化执行配置结构体
type EAutomationExecuteConfig struct {
	Rules []*EAutomationExecuteRule `json:"rules"` // 执行规则列表
}

// UnmarshalValue 实现数据转换方法
func (a *SAutomation) UnmarshalValue(value interface{}) error {
	if model, ok := value.(*s_db_model.SAutomationModel); ok {
		*a = SAutomation{
			Id:             model.Id,
			Name:           model.Name,
			TimeRangeType:  model.TimeRangeType,
			TimeRangeValue: model.TimeRangeValue,
			Enabled:        model.Enabled,
		}

		// 解析 TriggerRule JSON
		if model.TriggerRule != "" {
			var triggerCfg EAutomationTriggerConfig
			if err := json.Unmarshal([]byte(model.TriggerRule), &triggerCfg); err == nil {
				a.TriggerRule = &triggerCfg
			}
		}

		// 解析 ExecuteRule JSON
		if model.ExecuteRule != "" {
			var execCfg EAutomationExecuteConfig
			if err := json.Unmarshal([]byte(model.ExecuteRule), &execCfg); err == nil {
				a.ExecuteRule = &execCfg
			}
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
