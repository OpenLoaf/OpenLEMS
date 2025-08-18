package ess_boost_gold_v1

import "common/c_modbus"

type EssBoostGoldConfig struct {
	c_modbus.SModbusDeviceConfig
	RatedPower     int32  // 额定功率
	Capacity       uint32 // 容量
	MaxInputPower  uint32 // 最大输入功率
	MaxOutputPower uint32 // 最大输出功率
	UsePcsData     bool   // 是否使用PCS数据
}
