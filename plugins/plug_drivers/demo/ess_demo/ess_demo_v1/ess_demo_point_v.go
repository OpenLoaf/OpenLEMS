package ess_demo_v1

import (
	"common/c_base"
	"common/c_proto"
	"fmt"
	"github.com/shockerli/cvt"
	"time"
)

var ReadTask = &c_proto.SModbusTask{
	Name:        "ReadTask",
	DisplayName: "查询数据",
	Addr:        Status.Addr,
	Quantity:    ChargeEfficiency.Addr - Status.Addr + 1,
	Function:    c_proto.EMqHoldingRegisters,
	CycleMill:   500,
	Lifetime:    3 * time.Second,
	Metas:       []*c_base.Meta{Status, Power, SOC, GeneratedEnergy, ConsumedEnergy, MaxChargePower, MaxDischargePower, DeviceControl, TargetPower, PowerCapacity, EnergyCapacity, MinSOC, MaxSOC, ChargeEfficiency},
}

var (
	Status = &c_base.Meta{Name: "Status", Cn: "状态", Addr: 0xC8, ReadType: c_base.RInt16, StatusExplain: func(value any) string {
		if v, err := cvt.Uint8E(value); err == nil {
			switch v {
			case 0:
				return "关机"
			case 1:
				return "待机"
			case 2:
				return "充电中"
			case 3:
				return "放电中"
			case 4:
				return "故障"
			}
		}
		return fmt.Sprintf("未知值: %v", value)
	}}
	Power             = &c_base.Meta{Name: "Power", Cn: "功率", Addr: 0xC9, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	SOC               = &c_base.Meta{Name: "SOC", Cn: "当前SOC", Addr: 0xCA, ReadType: c_base.RInt16, Factor: 0.1, Unit: "%"}
	GeneratedEnergy   = &c_base.Meta{Name: "GeneratedEnergy", Cn: "累计放电量", Addr: 0xCB, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kWh"}
	ConsumedEnergy    = &c_base.Meta{Name: "ConsumedEnergy", Cn: "累计用电量", Addr: 0xCC, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kWh"}
	MaxChargePower    = &c_base.Meta{Name: "MaxChargePower", Cn: "最大允许充电功率", Addr: 0xCD, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	MaxDischargePower = &c_base.Meta{Name: "MaxDischargePower", Cn: "最大允许放电功率", Addr: 0xCE, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	DeviceControl     = &c_base.Meta{Name: "DeviceControl", Cn: "设备状态控制", Addr: 0xCF, ReadType: c_base.RInt16, StatusExplain: func(value any) string {
		if v, err := cvt.Uint8E(value); err == nil {
			switch v {
			case 0:
				return "关机"
			case 1:
				return "开机"
			}
		}
		return fmt.Sprintf("未知值: %v", value)
	}}
	TargetPower      = &c_base.Meta{Name: "TargetPower", Cn: "目标功率设置", Addr: 0xD0, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	PowerCapacity    = &c_base.Meta{Name: "PowerCapacity", Cn: "功率容量", Addr: 0xD1, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	EnergyCapacity   = &c_base.Meta{Name: "EnergyCapacity", Cn: "能量容量", Addr: 0xD2, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kWh"}
	MinSOC           = &c_base.Meta{Name: "MinSOC", Cn: "最小SOC", Addr: 0xD3, ReadType: c_base.RInt16, Factor: 0.1, Unit: "%"}
	MaxSOC           = &c_base.Meta{Name: "MaxSOC", Cn: "最大SOC", Addr: 0xD4, ReadType: c_base.RInt16, Factor: 0.1, Unit: "%"}
	ChargeEfficiency = &c_base.Meta{Name: "ChargeEfficiency", Cn: "充放电效率", Addr: 0xD5, ReadType: c_base.RInt16, Factor: 0.001}
)
