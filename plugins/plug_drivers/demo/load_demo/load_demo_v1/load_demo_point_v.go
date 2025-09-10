package load_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	Status      = &c_proto.SModbusPoint{Addr: 0x012C, SPoint: &c_base.SPoint{Key: "Status", Name: "设备状态字", Desc: "设备状态字"}, DataAccess: c_default.VDataAccessInt16}
	Power       = &c_proto.SModbusPoint{Addr: 0x012D, SPoint: &c_base.SPoint{Key: "Power", Name: "当前功率", Unit: "kW", Desc: "当前功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	Energy      = &c_proto.SModbusPoint{Addr: 0x012E, SPoint: &c_base.SPoint{Key: "Energy", Name: "累计用电量", Unit: "kWh", Desc: "累计用电量"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	MaxLoad     = &c_proto.SModbusPoint{Addr: 0x012F, SPoint: &c_base.SPoint{Key: "MaxLoad", Name: "最大负荷", Unit: "kW", Desc: "最大负荷"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PowerFactor = &c_proto.SModbusPoint{Addr: 0x0130, SPoint: &c_base.SPoint{Key: "PowerFactor", Name: "功率因数", Desc: "功率因数"}, DataAccess: c_default.VDataAccessInt16Scale001}
	LoadRate    = &c_proto.SModbusPoint{Addr: 0x0131, SPoint: &c_base.SPoint{Key: "LoadRate", Name: "当前负荷率", Unit: "%", Desc: "当前负荷率"}, DataAccess: c_default.VDataAccessInt16Scale01}
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  LoadRate.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points: []*c_proto.SModbusPoint{
		Status, Power, Energy, MaxLoad, PowerFactor, LoadRate,
	},
}
