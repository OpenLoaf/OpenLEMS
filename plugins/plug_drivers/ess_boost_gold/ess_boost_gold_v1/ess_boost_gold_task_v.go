package ess_boost_gold_v1

import (
	"common/c_base"
	"modbus/p_modbus"
)

var GroupBasic = &p_modbus.SModbusTask{
	Name:      "GroupBasic",
	Desc:      "基础信息",
	Function:  p_modbus.MqInputRegisters,
	Addr:      Pt.Addr,
	Quantity:  Status.Addr - Pt.Addr + 2,
	CycleMill: 100,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Pt, Qt, Vt, It, Soc, Temp, Diff, Vmax, Vmin, Status,
	},
}

var GroupController = &p_modbus.SModbusTask{
	Name:      "GroupController",
	Desc:      "状态控制",
	Function:  p_modbus.MqHoldingRegisters,
	Addr:      CONTROL_ON_OFF.Addr,
	Quantity:  1,
	CycleMill: 1000,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		CONTROL_ON_OFF,
	},
}

var GroupSetting = &p_modbus.SModbusTask{
	Name:      "GroupSetting",
	Desc:      "设置信息",
	Function:  p_modbus.MqHoldingRegisters,
	Addr:      Set_Ap_Power.Addr,
	Quantity:  2,
	CycleMill: 500,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Set_Ap_Power, Set_Rp_Power,
	},
}
