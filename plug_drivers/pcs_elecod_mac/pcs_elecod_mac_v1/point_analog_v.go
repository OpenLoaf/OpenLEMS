package pcs_elecod_mac_v1

import "common/c_base"

var (
	// 直流参数 (0x01)
	analogDcVoltage  = &c_base.Meta{Name: "AnalogDcVoltage", Cn: "直流电压", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogDcCurrent  = &c_base.Meta{Name: "AnalogDcCurrent", Cn: "直流电流", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogDcPower    = &c_base.Meta{Name: "AnalogDcPower", Cn: "直流功率", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogBusVoltage = &c_base.Meta{Name: "AnalogBusVoltage", Cn: "母线电压", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网电压 (0x02)
	analogGridVoltageA  = &c_base.Meta{Name: "AnalogGridVoltageA", Cn: "电网电压A(AB)", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGridVoltageB  = &c_base.Meta{Name: "AnalogGridVoltageB", Cn: "电网电压B(BC)", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGridVoltageC  = &c_base.Meta{Name: "AnalogGridVoltageC", Cn: "电网电压C(CA)", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogPowerTubeTemp = &c_base.Meta{Name: "AnalogPowerTubeTemp", Cn: "功率管温度", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 电网电流 (0x03)
	analogGridCurrentA = &c_base.Meta{Name: "AnalogGridCurrentA", Cn: "电网电流A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogGridCurrentB = &c_base.Meta{Name: "AnalogGridCurrentB", Cn: "电网电流B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogGridCurrentC = &c_base.Meta{Name: "AnalogGridCurrentC", Cn: "电网电流C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogBridgeTemp   = &c_base.Meta{Name: "AnalogBridgeTemp", Cn: "平衡桥温度", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 电网频率 (0x04)
	analogGridFrequencyA = &c_base.Meta{Name: "AnalogGridFrequencyA", Cn: "电网频率A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogGridFrequencyB = &c_base.Meta{Name: "AnalogGridFrequencyB", Cn: "电网频率B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogGridFrequencyC = &c_base.Meta{Name: "AnalogGridFrequencyC", Cn: "电网频率C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogAmbientTemp    = &c_base.Meta{Name: "AnalogAmbientTemp", Cn: "环境温度", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 有功功率 (0x05)
	analogActivePowerA     = &c_base.Meta{Name: "AnalogActivePowerA", Cn: "有功功率A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogActivePowerB     = &c_base.Meta{Name: "AnalogActivePowerB", Cn: "有功功率B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogActivePowerC     = &c_base.Meta{Name: "AnalogActivePowerC", Cn: "有功功率C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogTotalActivePower = &c_base.Meta{Name: "AnalogTotalActivePower", Cn: "总有功功率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 无功功率 (0x06)
	analogReactivePowerA     = &c_base.Meta{Name: "AnalogReactivePowerA", Cn: "无功功率A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogReactivePowerB     = &c_base.Meta{Name: "AnalogReactivePowerB", Cn: "无功功率B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogReactivePowerC     = &c_base.Meta{Name: "AnalogReactivePowerC", Cn: "无功功率C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogTotalReactivePower = &c_base.Meta{Name: "AnalogTotalReactivePower", Cn: "总无功功率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}

	// 功率因素 (0x07)
	analogPowerFactorA     = &c_base.Meta{Name: "AnalogPowerFactorA", Cn: "功率因素A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogPowerFactorB     = &c_base.Meta{Name: "AnalogPowerFactorB", Cn: "功率因素B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogPowerFactorC     = &c_base.Meta{Name: "AnalogPowerFactorC", Cn: "功率因素C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogTotalPowerFactor = &c_base.Meta{Name: "AnalogTotalPowerFactor", Cn: "总功率因素", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 视在功率 (0x08)
	analogApparentPowerA     = &c_base.Meta{Name: "AnalogApparentPowerA", Cn: "视在功率A", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogApparentPowerB     = &c_base.Meta{Name: "AnalogApparentPowerB", Cn: "视在功率B", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogApparentPowerC     = &c_base.Meta{Name: "AnalogApparentPowerC", Cn: "视在功率C", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogTotalApparentPower = &c_base.Meta{Name: "AnalogTotalApparentPower", Cn: "总视在功率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}

	// 逆变电压 (0x09)
	analogInverterVoltageA = &c_base.Meta{Name: "AnalogInverterVoltageA", Cn: "逆变电压A(AB)", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogInverterVoltageB = &c_base.Meta{Name: "AnalogInverterVoltageB", Cn: "逆变电压B(BC)", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogInverterVoltageC = &c_base.Meta{Name: "AnalogInverterVoltageC", Cn: "逆变电压C(CA)", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogReserved1        = &c_base.Meta{Name: "AnalogReserved1", Cn: "预留", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 对地阻抗和漏电流 (0x0A)
	analogPositiveGroundImpedance = &c_base.Meta{Name: "AnalogPositiveGroundImpedance", Cn: "正对地阻抗", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	analogNegativeGroundImpedance = &c_base.Meta{Name: "AnalogNegativeGroundImpedance", Cn: "负对地阻抗", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	analogLeakageCurrent          = &c_base.Meta{Name: "AnalogLeakageCurrent", Cn: "漏电流", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "A"}
	analogConverterEfficiency     = &c_base.Meta{Name: "AnalogConverterEfficiency", Cn: "变流器效率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "%"}

	// 充电量 (0x0B)
	analogTotalChargeL    = &c_base.Meta{Name: "AnalogTotalChargeL", Cn: "总充电量L", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kWh"}
	analogTotalChargeH    = &c_base.Meta{Name: "AnalogTotalChargeH", Cn: "总充电量H", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kWh"}
	analogTotalDischargeL = &c_base.Meta{Name: "AnalogTotalDischargeL", Cn: "总放电量L", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kWh"}
	analogTotalDischargeH = &c_base.Meta{Name: "AnalogTotalDischargeH", Cn: "总放电量H", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kWh"}

	// 母线电压 (0x0C)
	analogPositiveBusVoltage    = &c_base.Meta{Name: "AnalogPositiveBusVoltage", Cn: "正母线电压", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogNegativeBusVoltage    = &c_base.Meta{Name: "AnalogNegativeBusVoltage", Cn: "负母线电压", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGroundNegativeVoltage = &c_base.Meta{Name: "AnalogGroundNegativeVoltage", Cn: "地对负电压", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogReserved2             = &c_base.Meta{Name: "AnalogReserved2", Cn: "预留", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
)
