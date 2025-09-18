package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ToggleAutomationReq 开启/停用自动化任务请求
type ToggleAutomationReq struct {
	g.Meta `path:"/automation/{id}/toggle" method:"post" tags:"自动化相关" summary:"开启/停用自动化任务"`
	Id     int  `json:"id" v:"required|min:1" dc:"自动化任务ID"`
	Enable bool `json:"enable" v:"required" dc:"是否启用"`
}

// ToggleAutomationRes 开启/停用自动化任务响应
type ToggleAutomationRes struct {
	Enabled bool `json:"enabled" dc:"当前启用状态"`
}
