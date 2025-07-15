package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteProtocolReq struct {
	g.Meta     `path:"/protocol/delete" method:"delete" tags:"协议相关" summary:"删除协议"`
	ProtocolId string `query:"protocolId" dc:"协议ID"`
}

type DeleteProtocolRes struct {
	ProtocolId string `json:"protocolId" dc:"协议ID"`
}
