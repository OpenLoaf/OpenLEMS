package v1

import (
	"s_automation"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 触发规则类型采用 s_automation 的公开结构体

// CreateAutomationReq 创建自动化任务请求
type CreateAutomationReq struct {
	g.Meta         `path:"/automation" method:"post" tags:"自动化相关" summary:"创建自动化任务" role:"admin"`
	Name           string                                 `json:"name" v:"required" dc:"自动化任务名称"`
	StartTime      *gtime.Time                            `json:"startTime,omitempty" dc:"开始时间"`
	EndTime        *gtime.Time                            `json:"endTime,omitempty" dc:"结束时间"`
	TimeRangeType  *string                                `json:"timeRangeType,omitempty" dc:"时间范围类型"`
	TimeRangeValue *string                                `json:"timeRangeValue,omitempty" dc:"时间范围值"`
	TriggerRule    *s_automation.SAutomationTriggerConfig `json:"triggerRule" v:"required" dc:"触发配置"`
	ExecuteRule    string                                 `json:"executeRule" v:"required" dc:"执行规则（JSON格式）"`
	Enabled        bool                                   `json:"enabled" default:"true" dc:"是否启用"`
}

// CreateAutomationRes 创建自动化任务响应
type CreateAutomationRes struct {
	Id int `json:"id" dc:"创建的自动化任务ID"`
}
