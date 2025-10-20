package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetPriceListReq 获取电价列表请求
type GetPriceListReq struct {
	g.Meta   `path:"/price" method:"get" tags:"电价管理" summary:"获取电价列表"`
	Page     int     `json:"page" dc:"页码" v:"required|min:1"`
	PageSize int     `json:"pageSize" dc:"每页数量" v:"required|min:1|max:100"`
	Status   *string `json:"status" dc:"状态筛选" v:"in:Enable,Disable"`
	Keyword  *string `json:"keyword" dc:"关键词" v:"length:0,50"`
	Priority *int    `json:"priority" dc:"优先级" v:"between:1,5"`
}

// GetPriceListRes 获取电价列表响应
type GetPriceListRes struct {
	List  []*Price `json:"list" dc:"电价列表"`
	Total int      `json:"total" dc:"总数"`
}
