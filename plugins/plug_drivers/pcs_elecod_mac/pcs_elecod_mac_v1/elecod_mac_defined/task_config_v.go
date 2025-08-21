package elecod_mac_defined

import (
	"canbus/p_canbus"
	"common/c_base"
	elecod_canbus "pcs_elecod/elecod_canbus"
)

var (
	ConfigAllTasks = []*p_canbus.SCanbusTask{
		&configTotalPowerInfo,
		&configPhasePowerInfo,
		&configPhaseReactiveFreqInfo,
		&configPhasePowerFactorChangeRateInfo,
		&configControlBatteryTypeInfo,
		&configBatteryProtectInfo,
		&configBatteryLimitVoltageInfo,
		&configGridOvervoltProtectInfo,
		&configGridUndervoltOverfreqInfo,
		&configGridOverfreqProtectInfo,
		&configGridOverfreqUnderfreqInfo,
		&configGridUnderfreqDcBusInfo,
		&configGridOvervoltTime1Info,
		&configGridOvervoltTime2Info,
		&configGridUndervoltTime1Info,
		&configGridOverfreqTime1Info,
		&configGridUnderfreqTime1Info,
		&configGridUnderfreqTime2Info,
		&configLowPenetrationInfo,
		&configCapacityInertiaDampingInfo,
		&configActiveFreqRegulationInfo,
		&configReactiveVoltRegulationInfo,
	}
)

const (
	configMessageType = elecod_canbus.MessageTypeConfig
)

// isConfig returns true if the SCANFrameInfo is a config frame
func isConfig(info elecod_canbus.SCANFrameInfo) bool {
	return (info.SourceDeviceType == elecod_canbus.DeviceTypeMAC || info.SourceDeviceType == elecod_canbus.DeviceTypeScreen) &&
		(info.TargetDeviceType == elecod_canbus.DeviceTypeMAC || info.TargetDeviceType == elecod_canbus.DeviceTypeScreen) &&
		info.MessageType == elecod_canbus.MessageTypeConfig
}

