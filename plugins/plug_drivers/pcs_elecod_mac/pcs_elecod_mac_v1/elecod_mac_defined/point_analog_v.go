package elecod_mac_defined

import "common/c_base"

var (
	AnalogAcGroup    = &c_base.MetaGroup{GroupName: "交流侧", GroupSort: 10}
	AnalogDcGroup    = &c_base.MetaGroup{GroupName: "交流侧", GroupSort: 20}
	AnalogTotalGroup = &c_base.MetaGroup{GroupName: "累计", GroupSort: 30}
	AnalogOtherGroup = &c_base.MetaGroup{GroupName: "其他遥测", GroupSort: 40}
)

var (
	// 有功功率 (0x05) - 调整到最前面
	analogActivePowerA     = &c_base.Meta{Name: "AnalogActivePowerA", Group: AnalogAcGroup, Cn: "有功功率A", Addr: 0, Sort: 1, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogActivePowerB     = &c_base.Meta{Name: "AnalogActivePowerB", Group: AnalogAcGroup, Cn: "有功功率B", Addr: 2, Sort: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogActivePowerC     = &c_base.Meta{Name: "AnalogActivePowerC", Group: AnalogAcGroup, Cn: "有功功率C", Addr: 4, Sort: 3, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	AnalogTotalActivePower = &c_base.Meta{Name: "AnalogTotalActivePower", Group: AnalogAcGroup, Cn: "总有功功率", Debug: true, Addr: 6, Sort: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 直流参数 (0x01)
	analogDcVoltage  = &c_base.Meta{Name: "AnalogDcVoltage", Group: AnalogDcGroup, Cn: "直流电压", Addr: 0, Sort: 5, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogDcCurrent  = &c_base.Meta{Name: "AnalogDcCurrent", Group: AnalogDcGroup, Cn: "直流电流", Addr: 2, Sort: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogDcPower    = &c_base.Meta{Name: "AnalogDcPower", Group: AnalogDcGroup, Cn: "直流功率", Addr: 4, Sort: 7, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	analogBusVoltage = &c_base.Meta{Name: "AnalogBusVoltage", Group: AnalogDcGroup, Cn: "直流母线电压", Addr: 6, Sort: 8, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网电压 (0x02)
	analogGridVoltageA  = &c_base.Meta{Name: "AnalogGridVoltageA", Group: AnalogAcGroup, Cn: "电网电压A(AB)", Addr: 0, Sort: 9, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGridVoltageB  = &c_base.Meta{Name: "AnalogGridVoltageB", Group: AnalogAcGroup, Cn: "电网电压B(BC)", Addr: 2, Sort: 10, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGridVoltageC  = &c_base.Meta{Name: "AnalogGridVoltageC", Group: AnalogAcGroup, Cn: "电网电压C(CA)", Addr: 4, Sort: 11, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogPowerTubeTemp = &c_base.Meta{Name: "AnalogPowerTubeTemp", Group: AnalogAcGroup, Cn: "功率管温度", Addr: 6, Sort: 12, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 电网电流 (0x03)
	analogGridCurrentA = &c_base.Meta{Name: "AnalogGridCurrentA", Group: AnalogAcGroup, Cn: "电网电流A", Addr: 0, Sort: 13, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogGridCurrentB = &c_base.Meta{Name: "AnalogGridCurrentB", Group: AnalogAcGroup, Cn: "电网电流B", Addr: 2, Sort: 14, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogGridCurrentC = &c_base.Meta{Name: "AnalogGridCurrentC", Group: AnalogAcGroup, Cn: "电网电流C", Addr: 4, Sort: 15, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	analogBridgeTemp   = &c_base.Meta{Name: "AnalogBridgeTemp", Group: AnalogAcGroup, Cn: "平衡桥温度", Addr: 6, Sort: 16, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 电网频率 (0x04)
	analogGridFrequencyA = &c_base.Meta{Name: "AnalogGridFrequencyA", Group: AnalogAcGroup, Cn: "电网频率A", Addr: 0, Sort: 17, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogGridFrequencyB = &c_base.Meta{Name: "AnalogGridFrequencyB", Group: AnalogAcGroup, Cn: "电网频率B", Addr: 2, Sort: 18, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogGridFrequencyC = &c_base.Meta{Name: "AnalogGridFrequencyC", Group: AnalogAcGroup, Cn: "电网频率C", Addr: 4, Sort: 19, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	analogAmbientTemp    = &c_base.Meta{Name: "AnalogAmbientTemp", Group: AnalogAcGroup, Cn: "环境温度", Addr: 6, Sort: 20, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "℃"}

	// 无功功率 (0x06)
	analogReactivePowerA     = &c_base.Meta{Name: "AnalogReactivePowerA", Group: AnalogAcGroup, Cn: "无功功率A", Addr: 0, Sort: 21, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogReactivePowerB     = &c_base.Meta{Name: "AnalogReactivePowerB", Group: AnalogAcGroup, Cn: "无功功率B", Addr: 2, Sort: 22, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogReactivePowerC     = &c_base.Meta{Name: "AnalogReactivePowerC", Group: AnalogAcGroup, Cn: "无功功率C", Addr: 4, Sort: 23, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	analogTotalReactivePower = &c_base.Meta{Name: "AnalogTotalReactivePower", Group: AnalogAcGroup, Cn: "总无功功率", Addr: 6, Sort: 24, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}

	// 功率因素 (0x07)
	analogPowerFactorA     = &c_base.Meta{Name: "AnalogPowerFactorA", Group: AnalogAcGroup, Cn: "功率因素A", Addr: 0, Sort: 25, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogPowerFactorB     = &c_base.Meta{Name: "AnalogPowerFactorB", Group: AnalogAcGroup, Cn: "功率因素B", Addr: 2, Sort: 26, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogPowerFactorC     = &c_base.Meta{Name: "AnalogPowerFactorC", Group: AnalogAcGroup, Cn: "功率因素C", Addr: 4, Sort: 27, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	analogTotalPowerFactor = &c_base.Meta{Name: "AnalogTotalPowerFactor", Group: AnalogAcGroup, Cn: "总功率因素", Addr: 6, Sort: 28, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 视在功率 (0x08)
	analogApparentPowerA     = &c_base.Meta{Name: "AnalogApparentPowerA", Group: AnalogAcGroup, Cn: "视在功率A", Addr: 0, Sort: 29, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogApparentPowerB     = &c_base.Meta{Name: "AnalogApparentPowerB", Group: AnalogAcGroup, Cn: "视在功率B", Addr: 2, Sort: 30, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogApparentPowerC     = &c_base.Meta{Name: "AnalogApparentPowerC", Group: AnalogAcGroup, Cn: "视在功率C", Addr: 4, Sort: 31, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}
	analogTotalApparentPower = &c_base.Meta{Name: "AnalogTotalApparentPower", Group: AnalogAcGroup, Cn: "总视在功率", Addr: 6, Sort: 32, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kVA"}

	// 逆变电压 (0x09)
	analogInverterVoltageA = &c_base.Meta{Name: "AnalogInverterVoltageA", Group: AnalogAcGroup, Cn: "逆变电压A(AB)", Addr: 0, Sort: 33, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogInverterVoltageB = &c_base.Meta{Name: "AnalogInverterVoltageB", Group: AnalogAcGroup, Cn: "逆变电压B(BC)", Addr: 2, Sort: 34, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogInverterVoltageC = &c_base.Meta{Name: "AnalogInverterVoltageC", Group: AnalogAcGroup, Cn: "逆变电压C(CA)", Addr: 4, Sort: 35, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogReserved1        = &c_base.Meta{Name: "AnalogReserved1", Group: AnalogAcGroup, Cn: "预留", Addr: 6, Sort: 36, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 对地阻抗和漏电流 (0x0A)
	analogPositiveGroundImpedance = &c_base.Meta{Name: "AnalogPositiveGroundImpedance", Group: AnalogOtherGroup, Cn: "正对地阻抗", Addr: 0, Sort: 37, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	analogNegativeGroundImpedance = &c_base.Meta{Name: "AnalogNegativeGroundImpedance", Group: AnalogOtherGroup, Cn: "负对地阻抗", Addr: 2, Sort: 38, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	analogLeakageCurrent          = &c_base.Meta{Name: "AnalogLeakageCurrent", Group: AnalogOtherGroup, Cn: "漏电流", Addr: 4, Sort: 39, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "A"}
	analogConverterEfficiency     = &c_base.Meta{Name: "AnalogConverterEfficiency", Group: AnalogOtherGroup, Cn: "变流器效率", Addr: 6, Sort: 40, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "%"}

	// 充电量 (0x0B)
	analogTotalChargeL    = &c_base.Meta{Name: "AnalogTotalChargeL", Group: AnalogTotalGroup, Cn: "总充电量L", Addr: 0, Sort: 41, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kWh"}
	analogTotalChargeH    = &c_base.Meta{Name: "AnalogTotalChargeH", Group: AnalogTotalGroup, Cn: "总充电量H", Addr: 2, Sort: 42, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kWh"}
	analogTotalDischargeL = &c_base.Meta{Name: "AnalogTotalDischargeL", Group: AnalogTotalGroup, Cn: "总放电量L", Addr: 4, Sort: 43, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kWh"}
	analogTotalDischargeH = &c_base.Meta{Name: "AnalogTotalDischargeH", Group: AnalogTotalGroup, Cn: "总放电量H", Addr: 6, Sort: 44, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kWh"}

	// 母线电压 (0x0C)
	analogPositiveBusVoltage    = &c_base.Meta{Name: "AnalogPositiveBusVoltage", Group: AnalogDcGroup, Cn: "正母线电压", Addr: 0, Sort: 45, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogNegativeBusVoltage    = &c_base.Meta{Name: "AnalogNegativeBusVoltage", Group: AnalogDcGroup, Cn: "负母线电压", Addr: 2, Sort: 46, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogGroundNegativeVoltage = &c_base.Meta{Name: "AnalogGroundNegativeVoltage", Group: AnalogDcGroup, Cn: "地对负电压", Addr: 4, Sort: 47, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	analogReserved2             = &c_base.Meta{Name: "AnalogReserved2", Group: AnalogDcGroup, Cn: "预留", Addr: 6, Sort: 48, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
)
