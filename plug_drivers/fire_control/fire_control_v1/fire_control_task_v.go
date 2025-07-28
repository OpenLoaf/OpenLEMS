package fire_control_v1

import (
	"common/c_base"
	"modbus/p_modbus"
)

var (
	GroupBasic = &p_modbus.SModbusTask{
		Name:      "GroupBasic",
		Desc:      "基础信息",
		Addr:      BATT_VOLTAGE.Addr,
		Quantity:  BATT_MIN.Addr - BATT_VOLTAGE.Addr + 2,
		CycleMill: 1000,
		Function:  p_modbus.MqInputRegisters,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{
			BATT_VOLTAGE, BATT_CURRENT, RACK_SOC, RACK_SOH,
			RACK_R_PLUS, RACK_R_MINUS, RACK_STATUS, DI_STATUS, DO_STATUS,
			TEMP_NUM, TEMP_MAX, TEMP_MAX_RACK, TEMP_MAX_MODULE, TEMP_MIN, TEMP_MIN_RACK, TEMP_MIN_MODULE, TEMP_AVG,
			BATT_NUM, BATT_AVG, BATT_MAX,
			BATT_MAX_RACK, BATT_MAX_MODULE, BATT_MIN,
		},
	}
)

var (
	GroupStatistic = &p_modbus.SModbusTask{
		Name:      "GroupStatistic",
		Desc:      "统计信息",
		Addr:      BATT_VOLTAGE.Addr,
		Quantity:  TOTAL_CHARGE_ENERGY.Addr - LINKED_STATUS.Addr + 2,
		CycleMill: 1000,
		Function:  p_modbus.MqInputRegisters,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{
			TOTAL_CHARGE_ENERGY, TOTAL_DISCHARGE_ENERGY, SINGLE_CHARGE_ENERGY, SINGLE_DISCHARGE_ENERGY, LINKED_STATUS,
		},
	}
)
