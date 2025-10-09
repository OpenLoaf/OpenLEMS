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
	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetryPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "Power",
			Name:      "功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "当前功率",
		},
		MethodName: "GetPower",
	}

	telemetrySocPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "SOC",
			Name:      "当前SOC",
			Unit:      "%",
			ValueType: c_enum.EFloat32,
			Desc:      "当前SOC",
		},
		MethodName: "GetSOC",
	}

	telemetryGeneratedEnergyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "GeneratedEnergy",
			Name:      "累计放电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat32,
			Desc:      "累计放电量",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}

	telemetryConsumedEnergyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "ConsumedEnergy",
			Name:      "累计用电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat32,
			Desc:      "累计用电量",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	// 协议点位定义 - 直接创建，启动时验证SPoint字段
	Status = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Status",
				Name:      "状态",
				ValueType: c_enum.EInt16,
				Desc:      "设备状态",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0xC8,
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

	Power = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Power",
				Name:      "功率",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "当前功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xC9,
	}

	SOC = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "SOC",
				Name:      "当前SOC",
				Unit:      "%",
				ValueType: c_enum.EFloat32,
				Desc:      "当前SOC",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xCA,
	}

	GeneratedEnergy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "GeneratedEnergy",
				Name:      "累计放电量",
				Unit:      "kWh",
				ValueType: c_enum.EFloat32,
				Desc:      "累计放电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0xCB,
	}

	ConsumedEnergy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ConsumedEnergy",
				Name:      "累计用电量",
				Unit:      "kWh",
				ValueType: c_enum.EFloat32,
				Desc:      "累计用电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0xCC,
	}

	MaxChargePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "MaxChargePower",
				Name:      "最大允许充电功率",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "最大允许充电功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xCD,
	}

	MaxDischargePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "MaxDischargePower",
				Name:      "最大允许放电功率",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "最大允许放电功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xCE,
	}

	DeviceControl = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "DeviceControl",
				Name:      "设备状态控制",
				ValueType: c_enum.EInt16,
				Desc:      "设备状态控制",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0xCF,
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

	TargetPower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TargetPower",
				Name:      "目标功率设置",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "目标功率设置",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xD0,
	}

	PowerCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PowerCapacity",
				Name:      "功率容量",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "功率容量",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xD1,
	}

	EnergyCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "EnergyCapacity",
				Name:      "能量容量",
				Unit:      "kWh",
				ValueType: c_enum.EFloat32,
				Desc:      "能量容量",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xD2,
	}

	MinSOC = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "MinSOC",
				Name:      "最小SOC",
				Unit:      "%",
				ValueType: c_enum.EFloat32,
				Desc:      "最小SOC",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xD3,
	}

	MaxSOC = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "MaxSOC",
				Name:      "最大SOC",
				Unit:      "%",
				ValueType: c_enum.EFloat32,
				Desc:      "最大SOC",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0xD4,
	}

	ChargeEfficiency = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ChargeEfficiency",
				Name:      "充放电效率",
				ValueType: c_enum.EFloat32,
				Desc:      "充放电效率",
			},
			DataAccess: c_default.VDataAccessInt16Scale001,
		},
		Addr: 0xD5,
	}
)
