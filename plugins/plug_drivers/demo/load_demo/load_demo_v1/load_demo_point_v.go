package load_demo_v1

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

	telemetryEnergyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "Energy",
			Name:      "累计用电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat32,
			Desc:      "累计用电量",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	telemetryMaxLoadPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "MaxLoad",
			Name:      "最大负荷",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "最大负荷",
		},
		MethodName: "GetMaxInputPower",
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
		Addr: 0x012C,
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
		Addr: 0x012D,
	}

	Energy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Energy",
				Name:      "累计用电量",
				Unit:      "kWh",
				ValueType: c_enum.EFloat32,
				Desc:      "累计用电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x012E,
	}

	MaxLoad = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "MaxLoad",
				Name:      "最大负荷",
				Unit:      "kW",
				ValueType: c_enum.EFloat32,
				Desc:      "最大负荷",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x012F,
	}

	PowerFactor = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PowerFactor",
				Name:      "功率因数",
				ValueType: c_enum.EFloat32,
				Desc:      "功率因数",
			},
			DataAccess: c_default.VDataAccessInt16Scale001,
		},
		Addr: 0x0130,
	}

	LoadRate = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "LoadRate",
				Name:      "当前负荷率",
				Unit:      "%",
				ValueType: c_enum.EFloat32,
				Desc:      "当前负荷率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0131,
	}
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
