package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
)

var (
	configAllTasks = []*p_canbus.SCanbusTask{
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

// isConfig returns true if the CANFrameInfo is a config frame
func isConfig(info CANFrameInfo) bool {
	return (info.SourceDeviceType == DeviceTypeMAC || info.SourceDeviceType == DeviceTypeScreen) &&
		(info.TargetDeviceType == DeviceTypeMAC || info.TargetDeviceType == DeviceTypeScreen) &&
		info.MessageType == MessageTypeConfig
}

var (
	configTotalPowerInfo = p_canbus.SCanbusTask{
		Name: "总功率参数",
		Metas: []*c_base.Meta{
			confTotalActivePower, confTotalReactivePower, confTotalPowerFactor, confRatedPower,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x01
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configPhasePowerInfo = p_canbus.SCanbusTask{
		Name: "相功率参数",
		Metas: []*c_base.Meta{
			confActivePowerA, confActivePowerB, confActivePowerC, confRatedPhaseVoltage,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x02
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configPhaseReactiveFreqInfo = p_canbus.SCanbusTask{
		Name: "相无功功率和频率",
		Metas: []*c_base.Meta{
			confReactivePowerA, confReactivePowerB, confReactivePowerC, confRatedFrequency,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x03
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configPhasePowerFactorChangeRateInfo = p_canbus.SCanbusTask{
		Name: "相功率因数和有功变化率",
		Metas: []*c_base.Meta{
			confPowerFactorA, confPowerFactorB, confPowerFactorC, confActivePowerChangeRate,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x04
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configControlBatteryTypeInfo = p_canbus.SCanbusTask{
		Name: "控制位和电池类型",
		Metas: []*c_base.Meta{
			confControlBit1, confControlBit2, confControlBit3, confBatteryType,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x05
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configBatteryProtectInfo = p_canbus.SCanbusTask{
		Name: "电池保护点",
		Metas: []*c_base.Meta{
			confBatteryUndervoltProtect, confBatteryUndervoltRecover, confBatteryOvervoltProtect, confBatteryOvervoltRecover,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x06
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configBatteryLimitVoltageInfo = p_canbus.SCanbusTask{
		Name: "电池限流和电压",
		Metas: []*c_base.Meta{
			confBatteryChargeLimit, confBatteryDischargeLimit, confBatteryFloatVoltage, confBatteryEqualizeVoltage,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x07
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOvervoltProtectInfo = p_canbus.SCanbusTask{
		Name: "电网过压保护",
		Metas: []*c_base.Meta{
			confInsulationImpedanceThreshold, confGridOvervoltLevel1_2, confGridOvervoltLevel3_4, confGridOvervoltLevel5_Recover,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x08
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridUndervoltOverfreqInfo = p_canbus.SCanbusTask{
		Name: "电网欠压和过频保护",
		Metas: []*c_base.Meta{
			confGridUndervoltLevel1_2, confGridUndervoltLevel3_4, confGridUndervoltLevel5_Recover, confGridOverfreqLevel1,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x09
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOverfreqProtectInfo = p_canbus.SCanbusTask{
		Name: "电网过频保护",
		Metas: []*c_base.Meta{
			confGridOverfreqLevel2, confGridOverfreqLevel3, confGridOverfreqLevel4, confGridOverfreqLevel5,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0A
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOverfreqUnderfreqInfo = p_canbus.SCanbusTask{
		Name: "电网过频和欠频恢复",
		Metas: []*c_base.Meta{
			confGridOverfreqRecover, confGridUnderfreqLevel1, confGridUnderfreqLevel2, confGridUnderfreqLevel3,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0B
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridUnderfreqDcBusInfo = p_canbus.SCanbusTask{
		Name: "电网欠频恢复和直流母线",
		Metas: []*c_base.Meta{
			confGridUnderfreqLevel4, confGridUnderfreqLevel5, confGridUnderfreqRecover, confDcBusVoltageRef,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0C
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOvervoltTime1Info = p_canbus.SCanbusTask{
		Name: "电网过压一级保护时间",
		Metas: []*c_base.Meta{
			confGridOvervolt1TimeH_L, confGridOvervolt2TimeH_L, confGridOvervolt3TimeH_L, confGridOvervolt4TimeH_L,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0D
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOvervoltTime2Info = p_canbus.SCanbusTask{
		Name: "电网过压五级保护时间",
		Metas: []*c_base.Meta{
			confGridOvervolt5TimeH_L, confGridRecoverConfirmTimeH_L, confGridUndervolt1TimeH_L, confGridUndervolt2TimeH_L,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0E
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridUndervoltTime1Info = p_canbus.SCanbusTask{
		Name: "电网欠压三级保护时间",
		Metas: []*c_base.Meta{
			confGridUndervolt3TimeH_L, confGridUndervolt4TimeH_L, confGridUndervolt5TimeH_L, confGridOverfreq1TimeH_L,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x0F
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridOverfreqTime1Info = p_canbus.SCanbusTask{
		Name: "电网过频二级保护时间",
		Metas: []*c_base.Meta{
			confGridOverfreq2TimeH_L, confGridOverfreq3TimeH_L, confGridOverfreq4TimeH_L, confGridOverfreq5TimeH_L,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x10
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridUnderfreqTime1Info = p_canbus.SCanbusTask{
		Name: "电网欠频一级保护时间",
		Metas: []*c_base.Meta{
			confGridUnderfreq1TimeH_L, confGridUnderfreq2TimeH_L, confGridUnderfreq3TimeH_L, confGridUnderfreq4TimeH_L,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x11
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configGridUnderfreqTime2Info = p_canbus.SCanbusTask{
		Name: "电网欠频五级保护时间",
		Metas: []*c_base.Meta{
			confGridUnderfreq5TimeH_L, confHighPenetrationReactive, confHighPenetrationActive, confLowPenetrationReactive,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x12
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configLowPenetrationInfo = p_canbus.SCanbusTask{
		Name: "低穿参数和母线电压",
		Metas: []*c_base.Meta{
			confLowPenetrationActiveCurrent, confLowPenetrationActiveRecover, confDcBusOvervoltPoint, confAuthorizedCapacity,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x13
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configCapacityInertiaDampingInfo = p_canbus.SCanbusTask{
		Name: "授权容量和预留",
		Metas: []*c_base.Meta{
			confInertiaTimeConstant, confDampingCoefficient, confReserved3, confActiveFreqRegulationCoeff,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x14
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configActiveFreqRegulationInfo = p_canbus.SCanbusTask{
		Name: "有功调频参数",
		Metas: []*c_base.Meta{
			confFreqRegulationDeadZone, confFreqRegulationActiveUpper, confFreqRegulationActiveLower, confReactiveVoltRegulationCoeff,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x15
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	configReactiveVoltRegulationInfo = p_canbus.SCanbusTask{
		Name: "无功调压参数",
		Metas: []*c_base.Meta{
			confVoltRegulationDeadZone, confVoltRegulationReactiveUpper, confVoltRegulationReactiveLower,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := isConfig(info) && info.ServiceCode == 0x16
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}
)
