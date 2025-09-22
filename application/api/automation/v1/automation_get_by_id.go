package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// GetAutomationByIdReq 根据ID获取自动化任务请求
type GetAutomationByIdReq struct {
	g.Meta `path:"/automation/{id}" method:"get" tags:"自动化相关" summary:"根据ID获取自动化任务详情"`
	Id     int `json:"id" v:"required|min:1" dc:"自动化任务ID"`
}

// GetAutomationByIdRes 根据ID获取自动化任务响应
type GetAutomationByIdRes struct {
	Automation *entity.SAutomation `json:"automation" dc:"自动化任务详情"`
}
