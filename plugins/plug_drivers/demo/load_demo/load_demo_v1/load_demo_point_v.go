package load_demo_v1

import (
	"common/c_base"
	"common/c_proto"
	"time"
)

var (
	Status      = &c_base.Meta{Name: "Status", Cn: "设备状态字", Addr: 0x012C, ReadType: c_base.RInt16}
	Power       = &c_base.Meta{Name: "Power", Cn: "当前功率", Addr: 0x012D, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	Energy      = &c_base.Meta{Name: "Energy", Cn: "累计用电量", Addr: 0x012E, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kWh"}
	MaxLoad     = &c_base.Meta{Name: "MaxLoad", Cn: "最大负荷", Addr: 0x012F, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	PowerFactor = &c_base.Meta{Name: "PowerFactor", Cn: "功率因数", Addr: 0x0130, ReadType: c_base.RInt16, Factor: 0.01}
	LoadRate    = &c_base.Meta{Name: "LoadRate", Cn: "当前负荷率", Addr: 0x0131, ReadType: c_base.RInt16, Factor: 0.1, Unit: "%"}
)

var ReadTask = &c_proto.SModbusTask{
	Name:        "ReadTask",
	DisplayName: "查询数据",
	Addr:        Status.Addr,
	Quantity:    LoadRate.Addr - Status.Addr + 1,
	Function:    c_proto.EMqHoldingRegisters,
	CycleMill:   500,
	Lifetime:    3 * time.Second,
	Metas: []*c_base.Meta{
		Status, Power, Energy, MaxLoad, PowerFactor, LoadRate,
	},
}
