package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_func"
	"fmt"
	"github.com/shockerli/cvt"
)

var (
	StatusGroup = &c_base.SPointGroup{GroupName: "状态", GroupSort: 1}
)

// 设备信息
var (
// Pylon            = &c_base.SModbusPoint{Id: "Pylon", Addr: 0x1000, BitSize: 16 * 5, SystemType: c_base.SString, Desc: "设备信息"}
// MBms             = &c_base.SModbusPoint{Id: "MBms", Addr: 0x1005, BitSize: 16 * 5, SystemType: c_base.SString, Desc: "电池信息"}
// Version          = &c_base.SModbusPoint{Id: "Version", Addr: 0x100A,  ReadType: c_base.RInt16, Desc: "版本信息"}
// InternalVersion  = &c_base.SModbusPoint{Id: "InternalVersion", Addr: 0x100B,  ReadType: c_base.RInt16, Desc: "内部版本信息"}
// ParallelPilesQty = &c_base.SModbusPoint{Id: "ParallelPilesQty", Addr: 0x100C,  ReadType: c_base.RInt16, Desc: "并联电池数量"}
)

// 控制信息
var (
	SleepControl   = &c_base.SModbusPoint{Name: "SleepControl", Addr: 0x1090, ReadType: c_base.RInt16, Desc: "休眠控制, 写入0xAA进入休眠状态，写入0x55唤醒"}
	AllowCharge    = &c_base.SModbusPoint{Name: "AllowCharge", Addr: 0x1091, ReadType: c_base.RInt16, Desc: "允许充电控制, 写入0xAA允许充电，其他无效"}
	AllowDischarge = &c_base.SModbusPoint{Name: "AllowDischarge", Addr: 0x1092, ReadType: c_base.RInt16, Desc: "允许放电控制, 写入0xAA允许放电，其他无效"}
	TempShield     = &c_base.SModbusPoint{Name: "TempShield", Addr: 0x1093, ReadType: c_base.RInt16, Desc: "临时屏蔽“无通信时切离继电器功能”的请求"}
	AllowRun       = &c_base.SModbusPoint{Name: "AllowRun", Addr: 0x1094, ReadType: c_base.RInt16, Desc: "运行指令，通知系统开始并柜, 写入0xAA允许运行，其他无效"}
)

// 时间
var (
	Year   = &c_base.SModbusPoint{Name: "Year", Addr: 0x10E0, ReadType: c_base.RUint16, Desc: "年 00~99（表示 2000~2099）"}
	Month  = &c_base.SModbusPoint{Name: "Month", Addr: 0x10E1, ReadType: c_base.RUint16, Desc: "月 1~12"}
	Day    = &c_base.SModbusPoint{Name: "Day", Addr: 0x10E2, ReadType: c_base.RUint16, Desc: "日 1~31"}
	Hour   = &c_base.SModbusPoint{Name: "Hour", Addr: 0x10E3, ReadType: c_base.RUint16, Desc: "时 0~23"}
	Minute = &c_base.SModbusPoint{Name: "Minute", Addr: 0x10E4, ReadType: c_base.RUint16, Desc: "分 0~59"}
	Second = &c_base.SModbusPoint{Name: "Second", Addr: 0x10E5, ReadType: c_base.RUint16, Desc: "秒 0~59"}
)

