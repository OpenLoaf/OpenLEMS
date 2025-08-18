package bms_pylon_tech_us108_v1

import "common/c_modbus"

type PylonTechUs108BmsConfig struct {
	c_modbus.SModbusDeviceConfig
	SyncTime       bool   // 是否同步时间
	RatedPower     int32  // 额定功率
	Capacity       uint32 // 容量
	MaxInputPower  uint32 // 最大输入功率
	MaxOutputPower uint32 // 最大输出功率
}
