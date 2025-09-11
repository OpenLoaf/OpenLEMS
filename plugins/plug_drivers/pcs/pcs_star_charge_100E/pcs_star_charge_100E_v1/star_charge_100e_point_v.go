package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"

	"github.com/shockerli/cvt"
)

// 年月日，可写
var (
	Year   = &c_proto.SModbusPoint{Addr: 30297, SPoint: &c_base.SPoint{Key: "Year", Name: "年", Desc: "年"}, DataAccess: c_default.VDataAccessUInt16}
	Month  = &c_proto.SModbusPoint{Addr: 30298, SPoint: &c_base.SPoint{Key: "Month", Name: "月", Desc: "月 1~12"}, DataAccess: c_default.VDataAccessUInt16}
	Day    = &c_proto.SModbusPoint{Addr: 30299, SPoint: &c_base.SPoint{Key: "Day", Name: "日", Desc: "日 1~31"}, DataAccess: c_default.VDataAccessUInt16}
	Hour   = &c_proto.SModbusPoint{Addr: 30300, SPoint: &c_base.SPoint{Key: "Hour", Name: "时", Desc: "时 0~23"}, DataAccess: c_default.VDataAccessUInt16}
	Minute = &c_proto.SModbusPoint{Addr: 30301, SPoint: &c_base.SPoint{Key: "Minute", Name: "分", Desc: "分 0~59"}, DataAccess: c_default.VDataAccessUInt16}
	Second = &c_proto.SModbusPoint{Addr: 30302, SPoint: &c_base.SPoint{Key: "Second", Name: "秒", Desc: "秒 0~59"}, DataAccess: c_default.VDataAccessUInt16}
)

var (
	OnOffCommand = &c_proto.SModbusPoint{
		Addr:       30314,
		SPoint:     &c_base.SPoint{Key: "OnOffCommand", Name: "开关机指令", Desc: "On/off command: 0- Shutdown, 1- Startup, 2- Standby"},
		DataAccess: c_default.VDataAccessUInt16,
		StatusExplain: func(value any) (string, error) {
			switch cvt.Int8(value) {
			case 0:
				return "关机", nil
			case 1:
				return "已启动", nil
			case 2:
				return "待机", nil
			}
			return "未知值:" + cvt.String(value), nil
		},
	}
	ActivePowerSetting   = &c_proto.SModbusPoint{Addr: 30315, SPoint: &c_base.SPoint{Key: "ActivePowerSetting", Name: "有功功率设置", Desc: "Inverter active power setting, Positive power represents battery discharge, with power from the DC side to the AC side, Negative power represents battery charging, with power from the AC side to the DC side"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.001}}
	ReactivePowerSetting = &c_proto.SModbusPoint{Addr: 30317, SPoint: &c_base.SPoint{Key: "ReactivePowerSetting", Name: "无功功率设置", Desc: "Inverter reactive power setting"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.001}}
)

