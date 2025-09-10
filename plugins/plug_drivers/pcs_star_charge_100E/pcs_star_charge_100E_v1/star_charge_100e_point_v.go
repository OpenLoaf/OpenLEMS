package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"github.com/shockerli/cvt"
)

// 年月日，可写
var (
	Year   = &c_base.SModbusPoint{Name: "Year", Addr: 30297, ReadType: c_base.RUint16, Desc: "年"}
	Month  = &c_base.SModbusPoint{Name: "Month", Addr: 30298, ReadType: c_base.RUint16, Desc: "月 1~12"}
	Day    = &c_base.SModbusPoint{Name: "Day", Addr: 30299, ReadType: c_base.RUint16, Desc: "日 1~31"}
	Hour   = &c_base.SModbusPoint{Name: "Hour", Addr: 30300, ReadType: c_base.RUint16, Desc: "时 0~23"}
	Minute = &c_base.SModbusPoint{Name: "Minute", Addr: 30301, ReadType: c_base.RUint16, Desc: "分 0~59"}
	Second = &c_base.SModbusPoint{Name: "Second", Addr: 30302, ReadType: c_base.RUint16, Desc: "秒 0~59"}
)

var (
	OnOffCommand = &c_base.SModbusPoint{Name: "OnOffCommand", Cn: "开关机指令", Addr: 30314, ReadType: c_base.RUint16,
		StatusExplain: func(value any) string {
			switch cvt.Int8(value) {
			case 0:
				return "关机"
			case 1:
				return "已启动"
			case 2:
				return "待机"
			}
			return "未知值:" + cvt.String(value)
		}, Desc: "On/off command: 0- Shutdown, 1- Startup, 2- Standby"}
	ActivePowerSetting   = &c_base.SModbusPoint{Debug: false, Name: "ActivePowerSetting", Cn: "有功功率设置", Addr: 30315, ReadType: c_base.RInt32, Endianness: c_enum.EMiddleEndian, Factor: 0.001, Desc: "Inverter active power setting, Positive power represents battery discharge, with power from the DC side to the AC side, Negative power represents battery charging, with power from the AC side to the DC side"}
	ReactivePowerSetting = &c_base.SModbusPoint{Name: "ReactivePowerSetting", Cn: "无功功率设置", Addr: 30317, ReadType: c_base.RInt32, Endianness: c_enum.EMiddleEndian, Factor: 0.001, Desc: "Inverter reactive power setting"}
)

