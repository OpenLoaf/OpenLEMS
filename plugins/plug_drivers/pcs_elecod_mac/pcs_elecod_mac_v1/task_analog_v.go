package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
)

var (
	analogAllTasks = []*p_canbus.SCanbusTask{
		&analogDCInfo,
		&analogGridVoltageInfo,
		&analogGridCurrentInfo,
		&analogGridFrequencyInfo,
		&analogActivePowerInfo,
		&analogReactivePowerInfo,
	}
)

var (
	analogDCInfo = p_canbus.SCanbusTask{
		Name: "直流参数",
		Metas: []*c_base.Meta{
			analogDcVoltage, analogDcCurrent, analogDcPower, analogBusVoltage,
		},
		GetCanbusID: func(params map[string]any) uint32 {
			return buildCANbusID(&CANFrameInfo{})
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x01
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogGridVoltageInfo = p_canbus.SCanbusTask{
		Name: "电网电压",
		Metas: []*c_base.Meta{
			analogGridVoltageA, analogGridVoltageB, analogGridVoltageC, analogPowerTubeTemp,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x02
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogGridCurrentInfo = p_canbus.SCanbusTask{
		Name: "电网电流",
		Metas: []*c_base.Meta{
			analogGridCurrentA, analogGridCurrentB, analogGridCurrentC, analogBridgeTemp,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x03
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogGridFrequencyInfo = p_canbus.SCanbusTask{
		Name: "电网频率",
		Metas: []*c_base.Meta{
			analogGridFrequencyA, analogGridFrequencyB, analogGridFrequencyC, analogAmbientTemp,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x04
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogActivePowerInfo = p_canbus.SCanbusTask{
		Name: "有功功率",
		Metas: []*c_base.Meta{
			analogActivePowerA, analogActivePowerB, analogActivePowerC, analogTotalActivePower,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x05
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogReactivePowerInfo = p_canbus.SCanbusTask{
		Name: "无功功率",
		Metas: []*c_base.Meta{
			analogReactivePowerA, analogReactivePowerB, analogReactivePowerC, analogTotalReactivePower,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x06
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogPowerFactorInfo = p_canbus.SCanbusTask{
		Name: "功率因素",
		Metas: []*c_base.Meta{
			analogPowerFactorA, analogPowerFactorB, analogPowerFactorC, analogTotalPowerFactor,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x07
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogApparentPowerInfo = p_canbus.SCanbusTask{
		Name: "视在功率",
		Metas: []*c_base.Meta{
			analogApparentPowerA, analogApparentPowerB, analogApparentPowerC, analogTotalApparentPower,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x08
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogInverterVoltageInfo = p_canbus.SCanbusTask{
		Name: "逆变电压",
		Metas: []*c_base.Meta{
			analogInverterVoltageA, analogInverterVoltageB, analogInverterVoltageC, analogReserved1,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x09
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogGroundImpedanceInfo = p_canbus.SCanbusTask{
		Name: "对地阻抗和漏电流",
		Metas: []*c_base.Meta{
			analogPositiveGroundImpedance, analogNegativeGroundImpedance, analogLeakageCurrent, analogConverterEfficiency,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x0A
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogTotalChargeInfo = p_canbus.SCanbusTask{
		Name: "充电量",
		Metas: []*c_base.Meta{
			analogTotalChargeL, analogTotalChargeH, analogTotalDischargeL, analogTotalDischargeH,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x0B
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}

	analogBusVoltageInfo = p_canbus.SCanbusTask{
		Name: "母线电压",
		Metas: []*c_base.Meta{
			analogPositiveBusVoltage, analogNegativeBusVoltage, analogGroundNegativeVoltage, analogReserved2,
		},
		IDMatch: func(id uint32) bool {
			info := parseCANbusID(id)
			match := info.TargetDeviceType == DeviceTypeScreen &&
				info.SourceDeviceType == DeviceTypeMAC &&
				info.MessageType == MessageTypeAnalog &&
				info.ServiceCode == 0x0C
			if match {
				PrintCanFrame(id, info)
				return true
			}
			return false
		},
	}
)
