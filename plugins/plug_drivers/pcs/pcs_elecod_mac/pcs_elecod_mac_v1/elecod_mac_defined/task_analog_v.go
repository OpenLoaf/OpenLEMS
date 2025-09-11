package elecod_mac_defined

import (
	"canbus/p_canbus"
	"common/c_base"
	elecod_canbus "pcs_elecod/elecod_canbus"
)

var (
	AnalogAllTasks = []*p_canbus.SCanbusTask{
		&analogDCInfo,
		&analogGridVoltageInfo,
		&analogGridCurrentInfo,
		&analogGridFrequencyInfo,
		&analogActivePowerInfo,
		&analogReactivePowerInfo,
	}
)

const (
	analogMessageType = elecod_canbus.MessageTypeAnalog
)

var (
	analogDCInfo = p_canbus.SCanbusTask{
		Name: "直流参数",
		Metas: []*c_base.Meta{
			analogDcVoltage, analogDcCurrent, analogDcPower, analogBusVoltage,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x01, params)
		},
	}

	analogGridVoltageInfo = p_canbus.SCanbusTask{
		Name: "电网电压",
		Metas: []*c_base.Meta{
			analogGridVoltageA, analogGridVoltageB, analogGridVoltageC, analogPowerTubeTemp,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x02, params)
		},
	}

	analogGridCurrentInfo = p_canbus.SCanbusTask{
		Name: "电网电流",
		Metas: []*c_base.Meta{
			analogGridCurrentA, analogGridCurrentB, analogGridCurrentC, analogBridgeTemp,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x03, params)
		},
	}

	analogGridFrequencyInfo = p_canbus.SCanbusTask{
		Name: "电网频率",
		Metas: []*c_base.Meta{
			analogGridFrequencyA, analogGridFrequencyB, analogGridFrequencyC, analogAmbientTemp,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x04, params)
		},
	}

	analogActivePowerInfo = p_canbus.SCanbusTask{
		Name: "有功功率",
		Metas: []*c_base.Meta{
			analogActivePowerA, analogActivePowerB, analogActivePowerC, AnalogTotalActivePower,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x05, params)
		},
	}

	analogReactivePowerInfo = p_canbus.SCanbusTask{
		Name: "无功功率",
		Metas: []*c_base.Meta{
			analogReactivePowerA, analogReactivePowerB, analogReactivePowerC, analogTotalReactivePower,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x06, params)
		},
	}

	analogPowerFactorInfo = p_canbus.SCanbusTask{
		Name: "功率因素",
		Metas: []*c_base.Meta{
			analogPowerFactorA, analogPowerFactorB, analogPowerFactorC, analogTotalPowerFactor,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x07, params)
		},
	}

	analogApparentPowerInfo = p_canbus.SCanbusTask{
		Name: "视在功率",
		Metas: []*c_base.Meta{
			analogApparentPowerA, analogApparentPowerB, analogApparentPowerC, analogTotalApparentPower,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x08, params)
		},
	}

	analogInverterVoltageInfo = p_canbus.SCanbusTask{
		Name: "逆变电压",
		Metas: []*c_base.Meta{
			analogInverterVoltageA, analogInverterVoltageB, analogInverterVoltageC,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x09, params)
		},
	}

	analogGroundImpedanceInfo = p_canbus.SCanbusTask{
		Name: "对地阻抗和漏电流",
		Metas: []*c_base.Meta{
			analogPositiveGroundImpedance, analogNegativeGroundImpedance, analogLeakageCurrent, analogConverterEfficiency,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0A, params)
		},
	}

	analogTotalChargeInfo = p_canbus.SCanbusTask{
		Name: "充电量",
		Metas: []*c_base.Meta{
			analogTotalChargeL, analogTotalChargeH, analogTotalDischargeL, analogTotalDischargeH,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0B, params)
		},
	}

	analogBusVoltageInfo = p_canbus.SCanbusTask{
		Name: "母线电压",
		Metas: []*c_base.Meta{
			analogPositiveBusVoltage, analogNegativeBusVoltage, analogGroundNegativeVoltage,
		},
		GetCanbusID: func(params map[string]any) *uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0C, params)
		},
	}
)
