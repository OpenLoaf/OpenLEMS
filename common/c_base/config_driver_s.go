package c_base

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	Id              string
	Name            string     // 设备名称
	Driver          string     // 驱动名称，不需要带版本号
	Group           EGroupType // 组名称
	CabinetId       uint8      // 柜子ID, 不同柜子ID对应不同对柜子
	IsMaster        bool       // 是否是主机
	Enable          bool       // 是否启用
	LogLevel        string     // 日志等级
	PrintCacheValue bool       // 打印缓存值
}
