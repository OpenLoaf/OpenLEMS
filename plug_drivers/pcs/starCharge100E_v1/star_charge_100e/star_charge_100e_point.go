package star_charge_100e

import "ems-plan/c_base"

// 年月日，可写
var (
	Year   = &c_base.Meta{Name: "Year", Addr: 30297, ReadType: c_base.RUint16, Desc: "年"}
	Month  = &c_base.Meta{Name: "Month", Addr: 30298, ReadType: c_base.RUint16, Desc: "月 1~12"}
	Day    = &c_base.Meta{Name: "Day", Addr: 30299, ReadType: c_base.RUint16, Desc: "日 1~31"}
	Hour   = &c_base.Meta{Name: "Hour", Addr: 30300, ReadType: c_base.RUint16, Desc: "时 0~23"}
	Minute = &c_base.Meta{Name: "Minute", Addr: 30301, ReadType: c_base.RUint16, Desc: "分 0~59"}
	Second = &c_base.Meta{Name: "Second", Addr: 30302, ReadType: c_base.RUint16, Desc: "秒 0~59"}
)

var (
	OnOffCommand         = &c_base.Meta{Name: "OnOffCommand", Cn: "开关机指令", Addr: 30314, ReadType: c_base.RUint16, Desc: "On/off command: 0- Shutdown, 1- Startup, 2- Standby"}
	ActivePowerSetting   = &c_base.Meta{Debug: false, Name: "ActivePowerSetting", Cn: "有功功率设置", Addr: 30315, ReadType: c_base.RInt32, Endianness: c_base.EMiddleEndian, Factor: 0.001, Desc: "Inverter active power setting, Positive power represents battery discharge, with power from the DC side to the AC side, Negative power represents battery charging, with power from the AC side to the DC side"}
	ReactivePowerSetting = &c_base.Meta{Name: "ReactivePowerSetting", Cn: "无功功率设置", Addr: 30317, ReadType: c_base.RInt32, Endianness: c_base.EMiddleEndian, Factor: 0.001, Desc: "Inverter reactive power setting"}
)