var (
	PhaseAVoltageGridSide     = &c_proto.SModbusPoint{Addr: 30329, SPoint: &c_base.SPoint{Key: "PhaseAVoltageGridSide", Name: "电网侧A相电压", Unit: "V", Desc: "Phase A voltage of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseACurrentInverterSide = &c_proto.SModbusPoint{Addr: 30330, SPoint: &c_base.SPoint{Key: "PhaseACurrentInverterSide", Name: "逆变器测A相电流", Unit: "A", Desc: "Phase A current of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseACurrentGridSide     = &c_proto.SModbusPoint{Addr: 30331, SPoint: &c_base.SPoint{Key: "PhaseACurrentGridSide", Name: "电网侧A相电流", Unit: "A", Desc: "Phase A current of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseAPowerInverterSide   = &c_proto.SModbusPoint{Addr: 30332, SPoint: &c_base.SPoint{Key: "PhaseAPowerInverterSide", Name: "逆变器测A相功率", Unit: "kW", Desc: "Phase A power of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale0001}

	PhaseBVoltageGridSide     = &c_proto.SModbusPoint{Addr: 30334, SPoint: &c_base.SPoint{Key: "PhaseBVoltageGridSide", Name: "电网侧B相电压", Unit: "V", Desc: "Phase B voltage of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseBCurrentInverterSide = &c_proto.SModbusPoint{Addr: 30335, SPoint: &c_base.SPoint{Key: "PhaseBCurrentInverterSide", Name: "逆变器测B相电流", Unit: "A", Desc: "Phase B current of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseBCurrentGridSide     = &c_proto.SModbusPoint{Addr: 30336, SPoint: &c_base.SPoint{Key: "PhaseBCurrentGridSide", Name: "电网侧B相电流", Unit: "A", Desc: "Phase B current of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseBPowerInverterSide   = &c_proto.SModbusPoint{Addr: 30337, SPoint: &c_base.SPoint{Key: "PhaseBPowerInverterSide", Name: "逆变器测B相功率", Unit: "kW", Desc: "Phase B power of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale0001}

	PhaseCVoltageGridSide     = &c_proto.SModbusPoint{Addr: 30339, SPoint: &c_base.SPoint{Key: "PhaseCVoltageGridSide", Name: "电网侧C相电压", Unit: "V", Desc: "Phase C voltage of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseCCurrentInverterSide = &c_proto.SModbusPoint{Addr: 30340, SPoint: &c_base.SPoint{Key: "PhaseCCurrentInverterSide", Name: "逆变器测C相电流", Unit: "A", Desc: "Phase C current of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseCCurrentGridSide     = &c_proto.SModbusPoint{Addr: 30341, SPoint: &c_base.SPoint{Key: "PhaseCCurrentGridSide", Name: "电网侧C相电流", Unit: "A", Desc: "Phase C current of grid side"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PhaseCPowerInverterSide   = &c_proto.SModbusPoint{Addr: 30342, SPoint: &c_base.SPoint{Key: "PhaseCPowerInverterSide", Name: "逆变器测C相功率", Unit: "kW", Desc: "Phase C power of inverter side"}, DataAccess: c_default.VDataAccessUInt16Scale0001}

	CurrentBalancedBridge          = &c_proto.SModbusPoint{Addr: 30344, SPoint: &c_base.SPoint{Key: "CurrentBalancedBridge", Name: "平衡桥电流", Unit: "A", Desc: "Current of balanced bridge"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	VoltageLineAB                  = &c_proto.SModbusPoint{Addr: 30345, SPoint: &c_base.SPoint{Key: "VoltageLineAB", Name: "AB线电压", Unit: "V", Desc: "Voltage of line AB"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	VoltageLineBC                  = &c_proto.SModbusPoint{Addr: 30346, SPoint: &c_base.SPoint{Key: "VoltageLineBC", Name: "BC线电压", Unit: "V", Desc: "Voltage of line BC"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	VoltageLineCA                  = &c_proto.SModbusPoint{Addr: 30347, SPoint: &c_base.SPoint{Key: "VoltageLineCA", Name: "CA线电压", Unit: "V", Desc: "Voltage of line CA"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	AverageFrequency               = &c_proto.SModbusPoint{Addr: 30348, SPoint: &c_base.SPoint{Key: "AverageFrequency", Name: "平均频率", Unit: "Hz", Desc: "Average frequency"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	AveragePowerFactor             = &c_proto.SModbusPoint{Addr: 30349, SPoint: &c_base.SPoint{Key: "AveragePowerFactor", Name: "平均功率因数", Desc: "Average power factor"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	AverageVoltageBus              = &c_proto.SModbusPoint{Addr: 30350, SPoint: &c_base.SPoint{Key: "AverageVoltageBus", Name: "母线平均电压", Unit: "V", Desc: "Average voltage of bus"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	AverageVoltagePositive         = &c_proto.SModbusPoint{Addr: 30351, SPoint: &c_base.SPoint{Key: "AverageVoltagePositive", Name: "正母线平均电压", Unit: "V", Desc: "Average voltage of positive bus"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	AverageVoltageNegative         = &c_proto.SModbusPoint{Addr: 30352, SPoint: &c_base.SPoint{Key: "AverageVoltageNegative", Name: "负母线平均电压", Unit: "V", Desc: "Average voltage of negative bus"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	TotalActivePowerInverterSide   = &c_proto.SModbusPoint{Addr: 30353, SPoint: c_default.VPointP, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, ValueType: c_enum.EFloat64, Factor: 0.001}}
	TotalReactivePowerInverterSide = &c_proto.SModbusPoint{Addr: 30355, SPoint: c_default.VPointQ, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, ValueType: c_enum.EFloat64, Factor: 0.001}}
	TotalApparentPowerInverterSide = &c_proto.SModbusPoint{Addr: 30357, SPoint: c_default.VPointS, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, ValueType: c_enum.EFloat64, Factor: 0.001}}

	BatterySideVoltage      = &c_proto.SModbusPoint{Addr: 30359, SPoint: &c_base.SPoint{Key: "BatterySideVoltage", Name: "电池侧电压", Unit: "V", Desc: "Battery side voltage"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	BatterySideCurrent      = &c_proto.SModbusPoint{Addr: 30360, SPoint: &c_base.SPoint{Key: "BatterySideCurrent", Name: "电池侧电流", Unit: "A", Desc: "Battery side current"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	BatterySidePower        = &c_proto.SModbusPoint{Addr: 30361, SPoint: &c_base.SPoint{Key: "BatterySidePower", Name: "电池侧功率", Unit: "kW", Desc: "Battery side power"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	InverterOperationStatus = &c_proto.SModbusPoint{
		Addr:       30374,
		SPoint:     &c_base.SPoint{Key: "InverterOperationStatus", Name: "逆变器运行状态", Desc: "Inverter operation status: 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error"},
		DataAccess: c_default.VDataAccessUInt16,
		StatusExplain: func(value any) (string, error) {
			switch cvt.Int8(value) {
			case 0:
				return "等待设备启动", nil
			case 1:
				return "上电自检", nil
			case 2:
				return "并网运行", nil
			case 3:
				return "离网运行", nil
			case 4:
				return "保留", nil
			case 5:
				return "异常", nil
			}
			return "未知值:" + cvt.String(value), nil
		},
	}

	SerialNumber1 = &c_proto.SModbusPoint{Addr: 30407, SPoint: &c_base.SPoint{Key: "SerialNumber1", Name: "序列号1", Desc: "Serial number 1/5"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}
	SerialNumber2 = &c_proto.SModbusPoint{Addr: 30409, SPoint: &c_base.SPoint{Key: "SerialNumber2", Name: "序列号2", Desc: "Serial number 2/5"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}
	SerialNumber3 = &c_proto.SModbusPoint{Addr: 30411, SPoint: &c_base.SPoint{Key: "SerialNumber3", Name: "序列号3", Desc: "Serial number 3/5"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}
	SerialNumber4 = &c_proto.SModbusPoint{Addr: 30413, SPoint: &c_base.SPoint{Key: "SerialNumber4", Name: "序列号4", Desc: "Serial number 4/5"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}
	SerialNumber5 = &c_proto.SModbusPoint{Addr: 30415, SPoint: &c_base.SPoint{Key: "SerialNumber5", Name: "序列号5", Desc: "Serial number 5/5"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}

	GridModeSetting = &c_proto.SModbusPoint{Addr: 30443, SPoint: &c_base.SPoint{Key: "GridModeSetting", Name: "逆变器运行模式设置", Desc: "Inverter operation mode setting: 0- Initial state of power on (when internal communication of the inverter has not yet been established), 1- Grid connected mode (default mode for startup), 2- Off grid mode, 3- Reserved"}, DataAccess: c_default.VDataAccessUInt16}

	Module1Temperature         = &c_proto.SModbusPoint{Addr: 30454, SPoint: &c_base.SPoint{Key: "Module1Temperature", Name: "模块1温度", Unit: "℃", Desc: "Module 1 Temperature"}, DataAccess: c_default.VDataAccessInt16Scale001}
	Module2Temperature         = &c_proto.SModbusPoint{Addr: 30455, SPoint: &c_base.SPoint{Key: "Module2Temperature", Name: "模块2温度", Unit: "℃", Desc: "Module 2 Temperature"}, DataAccess: c_default.VDataAccessInt16Scale001}
	Module3Temperature         = &c_proto.SModbusPoint{Addr: 30456, SPoint: &c_base.SPoint{Key: "Module3Temperature", Name: "模块3温度", Unit: "℃", Desc: "Module 3 Temperature"}, DataAccess: c_default.VDataAccessInt16Scale001}
	RadiatorTemperature        = &c_proto.SModbusPoint{Addr: 30457, SPoint: &c_base.SPoint{Key: "RadiatorTemperature", Name: "散热器温度", Unit: "℃", Desc: "Radiator temperature"}, DataAccess: c_default.VDataAccessInt16Scale001}
	InternalAmbientTemperature = &c_proto.SModbusPoint{Addr: 30459, SPoint: &c_base.SPoint{Key: "InternalAmbientTemperature", Name: "内部环境温度", Unit: "℃", Desc: "Internal ambient temperature"}, DataAccess: c_default.VDataAccessInt16Scale001}

	RunTime = &c_proto.SModbusPoint{Addr: 31278, SPoint: &c_base.SPoint{Key: "RunTime", Name: "运行时间", Desc: "Run time (seconds) query"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}

	DailyBatteryChargeEnergy    = &c_proto.SModbusPoint{Addr: 31284, SPoint: &c_base.SPoint{Key: "DailyBatteryChargeEnergy", Name: "每日电池充电能量", Unit: "kWh", Desc: "Daily battery charge energy"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.0001}}
	TotalBatteryChargeEnergy    = &c_proto.SModbusPoint{Addr: 31286, SPoint: &c_base.SPoint{Key: "TotalBatteryChargeEnergy", Name: "总电池充电能量", Unit: "kWh", Desc: "CurrentTotal battery charge energy"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.0001}}
	DailyBatteryDischargeEnergy = &c_proto.SModbusPoint{Addr: 31288, SPoint: &c_base.SPoint{Key: "DailyBatteryDischargeEnergy", Name: "每日电池放电能量", Unit: "kWh", Desc: "Daily battery discharge energy"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.0001}}
	TotalBatteryDischargeEnergy = &c_proto.SModbusPoint{Addr: 31290, SPoint: &c_base.SPoint{Key: "TotalBatteryDischargeEnergy", Name: "总电池放电能量", Unit: "kWh", Desc: "CurrentTotal battery discharge energy"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.0001}}

	AuxiliaryPowerOnStatus = &c_proto.SModbusPoint{Addr: 32000, SPoint: &c_base.SPoint{Key: "AuxiliaryPowerOnStatus", Name: "辅助电源开启状态", Desc: "Auxiliary power on status：1-yes， 0-no"}, DataAccess: c_default.VDataAccessUInt16ToBool}

	BatteryChargeStatus    = &c_proto.SModbusPoint{Addr: 32006, SPoint: &c_base.SPoint{Key: "BatteryChargeStatus", Name: "电池充电状态", Desc: "Battery charge status： 1-yes， 0-no"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	BatteryDischargeStatus = &c_proto.SModbusPoint{Addr: 32007, SPoint: &c_base.SPoint{Key: "BatteryDischargeStatus", Name: "电池放电状态", Desc: "Battery discharge status：1-yes， 0-no"}, DataAccess: c_default.VDataAccessUInt16ToBool}

	DCPositiveRelayStatus = &c_proto.SModbusPoint{Addr: 32033, SPoint: &c_base.SPoint{Key: "DCPositiveRelayStatus", Name: "直流正继电器状态", Desc: "DC positive relay status： 1-On, 0-Off"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	DCNegativeRelayStatus = &c_proto.SModbusPoint{Addr: 32034, SPoint: &c_base.SPoint{Key: "DCNegativeRelayStatus", Name: "直流负继电器状态", Desc: "DC negative relay status： 1-On, 0-Off"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	ACRelayStatus         = &c_proto.SModbusPoint{Addr: 32035, SPoint: &c_base.SPoint{Key: "ACRelayStatus", Name: "交流继电器状态", Desc: "AC relay status： 1-On, 0-Off"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	GridOutageStatus      = &c_proto.SModbusPoint{Addr: 32036, SPoint: &c_base.SPoint{Key: "GridOutageStatus", Name: "电网断电状态", Desc: "Grid Outage Status：1-yes, 0-no"}, DataAccess: c_default.VDataAccessUInt16ToBool}

	SoftwareVersion = &c_proto.SModbusPoint{Addr: 33300, SPoint: &c_base.SPoint{Key: "SoftwareVersion", Name: "软件版本", Desc: "Software Version"}, DataAccess: &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32}}
)
