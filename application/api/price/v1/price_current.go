package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetCurrentPriceReq 获取当前激活电价请求
type GetCurrentPriceReq struct {
	g.Meta `path:"/price/current" method:"get" tags:"电价管理" summary:"获取当前激活电价"`
}

// GetCurrentPriceRes 获取当前激活电价响应
type GetCurrentPriceRes struct {
	Price *Price `json:"price" dc:"当前激活电价"`
}
