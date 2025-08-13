package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetNetworkInterfaceListReq struct {
	g.Meta       `path:"/network/interface/list" method:"get" tags:"网络相关" summary:"获取本机网络接口列表"`
	OnlyEthernet bool `json:"onlyEthernet" dc:"仅返回以太网(有线)接口，默认false"`
}

type GetNetworkInterfaceListRes struct {
	Interfaces []*entity.SNetworkInterface `json:"list" dc:"网络接口列表"`
	Total      int                         `json:"total" dc:"总数"`
}
