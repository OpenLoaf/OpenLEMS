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

	// 协议点位定义 - 使用构造函数创建
	Status = c_proto.NewModbusPointWithDesc(0x012C, "Status", "设备状态字", c_enum.EInt16, "", "设备状态字", c_default.VDataAccessInt16)

	Power = c_proto.NewModbusPointWithDesc(0x012D, "Power", "当前功率", c_enum.EFloat32, "kW", "当前功率", c_default.VDataAccessInt16Scale01)

	Energy = c_proto.NewModbusPointWithDesc(0x012E, "Energy", "累计用电量", c_enum.EFloat32, "kWh", "累计用电量", c_default.VDataAccessUInt16Scale01)

	MaxLoad = c_proto.NewModbusPointWithDesc(0x012F, "MaxLoad", "最大负荷", c_enum.EFloat32, "kW", "最大负荷", c_default.VDataAccessInt16Scale01)

	PowerFactor = c_proto.NewModbusPointWithDesc(0x0130, "PowerFactor", "功率因数", c_enum.EFloat32, "", "功率因数", c_default.VDataAccessInt16Scale001)

	LoadRate = c_proto.NewModbusPointWithDesc(0x0131, "LoadRate", "当前负荷率", c_enum.EFloat32, "%", "当前负荷率", c_default.VDataAccessInt16Scale01)
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
