package c_base

import "github.com/gogf/gf/v2/errors/gerror"

// SProtocolConfig 基础协议配置
type SProtocolConfig struct {
	Id       string            // 协议Id
	Protocol EProtocolType     // 协议
	Address  string            // 地址
	Timeout  int64             // 链接
	LogLevel string            // 日志等级
	Params   map[string]string // 配置
}

func (b *SProtocolConfig) GetProtocol() EProtocolType {
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

func (b *SProtocolConfig) Check() error {
	if b.Id == "" {
		return gerror.Newf("协议ID不能为空")
	}
	if b.Protocol == "" {
		return gerror.Newf("协议类型不能为空")
	}

	return nil
}
