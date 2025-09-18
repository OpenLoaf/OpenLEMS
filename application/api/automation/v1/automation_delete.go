package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DeleteAutomationReq 删除自动化任务请求
type DeleteAutomationReq struct {
	g.Meta `path:"/automation/{id}" method:"delete" tags:"自动化相关" summary:"删除自动化任务"`
	Id     int `json:"id" v:"required|min:1" dc:"自动化任务ID"`
}

// DeleteAutomationRes 删除自动化任务响应
type DeleteAutomationRes struct {
}
