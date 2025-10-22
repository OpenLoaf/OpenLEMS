package ess_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
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

	// 协议点位定义 - 使用选项模式创建（带 ValueExplain）
	Status = c_proto.NewModbusPointExt(0xC8,
		c_proto.WithPresetPoint(c_default.VPointStatus),
		c_proto.WithDataAccess(c_default.VDataAccessInt16),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "0", Value: "关机", FromParam: false, Color: "#9CBF30"},
			{Key: "1", Value: "待机", FromParam: false, Color: "#6967EE"},
			{Key: "2", Value: "充电中", FromParam: false, Color: "#29A634"},
			{Key: "3", Value: "放电中", FromParam: false, Color: "#0098FA"},
			{Key: "4", Value: "故障", FromParam: false, Color: "#FF5D5D"},
		}),
	)

	Power = c_proto.NewModbusPointWithDesc(0xC9, "Power", "功率", c_enum.EFloat32, "kW", "当前功率", c_default.VDataAccessInt16Scale01)

	SOC = c_proto.NewModbusPointWithDesc(0xCA, "SOC", "当前SOC", c_enum.EFloat32, "%", "当前SOC", c_default.VDataAccessInt16Scale01)

	GeneratedEnergy = c_proto.NewModbusPointWithDesc(0xCB, "GeneratedEnergy", "累计放电量", c_enum.EFloat32, "kWh", "累计放电量", c_default.VDataAccessUInt16Scale01)

	ConsumedEnergy = c_proto.NewModbusPointWithDesc(0xCC, "ConsumedEnergy", "累计用电量", c_enum.EFloat32, "kWh", "累计用电量", c_default.VDataAccessUInt16Scale01)

	MaxChargePower = c_proto.NewModbusPointWithDesc(0xCD, "MaxChargePower", "最大允许充电功率", c_enum.EFloat32, "kW", "最大允许充电功率", c_default.VDataAccessInt16Scale01)

	MaxDischargePower = c_proto.NewModbusPointWithDesc(0xCE, "MaxDischargePower", "最大允许放电功率", c_enum.EFloat32, "kW", "最大允许放电功率", c_default.VDataAccessInt16Scale01)

	DeviceControl = c_proto.NewModbusPointExt(0xCF,
		c_proto.WithKey("DeviceControl"),
		c_proto.WithName("设备状态控制"),
		c_proto.WithValueType(c_enum.EInt16),
		c_proto.WithDesc("设备状态控制"),
		c_proto.WithDataAccess(c_default.VDataAccessInt16),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "0", Value: "关机", FromParam: false, Color: "#9CBF30"},
			{Key: "1", Value: "开机", FromParam: false, Color: "#29A634"},
		}),
	)

	TargetPower = c_proto.NewModbusPointWithDesc(0xD0, "TargetPower", "目标功率设置", c_enum.EFloat32, "kW", "目标功率设置", c_default.VDataAccessInt16Scale01)

	PowerCapacity = c_proto.NewModbusPointWithDesc(0xD1, "PowerCapacity", "功率容量", c_enum.EFloat32, "kW", "功率容量", c_default.VDataAccessInt16Scale01)

	EnergyCapacity = c_proto.NewModbusPointWithDesc(0xD2, "EnergyCapacity", "能量容量", c_enum.EFloat32, "kWh", "能量容量", c_default.VDataAccessInt16Scale01)

	MinSOC = c_proto.NewModbusPointWithDesc(0xD3, "MinSOC", "最小SOC", c_enum.EFloat32, "%", "最小SOC", c_default.VDataAccessInt16Scale01)

	MaxSOC = c_proto.NewModbusPointWithDesc(0xD4, "MaxSOC", "最大SOC", c_enum.EFloat32, "%", "最大SOC", c_default.VDataAccessInt16Scale01)

	ChargeEfficiency = c_proto.NewModbusPointWithDesc(0xD5, "ChargeEfficiency", "充放电效率", c_enum.EFloat32, "", "充放电效率", c_default.VDataAccessInt16Scale001)
)
