package ess_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"fmt"
	"time"

	"github.com/shockerli/cvt"
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  ChargeEfficiency.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points:    []*c_proto.SModbusPoint{Status, Power, SOC, GeneratedEnergy, ConsumedEnergy, MaxChargePower, MaxDischargePower, DeviceControl, TargetPower, PowerCapacity, EnergyCapacity, MinSOC, MaxSOC, ChargeEfficiency},
}

var (
	Status = &c_proto.SModbusPoint{
		Addr:       0xC8,
		SPoint:     &c_base.SPoint{Key: "Status", Name: "状态", Desc: "设备状态"},
		DataAccess: c_default.VDataAccessInt16,
		StatusExplain: func(value any) (string, error) {
			if v, err := cvt.Uint8E(value); err == nil {
				switch v {
				case 0:
					return "关机", nil
				case 1:
					return "待机", nil
				case 2:
					return "充电中", nil
				case 3:
					return "放电中", nil
				case 4:
					return "故障", nil
				}
			}
			return fmt.Sprintf("未知值: %v", value), nil
		},
	}
	Power             = &c_proto.SModbusPoint{Addr: 0xC9, SPoint: &c_base.SPoint{Key: "Power", Name: "功率", Unit: "kW", Desc: "当前功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	SOC               = &c_proto.SModbusPoint{Addr: 0xCA, SPoint: &c_base.SPoint{Key: "SOC", Name: "当前SOC", Unit: "%", Desc: "当前SOC"}, DataAccess: c_default.VDataAccessInt16Scale01}
	GeneratedEnergy   = &c_proto.SModbusPoint{Addr: 0xCB, SPoint: &c_base.SPoint{Key: "GeneratedEnergy", Name: "累计放电量", Unit: "kWh", Desc: "累计放电量"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	ConsumedEnergy    = &c_proto.SModbusPoint{Addr: 0xCC, SPoint: &c_base.SPoint{Key: "ConsumedEnergy", Name: "累计用电量", Unit: "kWh", Desc: "累计用电量"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	MaxChargePower    = &c_proto.SModbusPoint{Addr: 0xCD, SPoint: &c_base.SPoint{Key: "MaxChargePower", Name: "最大允许充电功率", Unit: "kW", Desc: "最大允许充电功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	MaxDischargePower = &c_proto.SModbusPoint{Addr: 0xCE, SPoint: &c_base.SPoint{Key: "MaxDischargePower", Name: "最大允许放电功率", Unit: "kW", Desc: "最大允许放电功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	DeviceControl     = &c_proto.SModbusPoint{
		Addr:       0xCF,
		SPoint:     &c_base.SPoint{Key: "DeviceControl", Name: "设备状态控制", Desc: "设备状态控制"},
		DataAccess: c_default.VDataAccessInt16,
		StatusExplain: func(value any) (string, error) {
			if v, err := cvt.Uint8E(value); err == nil {
				switch v {
				case 0:
					return "关机", nil
				case 1:
					return "开机", nil
				}
			}
			return fmt.Sprintf("未知值: %v", value), nil
		},
	}
	TargetPower      = &c_proto.SModbusPoint{Addr: 0xD0, SPoint: &c_base.SPoint{Key: "TargetPower", Name: "目标功率设置", Unit: "kW", Desc: "目标功率设置"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PowerCapacity    = &c_proto.SModbusPoint{Addr: 0xD1, SPoint: &c_base.SPoint{Key: "PowerCapacity", Name: "功率容量", Unit: "kW", Desc: "功率容量"}, DataAccess: c_default.VDataAccessInt16Scale01}
	EnergyCapacity   = &c_proto.SModbusPoint{Addr: 0xD2, SPoint: &c_base.SPoint{Key: "EnergyCapacity", Name: "能量容量", Unit: "kWh", Desc: "能量容量"}, DataAccess: c_default.VDataAccessInt16Scale01}
	MinSOC           = &c_proto.SModbusPoint{Addr: 0xD3, SPoint: &c_base.SPoint{Key: "MinSOC", Name: "最小SOC", Unit: "%", Desc: "最小SOC"}, DataAccess: c_default.VDataAccessInt16Scale01}
	MaxSOC           = &c_proto.SModbusPoint{Addr: 0xD4, SPoint: &c_base.SPoint{Key: "MaxSOC", Name: "最大SOC", Unit: "%", Desc: "最大SOC"}, DataAccess: c_default.VDataAccessInt16Scale01}
	ChargeEfficiency = &c_proto.SModbusPoint{Addr: 0xD5, SPoint: &c_base.SPoint{Key: "ChargeEfficiency", Name: "充放电效率", Desc: "充放电效率"}, DataAccess: c_default.VDataAccessInt16Scale001}
)
