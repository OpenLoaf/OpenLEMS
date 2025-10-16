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

	// 协议点位定义 - 直接创建，启动时验证SPoint字段
	Status = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Status",
				Name:      "设备状态字",
				ValueType: c_enum.EInt16,
				Desc:      "设备状态字",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0064,
	}

	Power = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Power",
				Name:      "当前功率",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "当前功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0065,
	}

	GeneratedEnergy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "GeneratedEnergy",
				Name:      "累计发电量",
				Unit:      "kWh",
				ValueType: c_enum.EFloat32,
				Desc:      "累计发电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x0066,
	}

	OnOffState = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "OnOffState",
				Name:      "开机状态",
				ValueType: c_enum.EInt16,
				Desc:      "开机状态",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0067,
		ValueExplain: []*c_base.SFieldExplain{
			{Key: "0", Value: "关机", Color: "#d9d9d9"},
			{Key: "1", Value: "开机", Color: "#52c41a"},
		},
	}

	PowerLimit = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PowerLimit",
				Name:      "功率限制值",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "功率限制值",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0068,
	}

	InstalledCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "InstalledCapacity",
				Name:      "装机容量",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "装机容量",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0069,
	}

	Irradiance = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Irradiance",
				Name:      "当前辐照度",
				Unit:      "W/m²",
				ValueType: c_enum.EFloat32,
				Desc:      "当前辐照度",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x006A,
	}

	Temperature = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Temperature",
				Name:      "当前温度",
				Unit:      "°C",
				ValueType: c_enum.EFloat32,
				Desc:      "当前温度",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x006B,
	}

	Efficiency = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Efficiency",
				Name:      "转换效率",
				ValueType: c_enum.EFloat32,
				Desc:      "转换效率",
			},
			DataAccess: c_default.VDataAccessInt16Scale001,
		},
		Addr: 0x006C,
	}
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
