package elecod_mac_defined

import (
	"common/c_base"
	"common/c_default"
	"common/c_proto"
)

var (
	ConfigMainGroup   = &c_base.SPointGroup{GroupName: "主要配置", GroupSort: 25}
	ConfigBatterGroup = &c_base.SPointGroup{GroupName: "电池配置", GroupSort: 29}
	ConfigAcGroup     = &c_base.SPointGroup{GroupName: "交流参数配置", GroupSort: 28}
)

var (
	// 总功率参数（0x01）
	confTotalActivePower = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointP,
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigMainGroup,
	}
	confTotalReactivePower = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointQ,
			DataAccess: c_default.VDataAccessFloat32Byte1Scale01,
		},
		Group: ConfigMainGroup,
	}
	confTotalPowerFactor = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "TotalPowerFactor", Name: "总功率因数", Unit: "", Desc: "总功率因数", Sort: 3, Precise: 2, Group: ConfigMainGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigMainGroup,
	}
	confRatedPower = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "RatedPower", Name: "额定功率", Unit: "kW", Desc: "额定功率", Sort: 4, Precise: 2, Group: ConfigMainGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale01,
		},
		Group: ConfigMainGroup,
	}

	// 相功率参数（0x02）
	confActivePowerA = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPa,
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigMainGroup,
	}
	confActivePowerB = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPb,
			DataAccess: c_default.VDataAccessFloat32Byte1Scale01,
		},
		Group: ConfigMainGroup,
	}
	confActivePowerC = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPc,
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigMainGroup,
	}
	confRatedPhaseVoltage = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "RatedPhaseVoltage", Name: "额定相电压", Unit: "V", Desc: "额定相电压", Sort: 8, Precise: 1, Group: ConfigMainGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale01,
		},
		Group: ConfigMainGroup,
	}

	// 相无功功率和频率（0x03）
	confReactivePowerA = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointQa,
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigMainGroup,
	}
	confReactivePowerB = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointQb,
			DataAccess: c_default.VDataAccessFloat32Byte1Scale01,
		},
		Group: ConfigMainGroup,
	}
	confReactivePowerC = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointQc,
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigMainGroup,
	}
	confRatedFrequency = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointFreq,
			DataAccess: c_default.VDataAccessFloat32Byte3Scale001,
		},
		Group: ConfigMainGroup,
	}

	// 相功率因数和有功变化率（0x04）
	confPowerFactorA = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPFa,
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigMainGroup,
	}
	confPowerFactorB = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPFb,
			DataAccess: c_default.VDataAccessFloat32Byte1Scale001,
		},
		Group: ConfigMainGroup,
	}
	confPowerFactorC = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointPFc,
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigMainGroup,
	}
	confActivePowerChangeRate = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ActivePowerChangeRate", Name: "有功变化速率", Unit: "", Desc: "有功变化速率", Sort: 16, Precise: 0, Group: ConfigMainGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale1,
		},
		Group: ConfigMainGroup,
	}

	// 控制位和电池类型（0x05）
	confControlBit1 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ControlBit1", Name: "控制位1", Unit: "", Desc: "控制位1", Sort: 17, Precise: 0, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigBatterGroup,
	}
	confControlBit2 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ControlBit2", Name: "控制位2", Unit: "", Desc: "控制位2", Sort: 18, Precise: 0, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte1Scale1,
		},
		Group: ConfigBatterGroup,
	}
	confControlBit3 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ControlBit3", Name: "控制位3", Unit: "", Desc: "控制位3", Sort: 19, Precise: 0, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryType = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryType", Name: "电池类型", Unit: "", Desc: "电池类型", Sort: 20, Precise: 0, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale1,
		},
		Group: ConfigBatterGroup,
	}

	// 电池保护点（0x06）- 使用预定义点位
	confBatteryUndervoltProtect = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryUndervoltProtect", Name: "电池欠压保护点", Unit: "V", Desc: "电池欠压保护点", Sort: 21, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryUndervoltRecover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryUndervoltRecover", Name: "电池欠压恢复点", Unit: "V", Desc: "电池欠压恢复点", Sort: 22, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryOvervoltProtect = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryOvervoltProtect", Name: "电池过压保护点", Unit: "V", Desc: "电池过压保护点", Sort: 23, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryOvervoltRecover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryOvervoltRecover", Name: "电池过压恢复点", Unit: "V", Desc: "电池过压恢复点", Sort: 24, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale01,
		},
		Group: ConfigBatterGroup,
	}

	// 电池限流和电压（0x07）
	confBatteryChargeLimit = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryChargeLimit", Name: "电池充电限流", Unit: "A", Desc: "电池充电限流", Sort: 25, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryDischargeLimit = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryDischargeLimit", Name: "电池放电限流", Unit: "A", Desc: "电池放电限流", Sort: 26, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryFloatVoltage = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryFloatVoltage", Name: "电池浮充电压", Unit: "V", Desc: "电池浮充电压", Sort: 27, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigBatterGroup,
	}
	confBatteryEqualizeVoltage = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "BatteryEqualizeVoltage", Name: "电池均充电压", Unit: "V", Desc: "电池均充电压", Sort: 28, Precise: 1, Group: ConfigBatterGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale01,
		},
		Group: ConfigBatterGroup,
	}

	// 电网过压保护（0x08）
	confInsulationImpedanceThreshold = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "InsulationImpedanceThreshold", Name: "绝缘阻抗保护阈值", Unit: "kΩ", Desc: "绝缘阻抗保护阈值", Sort: 29, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervoltLevel1_2 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervoltLevel1_2", Name: "电网过压一级：电网过压二级", Unit: "", Desc: "电网过压一级：电网过压二级", Sort: 30, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte1Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervoltLevel3_4 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervoltLevel3_4", Name: "电网过压三级：电网过压四级", Unit: "", Desc: "电网过压三级：电网过压四级", Sort: 31, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervoltLevel5_Recover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervoltLevel5_Recover", Name: "电网过压五级：电网过压恢复", Unit: "", Desc: "电网过压五级：电网过压恢复", Sort: 32, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网欠压和过频保护（0x09）
	confGridUndervoltLevel1_2 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervoltLevel1_2", Name: "电网欠压一级：电网欠压二级", Unit: "", Desc: "电网欠压一级：电网欠压二级", Sort: 33, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervoltLevel3_4 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervoltLevel3_4", Name: "电网欠压三级：电网欠压四级", Unit: "", Desc: "电网欠压三级：电网欠压四级", Sort: 34, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte1Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervoltLevel5_Recover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervoltLevel5_Recover", Name: "电网欠压五级：电网欠压恢复", Unit: "", Desc: "电网欠压五级：电网欠压恢复", Sort: 35, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreqLevel1 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqLevel1", Name: "电网过频一级", Unit: "Hz", Desc: "电网过频一级保护点", Sort: 36, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte3Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 电网过频保护（0x0A）
	confGridOverfreqLevel2 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqLevel2", Name: "电网过频二级", Unit: "Hz", Desc: "电网过频二级保护点", Sort: 37, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreqLevel3 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqLevel3", Name: "电网过频三级", Unit: "Hz", Desc: "电网过频三级保护点", Sort: 38, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreqLevel4 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqLevel4", Name: "电网过频四级", Unit: "Hz", Desc: "电网过频四级保护点", Sort: 39, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreqLevel5 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqLevel5", Name: "电网过频五级", Unit: "Hz", Desc: "电网过频五级保护点", Sort: 40, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 电网过频和欠频恢复（0x0B）
	confGridOverfreqRecover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreqRecover", Name: "电网过频恢复", Unit: "Hz", Desc: "电网过频恢复点", Sort: 41, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreqLevel1 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqLevel1", Name: "电网欠频一级", Unit: "Hz", Desc: "电网欠频一级保护点", Sort: 42, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreqLevel2 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqLevel2", Name: "电网欠频二级", Unit: "Hz", Desc: "电网欠频二级保护点", Sort: 43, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreqLevel3 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqLevel3", Name: "电网欠频三级", Unit: "Hz", Desc: "电网欠频三级保护点", Sort: 44, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 电网欠频恢复和直流母线（0x0C）
	confGridUnderfreqLevel4 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqLevel4", Name: "电网欠频四级", Unit: "Hz", Desc: "电网欠频四级保护点", Sort: 45, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreqLevel5 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqLevel5", Name: "电网欠频五级", Unit: "Hz", Desc: "电网欠频五级保护点", Sort: 46, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreqRecover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreqRecover", Name: "电网欠频恢复", Unit: "Hz", Desc: "电网欠频恢复点", Sort: 47, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confDcBusVoltageRef = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "DcBusVoltageRef", Name: "直流母线电压参考", Unit: "V", Desc: "直流母线电压参考", Sort: 48, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigAcGroup,
	}

	// 电网过压一级保护时间（0x0D）
	confGridOvervolt1TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervolt1Time", Name: "电网过压一级保护时间", Unit: "ms", Desc: "电网过压一级保护时间", Sort: 49, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervolt2TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervolt2Time", Name: "电网过压二级保护时间", Unit: "ms", Desc: "电网过压二级保护时间", Sort: 50, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervolt3TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervolt3Time", Name: "电网过压三级保护时间", Unit: "ms", Desc: "电网过压三级保护时间", Sort: 51, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOvervolt4TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervolt4Time", Name: "电网过压四级保护时间", Unit: "ms", Desc: "电网过压四级保护时间", Sort: 52, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网过压五级保护时间（0x0E）
	confGridOvervolt5TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOvervolt5Time", Name: "电网过压五级保护时间", Unit: "ms", Desc: "电网过压五级保护时间", Sort: 53, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridRecoverConfirmTimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridRecoverConfirmTime", Name: "电网恢复确认时间", Unit: "ms", Desc: "电网恢复确认时间", Sort: 54, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervolt1TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervolt1Time", Name: "电网欠压一级保护时间", Unit: "ms", Desc: "电网欠压一级保护时间", Sort: 55, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervolt2TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervolt2Time", Name: "电网欠压二级保护时间", Unit: "ms", Desc: "电网欠压二级保护时间", Sort: 56, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网欠压三级保护时间（0x0F）
	confGridUndervolt3TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervolt3Time", Name: "电网欠压三级保护时间", Unit: "ms", Desc: "电网欠压三级保护时间", Sort: 57, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervolt4TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervolt4Time", Name: "电网欠压四级保护时间", Unit: "ms", Desc: "电网欠压四级保护时间", Sort: 58, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUndervolt5TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUndervolt5Time", Name: "电网欠压五级保护时间", Unit: "ms", Desc: "电网欠压五级保护时间", Sort: 59, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreq1TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreq1Time", Name: "电网过频一级保护时间", Unit: "ms", Desc: "电网过频一级保护时间", Sort: 60, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网过频二级保护时间（0x10）
	confGridOverfreq2TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreq2Time", Name: "电网过频二级保护时间", Unit: "ms", Desc: "电网过频二级保护时间", Sort: 61, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreq3TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreq3Time", Name: "电网过频三级保护时间", Unit: "ms", Desc: "电网过频三级保护时间", Sort: 62, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreq4TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreq4Time", Name: "电网过频四级保护时间", Unit: "ms", Desc: "电网过频四级保护时间", Sort: 63, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridOverfreq5TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridOverfreq5Time", Name: "电网过频五级保护时间", Unit: "ms", Desc: "电网过频五级保护时间", Sort: 64, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网欠频一级保护时间（0x11）
	confGridUnderfreq1TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreq1Time", Name: "电网欠频一级保护时间", Unit: "ms", Desc: "电网欠频一级保护时间", Sort: 65, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreq2TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreq2Time", Name: "电网欠频二级保护时间", Unit: "ms", Desc: "电网欠频二级保护时间", Sort: 66, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreq3TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreq3Time", Name: "电网欠频三级保护时间", Unit: "ms", Desc: "电网欠频三级保护时间", Sort: 67, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confGridUnderfreq4TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreq4Time", Name: "电网欠频四级保护时间", Unit: "ms", Desc: "电网欠频四级保护时间", Sort: 68, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale1,
		},
		Group: ConfigAcGroup,
	}

	// 电网欠频五级保护时间（0x12）
	confGridUnderfreq5TimeH_L = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "GridUnderfreq5Time", Name: "电网欠频五级保护时间", Unit: "ms", Desc: "电网欠频五级保护时间", Sort: 69, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confHighPenetrationReactive = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "HighPenetrationReactive", Name: "高穿无功电流系数", Unit: "", Desc: "高穿无功电流系数", Sort: 70, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confHighPenetrationActive = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "HighPenetrationActive", Name: "高穿有功系数", Unit: "", Desc: "高穿有功系数", Sort: 71, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confLowPenetrationReactive = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "LowPenetrationReactive", Name: "低穿无功电流系数", Unit: "", Desc: "低穿无功电流系数", Sort: 72, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 低穿参数和母线电压（0x13）
	confLowPenetrationActiveCurrent = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "LowPenetrationActiveCurrent", Name: "低穿有功电流系数", Unit: "", Desc: "低穿有功电流系数", Sort: 73, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confLowPenetrationActiveRecover = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "LowPenetrationActiveRecover", Name: "低穿有功恢复速率", Unit: "", Desc: "低穿有功恢复速率", Sort: 74, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confDcBusOvervoltPoint = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "DcBusOvervoltPoint", Name: "直流母线过压点", Unit: "V", Desc: "直流母线过压点", Sort: 75, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigAcGroup,
	}
	confAuthorizedCapacity = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "AuthorizedCapacity", Name: "授权容量", Unit: "kW", Desc: "授权容量", Sort: 76, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigAcGroup,
	}

	// 授权容量和预留（0x14）
	confInertiaTimeConstant = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "InertiaTimeConstant", Name: "惯性时间常数", Unit: "s", Desc: "惯性时间常数", Sort: 77, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confDampingCoefficient = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "DampingCoefficient", Name: "阻尼系数", Unit: "", Desc: "阻尼系数", Sort: 78, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}
	confReserved3 = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "Reserved3", Name: "预留", Unit: "", Desc: "预留", Sort: 79, Precise: 0, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale1,
		},
		Group: ConfigAcGroup,
	}
	confActiveFreqRegulationCoeff = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ActiveFreqRegulationCoeff", Name: "有功调频系数", Unit: "", Desc: "有功调频系数", Sort: 80, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 有功调频参数（0x15）
	confFreqRegulationDeadZone = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "FreqRegulationDeadZone", Name: "调频频率死区", Unit: "Hz", Desc: "调频频率死区", Sort: 81, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale001,
		},
		Group: ConfigAcGroup,
	}
	confFreqRegulationActiveUpper = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "FreqRegulationActiveUpper", Name: "调频有功上限", Unit: "kW", Desc: "调频有功上限", Sort: 82, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigAcGroup,
	}
	confFreqRegulationActiveLower = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "FreqRegulationActiveLower", Name: "调频有功下限", Unit: "kW", Desc: "调频有功下限", Sort: 83, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigAcGroup,
	}
	confReactiveVoltRegulationCoeff = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "ReactiveVoltRegulationCoeff", Name: "无功调压系数", Unit: "", Desc: "无功调压系数", Sort: 84, Precise: 2, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale001,
		},
		Group: ConfigAcGroup,
	}

	// 无功调压参数（0x16）
	confVoltRegulationDeadZone = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "VoltRegulationDeadZone", Name: "调压电压死区", Unit: "V", Desc: "调压电压死区", Sort: 85, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigAcGroup,
	}
	confVoltRegulationReactiveUpper = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "VoltRegulationReactiveUpper", Name: "调压无功上限", Unit: "kvar", Desc: "调压无功上限", Sort: 86, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte0Scale01,
		},
		Group: ConfigAcGroup,
	}
	confVoltRegulationReactiveLower = &c_proto.SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     &c_base.SPoint{Key: "VoltRegulationReactiveLower", Name: "调压无功下限", Unit: "kvar", Desc: "调压无功下限", Sort: 87, Precise: 1, Group: ConfigAcGroup},
			DataAccess: c_default.VDataAccessFloat32Byte2Scale01,
		},
		Group: ConfigAcGroup,
	}
)
