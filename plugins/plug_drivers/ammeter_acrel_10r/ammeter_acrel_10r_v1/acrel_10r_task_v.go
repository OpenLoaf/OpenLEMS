package ammeter_acrel_10r_v1

import (
	"common/c_base"
	"common/c_modbus"
	"time"
)

var (
	GDatetime = &c_modbus.SModbusTask{
		Name:      "GDatetime",
		Desc:      "时间",
		Addr:      Year.Addr,
		Quantity:  6,
		Function:  c_modbus.MqHoldingRegisters,
		CycleMill: 0,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Year, Month, Day, Hour, Minute, Second},
	}

	GRealtimeInfo = &c_modbus.SModbusTask{
		Name:      "RealtimeInfo",
		Desc:      "实时信息",
		Addr:      Ua.Addr,
		Quantity:  F.Addr - Ua.Addr + 2,
		Function:  c_modbus.MqHoldingRegisters,
		CycleMill: 100,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Ua, Ub, Uc, Uab, Ubc, Uca, Ia, Ib, Ic, Pa, Pb, Pc, Pt, Qa, Qb, Qc, Qt, Sa, Sb, Sc, St, Pfa, Pfb, Pfc, Pft, F},
	}

	GTotal = &c_modbus.SModbusTask{
		Name:      "Total",
		Desc:      "总电能",
		Addr:      Epi.Addr,
		Quantity:  Eqe.Addr - Epi.Addr + 2,
		CycleMill: 10000,
		Function:  c_modbus.MqHoldingRegisters,
		Lifetime:  15 * time.Second,
		Metas:     []*c_base.Meta{Epi, Eqe},
	}

	GSwitch = &c_modbus.SModbusTask{
		Name:      "Switch",
		Desc:      "开关量输入",
		Addr:      SwitchIn1.Addr,
		Quantity:  SwitchOut4.Addr - SwitchIn1.Addr + 1,
		Function:  c_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{SwitchIn1, SwitchIn2, SwitchIn3, SwitchIn4, SwitchIn5, SwitchIn6, SwitchIn7, SwitchIn8, SwitchOut1, SwitchOut2, SwitchOut3, SwitchOut4},
	}
)
