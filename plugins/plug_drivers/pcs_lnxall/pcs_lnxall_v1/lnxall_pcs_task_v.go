package pcs_lnxall_v1

import (
	"common/c_base"
	"common/c_proto"
	"time"
)

var GroupAcInfo = &c_proto.SModbusTask{
	Name:      "GroupAcInfo",
	Desc:      "交流电流电压功率",
	Addr:      Ua.Addr,
	Quantity:  PF.Addr - Ua.Addr + 2,
	Function:  c_proto.EMqInputRegisters,
	CycleMill: 100,
	Lifetime:  30 * time.Second,
	Metas: []*c_base.Meta{
		Ua, Ub, Uc, Ia, Ib, Ic, Freq,
		Pa, Pb, Pc, Sa, Sb, Sc, Qa, Qb, Qc,
		PTotal, QTotal, STotal, PF,
	},
}

var GroupDcInfo = &c_proto.SModbusTask{
	Name:      "GroupDcInfo",
	Desc:      "直流信息",
	Addr:      Vbatt.Addr,
	Quantity:  Vn.Addr - Vbatt.Addr + 2,
	Function:  c_proto.EMqInputRegisters,
	CycleMill: 0,                // 不需要定时读取，需要的时候读取
	Lifetime:  30 * time.Second, // 30s后过期
	Metas: []*c_base.Meta{
		Vbatt, Ibatt, Pdc, Idc, Vbus, Vp, Vn,
	},
}

var GroupOtherInfo = &c_proto.SModbusTask{
	Name:      "GroupOtherInfo",
	Desc:      "其他信息",
	Addr:      WorkState.Addr,
	Quantity:  OnlineState.Addr - WorkState.Addr + 2,
	Function:  c_proto.EMqInputRegisters,
	CycleMill: 1000,
	Lifetime:  30 * time.Second,
	Metas: []*c_base.Meta{
		WorkState, IGBTTemp, PaPF, PbPF, PcPF, DcHistoryCharge, DcDayCharge,
		DcHistoryDischarge, DcDayDischarge, AcHistoryCharge, AcDayCharge, AcHistoryDischarge,
		AcDayDischarge, OnlineState,
	},
}
