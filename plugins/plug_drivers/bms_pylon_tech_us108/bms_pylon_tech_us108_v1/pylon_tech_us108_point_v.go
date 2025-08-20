package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_util"
)

var (
	StatusGroup = &c_base.MetaGroup{GroupName: "状态", GroupSort: 1}
)

// 设备信息
var (
// Pylon            = &c_base.Meta{Id: "Pylon", Addr: 0x1000, BitSize: 16 * 5, SystemType: c_base.SString, Desc: "设备信息"}
// MBms             = &c_base.Meta{Id: "MBms", Addr: 0x1005, BitSize: 16 * 5, SystemType: c_base.SString, Desc: "电池信息"}
// Version          = &c_base.Meta{Id: "Version", Addr: 0x100A,  ReadType: c_base.RInt16, Desc: "版本信息"}
// InternalVersion  = &c_base.Meta{Id: "InternalVersion", Addr: 0x100B,  ReadType: c_base.RInt16, Desc: "内部版本信息"}
// ParallelPilesQty = &c_base.Meta{Id: "ParallelPilesQty", Addr: 0x100C,  ReadType: c_base.RInt16, Desc: "并联电池数量"}
)

// 控制信息
var (
	SleepControl   = &c_base.Meta{Name: "SleepControl", Addr: 0x1090, ReadType: c_base.RInt16, Desc: "休眠控制, 写入0xAA进入休眠状态，写入0x55唤醒"}
	AllowCharge    = &c_base.Meta{Name: "AllowCharge", Addr: 0x1091, ReadType: c_base.RInt16, Desc: "允许充电控制, 写入0xAA允许充电，其他无效"}
	AllowDischarge = &c_base.Meta{Name: "AllowDischarge", Addr: 0x1092, ReadType: c_base.RInt16, Desc: "允许放电控制, 写入0xAA允许放电，其他无效"}
	TempShield     = &c_base.Meta{Name: "TempShield", Addr: 0x1093, ReadType: c_base.RInt16, Desc: "临时屏蔽“无通信时切离继电器功能”的请求"}
	AllowRun       = &c_base.Meta{Name: "AllowRun", Addr: 0x1094, ReadType: c_base.RInt16, Desc: "运行指令，通知系统开始并柜, 写入0xAA允许运行，其他无效"}
)

// 时间
var (
	Year   = &c_base.Meta{Name: "Year", Addr: 0x10E0, ReadType: c_base.RUint16, Desc: "年 00~99（表示 2000~2099）"}
	Month  = &c_base.Meta{Name: "Month", Addr: 0x10E1, ReadType: c_base.RUint16, Desc: "月 1~12"}
	Day    = &c_base.Meta{Name: "Day", Addr: 0x10E2, ReadType: c_base.RUint16, Desc: "日 1~31"}
	Hour   = &c_base.Meta{Name: "Hour", Addr: 0x10E3, ReadType: c_base.RUint16, Desc: "时 0~23"}
	Minute = &c_base.Meta{Name: "Minute", Addr: 0x10E4, ReadType: c_base.RUint16, Desc: "分 0~59"}
	Second = &c_base.Meta{Name: "Second", Addr: 0x10E5, ReadType: c_base.RUint16, Desc: "秒 0~59"}
)