var (
	configTotalPowerInfo = p_canbus.SCanbusTask{
		Name: "总功率参数",
		Metas: []*c_base.Meta{
			confTotalActivePower, confTotalReactivePower, confTotalPowerFactor, confRatedPower,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x01, params)
		},
	}

	configPhasePowerInfo = p_canbus.SCanbusTask{
		Name: "相功率参数",
		Metas: []*c_base.Meta{
			confActivePowerA, confActivePowerB, confActivePowerC, confRatedPhaseVoltage,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x02, params)
		},
	}

	configPhaseReactiveFreqInfo = p_canbus.SCanbusTask{
		Name: "相无功功率和频率",
		Metas: []*c_base.Meta{
			confReactivePowerA, confReactivePowerB, confReactivePowerC, confRatedFrequency,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x03, params)
		},
	}

	configPhasePowerFactorChangeRateInfo = p_canbus.SCanbusTask{
		Name: "相功率因数和有功变化率",
		Metas: []*c_base.Meta{
			confPowerFactorA, confPowerFactorB, confPowerFactorC, confActivePowerChangeRate,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x04, params)
		},
	}

	configControlBatteryTypeInfo = p_canbus.SCanbusTask{
		Name: "控制位和电池类型",
		Metas: []*c_base.Meta{
			confControlBit1, confControlBit2, confControlBit3, confBatteryType,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x05, params)
		},
	}

	configBatteryProtectInfo = p_canbus.SCanbusTask{
		Name: "电池保护点",
		Metas: []*c_base.Meta{
			confBatteryUndervoltProtect, confBatteryUndervoltRecover, confBatteryOvervoltProtect, confBatteryOvervoltRecover,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x06, params)
		},
	}

	configBatteryLimitVoltageInfo = p_canbus.SCanbusTask{
		Name: "电池限流和电压",
		Metas: []*c_base.Meta{
			confBatteryChargeLimit, confBatteryDischargeLimit, confBatteryFloatVoltage, confBatteryEqualizeVoltage,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x07, params)
		},
	}

	configGridOvervoltProtectInfo = p_canbus.SCanbusTask{
		Name: "电网过压保护",
		Metas: []*c_base.Meta{
			confInsulationImpedanceThreshold, confGridOvervoltLevel1_2, confGridOvervoltLevel3_4, confGridOvervoltLevel5_Recover,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x08, params)
		},
	}

	configGridUndervoltOverfreqInfo = p_canbus.SCanbusTask{
		Name: "电网欠压和过频保护",
		Metas: []*c_base.Meta{
			confGridUndervoltLevel1_2, confGridUndervoltLevel3_4, confGridUndervoltLevel5_Recover, confGridOverfreqLevel1,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x09, params)
		},
	}

	configGridOverfreqProtectInfo = p_canbus.SCanbusTask{
		Name: "电网过频保护",
		Metas: []*c_base.Meta{
			confGridOverfreqLevel2, confGridOverfreqLevel3, confGridOverfreqLevel4, confGridOverfreqLevel5,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0A, params)
		},
	}

	configGridOverfreqUnderfreqInfo = p_canbus.SCanbusTask{
		Name: "电网过频和欠频恢复",
		Metas: []*c_base.Meta{
			confGridOverfreqRecover, confGridUnderfreqLevel1, confGridUnderfreqLevel2, confGridUnderfreqLevel3,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0B, params)
		},
	}

	configGridUnderfreqDcBusInfo = p_canbus.SCanbusTask{
		Name: "电网欠频恢复和直流母线",
		Metas: []*c_base.Meta{
			confGridUnderfreqLevel4, confGridUnderfreqLevel5, confGridUnderfreqRecover, confDcBusVoltageRef,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0C, params)
		},
	}

	configGridOvervoltTime1Info = p_canbus.SCanbusTask{
		Name: "电网过压一级保护时间",
		Metas: []*c_base.Meta{
			confGridOvervolt1TimeH_L, confGridOvervolt2TimeH_L, confGridOvervolt3TimeH_L, confGridOvervolt4TimeH_L,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0D, params)
		},
	}

	configGridOvervoltTime2Info = p_canbus.SCanbusTask{
		Name: "电网过压五级保护时间",
		Metas: []*c_base.Meta{
			confGridOvervolt5TimeH_L, confGridRecoverConfirmTimeH_L, confGridUndervolt1TimeH_L, confGridUndervolt2TimeH_L,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0E, params)
		},
	}

	configGridUndervoltTime1Info = p_canbus.SCanbusTask{
		Name: "电网欠压三级保护时间",
		Metas: []*c_base.Meta{
			confGridUndervolt3TimeH_L, confGridUndervolt4TimeH_L, confGridUndervolt5TimeH_L, confGridOverfreq1TimeH_L,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0F, params)
		},
	}

	configGridOverfreqTime1Info = p_canbus.SCanbusTask{
		Name: "电网过频二级保护时间",
		Metas: []*c_base.Meta{
			confGridOverfreq2TimeH_L, confGridOverfreq3TimeH_L, confGridOverfreq4TimeH_L, confGridOverfreq5TimeH_L,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x10, params)
		},
	}

	configGridUnderfreqTime1Info = p_canbus.SCanbusTask{
		Name: "电网欠频一级保护时间",
		Metas: []*c_base.Meta{
			confGridUnderfreq1TimeH_L, confGridUnderfreq2TimeH_L, confGridUnderfreq3TimeH_L, confGridUnderfreq4TimeH_L,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x11, params)
		},
	}

	configGridUnderfreqTime2Info = p_canbus.SCanbusTask{
		Name: "电网欠频五级保护时间",
		Metas: []*c_base.Meta{
			confGridUnderfreq5TimeH_L, confHighPenetrationReactive, confHighPenetrationActive, confLowPenetrationReactive,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x12, params)
		},
	}

	configLowPenetrationInfo = p_canbus.SCanbusTask{
		Name: "低穿参数和母线电压",
		Metas: []*c_base.Meta{
			confLowPenetrationActiveCurrent, confLowPenetrationActiveRecover, confDcBusOvervoltPoint, confAuthorizedCapacity,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x13, params)
		},
	}

	configCapacityInertiaDampingInfo = p_canbus.SCanbusTask{
		Name: "授权容量和预留",
		Metas: []*c_base.Meta{
			confInertiaTimeConstant, confDampingCoefficient, confActiveFreqRegulationCoeff,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x14, params)
		},
	}

	configActiveFreqRegulationInfo = p_canbus.SCanbusTask{
		Name: "有功调频参数",
		Metas: []*c_base.Meta{
			confFreqRegulationDeadZone, confFreqRegulationActiveUpper, confFreqRegulationActiveLower, confReactiveVoltRegulationCoeff,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x15, params)
		},
	}

	configReactiveVoltRegulationInfo = p_canbus.SCanbusTask{
		Name: "无功调压参数",
		Metas: []*c_base.Meta{
			confVoltRegulationDeadZone, confVoltRegulationReactiveUpper, confVoltRegulationReactiveLower,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x16, params)
		},
	}
)
