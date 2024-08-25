package p_modbus

import (
	"ems-plan/c_device"
)

type SModbusDeviceConfig struct {
	c_device.IConfig
	UnitId   uint8
	LogLevel string
}

func (s *SModbusDeviceConfig) GetLogLevel() string {
	if s.LogLevel == "" {
		// 默认日志等级为info
		s.LogLevel = "INFO"
	}
	return s.LogLevel
}
