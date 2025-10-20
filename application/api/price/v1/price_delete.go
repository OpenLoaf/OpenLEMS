package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DeletePriceReq 删除电价请求
type DeletePriceReq struct {
	g.Meta `path:"/price/{id}" method:"delete" tags:"电价管理" summary:"删除电价" role:"admin"`
	Id     int `json:"id" dc:"电价ID" v:"required|min:1"`
}

// DeletePriceRes 删除电价响应
type DeletePriceRes struct {
	// 空响应结构体（操作类接口）
}
