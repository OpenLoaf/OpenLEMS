package pv_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (

	// 协议点位定义 - 使用构造函数创建
	Status = c_proto.NewModbusPointExt(0x0064,
		c_proto.WithKey("Status"),
		c_proto.WithName("设备状态字"),
		c_proto.WithValueType(c_enum.EInt16),
		c_proto.WithDesc("设备状态字"),
		c_proto.WithDataAccess(c_default.VDataAccessInt16),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "0", Value: "故障", Color: "#f5222d"},
			{Key: "1", Value: "关机", Color: "#d9d9d9"},
			{Key: "3", Value: "运行中", Color: "#52c41a"},
		}),
	)

	Power = c_proto.NewModbusPointWithDesc(0x0065, "Power", "当前功率", c_enum.EFloat32, "kW", "当前功率", c_default.VDataAccessInt16Scale01)

	GeneratedEnergy = c_proto.NewModbusPointWithDesc(0x0066, "GeneratedEnergy", "累计发电量", c_enum.EFloat32, "kWh", "累计发电量", c_default.VDataAccessUInt16Scale01)

	OnOffState = c_proto.NewModbusPointExt(0x0067,
		c_proto.WithKey("OnOffState"),
		c_proto.WithName("开机状态"),
		c_proto.WithValueType(c_enum.EInt16),
		c_proto.WithDesc("开机状态"),
		c_proto.WithDataAccess(c_default.VDataAccessInt16Scale01),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "0", Value: "关机", Color: "#d9d9d9"},
			{Key: "1", Value: "开机", Color: "#52c41a"},
		}),
	)

	PowerLimit = c_proto.NewModbusPointWithDesc(0x0068, "PowerLimit", "功率限制值", c_enum.EFloat32, "kW", "功率限制值", c_default.VDataAccessInt16Scale01)

	InstalledCapacity = c_proto.NewModbusPointWithDesc(0x0069, "InstalledCapacity", "装机容量", c_enum.EFloat32, "kW", "装机容量", c_default.VDataAccessInt16Scale01)

	Irradiance = c_proto.NewModbusPointWithDesc(0x006A, "Irradiance", "当前辐照度", c_enum.EFloat32, "W/m²", "当前辐照度", c_default.VDataAccessInt16Scale01)

	Temperature = c_proto.NewModbusPointWithDesc(0x006B, "Temperature", "当前温度", c_enum.EFloat32, "°C", "当前温度", c_default.VDataAccessInt16Scale01)

	Efficiency = c_proto.NewModbusPointWithDesc(0x006C, "Efficiency", "转换效率", c_enum.EFloat32, "", "转换效率", c_default.VDataAccessInt16Scale001)
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  Efficiency.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points: []*c_proto.SModbusPoint{
		Status, Power, GeneratedEnergy, OnOffState, PowerLimit, InstalledCapacity, Irradiance, Temperature, Efficiency,
	},
}
