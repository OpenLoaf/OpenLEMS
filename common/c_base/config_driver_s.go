package c_base

import "github.com/gogf/gf/v2/errors/gerror"

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	Id                 string            `json:"id,omitempty"`             // 设备ID
	Name               string            `json:"name,omitempty"`           // 设备名称
	ProtocolId         string            `json:"protocolId,omitempty"`     // 协议配置ID,如果配置了肯定是实体设备
	Driver             string            `json:"driver,omitempty"`         // 驱动名称
	Enable             bool              `json:"enable"`                   // 是否启用
	StorageIntervalSec int32             `json:"storageIntervalSec"`       // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	LogLevel           string            `json:"logLevel,omitempty"`       // 日志等级
	Params             map[string]string `json:"params,omitempty"`         // 额外参数
	DeviceChildren     []*SDriverConfig  `json:"deviceChildren,omitempty"` // 子设备
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

func (s *SDriverConfig) IsVirtual() bool {
	return s.ProtocolId == ""
}
