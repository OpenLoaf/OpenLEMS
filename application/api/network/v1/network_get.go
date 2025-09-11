package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetNetworkInterfaceListReq struct {
	g.Meta          `path:"/network/interface/list" method:"get" tags:"网络相关" summary:"获取本机网络接口列表"`
	IncludeLoopback bool `json:"includeLoopback" dc:"是否包含回环接口，默认true"`
}

type GetNetworkInterfaceListRes struct {
	Interfaces []*entity.SNetworkInterface `json:"interfaces" dc:"网络接口列表"`
	Total      int                         `json:"total" dc:"总数"`
	DNS        []string                    `json:"dns" dc:"DNS服务器地址列表"`
}

type UpdateNetworkInterfaceReq struct {
	g.Meta      `path:"/network/interface/update" method:"post" tags:"网络相关" summary:"更新网络接口配置"`
	Name        string   `json:"name" v:"required#接口名称必填"`
	IPAddresses []string `json:"ipAddresses" v:"required#IP地址列表必填"`
	Netmask     string   `json:"netmask" v:"required#子网掩码必填"`
	Gateway     string   `json:"gateway" v:"ipv4#网关地址格式不正确"`
}
type UpdateNetworkInterfaceRes struct {
	// 空响应体，成功则返回空对象，失败通过错误码/信息返回
}

type UpdateDNSReq struct {
	g.Meta `path:"/network/dns/update" method:"post" tags:"网络相关" summary:"更新系统DNS服务器"`
	DNS    []string `json:"dns" v:"required#DNS不能为空"`
}

type UpdateDNSRes struct{}
