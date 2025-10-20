package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetPriceDetailReq 获取电价详情请求
type GetPriceDetailReq struct {
	g.Meta `path:"/price/{id}" method:"get" tags:"电价管理" summary:"获取电价详情"`
	Id     int `json:"id" dc:"电价ID" v:"required|min:1"`
}

// GetPriceDetailRes 获取电价详情响应
type GetPriceDetailRes struct {
	Price *Price `json:"price" dc:"电价详情"`
}