var (
	PhaseAVoltageGridSide     = &c_base.SModbusPoint{Name: "PhaseAVoltageGridSide", Cn: "电网侧A相电压", Addr: 30329, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase A voltage of grid side"}
	PhaseACurrentInverterSide = &c_base.SModbusPoint{Name: "PhaseACurrentInverterSide", Cn: "逆变器测A相电流", Addr: 30330, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase A current of inverter side"}
	PhaseACurrentGridSide     = &c_base.SModbusPoint{Name: "PhaseACurrentGridSide", Cn: "电网侧A相电流", Addr: 30331, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase A current of grid side"}
	PhaseAPowerInverterSide   = &c_base.SModbusPoint{Name: "PhaseAPowerInverterSide", Cn: "逆变器测A相功率", Addr: 30332, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase A power of inverter side"}

	PhaseBVoltageGridSide     = &c_base.SModbusPoint{Name: "PhaseBVoltageGridSide", Cn: "电网侧B相电压", Addr: 30334, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase B voltage of grid side"}
	PhaseBCurrentInverterSide = &c_base.SModbusPoint{Name: "PhaseBCurrentInverterSide", Cn: "逆变器测B相电流", Addr: 30335, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase B current of inverter side"}
	PhaseBCurrentGridSide     = &c_base.SModbusPoint{Name: "PhaseBCurrentGridSide", Cn: "电网侧B相电流", Addr: 30336, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase B current of grid side"}
	PhaseBPowerInverterSide   = &c_base.SModbusPoint{Name: "PhaseBPowerInverterSide", Cn: "逆变器测B相功率", Addr: 30337, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase B power of inverter side"}

	PhaseCVoltageGridSide     = &c_base.SModbusPoint{Name: "PhaseCVoltageGridSide", Cn: "电网侧C相电压", Addr: 30339, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase C voltage of grid side"}
	PhaseCCurrentInverterSide = &c_base.SModbusPoint{Name: "PhaseCCurrentInverterSide", Cn: "逆变器测C相电流", Addr: 30340, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase C current of inverter side"}
	PhaseCCurrentGridSide     = &c_base.SModbusPoint{Name: "PhaseCCurrentGridSide", Cn: "电网侧C相电流", Addr: 30341, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase C current of grid side"}
	PhaseCPowerInverterSide   = &c_base.SModbusPoint{Name: "PhaseCPowerInverterSide", Cn: "逆变器测C相功率", Addr: 30342, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase C power of inverter side"}

	CurrentBalancedBridge          = &c_base.SModbusPoint{Name: "CurrentBalancedBridge", Cn: "平衡桥电流", Addr: 30344, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Current of balanced bridge"}
	VoltageLineAB                  = &c_base.SModbusPoint{Name: "VoltageLineAB", Cn: "AB线电压", Addr: 30345, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line AB"}
	VoltageLineBC                  = &c_base.SModbusPoint{Name: "VoltageLineBC", Cn: "BC线电压", Addr: 30346, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line BC"}
	VoltageLineCA                  = &c_base.SModbusPoint{Name: "VoltageLineCA", Cn: "CA线电压", Addr: 30347, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line CA"}
	AverageFrequency               = &c_base.SModbusPoint{Name: "AverageFrequency", Cn: "平均频率", Addr: 30348, ReadType: c_base.RUint16, Factor: 0.01, Unit: "Hz", Desc: "Average frequency"}
	AveragePowerFactor             = &c_base.SModbusPoint{Name: "AveragePowerFactor", Cn: "平均功率因数", Addr: 30349, ReadType: c_base.RUint16, Factor: 0.01, Desc: "Average power factor"}
	AverageVoltageBus              = &c_base.SModbusPoint{Name: "AverageVoltageBus", Cn: "母线平均电压", Addr: 30350, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of bus"}
	AverageVoltagePositive         = &c_base.SModbusPoint{Name: "AverageVoltagePositive", Cn: "正母线平均电压", Addr: 30351, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of positive bus"}
	AverageVoltageNegative         = &c_base.SModbusPoint{Name: "AverageVoltageNegative", Cn: "负母线平均电压", Addr: 30352, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of negative bus"}
	TotalActivePowerInverterSide   = &c_base.SModbusPoint{Name: "TotalActivePowerInverterSide", Cn: "逆变器侧总有功功率", Addr: 30353, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.001, Unit: "kW", Desc: "CurrentTotal active power on the inverter side"}
	TotalReactivePowerInverterSide = &c_base.SModbusPoint{Name: "TotalReactivePowerInverterSide", Cn: "逆变器侧总无功功率", Addr: 30355, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.001, Unit: "kVar", Desc: "CurrentTotal reactive power on the inverter side"}
	TotalApparentPowerInverterSide = &c_base.SModbusPoint{Name: "TotalApparentPowerInverterSide", Cn: "逆变器侧总视在功率", Addr: 30357, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.001, Unit: "kVA", Desc: "CurrentTotal apparent power on the inverter side"}

	BatterySideVoltage      = &c_base.SModbusPoint{Name: "BatterySideVoltage", Cn: "电池侧电压", Addr: 30359, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Battery side voltage"}
	BatterySideCurrent      = &c_base.SModbusPoint{Name: "BatterySideCurrent", Cn: "电池侧电流", Addr: 30360, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Battery side current"}
	BatterySidePower        = &c_base.SModbusPoint{Name: "BatterySidePower", Cn: "电池侧功率", Addr: 30361, ReadType: c_base.RUint16, Factor: 0.01, Unit: "kW", Desc: "Battery side power"}
	InverterOperationStatus = &c_base.SModbusPoint{Name: "InverterOperationStatus", Cn: "逆变器运行状态", Addr: 30374, ReadType: c_base.RUint16, StatusExplain: func(value any) string {
		switch cvt.Int8(value) {
		case 0:
			return "等待设备启动"
		case 1:
			return "上电自检"
		case 2:
			return "并网运行"
		case 3:
			return "离网运行"
		case 4:
			return "保留"
		case 5:
			return "异常"
		}
		return "未知值:" + cvt.String(value)
	}, Desc: "Inverter operation status: 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error"}

	SerialNumber1 = &c_base.SModbusPoint{Name: "SerialNumber1", Cn: "序列号1", Addr: 30407, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Serial number 1/5"}
	SerialNumber2 = &c_base.SModbusPoint{Name: "SerialNumber2", Cn: "序列号2", Addr: 30409, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Serial number 2/5"}
	SerialNumber3 = &c_base.SModbusPoint{Name: "SerialNumber3", Cn: "序列号3", Addr: 30411, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Serial number 3/5"}
	SerialNumber4 = &c_base.SModbusPoint{Name: "SerialNumber4", Cn: "序列号4", Addr: 30413, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Serial number 4/5"}
	SerialNumber5 = &c_base.SModbusPoint{Name: "SerialNumber5", Cn: "序列号5", Addr: 30415, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Serial number 5/5"}

	GridModeSetting = &c_base.SModbusPoint{Name: "GridModeSetting", Cn: "逆变器运行模式设置", Addr: 30443, ReadType: c_base.RUint16, Desc: "Inverter operation mode setting: 0- Initial state of power on (when internal communication of the inverter has not yet been established), 1- Grid connected mode (default mode for startup), 2- Off grid mode, 3- Reserved"}

	Module1Temperature         = &c_base.SModbusPoint{Name: "Module1Temperature", Cn: "模块1温度", Addr: 30454, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 1 Temperature"}
	Module2Temperature         = &c_base.SModbusPoint{Name: "Module2Temperature", Cn: "模块2温度", Addr: 30455, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 2 Temperature"}
	Module3Temperature         = &c_base.SModbusPoint{Name: "Module3Temperature", Cn: "模块3温度", Addr: 30456, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 3 Temperature"}
	RadiatorTemperature        = &c_base.SModbusPoint{Name: "RadiatorTemperature", Cn: "散热器温度", Addr: 30457, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Radiator temperature"}
	InternalAmbientTemperature = &c_base.SModbusPoint{Name: "InternalAmbientTemperature", Cn: "内部环境温度", Addr: 30459, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Internal ambient temperature"}

	RunTime = &c_base.SModbusPoint{Name: "RunTime", Cn: "运行时间", Addr: 31278, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Run time (seconds) query"}

	DailyBatteryChargeEnergy    = &c_base.SModbusPoint{Name: "DailyBatteryChargeEnergy", Cn: "每日电池充电能量", Addr: 31284, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Daily battery charge energy"}
	TotalBatteryChargeEnergy    = &c_base.SModbusPoint{Name: "TotalBatteryChargeEnergy", Cn: "总电池充电能量", Addr: 31286, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "CurrentTotal battery charge energy"}
	DailyBatteryDischargeEnergy = &c_base.SModbusPoint{Name: "DailyBatteryDischargeEnergy", Cn: "每日电池放电能量", Addr: 31288, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Daily battery discharge energy"}
	TotalBatteryDischargeEnergy = &c_base.SModbusPoint{Name: "TotalBatteryDischargeEnergy", Cn: "总电池放电能量", Addr: 31290, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_enum.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "CurrentTotal battery discharge energy"}

	AuxiliaryPowerOnStatus = &c_base.SModbusPoint{Name: "AuxiliaryPowerOnStatus", Cn: "辅助电源开启状态", Addr: 32000, ReadType: c_base.RUint16, Desc: "Auxiliary power on status：1-yes， 0-no"}

	BatteryChargeStatus    = &c_base.SModbusPoint{Name: "BatteryChargeStatus", Cn: "电池充电状态", Addr: 32006, ReadType: c_base.RUint16, Desc: "Battery charge status： 1-yes， 0-no"}
	BatteryDischargeStatus = &c_base.SModbusPoint{Name: "BatteryDischargeStatus", Cn: "电池放电状态", Addr: 32007, ReadType: c_base.RUint16, Desc: "Battery discharge status：1-yes， 0-no"}

	DCPositiveRelayStatus = &c_base.SModbusPoint{Name: "DCPositiveRelayStatus", Cn: "直流正继电器状态", Addr: 32033, ReadType: c_base.RUint16, Desc: "DC positive relay status： 1-On, 0-Off"}
	DCNegativeRelayStatus = &c_base.SModbusPoint{Name: "DCNegativeRelayStatus", Cn: "直流负继电器状态", Addr: 32034, ReadType: c_base.RUint16, Desc: "DC negative relay status： 1-On, 0-Off"}
	ACRelayStatus         = &c_base.SModbusPoint{Name: "ACRelayStatus", Cn: "交流继电器状态", Addr: 32035, SystemType: c_base.SBool, Desc: "AC relay status： 1-On, 0-Off"}
	GridOutageStatus      = &c_base.SModbusPoint{Name: "GridOutageStatus", Cn: "电网断电状态", Addr: 32036, SystemType: c_base.SBool, Desc: "Grid Outage Status：1-yes, 0-no"}

	SoftwareVersion = &c_base.SModbusPoint{Name: "SoftwareVersion", Cn: "软件版本", Addr: 33300, ReadType: c_base.RUint32, Endianness: c_enum.EMiddleEndian, Desc: "Software Version"}
)
