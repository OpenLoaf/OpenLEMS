package pcs_enjoy_basic_v1

import (
	"common/c_base"
	"modbus/p_modbus"
)

var GroupAcInfo = &p_modbus.SModbusTask{
	Name:      "GroupAcInfo",
	Desc:      "交流信息",
	Addr:      Ac_history_charge.Addr,
	Quantity:  Qc.Addr - Ac_history_charge.Addr + 1,
	Function:  p_modbus.MqHoldingRegisters,
	CycleMill: 1000,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Ua, Ub, Uc, Ia, Ib, Ic, Freq, Grid_Seq, Pcs_Degrade, Pcs_Degrade_Flag, Pa, Pb, Pc, Sa, Sb, Sc, Qa, Qb, Qc,
		Ac_history_charge, Ac_today_charge, Ac_history_discharge, Ac_today_discharge,
	},
}

var GroupPowerInfo = &p_modbus.SModbusTask{
	Name:      "GroupPowerInfo",
	Desc:      "查询功率信息",
	Addr:      Pt.Addr,
	Quantity:  Pf.Addr - Pt.Addr + 1,
	Function:  p_modbus.MqHoldingRegisters,
	CycleMill: 100,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Pt,
		Qt,
		St,
		Pf,
	},
}

var GroupBasicInfo = &p_modbus.SModbusTask{
	Name:      "GroupBasicInfo",
	Desc:      "基本信息",
	Addr:      Pcs_Status.Addr,
	Quantity:  Env_Temp.Addr - Pcs_Status.Addr + 1,
	Function:  p_modbus.MqHoldingRegisters,
	CycleMill: 1000,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Pcs_Status, IGBT_Temp, Env_Temp,
	},
}

var GroupSetting = &p_modbus.SModbusTask{
	Name:      "GroupSetting",
	Desc:      "设置信息",
	Addr:      Set_Ap.Addr,
	Quantity:  Set_Qp.Addr - Set_Ap.Addr + 1,
	Function:  p_modbus.MqHoldingRegisters,
	CycleMill: 1000,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Set_Ap, Set_Qp,
	},
}
