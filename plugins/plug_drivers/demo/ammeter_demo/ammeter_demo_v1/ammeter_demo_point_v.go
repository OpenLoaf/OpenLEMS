package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_proto"
	"time"
)

var (
	// Voltage measurements
	Status        = &c_base.Meta{Name: "Status", Cn: "设备状态字", Addr: 0x0190, ReadType: c_base.RInt16}
	PhaseAVoltage = &c_base.Meta{Name: "PhaseAVoltage", Cn: "A相电压", Addr: 0x0191, ReadType: c_base.RInt16, Unit: "V"}
	PhaseBVoltage = &c_base.Meta{Name: "PhaseBVoltage", Cn: "B相电压", Addr: 0x0192, ReadType: c_base.RInt16, Unit: "V"}
	PhaseCVoltage = &c_base.Meta{Name: "PhaseCVoltage", Cn: "C相电压", Addr: 0x0193, ReadType: c_base.RInt16, Unit: "V"}

	// Current measurements
	PhaseACurrent = &c_base.Meta{Name: "PhaseACurrent", Cn: "A相电流", Addr: 0x0194, ReadType: c_base.RInt16, Unit: "A"}
	PhaseBCurrent = &c_base.Meta{Name: "PhaseBCurrent", Cn: "B相电流", Addr: 0x0195, ReadType: c_base.RInt16, Unit: "A"}
	PhaseCCurrent = &c_base.Meta{Name: "PhaseCCurrent", Cn: "C相电流", Addr: 0x0196, ReadType: c_base.RInt16, Unit: "A"}

	// Power measurements
	PhaseAActivePower   = &c_base.Meta{Name: "PhaseAActivePower", Cn: "A相有功功率", Addr: 0x0197, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseBActivePower   = &c_base.Meta{Name: "PhaseBActivePower", Cn: "B相有功功率", Addr: 0x0198, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseCActivePower   = &c_base.Meta{Name: "PhaseCActivePower", Cn: "C相有功功率", Addr: 0x0199, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseAReactivePower = &c_base.Meta{Name: "PhaseAReactivePower", Cn: "A相无功功率", Addr: 0x019A, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseBReactivePower = &c_base.Meta{Name: "PhaseBReactivePower", Cn: "B相无功功率", Addr: 0x019B, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseCReactivePower = &c_base.Meta{Name: "PhaseCReactivePower", Cn: "C相无功功率", Addr: 0x019C, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseAApparentPower = &c_base.Meta{Name: "PhaseAApparentPower", Cn: "A相视在功率", Addr: 0x019D, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	PhaseBApparentPower = &c_base.Meta{Name: "PhaseBApparentPower", Cn: "B相视在功率", Addr: 0x019E, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	PhaseCApparentPower = &c_base.Meta{Name: "PhaseCApparentPower", Cn: "C相视在功率", Addr: 0x019F, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}

	// Total power and energy
	TotalActivePower    = &c_base.Meta{Name: "TotalActivePower", Cn: "总有功功率", Addr: 0x01A0, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	TotalReactivePower  = &c_base.Meta{Name: "TotalReactivePower", Cn: "总无功功率", Addr: 0x01A1, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	TotalApparentPower  = &c_base.Meta{Name: "TotalApparentPower", Cn: "总视在功率", Addr: 0x01A2, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	ForwardActiveEnergy = &c_base.Meta{Name: "ForwardActiveEnergy", Cn: "正向有功电量", Addr: 0x01A3, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "kWh"}
	ReverseActiveEnergy = &c_base.Meta{Name: "ReverseActiveEnergy", Cn: "反向有功电量", Addr: 0x01A4, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "kWh"}

	// System parameters
	Frequency        = &c_base.Meta{Name: "Frequency", Cn: "频率", Addr: 0x01A5, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "Hz"}
	PowerFactor      = &c_base.Meta{Name: "PowerFactor", Cn: "功率因数", Addr: 0x01A6, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.001}
	RatedLineVoltage = &c_base.Meta{Name: "RatedLineVoltage", Cn: "额定线电压", Addr: 0x01A7, ReadType: c_base.RInt16, Unit: "V"}
	RatedFrequency   = &c_base.Meta{Name: "RatedFrequency", Cn: "额定频率", Addr: 0x01A8, ReadType: c_base.RInt16, Unit: "Hz"}
)

var ReadTask = &c_proto.SModbusTask{
	Name:        "ReadTask",
	DisplayName: "查询数据",
	Addr:        Status.Addr,
	Quantity:    RatedFrequency.Addr - Status.Addr + 1,
	Function:    c_proto.EMqHoldingRegisters,
	CycleMill:   500,
	Lifetime:    3 * time.Second,
	Metas: []*c_base.Meta{
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