var (
	PhaseAVoltageGridSide     = &c_base.Meta{Name: "PhaseAVoltageGridSide", Cn: "电网侧A相电压", Addr: 30329, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase A voltage of grid side"}
	PhaseACurrentInverterSide = &c_base.Meta{Name: "PhaseACurrentInverterSide", Cn: "逆变器测A相电流", Addr: 30330, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase A current of inverter side"}
	PhaseACurrentGridSide     = &c_base.Meta{Name: "PhaseACurrentGridSide", Cn: "电网侧A相电流", Addr: 30331, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase A current of grid side"}
	PhaseAPowerInverterSide   = &c_base.Meta{Name: "PhaseAPowerInverterSide", Cn: "逆变器测A相功率", Addr: 30332, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase A power of inverter side"}

	PhaseBVoltageGridSide     = &c_base.Meta{Name: "PhaseBVoltageGridSide", Cn: "电网侧B相电压", Addr: 30334, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase B voltage of grid side"}
	PhaseBCurrentInverterSide = &c_base.Meta{Name: "PhaseBCurrentInverterSide", Cn: "逆变器测B相电流", Addr: 30335, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase B current of inverter side"}
	PhaseBCurrentGridSide     = &c_base.Meta{Name: "PhaseBCurrentGridSide", Cn: "电网侧B相电流", Addr: 30336, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase B current of grid side"}
	PhaseBPowerInverterSide   = &c_base.Meta{Name: "PhaseBPowerInverterSide", Cn: "逆变器测B相功率", Addr: 30337, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase B power of inverter side"}

	PhaseCVoltageGridSide     = &c_base.Meta{Name: "PhaseCVoltageGridSide", Cn: "电网侧C相电压", Addr: 30339, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Phase C voltage of grid side"}
	PhaseCCurrentInverterSide = &c_base.Meta{Name: "PhaseCCurrentInverterSide", Cn: "逆变器测C相电流", Addr: 30340, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase C current of inverter side"}
	PhaseCCurrentGridSide     = &c_base.Meta{Name: "PhaseCCurrentGridSide", Cn: "电网侧C相电流", Addr: 30341, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Phase C current of grid side"}
	PhaseCPowerInverterSide   = &c_base.Meta{Name: "PhaseCPowerInverterSide", Cn: "逆变器测C相功率", Addr: 30342, ReadType: c_base.RFloat32, Factor: 0.001, Unit: "kW", Desc: "Phase C power of inverter side"}

	CurrentBalancedBridge          = &c_base.Meta{Name: "CurrentBalancedBridge", Cn: "平衡桥电流", Addr: 30344, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Current of balanced bridge"}
	VoltageLineAB                  = &c_base.Meta{Name: "VoltageLineAB", Cn: "AB线电压", Addr: 30345, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line AB"}
	VoltageLineBC                  = &c_base.Meta{Name: "VoltageLineBC", Cn: "BC线电压", Addr: 30346, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line BC"}
	VoltageLineCA                  = &c_base.Meta{Name: "VoltageLineCA", Cn: "CA线电压", Addr: 30347, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Voltage of line CA"}
	AverageFrequency               = &c_base.Meta{Name: "AverageFrequency", Cn: "平均频率", Addr: 30348, ReadType: c_base.RUint16, Factor: 0.01, Unit: "Hz", Desc: "Average frequency"}
	AveragePowerFactor             = &c_base.Meta{Name: "AveragePowerFactor", Cn: "平均功率因数", Addr: 30349, ReadType: c_base.RUint16, Factor: 0.01, Desc: "Average power factor"}
	AverageVoltageBus              = &c_base.Meta{Name: "AverageVoltageBus", Cn: "母线平均电压", Addr: 30350, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of bus"}
	AverageVoltagePositive         = &c_base.Meta{Name: "AverageVoltagePositive", Cn: "正母线平均电压", Addr: 30351, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of positive bus"}
	AverageVoltageNegative         = &c_base.Meta{Name: "AverageVoltageNegative", Cn: "负母线平均电压", Addr: 30352, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Average voltage of negative bus"}
	TotalActivePowerInverterSide   = &c_base.Meta{Name: "TotalActivePowerInverterSide", Cn: "逆变器侧总有功功率", Addr: 30353, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.001, Unit: "kW", Desc: "Total active power on the inverter side"}
	TotalReactivePowerInverterSide = &c_base.Meta{Name: "TotalReactivePowerInverterSide", Cn: "逆变器侧总无功功率", Addr: 30355, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.001, Unit: "kVar", Desc: "Total reactive power on the inverter side"}
	TotalApparentPowerInverterSide = &c_base.Meta{Name: "TotalApparentPowerInverterSide", Cn: "逆变器侧总视在功率", Addr: 30357, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.001, Unit: "kVA", Desc: "Total apparent power on the inverter side"}

	BatterySideVoltage      = &c_base.Meta{Name: "BatterySideVoltage", Cn: "电池侧电压", Addr: 30359, ReadType: c_base.RUint16, Factor: 0.1, Unit: "V", Desc: "Battery side voltage"}
	BatterySideCurrent      = &c_base.Meta{Name: "BatterySideCurrent", Cn: "电池侧电流", Addr: 30360, ReadType: c_base.RUint16, Factor: 0.1, Unit: "A", Desc: "Battery side current"}
	BatterySidePower        = &c_base.Meta{Name: "BatterySidePower", Cn: "电池侧功率", Addr: 30361, ReadType: c_base.RUint16, Factor: 0.01, Unit: "kW", Desc: "Battery side power"}
	InverterOperationStatus = &c_base.Meta{Name: "InverterOperationStatus", Cn: "逆变器运行状态", Addr: 30374, ReadType: c_base.RUint16, Desc: "Inverter operation status: 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error"}

	SerialNumber1 = &c_base.Meta{Name: "SerialNumber1", Cn: "序列号1", Addr: 30407, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Serial number 1/5"}
	SerialNumber2 = &c_base.Meta{Name: "SerialNumber2", Cn: "序列号2", Addr: 30409, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Serial number 2/5"}
	SerialNumber3 = &c_base.Meta{Name: "SerialNumber3", Cn: "序列号3", Addr: 30411, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Serial number 3/5"}
	SerialNumber4 = &c_base.Meta{Name: "SerialNumber4", Cn: "序列号4", Addr: 30413, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Serial number 4/5"}
	SerialNumber5 = &c_base.Meta{Name: "SerialNumber5", Cn: "序列号5", Addr: 30415, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Serial number 5/5"}

	GridModeSetting = &c_base.Meta{Name: "GridModeSetting", Cn: "逆变器运行模式设置", Addr: 30443, ReadType: c_base.RUint16, Desc: "Inverter operation mode setting: 0- Initial state of power on (when internal communication of the inverter has not yet been established), 1- Grid connected mode (default mode for startup), 2- Off grid mode, 3- Reserved"}

	Module1Temperature         = &c_base.Meta{Name: "Module1Temperature", Cn: "模块1温度", Addr: 30454, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 1 Temperature"}
	Module2Temperature         = &c_base.Meta{Name: "Module2Temperature", Cn: "模块2温度", Addr: 30455, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 2 Temperature"}
	Module3Temperature         = &c_base.Meta{Name: "Module3Temperature", Cn: "模块3温度", Addr: 30456, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Module 3 Temperature"}
	RadiatorTemperature        = &c_base.Meta{Name: "RadiatorTemperature", Cn: "散热器温度", Addr: 30457, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Radiator temperature"}
	InternalAmbientTemperature = &c_base.Meta{Name: "InternalAmbientTemperature", Cn: "内部环境温度", Addr: 30459, ReadType: c_base.RInt16, Factor: 0.01, Unit: "℃", Desc: "Internal ambient temperature"}

	RunTime = &c_base.Meta{Name: "RunTime", Cn: "运行时间", Addr: 31278, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Run time (seconds) query"}

	DailyBatteryChargeEnergy    = &c_base.Meta{Name: "DailyBatteryChargeEnergy", Cn: "每日电池充电能量", Addr: 31284, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Daily battery charge energy"}
	TotalBatteryChargeEnergy    = &c_base.Meta{Name: "TotalBatteryChargeEnergy", Cn: "总电池充电能量", Addr: 31286, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Total battery charge energy"}
	DailyBatteryDischargeEnergy = &c_base.Meta{Name: "DailyBatteryDischargeEnergy", Cn: "每日电池放电能量", Addr: 31288, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Daily battery discharge energy"}
	TotalBatteryDischargeEnergy = &c_base.Meta{Name: "TotalBatteryDischargeEnergy", Cn: "总电池放电能量", Addr: 31290, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Unit: "kWh", Desc: "Total battery discharge energy"}

	AuxiliaryPowerOnStatus = &c_base.Meta{Name: "AuxiliaryPowerOnStatus", Cn: "辅助电源开启状态", Addr: 32000, ReadType: c_base.RUint16, Desc: "Auxiliary power on status：1-yes， 0-no"}

	BatteryChargeStatus    = &c_base.Meta{Name: "BatteryChargeStatus", Cn: "电池充电状态", Addr: 32006, ReadType: c_base.RUint16, Desc: "Battery charge status： 1-yes， 0-no"}
	BatteryDischargeStatus = &c_base.Meta{Name: "BatteryDischargeStatus", Cn: "电池放电状态", Addr: 32007, ReadType: c_base.RUint16, Desc: "Battery discharge status：1-yes， 0-no"}

	DCPositiveRelayStatus = &c_base.Meta{Name: "DCPositiveRelayStatus", Cn: "直流正继电器状态", Addr: 32033, ReadType: c_base.RUint16, Desc: "DC positive relay status： 1-On, 0-Off"}
	DCNegativeRelayStatus = &c_base.Meta{Name: "DCNegativeRelayStatus", Cn: "直流负继电器状态", Addr: 32034, ReadType: c_base.RUint16, Desc: "DC negative relay status： 1-On, 0-Off"}
	ACRelayStatus         = &c_base.Meta{Name: "ACRelayStatus", Cn: "交流继电器状态", Addr: 32035, SystemType: c_base.SBool, Desc: "AC relay status： 1-On, 0-Off"}
	GridOutageStatus      = &c_base.Meta{Name: "GridOutageStatus", Cn: "电网断电状态", Addr: 32036, SystemType: c_base.SBool, Desc: "Grid Outage Status：1-yes, 0-no"}

	SoftwareVersion = &c_base.Meta{Name: "SoftwareVersion", Cn: "软件版本", Addr: 33300, ReadType: c_base.RUint32, Endianness: c_base.EMiddleEndian, Desc: "Software Version"}
)
