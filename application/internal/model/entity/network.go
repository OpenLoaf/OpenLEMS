package entity

type SNetworkInterface struct {
	Name        string   `json:"name" dc:"接口名称"`
	Type        string   `json:"type" dc:"接口类型"`
	Netmask     string   `json:"netmask" dc:"子网掩码"`
	Gateway     string   `json:"gateway" dc:"网关"`
	MAC         string   `json:"mac" dc:"MAC地址"`
	Up          bool     `json:"up" dc:"接口是否启用"`
	MTU         int      `json:"mtu" dc:"最大传输单元"`
	Index       int      `json:"index" dc:"接口索引"`
	IPAddresses []string `json:"ipAddresses" dc:"IPv4地址列表"`
	DHCP        bool     `json:"dhcp" dc:"是否使用DHCP"`
}
