package entity

type SProtocol struct {
	ProtocolId       string `json:"protocolId" dc:"协议ID"`
	ProtocolName     string `json:"protocolName" dc:"协议名称"`
	ProtocolType     string `json:"protocolType" dc:"协议类型"`
	ProtocolAddress  string `json:"protocolAddress" dc:"协议地址"`
	ProtocolPort     int    `json:"protocolPort" dc:"协议端口"`
	ProtocolTimeout  int    `json:"protocolTimeout" dc:"协议超时时间"`
	ProtocolLogLevel string `json:"protocolLogLevel" dc:"协议日志级别"`
	ProtocolParams   string `json:"protocolParams" dc:"协议参数"`
	ProtocolActive   bool   `json:"protocolActive" dc:"协议是否激活"`

	// GPIO协议绑定的设备信息（仅GPIO协议有值）
	BoundDeviceId   string `json:"boundDeviceId,omitempty" dc:"绑定的设备ID"`
	BoundDeviceName string `json:"boundDeviceName,omitempty" dc:"绑定的设备名称"`
}
