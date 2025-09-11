package entity

type SNetworkInterface struct {
	Name        string   `json:"name" dc:"接口名称"`
	Type        string   `json:"type" dc:"接口类型"`
	IPv4        string   `json:"ipv4" dc:"IPv4地址"`
	IPv6        string   `json:"ipv6" dc:"IPv6地址"`
	Netmask     string   `json:"netmask" dc:"子网掩码"`
	Gateway     string   `json:"gateway" dc:"网关"`
	MAC         string   `json:"mac" dc:"MAC地址"`
	Connected   bool     `json:"connected" dc:"是否已连接"`
	Up          bool     `json:"up" dc:"接口是否启用"`
	MTU         int      `json:"mtu" dc:"最大传输单元"`
	Index       int      `json:"index" dc:"接口索引"`
	IPAddresses []string `json:"ipAddresses" dc:"所有IP地址列表"`
}
