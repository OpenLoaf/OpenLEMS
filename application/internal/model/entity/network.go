package entity

type SNetworkInterface struct {
	Name       string   `json:"name" dc:"接口名称"`
	Type       string   `json:"type" dc:"接口类型"`
	IPv4       string   `json:"ip" dc:"IP地址"`
	Netmask    string   `json:"netmask" dc:"子网掩码"`
	Gateway    string   `json:"gateway" dc:"网关"`
	DNSServers []string `json:"dns" dc:"DNS服务器"`
	MAC        string   `json:"mac" dc:"MAC地址"`
	Connected  bool     `json:"connected" dc:"是否已连接"`
}
