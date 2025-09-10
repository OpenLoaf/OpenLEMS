package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_proto"
	"time"
)

var (
	// Voltage measurements
	Status        = &c_base.SModbusPoint{Name: "Status", Cn: "设备状态字", Addr: 0x0190, ReadType: c_base.RInt16}
	PhaseAVoltage = &c_base.SModbusPoint{Name: "PhaseAVoltage", Cn: "A相电压", Addr: 0x0191, ReadType: c_base.RInt16, Unit: "V"}
	PhaseBVoltage = &c_base.SModbusPoint{Name: "PhaseBVoltage", Cn: "B相电压", Addr: 0x0192, ReadType: c_base.RInt16, Unit: "V"}
	PhaseCVoltage = &c_base.SModbusPoint{Name: "PhaseCVoltage", Cn: "C相电压", Addr: 0x0193, ReadType: c_base.RInt16, Unit: "V"}

	// Current measurements
	PhaseACurrent = &c_base.SModbusPoint{Name: "PhaseACurrent", Cn: "A相电流", Addr: 0x0194, ReadType: c_base.RInt16, Unit: "A"}
	PhaseBCurrent = &c_base.SModbusPoint{Name: "PhaseBCurrent", Cn: "B相电流", Addr: 0x0195, ReadType: c_base.RInt16, Unit: "A"}
	PhaseCCurrent = &c_base.SModbusPoint{Name: "PhaseCCurrent", Cn: "C相电流", Addr: 0x0196, ReadType: c_base.RInt16, Unit: "A"}

	// Power measurements
	PhaseAActivePower   = &c_base.SModbusPoint{Name: "PhaseAActivePower", Cn: "A相有功功率", Addr: 0x0197, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseBActivePower   = &c_base.SModbusPoint{Name: "PhaseBActivePower", Cn: "B相有功功率", Addr: 0x0198, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseCActivePower   = &c_base.SModbusPoint{Name: "PhaseCActivePower", Cn: "C相有功功率", Addr: 0x0199, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	PhaseAReactivePower = &c_base.SModbusPoint{Name: "PhaseAReactivePower", Cn: "A相无功功率", Addr: 0x019A, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseBReactivePower = &c_base.SModbusPoint{Name: "PhaseBReactivePower", Cn: "B相无功功率", Addr: 0x019B, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseCReactivePower = &c_base.SModbusPoint{Name: "PhaseCReactivePower", Cn: "C相无功功率", Addr: 0x019C, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	PhaseAApparentPower = &c_base.SModbusPoint{Name: "PhaseAApparentPower", Cn: "A相视在功率", Addr: 0x019D, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	PhaseBApparentPower = &c_base.SModbusPoint{Name: "PhaseBApparentPower", Cn: "B相视在功率", Addr: 0x019E, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	PhaseCApparentPower = &c_base.SModbusPoint{Name: "PhaseCApparentPower", Cn: "C相视在功率", Addr: 0x019F, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}

	// Total power and energy
	TotalActivePower    = &c_base.SModbusPoint{Name: "TotalActivePower", Cn: "总有功功率", Addr: 0x01A0, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW"}
	TotalReactivePower  = &c_base.SModbusPoint{Name: "TotalReactivePower", Cn: "总无功功率", Addr: 0x01A1, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar"}
	TotalApparentPower  = &c_base.SModbusPoint{Name: "TotalApparentPower", Cn: "总视在功率", Addr: 0x01A2, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVA"}
	ForwardActiveEnergy = &c_base.SModbusPoint{Name: "ForwardActiveEnergy", Cn: "正向有功电量", Addr: 0x01A3, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "kWh"}
	ReverseActiveEnergy = &c_base.SModbusPoint{Name: "ReverseActiveEnergy", Cn: "反向有功电量", Addr: 0x01A4, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "kWh"}

	// System parameters
	Frequency        = &c_base.SModbusPoint{Name: "Frequency", Cn: "频率", Addr: 0x01A5, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "Hz"}
	PowerFactor      = &c_base.SModbusPoint{Name: "PowerFactor", Cn: "功率因数", Addr: 0x01A6, ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.001}
	RatedLineVoltage = &c_base.SModbusPoint{Name: "RatedLineVoltage", Cn: "额定线电压", Addr: 0x01A7, ReadType: c_base.RInt16, Unit: "V"}
	RatedFrequency   = &c_base.SModbusPoint{Name: "RatedFrequency", Cn: "额定频率", Addr: 0x01A8, ReadType: c_base.RInt16, Unit: "Hz"}
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:        "ReadTask",
	DisplayName: "查询数据",
	Addr:        Status.Addr,
	Quantity:    RatedFrequency.Addr - Status.Addr + 1,
	Function:    c_enum.EMqHoldingRegisters,
	CycleMill:   500,
	Lifetime:    3 * time.Second,
	Metas: []*c_base.SModbusPoint{
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
