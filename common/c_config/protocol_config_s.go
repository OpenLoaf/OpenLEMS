package c_config

import "ems-plan/c_base"

// SProtocolConfig 基础协议配置
type SProtocolConfig struct {
	Name     string               //名称
	Protocol c_base.EProtocolType // 协议
	Address  string               // 地址
	Timeout  int64                // 链接
	LogLevel string               // 日志等级
	Config   map[string]any       // 配置
	Enable   bool                 // 是否启用

	DeviceChildren []map[string]any // 设备列表
}

func (b *SProtocolConfig) GetProtocol() c_base.EProtocolType {
	return b.Protocol
}

func (b *SProtocolConfig) GetAddress() string {
	return b.Address
}

func (b *SProtocolConfig) GetTimeout() int64 {
	if b.Timeout == 0 {
		b.Timeout = 3000
	}
	return b.Timeout
}

func (b *SProtocolConfig) GetLogLevel() string {
	if b.LogLevel == "" {
		// 默认日志等级为info
		b.LogLevel = "INFO"
	}
	return b.LogLevel
}
