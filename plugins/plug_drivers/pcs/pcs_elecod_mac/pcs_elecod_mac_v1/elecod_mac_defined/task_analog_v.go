package elecod_mac_defined

import (
	"common/c_proto"
	"pcs_elecod/elecod_canbus"
)

var (
	AnalogAllTasks = []*c_proto.SCanbusTask{
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
	analogDCInfo = c_proto.SCanbusTask{
		Name: "直流参数",
		Points: []*c_proto.SCanbusPoint{
			analogDcVoltage, analogDcCurrent, analogDcPower, analogBusVoltage,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x01, params)
		},
	}

	analogGridVoltageInfo = c_proto.SCanbusTask{
		Name: "电网电压",
		Points: []*c_proto.SCanbusPoint{
			analogGridVoltageA, analogGridVoltageB, analogGridVoltageC, analogPowerTubeTemp,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x02, params)
		},
	}

	analogGridCurrentInfo = c_proto.SCanbusTask{
		Name: "电网电流",
		Points: []*c_proto.SCanbusPoint{
			analogGridCurrentA, analogGridCurrentB, analogGridCurrentC, analogBridgeTemp,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x03, params)
		},
	}

	analogGridFrequencyInfo = c_proto.SCanbusTask{
		Name: "电网频率",
		Points: []*c_proto.SCanbusPoint{
			AnalogGridFrequencyA, AnalogGridFrequencyB, AnalogGridFrequencyC, analogAmbientTemp,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x04, params)
		},
	}

	analogActivePowerInfo = c_proto.SCanbusTask{
		Name: "有功功率",
		Points: []*c_proto.SCanbusPoint{
			analogActivePowerA, analogActivePowerB, analogActivePowerC, AnalogTotalActivePower,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x05, params)
		},
	}

	analogReactivePowerInfo = c_proto.SCanbusTask{
		Name: "无功功率",
		Points: []*c_proto.SCanbusPoint{
			analogReactivePowerA, analogReactivePowerB, analogReactivePowerC, AnalogTotalReactivePower,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x06, params)
		},
	}

	analogPowerFactorInfo = c_proto.SCanbusTask{
		Name: "功率因素",
		Points: []*c_proto.SCanbusPoint{
			analogPowerFactorA, analogPowerFactorB, analogPowerFactorC, AnalogTotalPowerFactor,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x07, params)
		},
	}

	analogApparentPowerInfo = c_proto.SCanbusTask{
		Name: "视在功率",
		Points: []*c_proto.SCanbusPoint{
			analogApparentPowerA, analogApparentPowerB, analogApparentPowerC, AnalogTotalApparentPower,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x08, params)
		},
	}

	analogInverterVoltageInfo = c_proto.SCanbusTask{
		Name: "逆变电压",
		Points: []*c_proto.SCanbusPoint{
			analogInverterVoltageA, analogInverterVoltageB, analogInverterVoltageC,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x09, params)
		},
	}

	analogGroundImpedanceInfo = c_proto.SCanbusTask{
		Name: "对地阻抗和漏电流",
		Points: []*c_proto.SCanbusPoint{
			analogPositiveGroundImpedance, analogNegativeGroundImpedance, analogLeakageCurrent, analogConverterEfficiency,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0A, params)
		},
	}

	analogTotalChargeInfo = c_proto.SCanbusTask{
		Name: "充电量",
		Points: []*c_proto.SCanbusPoint{
			analogTotalChargeL, analogTotalChargeH, analogTotalDischargeL, analogTotalDischargeH,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0B, params)
		},
	}

	analogBusVoltageInfo = c_proto.SCanbusTask{
		Name: "母线电压",
		Points: []*c_proto.SCanbusPoint{
			analogPositiveBusVoltage, analogNegativeBusVoltage, analogGroundNegativeVoltage,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return elecod_canbus.BuildMacToScreenCanbusID(analogMessageType, 0x0C, params)
		},
	}
)
