package c_base

import "fmt"

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	RefId           string            // 引用ID
	Id              string            // 设备ID
	Name            string            // 设备名称
	Driver          string            // 驱动名称，不需要带版本号
	Type            EDeviceType       // 组名称
	IsMaster        bool              // 是否是主机
	Enable          bool              // 是否启用
	LogLevel        string            // 日志等级
	PrintCacheValue bool              // 打印缓存值
	Params          map[string]string // 额外参数
}

func (s *SDriverConfig) CheckTypeIs(tp EDeviceType) {
	if s.Type != tp {
		panic(fmt.Sprintf("设备ID: %s 名称: %s 类型不匹配，期望类型：%s, 实际类型：%s", s.Id, s.Name, tp, s.Type))
	}
}
