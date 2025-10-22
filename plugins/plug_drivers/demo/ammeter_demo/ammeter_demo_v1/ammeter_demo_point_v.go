package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetryPTotalPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "pTotal",
			Name:      "功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "总有功功率",
		},
		MethodName: "GetPTotal",
	}

	telemetryFrequencyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "frequency",
			Name:      "频率",
			Unit:      "Hz",
			ValueType: c_enum.EFloat32,
			Desc:      "系统频率",
		},
		MethodName: "GetFrequency",
	}

	telemetryPfTotalPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "pfTotal",
			Name:      "功率因素",
			ValueType: c_enum.EFloat32,
			Desc:      "总功率因数",
		},
		MethodName: "GetPfTotal",
	}

	telemetryHistoryIncomingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyIncomingQuantity",
			Name:      "正向总有功电能",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "正向总有功电能",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	telemetryHistoryOutgoingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyOutgoingQuantity",
			Name:      "反向总有功电能",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "反向总有功电能",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}

	// 协议点位定义 - 使用构造函数创建
	Status = c_proto.NewModbusPointWithDesc(0x0190, "Status", "设备状态字", c_enum.EInt16, "", "设备状态字", c_default.VDataAccessInt16)

	PhaseAVoltage = c_proto.NewModbusPointWithDesc(0x0191, "PhaseAVoltage", "A相电压", c_enum.EInt16, "V", "A相电压", c_default.VDataAccessInt16)

	PhaseBVoltage = c_proto.NewModbusPointWithDesc(0x0192, "PhaseBVoltage", "B相电压", c_enum.EInt16, "V", "B相电压", c_default.VDataAccessInt16)

	PhaseCVoltage = c_proto.NewModbusPointWithDesc(0x0193, "PhaseCVoltage", "C相电压", c_enum.EInt16, "V", "C相电压", c_default.VDataAccessInt16)

	PhaseACurrent = c_proto.NewModbusPointWithDesc(0x0194, "PhaseACurrent", "A相电流", c_enum.EInt16, "A", "A相电流", c_default.VDataAccessInt16)

	PhaseBCurrent = c_proto.NewModbusPointWithDesc(0x0195, "PhaseBCurrent", "B相电流", c_enum.EInt16, "A", "B相电流", c_default.VDataAccessInt16)

	PhaseCCurrent = c_proto.NewModbusPointWithDesc(0x0196, "PhaseCCurrent", "C相电流", c_enum.EInt16, "A", "C相电流", c_default.VDataAccessInt16)

	PhaseAActivePower = c_proto.NewModbusPointWithDesc(0x0197, "PhaseAActivePower", "A相有功功率", c_enum.EInt16, "kW", "A相有功功率", c_default.VDataAccessInt16Scale01)

	PhaseBActivePower = c_proto.NewModbusPointWithDesc(0x0198, "PhaseBActivePower", "B相有功功率", c_enum.EInt16, "kW", "B相有功功率", c_default.VDataAccessInt16Scale01)

	PhaseCActivePower = c_proto.NewModbusPointWithDesc(0x0199, "PhaseCActivePower", "C相有功功率", c_enum.EInt16, "kW", "C相有功功率", c_default.VDataAccessInt16Scale01)

	PhaseAReactivePower = c_proto.NewModbusPointWithDesc(0x019A, "PhaseAReactivePower", "A相无功功率", c_enum.EInt16, "kVar", "A相无功功率", c_default.VDataAccessInt16Scale01)

	PhaseBReactivePower = c_proto.NewModbusPointWithDesc(0x019B, "PhaseBReactivePower", "B相无功功率", c_enum.EInt16, "kVar", "B相无功功率", c_default.VDataAccessInt16Scale01)

	PhaseCReactivePower = c_proto.NewModbusPointWithDesc(0x019C, "PhaseCReactivePower", "C相无功功率", c_enum.EInt16, "kVar", "C相无功功率", c_default.VDataAccessInt16Scale01)

	PhaseAApparentPower = c_proto.NewModbusPointWithDesc(0x019D, "PhaseAApparentPower", "A相视在功率", c_enum.EInt16, "kVA", "A相视在功率", c_default.VDataAccessInt16Scale01)

	PhaseBApparentPower = c_proto.NewModbusPointWithDesc(0x019E, "PhaseBApparentPower", "B相视在功率", c_enum.EInt16, "kVA", "B相视在功率", c_default.VDataAccessInt16Scale01)

	PhaseCApparentPower = c_proto.NewModbusPointWithDesc(0x019F, "PhaseCApparentPower", "C相视在功率", c_enum.EInt16, "kVA", "C相视在功率", c_default.VDataAccessInt16Scale01)

	TotalActivePower = c_proto.NewModbusPointWithDesc(0x01A0, "TotalActivePower", "总有功功率", c_enum.EInt16, "kW", "总有功功率", c_default.VDataAccessInt16Scale01)

	TotalReactivePower = c_proto.NewModbusPointWithDesc(0x01A1, "TotalReactivePower", "总无功功率", c_enum.EInt16, "kVar", "总无功功率", c_default.VDataAccessInt16Scale01)

	TotalApparentPower = c_proto.NewModbusPointWithDesc(0x01A2, "TotalApparentPower", "总视在功率", c_enum.EInt16, "kVA", "总视在功率", c_default.VDataAccessInt16Scale01)

	ForwardActiveEnergy = c_proto.NewModbusPointWithDesc(0x01A3, "ForwardActiveEnergy", "正向有功电量", c_enum.EUint16, "kWh", "正向有功电量", c_default.VDataAccessUInt16Scale001)

	ReverseActiveEnergy = c_proto.NewModbusPointWithDesc(0x01A4, "ReverseActiveEnergy", "反向有功电量", c_enum.EUint16, "kWh", "反向有功电量", c_default.VDataAccessUInt16Scale001)

	Frequency = c_proto.NewModbusPointWithDesc(0x01A5, "Frequency", "频率", c_enum.EInt16, "Hz", "频率", c_default.VDataAccessInt16Scale001)

	PowerFactor = c_proto.NewModbusPointWithDesc(0x01A6, "PowerFactor", "功率因数", c_enum.EInt16, "", "功率因数", c_default.VDataAccessInt16Scale01)

	RatedLineVoltage = c_proto.NewModbusPointWithDesc(0x01A7, "RatedLineVoltage", "额定线电压", c_enum.EInt16, "V", "额定线电压", c_default.VDataAccessInt16)

	RatedFrequency = c_proto.NewModbusPointWithDesc(0x01A8, "RatedFrequency", "额定频率", c_enum.EInt16, "Hz", "额定频率", c_default.VDataAccessInt16)
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  RatedFrequency.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points: []*c_proto.SModbusPoint{
		Status, PhaseAVoltage, PhaseBVoltage, PhaseCVoltage,
		PhaseACurrent, PhaseBCurrent, PhaseCCurrent,
		PhaseAActivePower, PhaseBActivePower, PhaseCActivePower,
		PhaseAReactivePower, PhaseBReactivePower, PhaseCReactivePower,
		PhaseAApparentPower, PhaseBApparentPower, PhaseCApparentPower,
		TotalActivePower, TotalReactivePower, TotalApparentPower,
		ForwardActiveEnergy, ReverseActiveEnergy,
		Frequency, PowerFactor, RatedLineVoltage, RatedFrequency,
	},
}
