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
	DNS        []string                    `json:"dns" dc:"系统DNS服务器列表"`
}

type UpdateNetworkInterfaceReq struct {
	g.Meta  `path:"/network/interface/update" method:"post" tags:"网络相关" summary:"更新网络接口配置"`
	Name    string `json:"name"  v:"required#接口名称必填"`
	IP      string `json:"ip"    v:"required|ipv4#IP地址必填|IP地址格式不正确"`
	Netmask string `json:"netmask" v:"required|ipv4#子网掩码必填|子网掩码格式不正确"`
	Gateway string `json:"gateway" v:"ipv4#网关地址格式不正确"`
}
type UpdateNetworkInterfaceRes struct {
	// 空响应体，成功则返回空对象，失败通过错误码/信息返回
}

type UpdateDNSReq struct {
	g.Meta `path:"/network/dns/update" method:"post" tags:"网络相关" summary:"更新系统DNS服务器"`
	DNS    []string `json:"dns" v:"required#DNS不能为空"`
}

type UpdateDNSRes struct{}
