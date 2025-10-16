package pv_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetryPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "Power",
			Name:      "当前功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "当前功率",
		},
		MethodName: "GetPower",
	}

	telemetryGeneratedEnergyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "GeneratedEnergy",
			Name:      "累计发电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat32,
			Desc:      "累计发电量",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}

	telemetryPowerLimitPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "PowerLimit",
			Name:      "功率限制值",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "功率限制值",
		},
		MethodName: "GetTargetPower",
	}

	telemetryCapacityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "InstalledCapacity",
			Name:      "装机容量",
			Unit:      "kW",
			ValueType: c_enum.EUint32,
			Desc:      "装机容量",
		},
		MethodName: "GetCapacity",
	}

	telemetryTemperaturePoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "Temperature",
			Name:      "当前温度",
			Unit:      "°C",
			ValueType: c_enum.EUint32,
			Desc:      "当前温度",
		},
		MethodName: "GetTemperature",
	}

	telemetryIrradiancePoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "Irradiance",
			Name:      "当前辐照度",
			Unit:      "W/m²",
			ValueType: c_enum.EUint32,
			Desc:      "当前辐照度",
		},
		MethodName: "GetIrradiance",
	}

	// 协议点位定义 - 使用构造函数创建
	Status = c_proto.NewModbusPointWithDesc(0x0064, "Status", "设备状态字", c_enum.EInt16, "", "设备状态字", c_default.VDataAccessInt16)

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
