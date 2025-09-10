package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_func"
	"common/c_proto"
	"fmt"

	"github.com/shockerli/cvt"
)

var (
	StatusGroup = &c_base.SPointGroup{GroupName: "状态", GroupSort: 1}
)

// 设备信息
var (
// Pylon            = &c_proto.SModbusPoint{Addr: 0x1000, Id: "Pylon", BitSize: 16 * 5, SystemType: c_base.SString, Desc: "设备信息"}
// MBms             = &c_proto.SModbusPoint{Addr: 0x1005, Id: "MBms", BitSize: 16 * 5, SystemType: c_base.SString, Desc: "电池信息"}
// Version          = &c_proto.SModbusPoint{Addr: 0x100A, Id: "Version", ReadType: c_base.RInt16, Desc: "版本信息"}
// InternalVersion  = &c_proto.SModbusPoint{Addr: 0x100B, Id: "InternalVersion", ReadType: c_base.RInt16, Desc: "内部版本信息"}
// ParallelPilesQty = &c_proto.SModbusPoint{Addr: 0x100C, Id: "ParallelPilesQty", ReadType: c_base.RInt16, Desc: "并联电池数量"}
)

// 控制信息
var (
	SleepControl   = &c_proto.SModbusPoint{Addr: 0x1090, SPoint: &c_base.SPoint{Key: "SleepControl", Name: "SleepControl", Desc: "休眠控制, 写入0xAA进入休眠状态，写入0x55唤醒"}, DataAccess: c_default.VDataAccessInt16}
	AllowCharge    = &c_proto.SModbusPoint{Addr: 0x1091, SPoint: &c_base.SPoint{Key: "AllowCharge", Name: "AllowCharge", Desc: "允许充电控制, 写入0xAA允许充电，其他无效"}, DataAccess: c_default.VDataAccessInt16}
	AllowDischarge = &c_proto.SModbusPoint{Addr: 0x1092, SPoint: &c_base.SPoint{Key: "AllowDischarge", Name: "AllowDischarge", Desc: "允许放电控制, 写入0xAA允许放电，其他无效"}, DataAccess: c_default.VDataAccessInt16}
	TempShield     = &c_proto.SModbusPoint{Addr: 0x1093, SPoint: &c_base.SPoint{Key: "TempShield", Name: "TempShield", Desc: "临时屏蔽\"无通信时切离继电器功能\"的请求"}, DataAccess: c_default.VDataAccessInt16}
	AllowRun       = &c_proto.SModbusPoint{Addr: 0x1094, SPoint: &c_base.SPoint{Key: "AllowRun", Name: "AllowRun", Desc: "运行指令，通知系统开始并柜, 写入0xAA允许运行，其他无效"}, DataAccess: c_default.VDataAccessInt16}
)

// 时间
var (
	Year   = &c_proto.SModbusPoint{Addr: 0x10E0, SPoint: &c_base.SPoint{Key: "Year", Name: "Year", Desc: "年 00~99（表示 2000~2099）"}, DataAccess: c_default.VDataAccessUInt16}
	Month  = &c_proto.SModbusPoint{Addr: 0x10E1, SPoint: &c_base.SPoint{Key: "Month", Name: "Month", Desc: "月 1~12"}, DataAccess: c_default.VDataAccessUInt16}
	Day    = &c_proto.SModbusPoint{Addr: 0x10E2, SPoint: &c_base.SPoint{Key: "Day", Name: "Day", Desc: "日 1~31"}, DataAccess: c_default.VDataAccessUInt16}
	Hour   = &c_proto.SModbusPoint{Addr: 0x10E3, SPoint: &c_base.SPoint{Key: "Hour", Name: "Hour", Desc: "时 0~23"}, DataAccess: c_default.VDataAccessUInt16}
	Minute = &c_proto.SModbusPoint{Addr: 0x10E4, SPoint: &c_base.SPoint{Key: "Minute", Name: "Minute", Desc: "分 0~59"}, DataAccess: c_default.VDataAccessUInt16}
	Second = &c_proto.SModbusPoint{Addr: 0x10E5, SPoint: &c_base.SPoint{Key: "Second", Name: "Second", Desc: "秒 0~59"}, DataAccess: c_default.VDataAccessUInt16}
)

