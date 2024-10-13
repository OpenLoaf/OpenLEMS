package c_base

import "github.com/gogf/gf/v2/errors/gerror"

// SProtocolConfig 基础协议配置
type SProtocolConfig struct {
	Id       string            `json:"id,omitempty" orm:"id"`             // 协议Id
	Protocol EProtocolType     `json:"protocol,omitempty" orm:"protocol"` // 协议
	Address  string            `json:"address,omitempty" orm:"address"`   // 地址
	Timeout  int64             `json:"timeout,omitempty" orm:"timeout"`   // 超时时间
	LogLevel string            `json:"logLevel,omitempty" orm:"logLevel"` // 日志等级
	Params   map[string]string `json:"params,omitempty" orm:"params"`     // 配置
}

func (b *SProtocolConfig) GetProtocol() EProtocolType {
	return b.Protocol
}

func (b *SProtocolConfig) GetAddress() string {
	return b.Address
}

func (b *SProtocolConfig) GetTimeout() int64 {
	if b.Timeout == 0 {
		b.Timeout = 30
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

func (b *SProtocolConfig) Check() error {
	if b.Id == "" {
		return gerror.Newf("协议ID不能为空")
	}
	if b.Protocol == "" {
		return gerror.Newf("协议类型不能为空")
	}

	return nil
}
