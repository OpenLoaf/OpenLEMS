package v1

import (
	"s_price"

	"github.com/gogf/gf/v2/frame/g"
)

// CreatePriceReq 创建电价请求
type CreatePriceReq struct {
	g.Meta        `path:"/price" method:"post" tags:"电价管理" summary:"创建电价" role:"admin"`
	Description   string                   `json:"description" dc:"电价描述" v:"length:2,100"`
	Priority      int                      `json:"priority" dc:"优先级，数值越小优先级越高" v:"required|between:1,5"`
	Status        string                   `json:"status" dc:"启用状态" v:"required|in:Enable,Disable"`
	DateRange     *s_price.SDateRange      `json:"dateRange" dc:"日期范围配置" v:"required"`
	TimeRange     *s_price.STimeRange      `json:"timeRange" dc:"时间范围配置" v:"required"`
	PriceSegments []*s_price.SPriceSegment `json:"priceSegments" dc:"电价时段配置" v:"required"`
	RemoteId      *string                  `json:"remoteId" dc:"远程电价ID"`
}

// CreatePriceRes 创建电价响应
type CreatePriceRes struct {
	// 空响应结构体（操作类接口）
}
