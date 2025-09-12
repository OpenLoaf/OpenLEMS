package v1

import (
	"t_network_manager/public"

	"github.com/gogf/gf/v2/frame/g"
)

type GetNetworkInterfaceListReq struct {
	g.Meta          `path:"/network/interface/list" method:"get" tags:"网络相关" summary:"获取本机网络接口列表"`
	IncludeLoopback bool `json:"includeLoopback" dc:"是否包含回环接口，默认true"`
}

type GetNetworkInterfaceListRes struct {
	Interfaces []*public.InterfaceSummary `json:"interfaces" dc:"网络接口列表"`
}

type UpdateNetworkInterfaceReq struct {
	g.Meta `path:"/network/interface/update" method:"post" tags:"网络相关" summary:"更新网络接口配置"`
	Name   string `json:"name" v:"required#接口名称必填"`
	//DHCP        bool     `json:"dhcp" dc:"是否使用DHCP模式"`
	//IPAddresses []string `json:"ipAddresses" dc:"IP地址列表（DHCP模式下可为空）"`
	//Netmask     string   `json:"netmask" dc:"子网掩码（DHCP模式下可为空）"`
	//Gateway     string   `json:"gateway" v:"ipv4#网关地址格式不正确" dc:"网关地址（DHCP模式下可为空）"`
	//DNS         []string `json:"dns" dc:"DNS服务器地址列表（可选）"`
	Config *public.InterfaceConfig `json:"config"`
}
type UpdateNetworkInterfaceRes struct {
	// 空响应体，成功则返回空对象，失败通过错误码/信息返回
}

type SetInterfaceStateReq struct {
	g.Meta `path:"/network/interface/state" method:"post" tags:"网络相关" summary:"设置网络接口状态（启用/禁用）"`
	Name   string `json:"name" v:"required#接口名称必填"`
	Up     bool   `json:"up" dc:"是否启用接口，true为启用，false为禁用"`
}

type SetInterfaceStateRes struct {
	// 空响应体，成功则返回空对象，失败通过错误码/信息返回
}

type PingReq struct {
	g.Meta `path:"/network/ping" method:"post" tags:"网络相关" summary:"执行ping测试"`
	Target string `json:"target" v:"required#目标地址必填" dc:"目标IP地址或域名"`
}

type PingRes struct {
	Result *public.PingResult `json:"result" dc:"ping测试结果"`
}
