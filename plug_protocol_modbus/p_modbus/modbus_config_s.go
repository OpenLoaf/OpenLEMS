package p_modbus

import "ems-plan/c_base"

type SModbusDeviceConfig struct {
	c_base.SDriverConfig
	UnitId uint8 // 单元ID
}

func (s *SModbusDeviceConfig) GetLogLevel() string {
	if s.LogLevel == "" {
		// 默认日志等级为info
		s.LogLevel = "INFO"
	}
	return s.LogLevel
}
