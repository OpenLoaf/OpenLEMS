package p_modbus

import "ems-plan/c_base"

type SModbusDeviceConfig struct {
	Id     string
	Name   string            // 设备名称
	Driver string            // 驱动名称，不需要带版本号
	UnitId uint8             // 单元ID
	Group  c_base.EGroupType // 组名称
	//Location        Location // 是站级还是柜级
	CabinetId       uint8  // 柜子ID, 不同柜子ID对应不同对柜子
	IsMaster        bool   // 是否是主机
	Enable          bool   // 是否启用
	LogLevel        string // 日志等级
	PrintCacheValue bool   // 打印缓存值

}

func (s *SModbusDeviceConfig) GetLogLevel() string {
	if s.LogLevel == "" {
		// 默认日志等级为info
		s.LogLevel = "INFO"
	}
	return s.LogLevel
}

func (s *SModbusDeviceConfig) GetId() string {
	return s.Id
}
