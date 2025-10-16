package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AutomationTimeCondition 时间触发条件结构体
type AutomationTimeCondition struct {
	// 基础时间条件
	Hour   *int `json:"hour,omitempty" dc:"小时 (0-23)，nil 表示不限制"`
	Minute *int `json:"minute,omitempty" dc:"分钟 (0-59)，nil 表示不限制"`

	// 日期条件
	DayOfWeek  *int `json:"dayOfWeek,omitempty" dc:"星期几 (0-6，0=周日)，nil 表示不限制"`
	DayOfMonth *int `json:"dayOfMonth,omitempty" dc:"月的第几天 (1-31)，nil 表示不限制"`

	// 月份条件
	Month *int `json:"month,omitempty" dc:"月份 (1-12)，nil 表示不限制"`

	// 时间范围条件
	StartTime string `json:"startTime,omitempty" dc:"开始时间 (HH:MM 格式)"`
	EndTime   string `json:"endTime,omitempty" dc:"结束时间 (HH:MM 格式)"`
}

// AutomationDeviceCondition 设备触发条件结构体
type AutomationDeviceCondition struct {
	DeviceId string `json:"deviceId" dc:"设备ID"`
	From     string `json:"from" dc:"是从哪里去取值"`
	Rule     string `json:"rule" dc:"规则表达式，如 P>30, Ia<100"`
}

// AutomationTriggerCondition 自动化触发条件结构体
type AutomationTriggerCondition struct {
	// 设备值触发条件
	DeviceCondition *AutomationDeviceCondition `json:"deviceCondition,omitempty" dc:"设备触发条件"`

	// 时间触发条件
	TimeCondition *AutomationTimeCondition `json:"timeCondition,omitempty" dc:"时间触发条件"`
}

// AutomationTriggerConfig 自动化触发配置结构体
type AutomationTriggerConfig struct {
	AnyMatch          []*AutomationTriggerCondition `json:"anyMatch" dc:"任意匹配条件（OR 逻辑）"`
	SubMatch          []*AutomationTriggerCondition `json:"subMatch" dc:"子匹配条件"`
	SubMatchAll       *bool                         `json:"subMatchAll" dc:"子匹配是否全部满足"`
	ExecutionInterval int                           `json:"executionInterval" default:"0" dc:"执行间隔（秒），0表示实时执行"`
}

// CreateAutomationReq 创建自动化任务请求
type CreateAutomationReq struct {
	g.Meta         `path:"/automation" method:"post" tags:"自动化相关" summary:"创建自动化任务"`
	Name           string                   `json:"name" v:"required" dc:"自动化任务名称"`
	StartTime      *gtime.Time              `json:"startTime,omitempty" dc:"开始时间"`
	EndTime        *gtime.Time              `json:"endTime,omitempty" dc:"结束时间"`
	TimeRangeType  string                   `json:"timeRangeType,omitempty" dc:"时间范围类型"`
	TimeRangeValue string                   `json:"timeRangeValue,omitempty" dc:"时间范围值"`
	TriggerConfig  *AutomationTriggerConfig `json:"triggerConfig" v:"required" dc:"触发配置"`
	ExecuteRule    string                   `json:"executeRule" v:"required" dc:"执行规则（JSON格式）"`
	Enabled        bool                     `json:"enabled" default:"true" dc:"是否启用"`
}

// CreateAutomationRes 创建自动化任务响应
type CreateAutomationRes struct {
	Id int `json:"id" dc:"创建的自动化任务ID"`
}
