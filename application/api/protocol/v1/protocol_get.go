package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetProtocolListReq struct {
	g.Meta `path:"/protocol/list" method:"get" tags:"协议相关" summary:"获取协议列表"`
	Type   string `json:"type" dc:"协议类型"`
}

type GetProtocolListRes struct {
	ProtocolList []*entity.SProtocol `json:"list" dc:"协议列表"`
	Total        int                 `json:"total" dc:"总数"`
}
