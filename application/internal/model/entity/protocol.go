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
}
