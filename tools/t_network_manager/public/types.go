package public

// InterfaceID 代表网卡接口的唯一标识
type InterfaceID string

// InterfaceSummary 网卡接口的简要信息
type InterfaceSummary struct {
	ID      InterfaceID   `json:"id"`
	Name    string        `json:"name"`
	MAC     string        `json:"mac"`
	IsUp    bool          `json:"isUp"`
	IPv4    []*Ipv4Config `json:"ipv4"`
	Gateway []string      `json:"gateway"`
	DNS     []string      `json:"dns"`
	MTU     int           `json:"mtu"`
	DHCP    bool          `json:"dhcp" dc:"是否使用DHCP配置"`
}

// InterfaceConfig 用于统一更新接口配置（指针表示可选字段）
type InterfaceConfig struct {
	IPv4    []*Ipv4Config `json:"ipv4"`
	Gateway []string      `json:"gateway"`
	DNS     []string      `json:"dns"`
	DHCP    *bool         `json:"dhcp"`
}

type Ipv4Config struct {
	IPv4       string `json:"ipv4"`
	SubnetMask string `json:"subnetMask"`
}