var (
	//BasicStatus   = &c_base.SModbusPoint{Id: "BasicStatus", Cn: "基本状态", Addr: 0x1100, ReadType: c_base.RInt16, Desc: "基本状态"}
	BasicStatus = &c_base.SModbusPoint{Name: "BasicStatus", Cn: "基本状态", Addr: 0x1100, ReadType: c_base.RBit0, BitLength: 2, Desc: "基本状态:00：休眠 01：充电 02：放电 03：搁置", StatusExplain: func(value any) string {
		if v, err := cvt.Int8E(value); err == nil {
			switch v {
			case 0:
				return "休眠"
			case 1:
				return "充电"
			case 2:
				return "放电"
			case 3:
				return "搁置"
			}
		}
		return fmt.Sprintf("未知值: %v", value)
	}}
	SystemErrorProtection = &c_base.SModbusPoint{Name: "SystemErrorProtection", Group: StatusGroup, Cn: "系统故障保护", Addr: 0x1100, ReadType: c_base.RBit3, Level: c_enum.EAlarmLevelError, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	CurrentProtection     = &c_base.SModbusPoint{Name: "CurrentProtection", Group: StatusGroup, Cn: "电流保护", Addr: 0x1100, ReadType: c_base.RBit4, Level: c_enum.EAlarmLevelError, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	VoltageProtection     = &c_base.SModbusPoint{Name: "VoltageProtection", Group: StatusGroup, Cn: "电压保护", Addr: 0x1100, ReadType: c_base.RBit5, Level: c_enum.EAlarmLevelError, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	TemperatureProtection = &c_base.SModbusPoint{Name: "TemperatureProtection", Group: StatusGroup, Cn: "温度保护", Addr: 0x1100, ReadType: c_base.RBit6, Level: c_enum.EAlarmLevelError, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	VoltageAlarm          = &c_base.SModbusPoint{Name: "VoltageAlarm", Group: StatusGroup, Cn: "电压警告", Addr: 0x1100, ReadType: c_base.RBit7, Level: c_enum.EAlarmLevelAlarm, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	CurrentAlarm          = &c_base.SModbusPoint{Name: "CurrentAlarm", Group: StatusGroup, Cn: "电流警告", Addr: 0x1100, ReadType: c_base.RBit8, Level: c_enum.EAlarmLevelAlarm, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	TemperatureAlarm      = &c_base.SModbusPoint{Name: "TemperatureAlarm", Group: StatusGroup, Cn: "温度警告", Addr: 0x1100, ReadType: c_base.RBit9, Level: c_enum.EAlarmLevelAlarm, Desc: "0-正常；1-保护", StatusExplain: c_func.StatusExplainProtectFunc}
	PileSystemIdleStatus  = &c_base.SModbusPoint{Name: "PileSystemIdleStatus", Group: StatusGroup, Cn: "电池组搁置状态", Addr: 0x1100, ReadType: c_base.RBit10, Desc: "0-否，1-搁置", StatusExplain: func(value any) string {
		return c_func.BoolExplain(value, "搁置", "否")
	}}
	PileSystemChargeStatus = &c_base.SModbusPoint{Name: "PileSystemChargeStatus", Group: StatusGroup, Cn: "电池组充电状态", Addr: 0x1100, ReadType: c_base.RBit11, Desc: "0-否，1-充电", StatusExplain: func(value any) string {
		return c_func.BoolExplain(value, "充电", "否")
	}}
	PileSystemDischargeStatus = &c_base.SModbusPoint{Name: "PileSystemDischargeStatus", Group: StatusGroup, Cn: "电池组放电状态", Addr: 0x1100, ReadType: c_base.RBit12, Desc: "0-否，1-放电", StatusExplain: func(value any) string {
		return c_func.BoolExplain(value, "放电", "否")
	}}
	PileSystemSleepStatus = &c_base.SModbusPoint{Name: "PileSystemSleepStatus", Group: StatusGroup, Cn: "电池组休眠状态", Addr: 0x1100, ReadType: c_base.RBit13, Desc: "0-否，1-休眠", StatusExplain: func(value any) string {
		return c_func.BoolExplain(value, "休眠", "否")
	}}
	FanWarn = &c_base.SModbusPoint{Name: "FanWarn", Group: StatusGroup, Cn: "电池组休眠状态", Addr: 0x1100, ReadType: c_base.RBit14, Level: c_enum.EAlarmLevelAlarm, Desc: "0-无异常，1-有异常", StatusExplain: func(value any) string {
		return c_func.BoolExplain(value, "有异常", "无异常")
	}}

	Protection                  = &c_base.SModbusPoint{Name: "Protection", Addr: 0x1101, ReadType: c_base.RInt16, Cn: "保护状态"}
	AlarmStatus1                = &c_base.SModbusPoint{Name: "AlarmStatus1", Addr: 0x1102, ReadType: c_base.RInt16, Cn: "告警状态1", Trigger: c_func.IsNotZero}
	DCVoltage                   = &c_base.SModbusPoint{Name: "DCVoltage", Addr: 0x1103, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "总电压", Unit: "V"}
	DCCurrent                   = &c_base.SModbusPoint{Name: "DCCurrent", Addr: 0x1104, ReadType: c_base.RInt32, SystemType: c_base.SFloat32, Factor: 0.01, Cn: "电流", Unit: "A"}
	Temperature                 = &c_base.SModbusPoint{Name: "Temperature", Addr: 0x1106, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "温度", Unit: " ̊C"}
	SOC                         = &c_base.SModbusPoint{Name: "SOC", Addr: 0x1107, ReadType: c_base.RUint16, Factor: 1, Desc: "SOC", Unit: "%"}
	CycleCount                  = &c_base.SModbusPoint{Name: "CycleCount", Addr: 0x1108, ReadType: c_base.RUint16, Cn: "循环次数"}
	PileMaxV                    = &c_base.SModbusPoint{Name: "PileMaxV", Addr: 0x1109, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池组最大充电电压值", Unit: "V"}
	PileMaxI                    = &c_base.SModbusPoint{Name: "PileMaxI", Addr: 0x110A, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01, Cn: "电池组最大充电电流值"}
	PileMinV                    = &c_base.SModbusPoint{Name: "PileMinV", Addr: 0x110C, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池组最小放电电压值", Unit: "V"}
	PileMaxDI                   = &c_base.SModbusPoint{Name: "PileMaxDI", Addr: 0x110D, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01, Cn: "电池组最大放电电流值"}
	Switching                   = &c_base.SModbusPoint{Name: "Switching", Addr: 0x110F, ReadType: c_base.RInt16, Cn: "开关量指示"}
	BatteryCellMaxVoltage       = &c_base.SModbusPoint{Name: "BatteryCellMaxVoltage", Addr: 0x1110, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Cn: "电池最高电压", Unit: "V"}
	BatteryCellMinVoltage       = &c_base.SModbusPoint{Name: "BatteryCellMinVoltage", Addr: 0x1111, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Cn: "电池最低电压", Unit: "V"}
	BatteryCellMaxVoltageChanel = &c_base.SModbusPoint{Name: "BatteryCellMaxVoltageChanel", Addr: 0x1112, ReadType: c_base.RInt16, Cn: "电池最高电压通道"}
	BatteryCellMinVoltageChanel = &c_base.SModbusPoint{Name: "BatteryCellMinVoltageChanel", Addr: 0x1113, ReadType: c_base.RInt16, Cn: "电池最低电压通道"}
	BatteryCellMaxTemp          = &c_base.SModbusPoint{Name: "BatteryCellMaxTemp", Addr: 0x1114, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池最高温度", Unit: " ̊C"}
	BatteryCellMinTemp          = &c_base.SModbusPoint{Name: "BatteryCellMinTemp", Addr: 0x1115, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池最低温度", Unit: " ̊C"}
	BatteryCellMaxTempChanel    = &c_base.SModbusPoint{Name: "BatteryCellMaxTempChanel", Addr: 0x1116, ReadType: c_base.RInt16, Cn: "电池最高温度通道", Unit: " ̊C"}
	BatteryCellMinTempChanel    = &c_base.SModbusPoint{Name: "BatteryCellMinTempChanel", Addr: 0x1117, ReadType: c_base.RInt16, Cn: "电池最低温度通道", Unit: " ̊C"}
	ModuleMaxVoltage            = &c_base.SModbusPoint{Name: "ModuleMaxVoltage", Addr: 0x1118, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Cn: "模块最高电压", Unit: "V"}
	ModuleMinVoltage            = &c_base.SModbusPoint{Name: "ModuleMinVoltage", Addr: 0x1119, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Cn: "模块最低电压", Unit: "V"}
	ModuleMaxVoltageChanel      = &c_base.SModbusPoint{Name: "ModuleMaxVoltageChanel", Addr: 0x111A, ReadType: c_base.RInt16, Cn: "模块最高电压通道", Unit: "V"}
	ModuleMinVoltageChanel      = &c_base.SModbusPoint{Name: "ModuleMinVoltageChanel", Addr: 0x111B, ReadType: c_base.RInt16, Cn: "模块最低电压通道", Unit: "V"}

	ModuleMaxTemp                                   = &c_base.SModbusPoint{Name: "ModuleMaxTemp", Addr: 0x111C, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "模块最高温度", Unit: " ̊C"}
	ModuleMinTemp                                   = &c_base.SModbusPoint{Name: "ModuleMinTemp", Addr: 0x111D, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "模块最低温度", Unit: " ̊C"}
	ModuleMaxTempChanel                             = &c_base.SModbusPoint{Name: "ModuleMaxTempChanel", Addr: 0x111E, ReadType: c_base.RInt16, Cn: "模块最高温度通道", Unit: " ̊C"}
	ModuleMinTempChanel                             = &c_base.SModbusPoint{Name: "ModuleMinTempChanel", Addr: 0x111F, ReadType: c_base.RInt16, Cn: "模块最低温度通道", Unit: " ̊C"}
	SOH                                             = &c_base.SModbusPoint{Name: "SOH", Addr: 0x1120, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 1, Cn: "电池健康度百分比", Unit: "%"}
	RemainCapacity                                  = &c_base.SModbusPoint{Name: "RemainCapacity", Addr: 0x1121, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Cn: "电池可放电能量", Unit: "kWh"}
	ChargeCapacity                                  = &c_base.SModbusPoint{Name: "ChargeCapacity", Addr: 0x1123, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Cn: "蓄电池充电量", Unit: "kWh"}
	DischargeCapacity                               = &c_base.SModbusPoint{Name: "DischargeCapacity", Addr: 0x1125, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Cn: "蓄电池放电量", Unit: "kWh"}
	TodayCharge                                     = &c_base.SModbusPoint{Name: "TodayCharge", Addr: 0x1127, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Cn: "当日累积充电量", Unit: "kWh"}
	TodayDischarge                                  = &c_base.SModbusPoint{Name: "TodayDischarge", Addr: 0x1129, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Cn: "当日累积放电量", Unit: "kWh"}
	HistoryCharge                                   = &c_base.SModbusPoint{Name: "HistoryCharge", Addr: 0x112B, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 1, Cn: "历史累积充电量", Unit: "kWh"}
	HistoryDischarge                                = &c_base.SModbusPoint{Name: "HistoryDischarge", Addr: 0x112D, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 1, Cn: "历史累积放电量", Unit: "kWh"}
	RequestForceChargeMark                          = &c_base.SModbusPoint{Name: "RequestForceChargeMark", Addr: 0x112F, SystemType: c_base.SBool, Cn: "强充标志位"}
	RequestBalanceChargeMark                        = &c_base.SModbusPoint{Name: "RequestBalanceChargeMark", Addr: 0x1130, SystemType: c_base.SBool, Cn: "均衡充电标志位"}
	PileNumberInParallel                            = &c_base.SModbusPoint{Name: "PileNumberInParallel", Addr: 0x1131, ReadType: c_base.RInt16, Cn: "当前并联电池组数量"}
	ErrorCode1                                      = &c_base.SModbusPoint{Name: "ErrorCode1", Addr: 0x1132, ReadType: c_base.RInt16, Cn: "故障代码1", Trigger: c_func.IsNotZero}
	ErrorCode2                                      = &c_base.SModbusPoint{Name: "ErrorCode2", Addr: 0x1134, ReadType: c_base.RInt16, Cn: "故障代码2", Trigger: c_func.IsNotZero}
	NumberBatteryModulesInSeriesConnectionInOnePile = &c_base.SModbusPoint{Name: "NumberBatteryModulesInSeriesConnectionInOnePile", Addr: 0x1136, ReadType: c_base.RInt16, Cn: "电池组模块串联数"}
	NumberCellInSeriesInOnePile                     = &c_base.SModbusPoint{Name: "NumberCellInSeriesInOnePile", Addr: 0x1137, ReadType: c_base.RInt16, Cn: "电池组单体串联数"}

	ChargeForbiddenMark            = &c_base.SModbusPoint{Name: "ChargeForbiddenMark", Addr: 0x1138, ReadType: c_base.RUint16, SystemType: c_base.SBool, Cn: "禁止充电标志"}
	DischargeForbiddenMark         = &c_base.SModbusPoint{Name: "DischargeForbiddenMark", Addr: 0x1139, ReadType: c_base.RUint16, SystemType: c_base.SBool, Cn: "禁止放电标志"}
	SOC30Flag                      = &c_base.SModbusPoint{Name: "SOC30Flag", Addr: 0x113A, SystemType: c_base.SBool, Cn: "SOC<=30%标志"}
	SOE                            = &c_base.SModbusPoint{Name: "SOE", Addr: 0x113B, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 1, Cn: "SOE", Unit: "%"}
	HeartbeatSignal                = &c_base.SModbusPoint{Name: "HeartbeatSignal", Addr: 0x113C, ReadType: c_base.RInt16, Cn: "心跳信号值", Desc: ",范围为 0x0000~0x00FF，每次读取该值都会增加 1"}
	ModulePCBMaxTemp               = &c_base.SModbusPoint{Name: "ModulePCBMaxTemp", Addr: 0x113D, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池模块单板最高温度", Unit: " ̊C"}
	ModulePCBMinTemp               = &c_base.SModbusPoint{Name: "ModulePCBMinTemp", Addr: 0x113E, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "电池模块单板最低温度", Unit: " ̊C"}
	ModulePCBMaxTempChanel         = &c_base.SModbusPoint{Name: "ModulePCBMaxTempChanel", Addr: 0x113F, ReadType: c_base.RInt16, Cn: "电池模块单板最高温度通道", Unit: " ̊C"}
	ModulePCBMinTempChanel         = &c_base.SModbusPoint{Name: "ModulePCBMinTempChanel", Addr: 0x1140, ReadType: c_base.RInt16, Cn: "电池模块单板最低温度通道", Unit: " ̊C"}
	SystemOperationStatus          = &c_base.SModbusPoint{Name: "SystemOperationStatus", Addr: 0x1141, ReadType: c_base.RInt16, Cn: "系统运行状态Linked with 0x1094.\n0x11:Standby(self-inspection is \nfinished but relay didn`t close, \nrequest to command the system at \n0x1094 to close relay, the status will \nthen switch to ‘Run’)\n0x22:Run(close relay)"}
	InsulationResistance           = &c_base.SModbusPoint{Name: "InsulationResistance", Addr: 0x1148, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Cn: "绝缘电阻值", Unit: "KΩ"}
	InsulationResistanceErrorLevel = &c_base.SModbusPoint{Name: "InsulationResistanceErrorLevel", Addr: 0x1149, ReadType: c_base.RInt16, Cn: "绝缘电阻故障等级"}
	TerminalMaxTemp                = &c_base.SModbusPoint{Name: "TerminalMaxTemp", Addr: 0x114A, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "端子最高温度", Unit: " ̊C"}
	TerminalMinTemp                = &c_base.SModbusPoint{Name: "TerminalMinTemp", Addr: 0x114B, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "端子最低温度", Unit: " ̊C"}

	TerminalMaxTempChannel = &c_base.SModbusPoint{Name: "TerminalMaxTempChannel", Addr: 0x114C, ReadType: c_base.RInt16, Cn: "端子最高温度通道", Unit: " ̊C"}
	TerminalMinTempChannel = &c_base.SModbusPoint{Name: "TerminalMinTempChannel", Addr: 0x114D, ReadType: c_base.RInt16, Cn: "端子最低温度通道", Unit: " ̊C"}
	AlarmStatus2           = &c_base.SModbusPoint{Name: "AlarmStatus2", Addr: 0x114E, ReadType: c_base.RInt16, Cn: "告警状态2", Trigger: c_func.IsNotZero}
	MaxChargePower         = &c_base.SModbusPoint{Name: "MaxChargePower", Addr: 0x1150, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "20 分钟内最大持续充电功率预测值", Unit: "kW"}
	MaxDischargePower      = &c_base.SModbusPoint{Name: "MaxDischargePower", Addr: 0x1151, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Cn: "20 分钟内最大持续放电功率预测值", Unit: "kW"}
)

var (
	External485Protocol = &c_base.SModbusPoint{Name: "External485Protocol", Addr: 0x12F0, ReadType: c_base.RInt16, Cn: "对外 485 协议", Desc: ";0x01：pylon modbus\n0x02：无效\n0x03：kostal\n0x04：kaco\nothers：无效"}
	ExternalCanProtocol = &c_base.SModbusPoint{Name: "ExternalCanProtocol", Addr: 0x12F1, ReadType: c_base.RInt16, Cn: "对外 can 协议", Desc: ";0x01：auto\n0x02：pylon can\n0x03：solax\n0x04：sma\nothers：无效"}
	InputDryContact1    = &c_base.SModbusPoint{Name: "InputDryContact1", Addr: 0x12F2, ReadType: c_base.RInt16, Cn: "输入干接点", Desc: " 1;0x0：失能\n0x1：急停功能，\nOthers：无效\n注意:该点位当前只支持 FH 系列，其它\n产品无效"}
)
