package pcs_elecod_mac_v1

import "common/c_base"

var (
	// 总功率参数（0x01）
	confTotalActivePower   = &c_base.Meta{Name: "confTotalActivePower", Cn: "总有功功率", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confTotalReactivePower = &c_base.Meta{Name: "confTotalReactivePower", Cn: "总无功功率", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confTotalPowerFactor   = &c_base.Meta{Name: "confTotalPowerFactor", Cn: "总功率因数", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confRatedPower         = &c_base.Meta{Name: "confRatedPower", Cn: "额定功率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 相功率参数（0x02）
	confActivePowerA      = &c_base.Meta{Name: "confActivePowerA", Cn: "A相有功功率", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confActivePowerB      = &c_base.Meta{Name: "confActivePowerB", Cn: "B相有功功率", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confActivePowerC      = &c_base.Meta{Name: "confActivePowerC", Cn: "C相有功功率", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confRatedPhaseVoltage = &c_base.Meta{Name: "confRatedPhaseVoltage", Cn: "额定相电压", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 相无功功率和频率（0x03）
	confReactivePowerA = &c_base.Meta{Name: "confReactivePowerA", Cn: "A相无功功率", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confReactivePowerB = &c_base.Meta{Name: "confReactivePowerB", Cn: "B相无功功率", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confReactivePowerC = &c_base.Meta{Name: "confReactivePowerC", Cn: "C相无功功率", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confRatedFrequency = &c_base.Meta{Name: "confRatedFrequency", Cn: "额定频率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 相功率因数和有功变化率（0x04）
	confPowerFactorA          = &c_base.Meta{Name: "confPowerFactorA", Cn: "A相功率因数", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confPowerFactorB          = &c_base.Meta{Name: "confPowerFactorB", Cn: "B相功率因数", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confPowerFactorC          = &c_base.Meta{Name: "confPowerFactorC", Cn: "C相功率因数", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confActivePowerChangeRate = &c_base.Meta{Name: "confActivePowerChangeRate", Cn: "有功变化速率", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 控制位和电池类型（0x05）
	confControlBit1 = &c_base.Meta{Name: "confControlBit1", Cn: "控制位1", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confControlBit2 = &c_base.Meta{Name: "confControlBit2", Cn: "控制位2", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confControlBit3 = &c_base.Meta{Name: "confControlBit3", Cn: "控制位3", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confBatteryType = &c_base.Meta{Name: "confBatteryType", Cn: "电池类型", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}

	// 电池保护点（0x06）
	confBatteryUndervoltProtect = &c_base.Meta{Name: "confBatteryUndervoltProtect", Cn: "电池欠压保护点", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryUndervoltRecover = &c_base.Meta{Name: "confBatteryUndervoltRecover", Cn: "电池欠压恢复点", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryOvervoltProtect  = &c_base.Meta{Name: "confBatteryOvervoltProtect", Cn: "电池过压保护点", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryOvervoltRecover  = &c_base.Meta{Name: "confBatteryOvervoltRecover", Cn: "电池过压恢复点", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电池限流和电压（0x07）
	confBatteryChargeLimit     = &c_base.Meta{Name: "confBatteryChargeLimit", Cn: "电池充电限流", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	confBatteryDischargeLimit  = &c_base.Meta{Name: "confBatteryDischargeLimit", Cn: "电池放电限流", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "A"}
	confBatteryFloatVoltage    = &c_base.Meta{Name: "confBatteryFloatVoltage", Cn: "电池浮充电压", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confBatteryEqualizeVoltage = &c_base.Meta{Name: "confBatteryEqualizeVoltage", Cn: "电池均充电压", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网过压保护（0x08）
	confInsulationImpedanceThreshold = &c_base.Meta{Name: "confInsulationImpedanceThreshold", Cn: "绝缘阻抗保护阈值", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: "kΩ"}
	confGridOvervoltLevel1_2         = &c_base.Meta{Name: "confGridOvervoltLevel1_2", Cn: "电网过压一级：电网过压二级", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOvervoltLevel3_4         = &c_base.Meta{Name: "confGridOvervoltLevel3_4", Cn: "电网过压三级：电网过压四级", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOvervoltLevel5_Recover   = &c_base.Meta{Name: "confGridOvervoltLevel5_Recover", Cn: "电网过压五级：电网过压恢复", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}

	// 电网欠压和过频保护（0x09）
	confGridUndervoltLevel1_2       = &c_base.Meta{Name: "confGridUndervoltLevel1_2", Cn: "电网欠压一级：电网欠压二级", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridUndervoltLevel3_4       = &c_base.Meta{Name: "confGridUndervoltLevel3_4", Cn: "电网欠压三级：电网欠压四级", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridUndervoltLevel5_Recover = &c_base.Meta{Name: "confGridUndervoltLevel5_Recover", Cn: "电网欠压五级：电网欠压恢复", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: ""}
	confGridOverfreqLevel1          = &c_base.Meta{Name: "confGridOverfreqLevel1", Cn: "电网过频一级", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网过频保护（0x0A）
	confGridOverfreqLevel2 = &c_base.Meta{Name: "confGridOverfreqLevel2", Cn: "电网过频二级", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel3 = &c_base.Meta{Name: "confGridOverfreqLevel3", Cn: "电网过频三级", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel4 = &c_base.Meta{Name: "confGridOverfreqLevel4", Cn: "电网过频四级", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridOverfreqLevel5 = &c_base.Meta{Name: "confGridOverfreqLevel5", Cn: "电网过频五级", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网过频和欠频恢复（0x0B）
	confGridOverfreqRecover = &c_base.Meta{Name: "confGridOverfreqRecover", Cn: "电网过频恢复", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel1 = &c_base.Meta{Name: "confGridUnderfreqLevel1", Cn: "电网欠频一级", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel2 = &c_base.Meta{Name: "confGridUnderfreqLevel2", Cn: "电网欠频二级", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel3 = &c_base.Meta{Name: "confGridUnderfreqLevel3", Cn: "电网欠频三级", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}

	// 电网欠频恢复和直流母线（0x0C）
	confGridUnderfreqLevel4  = &c_base.Meta{Name: "confGridUnderfreqLevel4", Cn: "电网欠频四级", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqLevel5  = &c_base.Meta{Name: "confGridUnderfreqLevel5", Cn: "电网欠频五级", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confGridUnderfreqRecover = &c_base.Meta{Name: "confGridUnderfreqRecover", Cn: "电网欠频恢复", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confDcBusVoltageRef      = &c_base.Meta{Name: "confDcBusVoltageRef", Cn: "直流母线电压参考", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}

	// 电网过压一级保护时间（0x0D）
	confGridOvervolt1TimeH_L = &c_base.Meta{Name: "confGridOvervolt1Time", Cn: "电网过压一级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt2TimeH_L = &c_base.Meta{Name: "confGridOvervolt2Time", Cn: "电网过压二级保护时间", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt3TimeH_L = &c_base.Meta{Name: "confGridOvervolt3Time", Cn: "电网过压三级保护时间", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOvervolt4TimeH_L = &c_base.Meta{Name: "confGridOvervolt4Time", Cn: "电网过压四级保护时间", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网过压五级保护时间（0x0E）
	confGridOvervolt5TimeH_L      = &c_base.Meta{Name: "confGridOvervolt5Time", Cn: "电网过压五级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridRecoverConfirmTimeH_L = &c_base.Meta{Name: "confGridRecoverConfirmTime", Cn: "电网恢复确认时间", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt1TimeH_L     = &c_base.Meta{Name: "confGridUndervolt1Time", Cn: "电网欠压一级保护时间", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt2TimeH_L     = &c_base.Meta{Name: "confGridUndervolt2Time", Cn: "电网欠压二级保护时间", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠压三级保护时间（0x0F）
	confGridUndervolt3TimeH_L = &c_base.Meta{Name: "confGridUndervolt3Time", Cn: "电网欠压三级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt4TimeH_L = &c_base.Meta{Name: "confGridUndervolt4Time", Cn: "电网欠压四级保护时间", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUndervolt5TimeH_L = &c_base.Meta{Name: "confGridUndervolt5Time", Cn: "电网欠压五级保护时间", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq1TimeH_L  = &c_base.Meta{Name: "confGridOverfreq1Time", Cn: "电网过频一级保护时间", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网过频二级保护时间（0x10）
	confGridOverfreq2TimeH_L = &c_base.Meta{Name: "confGridOverfreq2Time", Cn: "电网过频二级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq3TimeH_L = &c_base.Meta{Name: "confGridOverfreq3Time", Cn: "电网过频三级保护时间", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq4TimeH_L = &c_base.Meta{Name: "confGridOverfreq4Time", Cn: "电网过频四级保护时间", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridOverfreq5TimeH_L = &c_base.Meta{Name: "confGridOverfreq5Time", Cn: "电网过频五级保护时间", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠频一级保护时间（0x11）
	confGridUnderfreq1TimeH_L = &c_base.Meta{Name: "confGridUnderfreq1Time", Cn: "电网欠频一级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq2TimeH_L = &c_base.Meta{Name: "confGridUnderfreq2Time", Cn: "电网欠频二级保护时间", Addr: 1, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq3TimeH_L = &c_base.Meta{Name: "confGridUnderfreq3Time", Cn: "电网欠频三级保护时间", Addr: 2, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confGridUnderfreq4TimeH_L = &c_base.Meta{Name: "confGridUnderfreq4Time", Cn: "电网欠频四级保护时间", Addr: 3, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}

	// 电网欠频五级保护时间（0x12）
	confGridUnderfreq5TimeH_L   = &c_base.Meta{Name: "confGridUnderfreq5Time", Cn: "电网欠频五级保护时间", Addr: 0, ReadType: c_base.RInt32, SystemType: c_base.SInt32, Factor: 1, Unit: "ms"}
	confHighPenetrationReactive = &c_base.Meta{Name: "confHighPenetrationReactive", Cn: "高穿无功电流系数", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confHighPenetrationActive   = &c_base.Meta{Name: "confHighPenetrationActive", Cn: "高穿有功系数", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confLowPenetrationReactive  = &c_base.Meta{Name: "confLowPenetrationReactive", Cn: "低穿无功电流系数", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 低穿参数和母线电压（0x13）
	confLowPenetrationActiveCurrent = &c_base.Meta{Name: "confLowPenetrationActiveCurrent", Cn: "低穿有功电流系数", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confLowPenetrationActiveRecover = &c_base.Meta{Name: "confLowPenetrationActiveRecover", Cn: "低穿有功恢复速率", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confDcBusOvervoltPoint          = &c_base.Meta{Name: "confDcBusOvervoltPoint", Cn: "直流母线过压点", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confAuthorizedCapacity          = &c_base.Meta{Name: "confAuthorizedCapacity", Cn: "授权容量", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}

	// 授权容量和预留（0x14）
	confInertiaTimeConstant       = &c_base.Meta{Name: "confInertiaTimeConstant", Cn: "惯性时间常数", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "s"}
	confDampingCoefficient        = &c_base.Meta{Name: "confDampingCoefficient", Cn: "阻尼系数", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}
	confReserved3                 = &c_base.Meta{Name: "confReserved3", Cn: "预留", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 1, Unit: ""}
	confActiveFreqRegulationCoeff = &c_base.Meta{Name: "confActiveFreqRegulationCoeff", Cn: "有功调频系数", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 有功调频参数（0x15）
	confFreqRegulationDeadZone      = &c_base.Meta{Name: "confFreqRegulationDeadZone", Cn: "调频频率死区", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: "Hz"}
	confFreqRegulationActiveUpper   = &c_base.Meta{Name: "confFreqRegulationActiveUpper", Cn: "调频有功上限", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confFreqRegulationActiveLower   = &c_base.Meta{Name: "confFreqRegulationActiveLower", Cn: "调频有功下限", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kW"}
	confReactiveVoltRegulationCoeff = &c_base.Meta{Name: "confReactiveVoltRegulationCoeff", Cn: "无功调压系数", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.01, Unit: ""}

	// 无功调压参数（0x16）
	confVoltRegulationDeadZone      = &c_base.Meta{Name: "confVoltRegulationDeadZone", Cn: "调压电压死区", Addr: 0, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "V"}
	confVoltRegulationReactiveUpper = &c_base.Meta{Name: "confVoltRegulationReactiveUpper", Cn: "调压无功上限", Addr: 1, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
	confVoltRegulationReactiveLower = &c_base.Meta{Name: "confVoltRegulationReactiveLower", Cn: "调压无功下限", Addr: 2, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Factor: 0.1, Unit: "kvar"}
)
