package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpdateAutomationReq 更新自动化任务请求
type UpdateAutomationReq struct {
	g.Meta         `path:"/automation/{id}" method:"put" tags:"自动化相关" summary:"更新自动化任务"`
	Id             int                      `json:"id" v:"required|min:1" dc:"自动化任务ID"`
	Name           string                   `json:"name,omitempty" dc:"自动化任务名称"`
	StartTime      *gtime.Time              `json:"startTime,omitempty" dc:"开始时间"`
	EndTime        *gtime.Time              `json:"endTime,omitempty" dc:"结束时间"`
	TimeRangeType  string                   `json:"timeRangeType,omitempty" dc:"时间范围类型"`
	TimeRangeValue string                   `json:"timeRangeValue,omitempty" dc:"时间范围值"`
	TriggerConfig  *AutomationTriggerConfig `json:"triggerConfig,omitempty" dc:"触发配置"`
	ExecuteRule    string                   `json:"executeRule,omitempty" dc:"执行规则（JSON格式）"`
	Enabled        *bool                    `json:"enabled,omitempty" dc:"是否启用"`
}

// UpdateAutomationRes 更新自动化任务响应
type UpdateAutomationRes struct {
}
