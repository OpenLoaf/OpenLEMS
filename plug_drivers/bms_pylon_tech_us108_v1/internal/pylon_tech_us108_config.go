package internal

import "plug_protocol_modbus/p_modbus"

type PylonTechUs108BmsConfig struct {
	p_modbus.SModbusDeviceConfig
	SyncTime       bool   // 是否同步时间
	RatedPower     uint32 // 额定功率
	Capacity       uint32 // 容量
	MaxInputPower  uint32 // 最大输入功率
	MaxOutputPower uint32 // 最大输出功率
}