var (
	//BasicStatus   = &c_base.Meta{Id: "BasicStatus", Cn: "基本状态", Addr: 0x1100, ReadType: c_base.RInt16, Desc: "基本状态"}
	BasicStatus = &c_base.Meta{Name: "BasicStatus", Cn: "基本状态", Addr: 0x1100, ReadType: c_base.RBit0, BitLength: 2, Desc: "基本状态:00：休眠 01：充电 02：放电 03：搁置", StatusExplain: func(value any) string {
		if v, err := c_util.ToInt8(value); err == nil {
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
		return "未知"
	}}
	SystemErrorProtection     = &c_base.Meta{Name: "SystemErrorProtection", Group: StatusGroup, Cn: "系统故障保护", Addr: 0x1100, ReadType: c_base.RBit3, Level: c_base.EError, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	CurrentProtection         = &c_base.Meta{Name: "CurrentProtection", Group: StatusGroup, Cn: "电流保护", Addr: 0x1100, ReadType: c_base.RBit4, Level: c_base.EError, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	VoltageProtection         = &c_base.Meta{Name: "VoltageProtection", Group: StatusGroup, Cn: "电压保护", Addr: 0x1100, ReadType: c_base.RBit5, Level: c_base.EError, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	TemperatureProtection     = &c_base.Meta{Name: "TemperatureProtection", Group: StatusGroup, Cn: "温度保护", Addr: 0x1100, ReadType: c_base.RBit6, Level: c_base.EError, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	VoltageAlarm              = &c_base.Meta{Name: "VoltageAlarm", Group: StatusGroup, Cn: "电压警告", Addr: 0x1100, ReadType: c_base.RBit7, Level: c_base.EAlarm, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	CurrentAlarm              = &c_base.Meta{Name: "VoltageAlarm", Group: StatusGroup, Cn: "电流警告", Addr: 0x1100, ReadType: c_base.RBit8, Level: c_base.EAlarm, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	TemperatureAlarm          = &c_base.Meta{Name: "VoltageAlarm", Group: StatusGroup, Cn: "温度警告", Addr: 0x1100, ReadType: c_base.RBit9, Level: c_base.EAlarm, Desc: "0-正常；1-保护", StatusExplain: c_base.StatusExplainProtectFunc}
	PileSystemIdleStatus      = &c_base.Meta{Name: "PileSystemIdleStatus", Group: StatusGroup, Cn: "电池组搁置状态", Addr: 0x1100, ReadType: c_base.RBit10, Level: c_base.EAlarm, Desc: "0-否，1-搁置"}
	PileSystemChargeStatus    = &c_base.Meta{Name: "PileSystemChargeStatus", Group: StatusGroup, Cn: "电池组充电状态", Addr: 0x1100, ReadType: c_base.RBit11, Level: c_base.EAlarm, Desc: "0-否，1-充电"}
	PileSystemDischargeStatus = &c_base.Meta{Name: "PileSystemDischargeStatus", Group: StatusGroup, Cn: "电池组放电状态", Addr: 0x1100, ReadType: c_base.RBit12, Level: c_base.EAlarm, Desc: "0-否，1-充电"}
	PileSystemSleepStatus     = &c_base.Meta{Name: "PileSystemSleepStatus", Group: StatusGroup, Cn: "电池组休眠状态", Addr: 0x1100, ReadType: c_base.RBit13, Level: c_base.EAlarm, Desc: "0-否，1-休眠"}
	FanWarn                   = &c_base.Meta{Name: "FanWarn", Group: StatusGroup, Cn: "电池组休眠状态", Addr: 0x1100, ReadType: c_base.RBit14, Level: c_base.EAlarm, Desc: "0-否，1-有异常"}

	Protection                  = &c_base.Meta{Name: "Protection", Addr: 0x1101, ReadType: c_base.RInt16, Desc: "保护状态"}
	AlarmStatus1                = &c_base.Meta{Name: "AlarmStatus1", Addr: 0x1102, ReadType: c_base.RInt16, Desc: "告警状态1", Trigger: c_base.IsNotZero}
	DCVoltage                   = &c_base.Meta{Name: "DCVoltage", Addr: 0x1103, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "总电压", Unit: "V"}
	DCCurrent                   = &c_base.Meta{Name: "DCCurrent", Addr: 0x1104, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01, Desc: "电流", Unit: "A"}
	Temperature                 = &c_base.Meta{Name: "Temperature", Addr: 0x1106, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "温度", Unit: " ̊C"}
	SOC                         = &c_base.Meta{Name: "SOC", Addr: 0x1107, ReadType: c_base.RUint16, Factor: 1, Desc: "SOC", Unit: "%"}
	CycleCount                  = &c_base.Meta{Name: "CycleCount", Addr: 0x1108, ReadType: c_base.RUint16, Desc: "循环次数"}
	PileMaxV                    = &c_base.Meta{Name: "PileMaxV", Addr: 0x1109, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池组最大充电电压值", Unit: "V"}
	PileMaxI                    = &c_base.Meta{Name: "PileMaxI", Addr: 0x110A, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01, Desc: "电池组最大充电电流值"}
	PileMinV                    = &c_base.Meta{Name: "PileMinV", Addr: 0x110C, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池组最小放电电压值", Unit: "V"}
	PileMaxDI                   = &c_base.Meta{Name: "PileMaxDI", Addr: 0x110D, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01, Desc: "电池组最大放电电流值"}
	Switching                   = &c_base.Meta{Name: "Switching", Addr: 0x110F, ReadType: c_base.RInt16, Desc: "开关量指示"}
	BatteryCellMaxVoltage       = &c_base.Meta{Name: "BatteryCellMaxVoltage", Addr: 0x1110, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Desc: "电池最高电压", Unit: "V"}
	BatteryCellMinVoltage       = &c_base.Meta{Name: "BatteryCellMinVoltage", Addr: 0x1111, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Desc: "电池最低电压", Unit: "V"}
	BatteryCellMaxVoltageChanel = &c_base.Meta{Name: "BatteryCellMaxVoltageChanel", Addr: 0x1112, ReadType: c_base.RInt16, Desc: "电池最高电压通道"}
	BatteryCellMinVoltageChanel = &c_base.Meta{Name: "BatteryCellMinVoltageChanel", Addr: 0x1113, ReadType: c_base.RInt16, Desc: "电池最低电压通道"}
	BatteryCellMaxTemp          = &c_base.Meta{Name: "BatteryCellMaxTemp", Addr: 0x1114, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池最高温度", Unit: " ̊C"}
	BatteryCellMinTemp          = &c_base.Meta{Name: "BatteryCellMinTemp", Addr: 0x1115, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池最低温度", Unit: " ̊C"}
	BatteryCellMaxTempChanel    = &c_base.Meta{Name: "BatteryCellMaxTempChanel", Addr: 0x1116, ReadType: c_base.RInt16, Desc: "电池最高温度通道", Unit: " ̊C"}
	BatteryCellMinTempChanel    = &c_base.Meta{Name: "BatteryCellMinTempChanel", Addr: 0x1117, ReadType: c_base.RInt16, Desc: "电池最低温度通道", Unit: " ̊C"}
	ModuleMaxVoltage            = &c_base.Meta{Name: "ModuleMaxVoltage", Addr: 0x1118, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Desc: "模块最高电压", Unit: "V"}
	ModuleMinVoltage            = &c_base.Meta{Name: "ModuleMinVoltage", Addr: 0x1119, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.01, Desc: "模块最低电压", Unit: "V"}
	ModuleMaxVoltageChanel      = &c_base.Meta{Name: "ModuleMaxVoltageChanel", Addr: 0x111A, ReadType: c_base.RInt16, Desc: "模块最高电压通道", Unit: "V"}
	ModuleMinVoltageChanel      = &c_base.Meta{Name: "ModuleMinVoltageChanel", Addr: 0x111B, ReadType: c_base.RInt16, Desc: "模块最低电压通道", Unit: "V"}

	ModuleMaxTemp                                   = &c_base.Meta{Name: "ModuleMaxTemp", Addr: 0x111C, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "模块最高温度", Unit: " ̊C"}
	ModuleMinTemp                                   = &c_base.Meta{Name: "ModuleMinTemp", Addr: 0x111D, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "模块最低温度", Unit: " ̊C"}
	ModuleMaxTempChanel                             = &c_base.Meta{Name: "ModuleMaxTempChanel", Addr: 0x111E, ReadType: c_base.RInt16, Desc: "模块最高温度通道", Unit: " ̊C"}
	ModuleMinTempChanel                             = &c_base.Meta{Name: "ModuleMinTempChanel", Addr: 0x111F, ReadType: c_base.RInt16, Desc: "模块最低温度通道", Unit: " ̊C"}
	SOH                                             = &c_base.Meta{Name: "SOH", Addr: 0x1120, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 1, Desc: "电池健康度百分比", Unit: "%"}
	RemainCapacity                                  = &c_base.Meta{Name: "RemainCapacity", Addr: 0x1121, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Desc: "电池可放电能量", Unit: "kWh"}
	ChargeCapacity                                  = &c_base.Meta{Name: "ChargeCapacity", Addr: 0x1123, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Desc: "蓄电池充电量", Unit: "kWh"}
	DischargeCapacity                               = &c_base.Meta{Name: "DischargeCapacity", Addr: 0x1125, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Desc: "蓄电池放电量", Unit: "kWh"}
	TodayCharge                                     = &c_base.Meta{Name: "TodayCharge", Addr: 0x1127, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Desc: "当日累积充电量", Unit: "kWh"}
	TodayDischarge                                  = &c_base.Meta{Name: "TodayDischarge", Addr: 0x1129, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 0.001, Desc: "当日累积放电量", Unit: "kWh"}
	HistoryCharge                                   = &c_base.Meta{Name: "HistoryCharge", Addr: 0x112B, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 1, Desc: "历史累积充电量", Unit: "kWh"}
	HistoryDischarge                                = &c_base.Meta{Name: "HistoryDischarge", Addr: 0x112D, ReadType: c_base.RUint32, SystemType: c_base.SFloat64, Factor: 1, Desc: "历史累积放电量", Unit: "kWh"}
	RequestForceChargeMark                          = &c_base.Meta{Name: "RequestForceChargeMark", Addr: 0x112F, SystemType: c_base.SBool, Desc: "强充标志位"}
	RequestBalanceChargeMark                        = &c_base.Meta{Name: "RequestBalanceChargeMark", Addr: 0x1130, SystemType: c_base.SBool, Desc: "均衡充电标志位"}
	PileNumberInParallel                            = &c_base.Meta{Name: "PileNumberInParallel", Addr: 0x1131, ReadType: c_base.RInt16, Desc: "当前并联电池组数量"}
	ErrorCode1                                      = &c_base.Meta{Name: "ErrorCode1", Addr: 0x1132, ReadType: c_base.RInt16, Desc: "故障代码1", Trigger: c_base.IsNotZero}
	ErrorCode2                                      = &c_base.Meta{Name: "ErrorCode2", Addr: 0x1134, ReadType: c_base.RInt16, Desc: "故障代码2", Trigger: c_base.IsNotZero}
	NumberBatteryModulesInSeriesConnectionInOnePile = &c_base.Meta{Name: "NumberBatteryModulesInSeriesConnectionInOnePile", Addr: 0x1136, ReadType: c_base.RInt16, Desc: "电池组模块串联数"}
	NumberCellInSeriesInOnePile                     = &c_base.Meta{Name: "NumberCellInSeriesInOnePile", Addr: 0x1137, ReadType: c_base.RInt16, Desc: "电池组单体串联数"}

	ChargeForbiddenMark            = &c_base.Meta{Name: "ChargeForbiddenMark", Addr: 0x1138, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "禁止充电标志"}
	DischargeForbiddenMark         = &c_base.Meta{Name: "DischargeForbiddenMark", Addr: 0x1139, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "禁止放电标志"}
	SOC30Flag                      = &c_base.Meta{Name: "SOC30Flag", Addr: 0x113A, SystemType: c_base.SBool, Desc: "SOC<=30%标志"}
	SOE                            = &c_base.Meta{Name: "SOE", Addr: 0x113B, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 1, Desc: "SOE", Unit: "%"}
	HeartbeatSignal                = &c_base.Meta{Name: "HeartbeatSignal", Addr: 0x113C, ReadType: c_base.RInt16, Desc: "心跳信号值,范围为 0x0000~0x00FF，每次读取该值都会增加 1"}
	ModulePCBMaxTemp               = &c_base.Meta{Name: "ModulePCBMaxTemp", Addr: 0x113D, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池模块单板最高温度", Unit: " ̊C"}
	ModulePCBMinTemp               = &c_base.Meta{Name: "ModulePCBMinTemp", Addr: 0x113E, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "电池模块单板最低温度", Unit: " ̊C"}
	ModulePCBMaxTempChanel         = &c_base.Meta{Name: "ModulePCBMaxTempChanel", Addr: 0x113F, ReadType: c_base.RInt16, Desc: "电池模块单板最高温度通道", Unit: " ̊C"}
	ModulePCBMinTempChanel         = &c_base.Meta{Name: "ModulePCBMinTempChanel", Addr: 0x1140, ReadType: c_base.RInt16, Desc: "电池模块单板最低温度通道", Unit: " ̊C"}
	SystemOperationStatus          = &c_base.Meta{Name: "SystemOperationStatus", Addr: 0x1141, ReadType: c_base.RInt16, Desc: "系统运行状态Linked with 0x1094.\n0x11:Standby(self-inspection is \nfinished but relay didn`t close, \nrequest to command the system at \n0x1094 to close relay, the status will \nthen switch to ‘Run’)\n0x22:Run(close relay)"}
	InsulationResistance           = &c_base.Meta{Name: "InsulationResistance", Addr: 0x1148, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Desc: "绝缘电阻值", Unit: "KΩ"}
	InsulationResistanceErrorLevel = &c_base.Meta{Name: "InsulationResistanceErrorLevel", Addr: 0x1149, ReadType: c_base.RInt16, Desc: "绝缘电阻故障等级"}
	TerminalMaxTemp                = &c_base.Meta{Name: "TerminalMaxTemp", Addr: 0x114A, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "端子最高温度", Unit: " ̊C"}
	TerminalMinTemp                = &c_base.Meta{Name: "TerminalMinTemp", Addr: 0x114B, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "端子最低温度", Unit: " ̊C"}

	TerminalMaxTempChannel = &c_base.Meta{Name: "TerminalMaxTempChannel", Addr: 0x114C, ReadType: c_base.RInt16, Desc: "端子最高温度通道", Unit: " ̊C"}
	TerminalMinTempChannel = &c_base.Meta{Name: "TerminalMinTempChannel", Addr: 0x114D, ReadType: c_base.RInt16, Desc: "端子最低温度通道", Unit: " ̊C"}
	AlarmStatus2           = &c_base.Meta{Name: "AlarmStatus2", Addr: 0x114E, ReadType: c_base.RInt16, Desc: "告警状态2", Trigger: c_base.IsNotZero}
	MaxChargePower         = &c_base.Meta{Name: "MaxChargePower", Addr: 0x1150, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "20 分钟内最大持续充电功率预测值", Unit: "kW"}
	MaxDischargePower      = &c_base.Meta{Name: "MaxDischargePower", Addr: 0x1151, ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.1, Desc: "20 分钟内最大持续放电功率预测值", Unit: "kW"}
)

var (
	External485Protocol = &c_base.Meta{Name: "External485Protocol", Addr: 0x12F0, ReadType: c_base.RInt16, Desc: "对外 485 协议;0x01：pylon modbus\n0x02：无效\n0x03：kostal\n0x04：kaco\nothers：无效"}
	ExternalCanProtocol = &c_base.Meta{Name: "ExternalCanProtocol", Addr: 0x12F1, ReadType: c_base.RInt16, Desc: "对外 can 协议;0x01：auto\n0x02：pylon can\n0x03：solax\n0x04：sma\nothers：无效"}
	InputDryContact1    = &c_base.Meta{Name: "InputDryContact1", Addr: 0x12F2, ReadType: c_base.RInt16, Desc: "输入干接点 1;0x0：失能\n0x1：急停功能，\nOthers：无效\n注意:该点位当前只支持 FH 系列，其它\n产品无效"}
)
