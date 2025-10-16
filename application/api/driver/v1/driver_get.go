package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDriverListReq struct {
	g.Meta `path:"/driver/list" method:"get" tags:"驱动相关" summary:"获取驱动列表" role:"user"`
}

type GetDriverListRes struct {
	DriverList []*entity.SDriver `json:"list" dc:"驱动列表"`
	Total      int               `json:"total" dc:"总数"`
}

type GetDriverReq struct {
	g.Meta     `path:"/driver/get" method:"get" tags:"驱动相关" summary:"获取驱动详情"`
	DriverName string `json:"driverName" dc:"驱动名称" v:"required"`
}

type GetDriverRes struct {
	Driver *entity.SDriver `json:"driver" dc:"驱动"`
}
