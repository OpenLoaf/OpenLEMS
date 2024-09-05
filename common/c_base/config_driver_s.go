package c_base

import "github.com/gogf/gf/v2/errors/gerror"

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	Id             string            `json:"id,omitempty"`             // 设备ID
	Name           string            `json:"name,omitempty"`           // 设备名称
	ProtocolId     string            `json:"protocolId,omitempty"`     // 协议配置ID,如果配置了肯定是实体设备
	Driver         string            `json:"driver,omitempty"`         // 驱动名称，不需要带版本号。
	Type           EDeviceType       `json:"type,omitempty"`           // 设备类型
	IsVirtual      bool              `json:"isVirtual"`                // 是否是虚拟设备
	Enable         bool              `json:"enable"`                   // 是否启用
	LogLevel       string            `json:"logLevel,omitempty"`       // 日志等级
	Params         map[string]string `json:"params,omitempty"`         // 额外参数
	DeviceChildren []*SDriverConfig  `json:"deviceChildren,omitempty"` // 子设备
}

func (s *SDriverConfig) CheckTypeIs(tp EDeviceType) {
	if s.Type != tp {
		panic(gerror.Newf("设备ID: %s 名称: %s 类型不匹配，期望类型：%s, 实际类型：%s", s.Id, s.Name, tp, s.Type))
	}
}

func (s *SDriverConfig) Check() error {
	if s.Id == "" {
		return gerror.Newf("设备ID不能为空")
	}
	if s.Name == "" {
		return gerror.Newf("设备名称不能为空")
	}
	return nil
}
