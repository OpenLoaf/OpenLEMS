package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	// Voltage measurements
	Status        = &c_proto.SModbusPoint{Addr: 0x0190, SPoint: &c_base.SPoint{Key: "Status", Name: "设备状态字", Desc: "设备状态字"}, DataAccess: c_default.VDataAccessInt16}
	PhaseAVoltage = &c_proto.SModbusPoint{Addr: 0x0191, SPoint: &c_base.SPoint{Key: "PhaseAVoltage", Name: "A相电压", Unit: "V", Desc: "A相电压"}, DataAccess: c_default.VDataAccessInt16}
	PhaseBVoltage = &c_proto.SModbusPoint{Addr: 0x0192, SPoint: &c_base.SPoint{Key: "PhaseBVoltage", Name: "B相电压", Unit: "V", Desc: "B相电压"}, DataAccess: c_default.VDataAccessInt16}
	PhaseCVoltage = &c_proto.SModbusPoint{Addr: 0x0193, SPoint: &c_base.SPoint{Key: "PhaseCVoltage", Name: "C相电压", Unit: "V", Desc: "C相电压"}, DataAccess: c_default.VDataAccessInt16}

	// Current measurements
	PhaseACurrent = &c_proto.SModbusPoint{Addr: 0x0194, SPoint: &c_base.SPoint{Key: "PhaseACurrent", Name: "A相电流", Unit: "A", Desc: "A相电流"}, DataAccess: c_default.VDataAccessInt16}
	PhaseBCurrent = &c_proto.SModbusPoint{Addr: 0x0195, SPoint: &c_base.SPoint{Key: "PhaseBCurrent", Name: "B相电流", Unit: "A", Desc: "B相电流"}, DataAccess: c_default.VDataAccessInt16}
	PhaseCCurrent = &c_proto.SModbusPoint{Addr: 0x0196, SPoint: &c_base.SPoint{Key: "PhaseCCurrent", Name: "C相电流", Unit: "A", Desc: "C相电流"}, DataAccess: c_default.VDataAccessInt16}

	// Power measurements
	PhaseAActivePower   = &c_proto.SModbusPoint{Addr: 0x0197, SPoint: &c_base.SPoint{Key: "PhaseAActivePower", Name: "A相有功功率", Unit: "kW", Desc: "A相有功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseBActivePower   = &c_proto.SModbusPoint{Addr: 0x0198, SPoint: &c_base.SPoint{Key: "PhaseBActivePower", Name: "B相有功功率", Unit: "kW", Desc: "B相有功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseCActivePower   = &c_proto.SModbusPoint{Addr: 0x0199, SPoint: &c_base.SPoint{Key: "PhaseCActivePower", Name: "C相有功功率", Unit: "kW", Desc: "C相有功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseAReactivePower = &c_proto.SModbusPoint{Addr: 0x019A, SPoint: &c_base.SPoint{Key: "PhaseAReactivePower", Name: "A相无功功率", Unit: "kVar", Desc: "A相无功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseBReactivePower = &c_proto.SModbusPoint{Addr: 0x019B, SPoint: &c_base.SPoint{Key: "PhaseBReactivePower", Name: "B相无功功率", Unit: "kVar", Desc: "B相无功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseCReactivePower = &c_proto.SModbusPoint{Addr: 0x019C, SPoint: &c_base.SPoint{Key: "PhaseCReactivePower", Name: "C相无功功率", Unit: "kVar", Desc: "C相无功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseAApparentPower = &c_proto.SModbusPoint{Addr: 0x019D, SPoint: &c_base.SPoint{Key: "PhaseAApparentPower", Name: "A相视在功率", Unit: "kVA", Desc: "A相视在功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseBApparentPower = &c_proto.SModbusPoint{Addr: 0x019E, SPoint: &c_base.SPoint{Key: "PhaseBApparentPower", Name: "B相视在功率", Unit: "kVA", Desc: "B相视在功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	PhaseCApparentPower = &c_proto.SModbusPoint{Addr: 0x019F, SPoint: &c_base.SPoint{Key: "PhaseCApparentPower", Name: "C相视在功率", Unit: "kVA", Desc: "C相视在功率"}, DataAccess: c_default.VDataAccessInt16Scale01}

	// Total power and energy
	TotalActivePower    = &c_proto.SModbusPoint{Addr: 0x01A0, SPoint: &c_base.SPoint{Key: "TotalActivePower", Name: "总有功功率", Unit: "kW", Desc: "总有功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	TotalReactivePower  = &c_proto.SModbusPoint{Addr: 0x01A1, SPoint: &c_base.SPoint{Key: "TotalReactivePower", Name: "总无功功率", Unit: "kVar", Desc: "总无功功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	TotalApparentPower  = &c_proto.SModbusPoint{Addr: 0x01A2, SPoint: &c_base.SPoint{Key: "TotalApparentPower", Name: "总视在功率", Unit: "kVA", Desc: "总视在功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	ForwardActiveEnergy = &c_proto.SModbusPoint{Addr: 0x01A3, SPoint: &c_base.SPoint{Key: "ForwardActiveEnergy", Name: "正向有功电量", Unit: "kWh", Desc: "正向有功电量"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	ReverseActiveEnergy = &c_proto.SModbusPoint{Addr: 0x01A4, SPoint: &c_base.SPoint{Key: "ReverseActiveEnergy", Name: "反向有功电量", Unit: "kWh", Desc: "反向有功电量"}, DataAccess: c_default.VDataAccessUInt16Scale001}

	// System parameters
	Frequency        = &c_proto.SModbusPoint{Addr: 0x01A5, SPoint: &c_base.SPoint{Key: "Frequency", Name: "频率", Unit: "Hz", Desc: "频率"}, DataAccess: c_default.VDataAccessInt16Scale001}
	PowerFactor      = &c_proto.SModbusPoint{Addr: 0x01A6, SPoint: &c_base.SPoint{Key: "PowerFactor", Name: "功率因数", Desc: "功率因数"}, DataAccess: c_default.VDataAccessInt16Scale001}
	RatedLineVoltage = &c_proto.SModbusPoint{Addr: 0x01A7, SPoint: &c_base.SPoint{Key: "RatedLineVoltage", Name: "额定线电压", Unit: "V", Desc: "额定线电压"}, DataAccess: c_default.VDataAccessInt16}
	RatedFrequency   = &c_proto.SModbusPoint{Addr: 0x01A8, SPoint: &c_base.SPoint{Key: "RatedFrequency", Name: "额定频率", Unit: "Hz", Desc: "额定频率"}, DataAccess: c_default.VDataAccessInt16}
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
