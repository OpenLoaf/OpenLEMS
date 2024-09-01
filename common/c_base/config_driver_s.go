package c_base

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	Id              string
	Name            string            // 设备名称
	Driver          string            // 驱动名称，不需要带版本号
	StationType     EStationType      // 组名称
	CabinetId       uint8             // 柜子ID, 不同柜子ID对应不同对柜子
	IsMaster        bool              // 是否是主机
	Enable          bool              // 是否启用
	LogLevel        string            // 日志等级
	PrintCacheValue bool              // 打印缓存值
	Params          map[string]string // 额外参数
}

func (s *SDriverConfig) GetId() string {
	return s.Id
}

func (s *SDriverConfig) GetName() string {
	return s.Name
}

func (s *SDriverConfig) GetIsMaster() bool {
	return s.IsMaster
}

func (s *SDriverConfig) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s *SDriverConfig) IsEnable() bool {
	return s.Enable
}

func (s *SDriverConfig) GetStationType() EStationType {
	return s.StationType
}
