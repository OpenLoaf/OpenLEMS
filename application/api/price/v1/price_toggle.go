package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TogglePriceReq 启用/停用电价请求
type TogglePriceReq struct {
	g.Meta `path:"/price/{id}/toggle" method:"put" tags:"电价管理" summary:"启用/停用电价" role:"admin"`
	Id     int    `json:"id" dc:"电价ID" v:"required|min:1"`
	Status string `json:"status" dc:"状态" v:"required|in:Enable,Disable"`
}

// TogglePriceRes 启用/停用电价响应
type TogglePriceRes struct {
	// 空响应结构体（操作类接口）
}
