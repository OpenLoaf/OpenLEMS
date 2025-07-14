package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDriverListReq struct {
	g.Meta `path:"/driver/list" method:"get" tags:"驱动相关" summary:"获取驱动列表"`
}

type GetDriverListRes struct {
	DriverList []*entity.SDriver `json:"list" dc:"驱动列表"`
	Total      int               `json:"total" dc:"总数"`
}
