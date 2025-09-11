package elecod_mac_defined

import (
	"common/c_base"
	"common/c_default"
	"common/c_proto"
)

var (
	AnalogAcGroup    = &c_base.SPointGroup{GroupName: "交流侧", GroupSort: 10}
	AnalogDcGroup    = &c_base.SPointGroup{GroupName: "直流侧", GroupSort: 20}
	AnalogTotalGroup = &c_base.SPointGroup{GroupName: "累计", GroupSort: 30}
	AnalogOtherGroup = &c_base.SPointGroup{GroupName: "其他遥测", GroupSort: 40}
)

var (
	// 有功功率 (0x05) - 调整到最前面
	analogActivePowerA     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPa, DataAccess: c_default.VDataAccessFloat32Byte0Scale01, Group: AnalogAcGroup}
	analogActivePowerB     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPb, DataAccess: c_default.VDataAccessFloat32Byte1Scale01, Group: AnalogAcGroup}
	analogActivePowerC     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPc, DataAccess: c_default.VDataAccessFloat32Byte2Scale01, Group: AnalogAcGroup}
	AnalogTotalActivePower = &c_proto.SCanbusPoint{SPoint: c_default.VPointP, DataAccess: c_default.VDataAccessFloat32Byte3Scale01, Group: AnalogAcGroup}

	// 直流参数 (0x01)
	analogDcVoltage  = &c_proto.SCanbusPoint{SPoint: c_default.VPointDcVoltage, DataAccess: c_default.VDataAccessFloat32Byte0Scale01, Group: AnalogDcGroup}
	analogDcCurrent  = &c_proto.SCanbusPoint{SPoint: c_default.VPointDcCurrent, DataAccess: c_default.VDataAccessFloat32Byte1Scale01, Group: AnalogDcGroup}
	analogDcPower    = &c_proto.SCanbusPoint{SPoint: c_default.VPointDcPower, DataAccess: c_default.VDataAccessFloat32Byte2Scale01, Group: AnalogDcGroup}
	analogBusVoltage = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "BusVoltage", Name: "直流母线电压", Unit: "V", Desc: "直流母线电压", Sort: 8, Precise: 1, Group: AnalogDcGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale01}

	// 电网电压 (0x02)
	analogGridVoltageA  = &c_proto.SCanbusPoint{SPoint: c_default.VPointUa, DataAccess: c_default.VDataAccessFloat32Byte0Scale01, Group: AnalogAcGroup}
	analogGridVoltageB  = &c_proto.SCanbusPoint{SPoint: c_default.VPointUb, DataAccess: c_default.VDataAccessFloat32Byte1Scale01, Group: AnalogAcGroup}
	analogGridVoltageC  = &c_proto.SCanbusPoint{SPoint: c_default.VPointUc, DataAccess: c_default.VDataAccessFloat32Byte2Scale01, Group: AnalogAcGroup}
	analogPowerTubeTemp = &c_proto.SCanbusPoint{SPoint: c_default.VPointIGBTTemp, DataAccess: c_default.VDataAccessFloat32Byte3Scale1, Group: AnalogAcGroup}

	// 电网电流 (0x03)
	analogGridCurrentA = &c_proto.SCanbusPoint{SPoint: c_default.VPointIa, DataAccess: c_default.VDataAccessFloat32Byte0Scale01, Group: AnalogAcGroup}
	analogGridCurrentB = &c_proto.SCanbusPoint{SPoint: c_default.VPointIb, DataAccess: c_default.VDataAccessFloat32Byte1Scale01, Group: AnalogAcGroup}
	analogGridCurrentC = &c_proto.SCanbusPoint{SPoint: c_default.VPointIc, DataAccess: c_default.VDataAccessFloat32Byte2Scale01, Group: AnalogAcGroup}
	analogBridgeTemp   = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "BridgeTemp", Name: "平衡桥温度", Unit: "℃", Desc: "平衡桥温度", Sort: 16, Precise: 1, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale1}

	// 电网频率 (0x04)
	AnalogGridFrequencyA = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "FreqA", Name: "电网频率A", Unit: "Hz", Desc: "电网频率A", Sort: 17, Precise: 2, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte0Scale001}
	AnalogGridFrequencyB = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "FreqB", Name: "电网频率B", Unit: "Hz", Desc: "电网频率B", Sort: 18, Precise: 2, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte1Scale001}
	AnalogGridFrequencyC = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "FreqC", Name: "电网频率C", Unit: "Hz", Desc: "电网频率C", Sort: 19, Precise: 2, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte2Scale001}
	analogAmbientTemp    = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "AmbientTemp", Name: "环境温度", Unit: "℃", Desc: "环境温度", Sort: 20, Precise: 1, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale1}

	// 无功功率 (0x06)
	analogReactivePowerA     = &c_proto.SCanbusPoint{SPoint: c_default.VPointQa, DataAccess: c_default.VDataAccessFloat32Byte0Scale01}
	analogReactivePowerB     = &c_proto.SCanbusPoint{SPoint: c_default.VPointQb, DataAccess: c_default.VDataAccessFloat32Byte1Scale01}
	analogReactivePowerC     = &c_proto.SCanbusPoint{SPoint: c_default.VPointQc, DataAccess: c_default.VDataAccessFloat32Byte2Scale01}
	AnalogTotalReactivePower = &c_proto.SCanbusPoint{SPoint: c_default.VPointQ, DataAccess: c_default.VDataAccessFloat32Byte3Scale01}

	// 视在功率 (0x08) - 使用预定义点位
	analogApparentPowerA     = &c_proto.SCanbusPoint{SPoint: c_default.VPointSa, DataAccess: c_default.VDataAccessFloat32Byte0Scale01}
	analogApparentPowerB     = &c_proto.SCanbusPoint{SPoint: c_default.VPointSb, DataAccess: c_default.VDataAccessFloat32Byte1Scale01}
	analogApparentPowerC     = &c_proto.SCanbusPoint{SPoint: c_default.VPointSc, DataAccess: c_default.VDataAccessFloat32Byte2Scale01}
	AnalogTotalApparentPower = &c_proto.SCanbusPoint{SPoint: c_default.VPointS, DataAccess: c_default.VDataAccessFloat32Byte3Scale01}

	// 功率因数 (0x07)
	analogPowerFactorA     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPFa, DataAccess: c_default.VDataAccessFloat32Byte0Scale001}
	analogPowerFactorB     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPFb, DataAccess: c_default.VDataAccessFloat32Byte1Scale001}
	analogPowerFactorC     = &c_proto.SCanbusPoint{SPoint: c_default.VPointPFc, DataAccess: c_default.VDataAccessFloat32Byte2Scale001}
	AnalogTotalPowerFactor = &c_proto.SCanbusPoint{SPoint: c_default.VPointPF, DataAccess: c_default.VDataAccessFloat32Byte3Scale001}

	// 逆变电压 (0x09)
	analogInverterVoltageA = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "InverterVoltageA", Name: "逆变电压A(AB)", Unit: "V", Desc: "逆变电压A(AB)", Sort: 33, Precise: 1, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte0Scale01}
	analogInverterVoltageB = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "InverterVoltageB", Name: "逆变电压B(BC)", Unit: "V", Desc: "逆变电压B(BC)", Sort: 34, Precise: 1, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte1Scale01}
	analogInverterVoltageC = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "InverterVoltageC", Name: "逆变电压C(CA)", Unit: "V", Desc: "逆变电压C(CA)", Sort: 35, Precise: 1, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte2Scale01}
	analogReserved1        = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "Reserved1", Name: "预留", Unit: "", Desc: "预留", Sort: 36, Precise: 0, Group: AnalogAcGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale1}

	// 对地阻抗和漏电流 (0x0A)
	analogPositiveGroundImpedance = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "PositiveGroundImpedance", Name: "正对地阻抗", Unit: "kΩ", Desc: "正对地阻抗", Sort: 37, Precise: 0, Group: AnalogOtherGroup}, DataAccess: c_default.VDataAccessFloat32Byte0Scale1}
	analogNegativeGroundImpedance = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "NegativeGroundImpedance", Name: "负对地阻抗", Unit: "kΩ", Desc: "负对地阻抗", Sort: 38, Precise: 0, Group: AnalogOtherGroup}, DataAccess: c_default.VDataAccessFloat32Byte1Scale1}
	analogLeakageCurrent          = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "LeakageCurrent", Name: "漏电流", Unit: "A", Desc: "漏电流", Sort: 39, Precise: 2, Group: AnalogOtherGroup}, DataAccess: c_default.VDataAccessFloat32Byte2Scale001}
	analogConverterEfficiency     = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "ConverterEfficiency", Name: "变流器效率", Unit: "%", Desc: "变流器效率", Sort: 40, Precise: 1, Group: AnalogOtherGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale01}

	// 充电量 (0x0B) - 使用预定义点位
	analogTotalChargeL    = &c_proto.SCanbusPoint{SPoint: c_default.VPointTotalCharge, DataAccess: c_default.VDataAccessFloat32Byte0Scale01, Group: AnalogTotalGroup}
	analogTotalChargeH    = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "TotalChargeH", Name: "总充电量H", Unit: "kWh", Desc: "总充电量H", Sort: 42, Precise: 2, Group: AnalogTotalGroup}, DataAccess: c_default.VDataAccessFloat32Byte1Scale1}
	analogTotalDischargeL = &c_proto.SCanbusPoint{SPoint: c_default.VPointTotalDischarge, DataAccess: c_default.VDataAccessFloat32Byte2Scale01, Group: AnalogTotalGroup}
	analogTotalDischargeH = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "TotalDischargeH", Name: "总放电量H", Unit: "kWh", Desc: "总放电量H", Sort: 44, Precise: 2, Group: AnalogTotalGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale1}

	// 母线电压 (0x0C)
	analogPositiveBusVoltage    = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "PositiveBusVoltage", Name: "正母线电压", Unit: "V", Desc: "正母线电压", Sort: 45, Precise: 1, Group: AnalogDcGroup}, DataAccess: c_default.VDataAccessFloat32Byte0Scale01}
	analogNegativeBusVoltage    = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "NegativeBusVoltage", Name: "负母线电压", Unit: "V", Desc: "负母线电压", Sort: 46, Precise: 1, Group: AnalogDcGroup}, DataAccess: c_default.VDataAccessFloat32Byte1Scale01}
	analogGroundNegativeVoltage = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "GroundNegativeVoltage", Name: "地对负电压", Unit: "V", Desc: "地对负电压", Sort: 47, Precise: 1, Group: AnalogDcGroup}, DataAccess: c_default.VDataAccessFloat32Byte2Scale01}
	analogReserved2             = &c_proto.SCanbusPoint{SPoint: &c_base.SPoint{Key: "Reserved2", Name: "预留", Unit: "", Desc: "预留", Sort: 48, Precise: 0, Group: AnalogDcGroup}, DataAccess: c_default.VDataAccessFloat32Byte3Scale1}
)
