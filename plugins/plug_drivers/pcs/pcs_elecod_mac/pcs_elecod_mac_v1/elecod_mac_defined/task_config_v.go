package elecod_mac_defined

import (
	"common/c_proto"
	"pcs_elecod/elecod_canbus"
)

var (
	ConfigAllTasks = []*c_proto.SCanbusTask{
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
	configTotalPowerInfo = c_proto.SCanbusTask{
		Name: "总功率参数",
		Points: []*c_proto.SCanbusPoint{
			confTotalActivePower, confTotalReactivePower, confTotalPowerFactor, confRatedPower,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x01, params)
		},
	}

	configPhasePowerInfo = c_proto.SCanbusTask{
		Name: "相功率参数",
		Points: []*c_proto.SCanbusPoint{
			confActivePowerA, confActivePowerB, confActivePowerC, confRatedPhaseVoltage,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x02, params)
		},
	}

	configPhaseReactiveFreqInfo = c_proto.SCanbusTask{
		Name: "相无功功率和频率",
		Points: []*c_proto.SCanbusPoint{
			confReactivePowerA, confReactivePowerB, confReactivePowerC, confRatedFrequency,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x03, params)
		},
	}

	configPhasePowerFactorChangeRateInfo = c_proto.SCanbusTask{
		Name: "相功率因数和有功变化率",
		Points: []*c_proto.SCanbusPoint{
			confPowerFactorA, confPowerFactorB, confPowerFactorC, confActivePowerChangeRate,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x04, params)
		},
	}

	configControlBatteryTypeInfo = c_proto.SCanbusTask{
		Name: "控制位和电池类型",
		Points: []*c_proto.SCanbusPoint{
			confControlBit1, confControlBit2, confControlBit3, confBatteryType,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x05, params)
		},
	}

	configBatteryProtectInfo = c_proto.SCanbusTask{
		Name: "电池保护点",
		Points: []*c_proto.SCanbusPoint{
			confBatteryUndervoltProtect, confBatteryUndervoltRecover, confBatteryOvervoltProtect, confBatteryOvervoltRecover,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x06, params)
		},
	}

	configBatteryLimitVoltageInfo = c_proto.SCanbusTask{
		Name: "电池限流和电压",
		Points: []*c_proto.SCanbusPoint{
			confBatteryChargeLimit, confBatteryDischargeLimit, confBatteryFloatVoltage, confBatteryEqualizeVoltage,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x07, params)
		},
	}

	configGridOvervoltProtectInfo = c_proto.SCanbusTask{
		Name: "电网过压保护",
		Points: []*c_proto.SCanbusPoint{
			confInsulationImpedanceThreshold, confGridOvervoltLevel1_2, confGridOvervoltLevel3_4, confGridOvervoltLevel5_Recover,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x08, params)
		},
	}

	configGridUndervoltOverfreqInfo = c_proto.SCanbusTask{
		Name: "电网欠压和过频保护",
		Points: []*c_proto.SCanbusPoint{
			confGridUndervoltLevel1_2, confGridUndervoltLevel3_4, confGridUndervoltLevel5_Recover, confGridOverfreqLevel1,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x09, params)
		},
	}

	configGridOverfreqProtectInfo = c_proto.SCanbusTask{
		Name: "电网过频保护",
		Points: []*c_proto.SCanbusPoint{
			confGridOverfreqLevel2, confGridOverfreqLevel3, confGridOverfreqLevel4, confGridOverfreqLevel5,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0A, params)
		},
	}

	configGridOverfreqUnderfreqInfo = c_proto.SCanbusTask{
		Name: "电网过频和欠频恢复",
		Points: []*c_proto.SCanbusPoint{
			confGridOverfreqRecover, confGridUnderfreqLevel1, confGridUnderfreqLevel2, confGridUnderfreqLevel3,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0B, params)
		},
	}

	configGridUnderfreqDcBusInfo = c_proto.SCanbusTask{
		Name: "电网欠频恢复和直流母线",
		Points: []*c_proto.SCanbusPoint{
			confGridUnderfreqLevel4, confGridUnderfreqLevel5, confGridUnderfreqRecover, confDcBusVoltageRef,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0C, params)
		},
	}

	configGridOvervoltTime1Info = c_proto.SCanbusTask{
		Name: "电网过压一级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridOvervolt1TimeH_L, confGridOvervolt2TimeH_L, confGridOvervolt3TimeH_L, confGridOvervolt4TimeH_L,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0D, params)
		},
	}

	configGridOvervoltTime2Info = c_proto.SCanbusTask{
		Name: "电网过压五级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridOvervolt5TimeH_L, confGridRecoverConfirmTimeH_L, confGridUndervolt1TimeH_L, confGridUndervolt2TimeH_L,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0E, params)
		},
	}

	configGridUndervoltTime1Info = c_proto.SCanbusTask{
		Name: "电网欠压三级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridUndervolt3TimeH_L, confGridUndervolt4TimeH_L, confGridUndervolt5TimeH_L, confGridOverfreq1TimeH_L,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x0F, params)
		},
	}

	configGridOverfreqTime1Info = c_proto.SCanbusTask{
		Name: "电网过频二级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridOverfreq2TimeH_L, confGridOverfreq3TimeH_L, confGridOverfreq4TimeH_L, confGridOverfreq5TimeH_L,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x10, params)
		},
	}

	configGridUnderfreqTime1Info = c_proto.SCanbusTask{
		Name: "电网欠频一级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridUnderfreq1TimeH_L, confGridUnderfreq2TimeH_L, confGridUnderfreq3TimeH_L,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x11, params)
		},
	}

	configGridUnderfreqTime2Info = c_proto.SCanbusTask{
		Name: "电网欠频五级保护时间",
		Points: []*c_proto.SCanbusPoint{
			confGridUnderfreq5TimeH_L, confHighPenetrationReactive, confHighPenetrationActive, confLowPenetrationReactive,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x12, params)
		},
	}

	configLowPenetrationInfo = c_proto.SCanbusTask{
		Name: "低穿参数和母线电压",
		Points: []*c_proto.SCanbusPoint{
			confLowPenetrationActiveCurrent, confLowPenetrationActiveRecover, confDcBusOvervoltPoint, confAuthorizedCapacity,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x13, params)
		},
	}

	configCapacityInertiaDampingInfo = c_proto.SCanbusTask{
		Name: "授权容量和预留",
		Points: []*c_proto.SCanbusPoint{
			confInertiaTimeConstant, confDampingCoefficient, confActiveFreqRegulationCoeff,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x14, params)
		},
	}

	configActiveFreqRegulationInfo = c_proto.SCanbusTask{
		Name: "有功调频参数",
		Points: []*c_proto.SCanbusPoint{
			confFreqRegulationDeadZone, confFreqRegulationActiveUpper, confFreqRegulationActiveLower, confReactiveVoltRegulationCoeff,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x15, params)
		},
	}

	configReactiveVoltRegulationInfo = c_proto.SCanbusTask{
		Name: "无功调压参数",
		Points: []*c_proto.SCanbusPoint{
			confVoltRegulationDeadZone, confVoltRegulationReactiveUpper, confVoltRegulationReactiveLower,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(configMessageType, 0x16, params)
		},
	}
)
