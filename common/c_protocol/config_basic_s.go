package c_protocol

type ConfigBasic struct {
	Name           string           //名称
	Protocol       EType            // 协议
	Address        string           // 地址
	Timeout        int64            // 链接
	Enable         bool             // 是否启用
	LogLevel       string           // 日志等级
	Config         map[string]any   // 配置
	DeviceChildren []map[string]any // 设备列表
}
