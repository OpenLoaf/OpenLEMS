package pcs_elecod_mac_v1

import "common/c_base"

var (
	ConfigMainGroup   = &c_base.MetaGroup{GroupName: "主要配置", GroupSort: 25}
	ConfigBatterGroup = &c_base.MetaGroup{GroupName: "电池配置", GroupSort: 29}
	ConfigAcGroup     = &c_base.MetaGroup{GroupName: "交流参数配置", GroupSort: 28}
)

var (
	// 总功率参数（0x01）
	confTotalActivePower   = &c_base.Meta{Name: "confTotalActivePower", Group: ConfigMainGroup, Cn: "总有功功率", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confTotalReactivePower = &c_base.Meta{Name: "confTotalReactivePower", Group: ConfigMainGroup, Cn: "总无功功率", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confTotalPowerFactor   = &c_base.Meta{Name: "confTotalPowerFactor", Group: ConfigMainGroup, Cn: "总功率因数", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confRatedPower         = &c_base.Meta{Name: "confRatedPower", Group: ConfigMainGroup, Cn: "额定功率", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 相功率参数（0x02）
	confActivePowerA      = &c_base.Meta{Name: "confActivePowerA", Group: ConfigMainGroup, Cn: "A相有功功率", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confActivePowerB      = &c_base.Meta{Name: "confActivePowerB", Group: ConfigMainGroup, Cn: "B相有功功率", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confActivePowerC      = &c_base.Meta{Name: "confActivePowerC", Group: ConfigMainGroup, Cn: "C相有功功率", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confRatedPhaseVoltage = &c_base.Meta{Name: "confRatedPhaseVoltage", Group: ConfigMainGroup, Cn: "额定相电压", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 相无功功率和频率（0x03）
	confReactivePowerA = &c_base.Meta{Name: "confReactivePowerA", Group: ConfigMainGroup, Cn: "A相无功功率", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confReactivePowerB = &c_base.Meta{Name: "confReactivePowerB", Group: ConfigMainGroup, Cn: "B相无功功率", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confReactivePowerC = &c_base.Meta{Name: "confReactivePowerC", Group: ConfigMainGroup, Cn: "C相无功功率", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confRatedFrequency = &c_base.Meta{Name: "confRatedFrequency", Group: ConfigMainGroup, Cn: "额定频率", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 相功率因数和有功变化率（0x04）
	confPowerFactorA          = &c_base.Meta{Name: "confPowerFactorA", Group: ConfigMainGroup, Cn: "A相功率因数", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confPowerFactorB          = &c_base.Meta{Name: "confPowerFactorB", Group: ConfigMainGroup, Cn: "B相功率因数", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confPowerFactorC          = &c_base.Meta{Name: "confPowerFactorC", Group: ConfigMainGroup, Cn: "C相功率因数", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confActivePowerChangeRate = &c_base.Meta{Name: "confActivePowerChangeRate", Group: ConfigMainGroup, Cn: "有功变化速率", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 控制位和电池类型（0x05）
	confControlBit1 = &c_base.Meta{Name: "confControlBit1", Group: ConfigBatterGroup, Cn: "控制位1", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confControlBit2 = &c_base.Meta{Name: "confControlBit2", Group: ConfigBatterGroup, Cn: "控制位2", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confControlBit3 = &c_base.Meta{Name: "confControlBit3", Group: ConfigBatterGroup, Cn: "控制位3", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confBatteryType = &c_base.Meta{Name: "confBatteryType", Group: ConfigBatterGroup, Cn: "电池类型", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 电池保护点（0x06）
	confBatteryUndervoltProtect = &c_base.Meta{Name: "confBatteryUndervoltProtect", Group: ConfigBatterGroup, Cn: "电池欠压保护点", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryUndervoltRecover = &c_base.Meta{Name: "confBatteryUndervoltRecover", Group: ConfigBatterGroup, Cn: "电池欠压恢复点", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryOvervoltProtect  = &c_base.Meta{Name: "confBatteryOvervoltProtect", Group: ConfigBatterGroup, Cn: "电池过压保护点", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryOvervoltRecover  = &c_base.Meta{Name: "confBatteryOvervoltRecover", Group: ConfigBatterGroup, Cn: "电池过压恢复点", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电池限流和电压（0x07）
	confBatteryChargeLimit     = &c_base.Meta{Name: "confBatteryChargeLimit", Group: ConfigBatterGroup, Cn: "电池充电限流", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	confBatteryDischargeLimit  = &c_base.Meta{Name: "confBatteryDischargeLimit", Group: ConfigBatterGroup, Cn: "电池放电限流", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	confBatteryFloatVoltage    = &c_base.Meta{Name: "confBatteryFloatVoltage", Group: ConfigBatterGroup, Cn: "电池浮充电压", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryEqualizeVoltage = &c_base.Meta{Name: "confBatteryEqualizeVoltage", Group: ConfigBatterGroup, Cn: "电池均充电压", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网过压保护（0x08）
	confInsulationImpedanceThreshold = &c_base.Meta{Name: "confInsulationImpedanceThreshold", Group: ConfigAcGroup, Cn: "绝缘阻抗保护阈值", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	confGridOvervoltLevel1_2         = &c_base.Meta{Name: "confGridOvervoltLevel1_2", Group: ConfigAcGroup, Cn: "电网过压一级：电网过压二级", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOvervoltLevel3_4         = &c_base.Meta{Name: "confGridOvervoltLevel3_4", Group: ConfigAcGroup, Cn: "电网过压三级：电网过压四级", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOvervoltLevel5_Recover   = &c_base.Meta{Name: "confGridOvervoltLevel5_Recover", Group: ConfigAcGroup, Cn: "电网过压五级：电网过压恢复", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}

	// 电网欠压和过频保护（0x09）
	confGridUndervoltLevel1_2       = &c_base.Meta{Name: "confGridUndervoltLevel1_2", Group: ConfigAcGroup, Cn: "电网欠压一级：电网欠压二级", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridUndervoltLevel3_4       = &c_base.Meta{Name: "confGridUndervoltLevel3_4", Group: ConfigAcGroup, Cn: "电网欠压三级：电网欠压四级", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridUndervoltLevel5_Recover = &c_base.Meta{Name: "confGridUndervoltLevel5_Recover", Group: ConfigAcGroup, Cn: "电网欠压五级：电网欠压恢复", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOverfreqLevel1          = &c_base.Meta{Name: "confGridOverfreqLevel1", Group: ConfigAcGroup, Cn: "电网过频一级", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网过频保护（0x0A）
	confGridOverfreqLevel2 = &c_base.Meta{Name: "confGridOverfreqLevel2", Group: ConfigAcGroup, Cn: "电网过频二级", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel3 = &c_base.Meta{Name: "confGridOverfreqLevel3", Group: ConfigAcGroup, Cn: "电网过频三级", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel4 = &c_base.Meta{Name: "confGridOverfreqLevel4", Group: ConfigAcGroup, Cn: "电网过频四级", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel5 = &c_base.Meta{Name: "confGridOverfreqLevel5", Group: ConfigAcGroup, Cn: "电网过频五级", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网过频和欠频恢复（0x0B）
	confGridOverfreqRecover = &c_base.Meta{Name: "confGridOverfreqRecover", Group: ConfigAcGroup, Cn: "电网过频恢复", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel1 = &c_base.Meta{Name: "confGridUnderfreqLevel1", Group: ConfigAcGroup, Cn: "电网欠频一级", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel2 = &c_base.Meta{Name: "confGridUnderfreqLevel2", Group: ConfigAcGroup, Cn: "电网欠频二级", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel3 = &c_base.Meta{Name: "confGridUnderfreqLevel3", Group: ConfigAcGroup, Cn: "电网欠频三级", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网欠频恢复和直流母线（0x0C）
	confGridUnderfreqLevel4  = &c_base.Meta{Name: "confGridUnderfreqLevel4", Group: ConfigAcGroup, Cn: "电网欠频四级", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel5  = &c_base.Meta{Name: "confGridUnderfreqLevel5", Group: ConfigAcGroup, Cn: "电网欠频五级", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqRecover = &c_base.Meta{Name: "confGridUnderfreqRecover", Group: ConfigAcGroup, Cn: "电网欠频恢复", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confDcBusVoltageRef      = &c_base.Meta{Name: "confDcBusVoltageRef", Group: ConfigAcGroup, Cn: "直流母线电压参考", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网过压一级保护时间（0x0D）
	confGridOvervolt1TimeH_L = &c_base.Meta{Name: "confGridOvervolt1Time", Group: ConfigAcGroup, Cn: "电网过压一级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt2TimeH_L = &c_base.Meta{Name: "confGridOvervolt2Time", Group: ConfigAcGroup, Cn: "电网过压二级保护时间", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt3TimeH_L = &c_base.Meta{Name: "confGridOvervolt3Time", Group: ConfigAcGroup, Cn: "电网过压三级保护时间", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt4TimeH_L = &c_base.Meta{Name: "confGridOvervolt4Time", Group: ConfigAcGroup, Cn: "电网过压四级保护时间", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网过压五级保护时间（0x0E）
	confGridOvervolt5TimeH_L      = &c_base.Meta{Name: "confGridOvervolt5Time", Group: ConfigAcGroup, Cn: "电网过压五级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridRecoverConfirmTimeH_L = &c_base.Meta{Name: "confGridRecoverConfirmTime", Group: ConfigAcGroup, Cn: "电网恢复确认时间", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt1TimeH_L     = &c_base.Meta{Name: "confGridUndervolt1Time", Group: ConfigAcGroup, Cn: "电网欠压一级保护时间", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt2TimeH_L     = &c_base.Meta{Name: "confGridUndervolt2Time", Group: ConfigAcGroup, Cn: "电网欠压二级保护时间", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠压三级保护时间（0x0F）
	confGridUndervolt3TimeH_L = &c_base.Meta{Name: "confGridUndervolt3Time", Group: ConfigAcGroup, Cn: "电网欠压三级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt4TimeH_L = &c_base.Meta{Name: "confGridUndervolt4Time", Group: ConfigAcGroup, Cn: "电网欠压四级保护时间", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt5TimeH_L = &c_base.Meta{Name: "confGridUndervolt5Time", Group: ConfigAcGroup, Cn: "电网欠压五级保护时间", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq1TimeH_L  = &c_base.Meta{Name: "confGridOverfreq1Time", Group: ConfigAcGroup, Cn: "电网过频一级保护时间", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网过频二级保护时间（0x10）
	confGridOverfreq2TimeH_L = &c_base.Meta{Name: "confGridOverfreq2Time", Group: ConfigAcGroup, Cn: "电网过频二级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq3TimeH_L = &c_base.Meta{Name: "confGridOverfreq3Time", Group: ConfigAcGroup, Cn: "电网过频三级保护时间", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq4TimeH_L = &c_base.Meta{Name: "confGridOverfreq4Time", Group: ConfigAcGroup, Cn: "电网过频四级保护时间", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq5TimeH_L = &c_base.Meta{Name: "confGridOverfreq5Time", Group: ConfigAcGroup, Cn: "电网过频五级保护时间", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠频一级保护时间（0x11）
	confGridUnderfreq1TimeH_L = &c_base.Meta{Name: "confGridUnderfreq1Time", Group: ConfigAcGroup, Cn: "电网欠频一级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq2TimeH_L = &c_base.Meta{Name: "confGridUnderfreq2Time", Group: ConfigAcGroup, Cn: "电网欠频二级保护时间", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq3TimeH_L = &c_base.Meta{Name: "confGridUnderfreq3Time", Group: ConfigAcGroup, Cn: "电网欠频三级保护时间", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq4TimeH_L = &c_base.Meta{Name: "confGridUnderfreq4Time", Group: ConfigAcGroup, Cn: "电网欠频四级保护时间", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠频五级保护时间（0x12）
	confGridUnderfreq5TimeH_L   = &c_base.Meta{Name: "confGridUnderfreq5Time", Group: ConfigAcGroup, Cn: "电网欠频五级保护时间", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confHighPenetrationReactive = &c_base.Meta{Name: "confHighPenetrationReactive", Group: ConfigAcGroup, Cn: "高穿无功电流系数", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confHighPenetrationActive   = &c_base.Meta{Name: "confHighPenetrationActive", Group: ConfigAcGroup, Cn: "高穿有功系数", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confLowPenetrationReactive  = &c_base.Meta{Name: "confLowPenetrationReactive", Group: ConfigAcGroup, Cn: "低穿无功电流系数", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 低穿参数和母线电压（0x13）
	confLowPenetrationActiveCurrent = &c_base.Meta{Name: "confLowPenetrationActiveCurrent", Group: ConfigAcGroup, Cn: "低穿有功电流系数", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confLowPenetrationActiveRecover = &c_base.Meta{Name: "confLowPenetrationActiveRecover", Group: ConfigAcGroup, Cn: "低穿有功恢复速率", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confDcBusOvervoltPoint          = &c_base.Meta{Name: "confDcBusOvervoltPoint", Group: ConfigAcGroup, Cn: "直流母线过压点", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confAuthorizedCapacity          = &c_base.Meta{Name: "confAuthorizedCapacity", Group: ConfigAcGroup, Cn: "授权容量", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 授权容量和预留（0x14）
	confInertiaTimeConstant       = &c_base.Meta{Name: "confInertiaTimeConstant", Group: ConfigAcGroup, Cn: "惯性时间常数", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "s"}
	confDampingCoefficient        = &c_base.Meta{Name: "confDampingCoefficient", Group: ConfigAcGroup, Cn: "阻尼系数", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confReserved3                 = &c_base.Meta{Name: "confReserved3", Group: ConfigAcGroup, Cn: "预留", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confActiveFreqRegulationCoeff = &c_base.Meta{Name: "confActiveFreqRegulationCoeff", Group: ConfigAcGroup, Cn: "有功调频系数", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 有功调频参数（0x15）
	confFreqRegulationDeadZone      = &c_base.Meta{Name: "confFreqRegulationDeadZone", Group: ConfigAcGroup, Cn: "调频频率死区", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confFreqRegulationActiveUpper   = &c_base.Meta{Name: "confFreqRegulationActiveUpper", Group: ConfigAcGroup, Cn: "调频有功上限", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confFreqRegulationActiveLower   = &c_base.Meta{Name: "confFreqRegulationActiveLower", Group: ConfigAcGroup, Cn: "调频有功下限", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confReactiveVoltRegulationCoeff = &c_base.Meta{Name: "confReactiveVoltRegulationCoeff", Group: ConfigAcGroup, Cn: "无功调压系数", Addr: 6, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 无功调压参数（0x16）
	confVoltRegulationDeadZone      = &c_base.Meta{Name: "confVoltRegulationDeadZone", Group: ConfigAcGroup, Cn: "调压电压死区", Addr: 0, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confVoltRegulationReactiveUpper = &c_base.Meta{Name: "confVoltRegulationReactiveUpper", Group: ConfigAcGroup, Cn: "调压无功上限", Addr: 2, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confVoltRegulationReactiveLower = &c_base.Meta{Name: "confVoltRegulationReactiveLower", Group: ConfigAcGroup, Cn: "调压无功下限", Addr: 4, Endianness: c_base.EMiddleEndian, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
)
