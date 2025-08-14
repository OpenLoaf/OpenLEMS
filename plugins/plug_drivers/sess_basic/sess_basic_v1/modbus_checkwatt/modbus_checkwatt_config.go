package modbus_checkwatt

import "time"

type SSessBasicConfig struct {
	ModbusServerEnable bool   // 是否启用modbus
	ModbusUrl          string // 长度 e.q. tcp://0.0.0.0:1506
	Timeout            int64  // 超时时间
	MaxClients         uint
}

func (s *SSessBasicConfig) GetModbusUrl() string {
	if s.ModbusUrl == "" {
		s.ModbusUrl = "tcp://0.0.0.0:1506"
	}
	return s.ModbusUrl
}

func (s *SSessBasicConfig) GetTimeout() time.Duration {
	if s.Timeout == 0 {
		s.Timeout = 30
	}
	return time.Duration(s.Timeout)
}

func (s *SSessBasicConfig) GetMaxClients() uint {
	if s.MaxClients == 0 {
		s.MaxClients = 5
	}
	return s.MaxClients
}