var (
	BasicStatus = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "BasicStatus", Name: "基本状态", Desc: "基本状态:00：休眠 01：充电 02：放电 03：搁置"}, DataAccess: &c_base.SDataAccess{BitIndex: 0, BitLength: 2, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16}, StatusExplain: func(value any) (string, error) {
		if v, err := cvt.Int8E(value); err == nil {
			switch v {
			case 0:
				return "休眠", nil
			case 1:
				return "充电", nil
			case 2:
				return "放电", nil
			case 3:
				return "搁置", nil
			}
		}
		return fmt.Sprintf("未知值: %v", value), nil
	}}
	SystemErrorProtection = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "SystemErrorProtection", Name: "系统故障保护", Group: StatusGroup, Trigger: c_default.FAlarmTriggerErrorBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 3, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	CurrentProtection     = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "CurrentProtection", Name: "电流保护", Group: StatusGroup, Trigger: c_default.FAlarmTriggerErrorBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 4, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	VoltageProtection     = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "VoltageProtection", Name: "电压保护", Group: StatusGroup, Trigger: c_default.FAlarmTriggerErrorBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 5, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	TemperatureProtection = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "TemperatureProtection", Name: "温度保护", Group: StatusGroup, Trigger: c_default.FAlarmTriggerErrorBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 6, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	VoltageAlarm          = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "VoltageAlarm", Name: "电压警告", Group: StatusGroup, Trigger: c_default.FAlarmTriggerAlertBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 7, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	CurrentAlarm          = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "CurrentAlarm", Name: "电流警告", Group: StatusGroup, Trigger: c_default.FAlarmTriggerAlertBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 8, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	TemperatureAlarm      = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "TemperatureAlarm", Name: "温度警告", Group: StatusGroup, Trigger: c_default.FAlarmTriggerAlertBool, Desc: "0-正常；1-保护"}, DataAccess: &c_base.SDataAccess{BitIndex: 9, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: c_func.StatusExplainProtectFunc}
	PileSystemIdleStatus  = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "PileSystemIdleStatus", Name: "电池组搁置状态", Group: StatusGroup, Desc: "0-否，1-搁置"}, DataAccess: &c_base.SDataAccess{BitIndex: 10, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: func(value any) (string, error) {
		return c_func.BoolExplain(value, "搁置", "否")
	}}
	PileSystemChargeStatus = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "PileSystemChargeStatus", Name: "电池组充电状态", Group: StatusGroup, Desc: "0-否，1-充电"}, DataAccess: &c_base.SDataAccess{BitIndex: 11, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: func(value any) (string, error) {
		return c_func.BoolExplain(value, "充电", "否")
	}}
	PileSystemDischargeStatus = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "PileSystemDischargeStatus", Name: "电池组放电状态", Group: StatusGroup, Desc: "0-否，1-放电"}, DataAccess: &c_base.SDataAccess{BitIndex: 12, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: func(value any) (string, error) {
		return c_func.BoolExplain(value, "放电", "否")
	}}
	PileSystemSleepStatus = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "PileSystemSleepStatus", Name: "电池组休眠状态", Group: StatusGroup, Desc: "0-否，1-休眠"}, DataAccess: &c_base.SDataAccess{BitIndex: 13, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: func(value any) (string, error) {
		return c_func.BoolExplain(value, "休眠", "否")
	}}
	FanWarn = &c_proto.SModbusPoint{Addr: 0x1100, SPoint: &c_base.SPoint{Key: "FanWarn", Name: "电池组休眠状态", Group: StatusGroup, Trigger: c_default.FAlarmTriggerAlertBool, Desc: "0-无异常，1-有异常"}, DataAccess: &c_base.SDataAccess{BitIndex: 14, BitLength: 1, DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}, StatusExplain: func(value any) (string, error) {
		return c_func.BoolExplain(value, "有异常", "无异常")
	}}

	Protection                  = &c_proto.SModbusPoint{Addr: 0x1101, SPoint: &c_base.SPoint{Key: "Protection", Name: "保护状态"}, DataAccess: c_default.VDataAccessInt16}
	AlarmStatus1                = &c_proto.SModbusPoint{Addr: 0x1102, SPoint: &c_base.SPoint{Key: "AlarmStatus1", Name: "告警状态1"}, DataAccess: c_default.VDataAccessInt16, Trigger: c_default.FAlarmTriggerAlertNotZero}
	DCVoltage                   = &c_proto.SModbusPoint{Addr: 0x1103, SPoint: c_default.VPointDcVoltage, DataAccess: c_default.VDataAccessUInt16Scale01}
	DCCurrent                   = &c_proto.SModbusPoint{Addr: 0x1104, SPoint: c_default.VPointDcCurrent, DataAccess: c_default.VDataAccessInt32Scale001}
	Temperature                 = &c_proto.SModbusPoint{Addr: 0x1106, SPoint: c_default.VPointTemp, DataAccess: c_default.VDataAccessUInt16Scale01}
	SOC                         = &c_proto.SModbusPoint{Addr: 0x1107, SPoint: c_default.VPointSOC, DataAccess: c_default.VDataAccessUInt16}
	CycleCount                  = &c_proto.SModbusPoint{Addr: 0x1108, SPoint: &c_base.SPoint{Key: "CycleCount", Name: "循环次数"}, DataAccess: c_default.VDataAccessUInt16}
	PileMaxV                    = &c_proto.SModbusPoint{Addr: 0x1109, SPoint: &c_base.SPoint{Key: "PileMaxV", Name: "电池组最大充电电压值", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PileMaxI                    = &c_proto.SModbusPoint{Addr: 0x110A, SPoint: &c_base.SPoint{Key: "PileMaxI", Name: "电池组最大充电电流值"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	PileMinV                    = &c_proto.SModbusPoint{Addr: 0x110C, SPoint: &c_base.SPoint{Key: "PileMinV", Name: "电池组最小放电电压值", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	PileMaxDI                   = &c_proto.SModbusPoint{Addr: 0x110D, SPoint: &c_base.SPoint{Key: "PileMaxDI", Name: "电池组最大放电电流值"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	Switching                   = &c_proto.SModbusPoint{Addr: 0x110F, SPoint: &c_base.SPoint{Key: "Switching", Name: "开关量指示"}, DataAccess: c_default.VDataAccessInt16}
	BatteryCellMaxVoltage       = &c_proto.SModbusPoint{Addr: 0x1110, SPoint: &c_base.SPoint{Key: "BatteryCellMaxVoltage", Name: "电池最高电压", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale0001}
	BatteryCellMinVoltage       = &c_proto.SModbusPoint{Addr: 0x1111, SPoint: &c_base.SPoint{Key: "BatteryCellMinVoltage", Name: "电池最低电压", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale0001}
	BatteryCellMaxVoltageChanel = &c_proto.SModbusPoint{Addr: 0x1112, SPoint: &c_base.SPoint{Key: "BatteryCellMaxVoltageChanel", Name: "电池最高电压通道"}, DataAccess: c_default.VDataAccessInt16}
	BatteryCellMinVoltageChanel = &c_proto.SModbusPoint{Addr: 0x1113, SPoint: &c_base.SPoint{Key: "BatteryCellMinVoltageChanel", Name: "电池最低电压通道"}, DataAccess: c_default.VDataAccessInt16}
	BatteryCellMaxTemp          = &c_proto.SModbusPoint{Addr: 0x1114, SPoint: c_default.VPointTempMax, DataAccess: c_default.VDataAccessUInt16Scale01}
	BatteryCellMinTemp          = &c_proto.SModbusPoint{Addr: 0x1115, SPoint: c_default.VPointTempMin, DataAccess: c_default.VDataAccessUInt16Scale01}
	BatteryCellMaxTempChanel    = &c_proto.SModbusPoint{Addr: 0x1116, SPoint: &c_base.SPoint{Key: "BatteryCellMaxTempChanel", Name: "电池最高温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	BatteryCellMinTempChanel    = &c_proto.SModbusPoint{Addr: 0x1117, SPoint: &c_base.SPoint{Key: "BatteryCellMinTempChanel", Name: "电池最低温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	ModuleMaxVoltage            = &c_proto.SModbusPoint{Addr: 0x1118, SPoint: &c_base.SPoint{Key: "ModuleMaxVoltage", Name: "模块最高电压", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	ModuleMinVoltage            = &c_proto.SModbusPoint{Addr: 0x1119, SPoint: &c_base.SPoint{Key: "ModuleMinVoltage", Name: "模块最低电压", Unit: "V"}, DataAccess: c_default.VDataAccessUInt16Scale001}
	ModuleMaxVoltageChanel      = &c_proto.SModbusPoint{Addr: 0x111A, SPoint: &c_base.SPoint{Key: "ModuleMaxVoltageChanel", Name: "模块最高电压通道", Unit: "V"}, DataAccess: c_default.VDataAccessInt16}
	ModuleMinVoltageChanel      = &c_proto.SModbusPoint{Addr: 0x111B, SPoint: &c_base.SPoint{Key: "ModuleMinVoltageChanel", Name: "模块最低电压通道", Unit: "V"}, DataAccess: c_default.VDataAccessInt16}

	ModuleMaxTemp                                   = &c_proto.SModbusPoint{Addr: 0x111C, SPoint: &c_base.SPoint{Key: "ModuleMaxTemp", Name: "模块最高温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	ModuleMinTemp                                   = &c_proto.SModbusPoint{Addr: 0x111D, SPoint: &c_base.SPoint{Key: "ModuleMinTemp", Name: "模块最低温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	ModuleMaxTempChanel                             = &c_proto.SModbusPoint{Addr: 0x111E, SPoint: &c_base.SPoint{Key: "ModuleMaxTempChanel", Name: "模块最高温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	ModuleMinTempChanel                             = &c_proto.SModbusPoint{Addr: 0x111F, SPoint: &c_base.SPoint{Key: "ModuleMinTempChanel", Name: "模块最低温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	SOH                                             = &c_proto.SModbusPoint{Addr: 0x1120, SPoint: c_default.VPointSOH, DataAccess: c_default.VDataAccessUInt16}
	RemainCapacity                                  = &c_proto.SModbusPoint{Addr: 0x1121, SPoint: &c_base.SPoint{Key: "RemainCapacity", Name: "电池可放电能量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	ChargeCapacity                                  = &c_proto.SModbusPoint{Addr: 0x1123, SPoint: &c_base.SPoint{Key: "ChargeCapacity", Name: "蓄电池充电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	DischargeCapacity                               = &c_proto.SModbusPoint{Addr: 0x1125, SPoint: &c_base.SPoint{Key: "DischargeCapacity", Name: "蓄电池放电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	TodayCharge                                     = &c_proto.SModbusPoint{Addr: 0x1127, SPoint: &c_base.SPoint{Key: "TodayCharge", Name: "当日累积充电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	TodayDischarge                                  = &c_proto.SModbusPoint{Addr: 0x1129, SPoint: &c_base.SPoint{Key: "TodayDischarge", Name: "当日累积放电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32Scale001}
	HistoryCharge                                   = &c_proto.SModbusPoint{Addr: 0x112B, SPoint: &c_base.SPoint{Key: "HistoryCharge", Name: "历史累积充电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32}
	HistoryDischarge                                = &c_proto.SModbusPoint{Addr: 0x112D, SPoint: &c_base.SPoint{Key: "HistoryDischarge", Name: "历史累积放电量", Unit: "kWh"}, DataAccess: c_default.VDataAccessUInt32}
	RequestForceChargeMark                          = &c_proto.SModbusPoint{Addr: 0x112F, SPoint: &c_base.SPoint{Key: "RequestForceChargeMark", Name: "强充标志位"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	RequestBalanceChargeMark                        = &c_proto.SModbusPoint{Addr: 0x1130, SPoint: &c_base.SPoint{Key: "RequestBalanceChargeMark", Name: "均衡充电标志位"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	PileNumberInParallel                            = &c_proto.SModbusPoint{Addr: 0x1131, SPoint: &c_base.SPoint{Key: "PileNumberInParallel", Name: "当前并联电池组数量"}, DataAccess: c_default.VDataAccessInt16}
	ErrorCode1                                      = &c_proto.SModbusPoint{Addr: 0x1132, SPoint: c_default.VPointErrorCode, DataAccess: c_default.VDataAccessInt16, Trigger: c_default.FAlarmTriggerErrorNotZero}
	ErrorCode2                                      = &c_proto.SModbusPoint{Addr: 0x1134, SPoint: c_default.VPointErrorCode, DataAccess: c_default.VDataAccessInt16, Trigger: c_default.FAlarmTriggerErrorNotZero}
	NumberBatteryModulesInSeriesConnectionInOnePile = &c_proto.SModbusPoint{Addr: 0x1136, SPoint: &c_base.SPoint{Key: "NumberBatteryModulesInSeriesConnectionInOnePile", Name: "电池组模块串联数"}, DataAccess: c_default.VDataAccessInt16}
	NumberCellInSeriesInOnePile                     = &c_proto.SModbusPoint{Addr: 0x1137, SPoint: &c_base.SPoint{Key: "NumberCellInSeriesInOnePile", Name: "电池组单体串联数"}, DataAccess: c_default.VDataAccessInt16}

	ChargeForbiddenMark            = &c_proto.SModbusPoint{Addr: 0x1138, SPoint: &c_base.SPoint{Key: "ChargeForbiddenMark", Name: "禁止充电标志"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	DischargeForbiddenMark         = &c_proto.SModbusPoint{Addr: 0x1139, SPoint: &c_base.SPoint{Key: "DischargeForbiddenMark", Name: "禁止放电标志"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	SOC30Flag                      = &c_proto.SModbusPoint{Addr: 0x113A, SPoint: &c_base.SPoint{Key: "SOC30Flag", Name: "SOC<=30%标志"}, DataAccess: c_default.VDataAccessUInt16ToBool}
	SOE                            = &c_proto.SModbusPoint{Addr: 0x113B, SPoint: &c_base.SPoint{Key: "SOE", Name: "SOE", Unit: "%"}, DataAccess: c_default.VDataAccessUInt16}
	HeartbeatSignal                = &c_proto.SModbusPoint{Addr: 0x113C, SPoint: &c_base.SPoint{Key: "HeartbeatSignal", Name: "心跳信号值", Desc: ",范围为 0x0000~0x00FF，每次读取该值都会增加 1"}, DataAccess: c_default.VDataAccessInt16}
	ModulePCBMaxTemp               = &c_proto.SModbusPoint{Addr: 0x113D, SPoint: &c_base.SPoint{Key: "ModulePCBMaxTemp", Name: "电池模块单板最高温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	ModulePCBMinTemp               = &c_proto.SModbusPoint{Addr: 0x113E, SPoint: &c_base.SPoint{Key: "ModulePCBMinTemp", Name: "电池模块单板最低温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	ModulePCBMaxTempChanel         = &c_proto.SModbusPoint{Addr: 0x113F, SPoint: &c_base.SPoint{Key: "ModulePCBMaxTempChanel", Name: "电池模块单板最高温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	ModulePCBMinTempChanel         = &c_proto.SModbusPoint{Addr: 0x1140, SPoint: &c_base.SPoint{Key: "ModulePCBMinTempChanel", Name: "电池模块单板最低温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	SystemOperationStatus          = &c_proto.SModbusPoint{Addr: 0x1141, SPoint: &c_base.SPoint{Key: "SystemOperationStatus", Name: "系统运行状态Linked with 0x1094.\n0x11:Standby(self-inspection is \nfinished but relay didn`t close, \nrequest to command the system at \n0x1094 to close relay, the status will \nthen switch to 'Run')\n0x22:Run(close relay)"}, DataAccess: c_default.VDataAccessInt16}
	InsulationResistance           = &c_proto.SModbusPoint{Addr: 0x1148, SPoint: &c_base.SPoint{Key: "InsulationResistance", Name: "绝缘电阻值", Unit: "KΩ"}, DataAccess: c_default.VDataAccessUInt16Scale0001}
	InsulationResistanceErrorLevel = &c_proto.SModbusPoint{Addr: 0x1149, SPoint: &c_base.SPoint{Key: "InsulationResistanceErrorLevel", Name: "绝缘电阻故障等级"}, DataAccess: c_default.VDataAccessInt16}
	TerminalMaxTemp                = &c_proto.SModbusPoint{Addr: 0x114A, SPoint: &c_base.SPoint{Key: "TerminalMaxTemp", Name: "端子最高温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	TerminalMinTemp                = &c_proto.SModbusPoint{Addr: 0x114B, SPoint: &c_base.SPoint{Key: "TerminalMinTemp", Name: "端子最低温度", Unit: "℃"}, DataAccess: c_default.VDataAccessUInt16Scale01}

	TerminalMaxTempChannel = &c_proto.SModbusPoint{Addr: 0x114C, SPoint: &c_base.SPoint{Key: "TerminalMaxTempChannel", Name: "端子最高温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	TerminalMinTempChannel = &c_proto.SModbusPoint{Addr: 0x114D, SPoint: &c_base.SPoint{Key: "TerminalMinTempChannel", Name: "端子最低温度通道", Unit: "℃"}, DataAccess: c_default.VDataAccessInt16}
	AlertStatus2           = &c_proto.SModbusPoint{Addr: 0x114E, SPoint: &c_base.SPoint{Key: "AlarmStatus2", Name: "告警状态2"}, DataAccess: c_default.VDataAccessInt16, Trigger: c_default.FAlarmTriggerAlertNotZero}
	MaxChargePower         = &c_proto.SModbusPoint{Addr: 0x1150, SPoint: &c_base.SPoint{Key: "MaxChargePower", Name: "20 分钟内最大持续充电功率预测值", Unit: "kW"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	MaxDischargePower      = &c_proto.SModbusPoint{Addr: 0x1151, SPoint: &c_base.SPoint{Key: "MaxDischargePower", Name: "20 分钟内最大持续放电功率预测值", Unit: "kW"}, DataAccess: c_default.VDataAccessUInt16Scale01}
)

var (
	External485Protocol = &c_proto.SModbusPoint{Addr: 0x12F0, SPoint: &c_base.SPoint{Key: "External485Protocol", Name: "对外 485 协议", Desc: ";0x01：pylon modbus\n0x02：无效\n0x03：kostal\n0x04：kaco\nothers：无效"}, DataAccess: c_default.VDataAccessInt16}
	ExternalCanProtocol = &c_proto.SModbusPoint{Addr: 0x12F1, SPoint: &c_base.SPoint{Key: "ExternalCanProtocol", Name: "对外 can 协议", Desc: ";0x01：auto\n0x02：pylon can\n0x03：solax\n0x04：sma\nothers：无效"}, DataAccess: c_default.VDataAccessInt16}
	InputDryContact1    = &c_proto.SModbusPoint{Addr: 0x12F2, SPoint: &c_base.SPoint{Key: "InputDryContact1", Name: "输入干接点", Desc: " 1;0x0：失能\n0x1：急停功能，\nOthers：无效\n注意:该点位当前只支持 FH 系列，其它\n产品无效"}, DataAccess: c_default.VDataAccessInt16}
)
