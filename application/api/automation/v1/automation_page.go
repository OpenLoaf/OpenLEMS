package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// GetAutomationPageReq 获取自动化分页列表请求
type GetAutomationPageReq struct {
	g.Meta   `path:"/automation/page" method:"get" tags:"自动化相关" summary:"获取自动化分页列表"`
	Page     int    `json:"page" v:"min:1" default:"1" dc:"页码"`
	PageSize int    `json:"pageSize" v:"min:1|max:100" default:"10" dc:"每页数量"`
	DeviceId string `json:"deviceId,omitempty" dc:"设备ID，可选过滤条件"`
	Enabled  *bool  `json:"enabled,omitempty" dc:"是否启用，可选过滤条件"`
}

// GetAutomationPageRes 获取自动化分页列表响应
type GetAutomationPageRes struct {
	List     []*entity.SAutomation `json:"list" dc:"自动化任务列表"`
	Total    int                   `json:"total" dc:"总数量"`
	Page     int                   `json:"page" dc:"当前页码"`
	PageSize int                   `json:"pageSize" dc:"每页数量"`
}
