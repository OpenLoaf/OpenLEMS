package network

// UpdateInterfaceRequest 更新网络接口请求
type UpdateInterfaceRequest struct {
	Name        string   `json:"name" v:"required" dc:"网络接口名称"`
	DHCP        bool     `json:"dhcp" dc:"是否使用DHCP"`
	IPAddresses []string `json:"ipAddresses" dc:"IP地址列表"`
	Netmask     string   `json:"netmask" dc:"子网掩码（十六进制格式）"`
	Gateway     string   `json:"gateway" dc:"网关地址"`
	DNS         []string `json:"dns" dc:"DNS服务器地址列表"`
}

// ConnectionInfo nmcli连接信息
type ConnectionInfo struct {
	Name        string   `json:"name"`        // 连接名称
	Device      string   `json:"device"`      // 设备名称
	Type        string   `json:"type"`        // 连接类型
	AutoConnect bool     `json:"autoConnect"` // 自动连接
	Active      bool     `json:"active"`      // 是否活动
	State       string   `json:"state"`       // 连接状态
	Method      string   `json:"method"`      // IP方法（auto/manual）
	IPAddresses []string `json:"ipAddresses"` // IP地址列表
	Netmask     string   `json:"netmask"`     // 子网掩码
	Gateway     string   `json:"gateway"`     // 网关地址
	DNS         []string `json:"dns"`         // DNS服务器地址列表
}

// DeviceInfo nmcli设备信息
type DeviceInfo struct {
	Name       string `json:"name"`       // 设备名称
	Type       string `json:"type"`       // 设备类型
	State      string `json:"state"`      // 设备状态
	Connection string `json:"connection"` // 连接名称
	MAC        string `json:"mac"`        // MAC地址
	MTU        int    `json:"mtu"`        // MTU值
}

// ValidationError 验证错误
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// NetworkError 网络操作错误
type NetworkError struct {
	Operation string `json:"operation"`
	Command   string `json:"command"`
	Output    string `json:"output"`
	Err       error  `json:"error"`
}

func (e *NetworkError) Error() string {
	return e.Err.Error()
}
