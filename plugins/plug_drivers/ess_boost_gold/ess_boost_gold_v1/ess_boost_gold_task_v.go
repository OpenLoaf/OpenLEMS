package ess_boost_gold_v1

import (
	"common/c_base"
	"common/c_proto"
)

var GroupBasic = &c_proto.SModbusTask{
	Name:      "GroupBasic",
	Desc:      "基础信息",
	Function:  c_proto.EMqInputRegisters,
	Addr:      Pt.Addr,
	Quantity:  Status.Addr - Pt.Addr + 2,
	CycleMill: 100,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Pt, Qt, Vt, It, Soc, Temp, Diff, Vmax, Vmin, Status,
	},
}

var GroupController = &c_proto.SModbusTask{
	Name:      "GroupController",
	Desc:      "状态控制",
	Function:  c_proto.EMqHoldingRegisters,
	Addr:      CONTROL_ON_OFF.Addr,
	Quantity:  1,
	CycleMill: 1000,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		CONTROL_ON_OFF,
	},
}

var GroupSetting = &c_proto.SModbusTask{
	Name:      "GroupSetting",
	Desc:      "设置信息",
	Function:  c_proto.EMqHoldingRegisters,
	Addr:      Set_Ap_Power.Addr,
	Quantity:  2,
	CycleMill: 500,
	Lifetime:  c_base.DefaultCacheLifeTime,
	Metas: []*c_base.Meta{
		Set_Ap_Power, Set_Rp_Power,
	},
}
