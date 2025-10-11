package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"fmt"

	"github.com/shockerli/cvt"
)

var (
	StatusGroup = &c_base.SPointGroup{GroupName: "状态", GroupSort: 1}

	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetryBmsStatusPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "bmsStatus",
			Name:      "电池状态",
			ValueType: c_enum.EInt16,
			Desc:      "电池状态",
			ValueExplain: []*c_base.SFieldExplain{
				{Key: "0", Value: "未知", Color: "#FF5D5D", FromParam: false},
				{Key: "1", Value: "关机", Color: "#9CBF30", FromParam: false},
				{Key: "2", Value: "待机", Color: "#6967EE", FromParam: false},
				{Key: "3", Value: "充电中", Color: "#29A634", FromParam: false},
				{Key: "4", Value: "放电中", Color: "#0098FA", FromParam: false},
				{Key: "5", Value: "故障", Color: "#FF5D5D", FromParam: false},
			},
		},
		MethodName: "GetBmsStatus",
	}

	telemetrySocPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "soc",
			Name:      "SOC",
			Unit:      "%",
			ValueType: c_enum.EFloat32,
			Desc:      "电池电量",
		},
		MethodName: "GetSoc",
	}

	telemetryCapacityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "capacity",
			Name:      "电池容量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat32,
			Desc:      "电池容量",
		},
		MethodName: "GetCapacity",
	}

	telemetryDcPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "dcPower",
			Name:      "直流功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat64,
			Desc:      "直流功率",
		},
		MethodName: "GetDcPower",
	}

	telemetryDcVoltagePoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "dcVoltage",
			Name:      "直流电压",
			Unit:      "V",
			ValueType: c_enum.EFloat64,
			Desc:      "直流电压",
		},
		MethodName: "GetDcVoltage",
	}

	telemetryDcCurrentPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "dcCurrent",
			Name:      "直流电流",
			Unit:      "A",
			ValueType: c_enum.EFloat64,
			Desc:      "直流电流",
		},
		MethodName: "GetDcCurrent",
	}

	telemetryHistoryIncomingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyIncomingQuantity",
			Name:      "充电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "充电量",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	telemetryHistoryOutgoingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyOutgoingQuantity",
			Name:      "放电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "放电量",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}
)

// 控制信息
var (
	SleepControl = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "SleepControl",
				Name:      "SleepControl",
				ValueType: c_enum.EInt16,
				Desc:      "休眠控制, 写入0xAA进入休眠状态，写入0x55唤醒",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1090,
	}

	AllowCharge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "AllowCharge",
				Name:      "AllowCharge",
				ValueType: c_enum.EInt16,
				Desc:      "允许充电控制, 写入0xAA允许充电，其他无效",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1091,
	}

	AllowDischarge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "AllowDischarge",
				Name:      "AllowDischarge",
				ValueType: c_enum.EInt16,
				Desc:      "允许放电控制, 写入0xAA允许放电，其他无效",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1092,
	}

	TempShield = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TempShield",
				Name:      "TempShield",
				ValueType: c_enum.EInt16,
				Desc:      "临时屏蔽\"无通信时切离继电器功能\"的请求",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1093,
	}

	AllowRun = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "AllowRun",
				Name:      "AllowRun",
				ValueType: c_enum.EInt16,
				Desc:      "运行指令，通知系统开始并柜, 写入0xAA允许运行，其他无效",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1094,
	}
)

// 时间
var (
	Year = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Year",
				Name:      "Year",
				ValueType: c_enum.EUint16,
				Desc:      "年 00~99（表示 2000~2099）",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E0,
	}

	Month = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Month",
				Name:      "Month",
				ValueType: c_enum.EUint16,
				Desc:      "月 1~12",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E1,
	}

	Day = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Day",
				Name:      "Day",
				ValueType: c_enum.EUint16,
				Desc:      "日 1~31",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E2,
	}

	Hour = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Hour",
				Name:      "Hour",
				ValueType: c_enum.EUint16,
				Desc:      "时 0~23",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E3,
	}

	Minute = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Minute",
				Name:      "Minute",
				ValueType: c_enum.EUint16,
				Desc:      "分 0~59",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E4,
	}

	Second = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Second",
				Name:      "Second",
				ValueType: c_enum.EUint16,
				Desc:      "秒 0~59",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x10E5,
	}
)

// 主要状态和参数点位
var (
	BasicStatus = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "BasicStatus",
				Name:      "基本状态",
				ValueType: c_enum.EUint16,
				Desc:      "基本状态:00：休眠 01：充电 02：放电 03：搁置",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   0,
				BitLength:  2,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
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
		},
	}

	DCVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointDcVoltage,
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x1103,
	}

	DCCurrent = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointDcCurrent,
			DataAccess: c_default.VDataAccessInt32Scale001,
		},
		Addr: 0x1104,
	}

	Temperature = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointTemp,
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x1106,
	}

	SOC = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointSOC,
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x1107,
	}

	CycleCount = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "CycleCount",
				Name:      "循环次数",
				ValueType: c_enum.EUint16,
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x1108,
	}

	PileMaxV = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileMaxV",
				Name:      "电池组最大充电电压值",
				ValueType: c_enum.EFloat32,
				Unit:      "V",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x1109,
	}

	PileMaxI = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileMaxI",
				Name:      "电池组最大充电电流值",
				ValueType: c_enum.EFloat64,
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x110A,
	}

	PileMinV = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileMinV",
				Name:      "电池组最小放电电压值",
				ValueType: c_enum.EFloat32,
				Unit:      "V",
			},
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x110C,
	}

	PileMaxDI = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileMaxDI",
				Name:      "电池组最大放电电流值",
				ValueType: c_enum.EFloat64,
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x110D,
	}

	BatteryCellMaxVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "BatteryCellMaxVoltage",
				Name:      "电池最高电压",
				ValueType: c_enum.EFloat32,
				Unit:      "V",
			},
			DataAccess: c_default.VDataAccessUInt16Scale0001,
		},
		Addr: 0x1110,
	}

	BatteryCellMinVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "BatteryCellMinVoltage",
				Name:      "电池最低电压",
				ValueType: c_enum.EFloat32,
				Unit:      "V",
			},
			DataAccess: c_default.VDataAccessUInt16Scale0001,
		},
		Addr: 0x1111,
	}

	BatteryCellMaxTemp = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointTempMax,
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x1114,
	}

	BatteryCellMinTemp = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointTempMin,
			DataAccess: c_default.VDataAccessUInt16Scale01,
		},
		Addr: 0x1115,
	}

	SOH = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     c_default.VPointSOH,
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x1120,
	}

	RemainCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "RemainCapacity",
				Name:      "电池可放电能量",
				ValueType: c_enum.EFloat64,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x1121,
	}

	ChargeCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ChargeCapacity",
				Name:      "蓄电池充电量",
				ValueType: c_enum.EFloat64,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x1123,
	}

	DischargeCapacity = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "DischargeCapacity",
				Name:      "蓄电池放电量",
				ValueType: c_enum.EFloat64,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x1125,
	}

	TodayCharge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TodayCharge",
				Name:      "当日累积充电量",
				ValueType: c_enum.EFloat64,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x1127,
	}

	TodayDischarge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TodayDischarge",
				Name:      "当日累积放电量",
				ValueType: c_enum.EFloat64,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32Scale001,
		},
		Addr: 0x1129,
	}

	HistoryCharge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "HistoryCharge",
				Name:      "历史累积充电量",
				ValueType: c_enum.EUint32,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32,
		},
		Addr: 0x112B,
	}

	HistoryDischarge = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "HistoryDischarge",
				Name:      "历史累积放电量",
				ValueType: c_enum.EUint32,
				Unit:      "kWh",
			},
			DataAccess: c_default.VDataAccessUInt32,
		},
		Addr: 0x112D,
	}

	ChargeForbiddenMark = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ChargeForbiddenMark",
				Name:      "禁止充电标志",
				ValueType: c_enum.EBool,
			},
			DataAccess: c_default.VDataAccessUInt16ToBool,
		},
		Addr: 0x1138,
	}

	DischargeForbiddenMark = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "DischargeForbiddenMark",
				Name:      "禁止放电标志",
				ValueType: c_enum.EBool,
			},
			DataAccess: c_default.VDataAccessUInt16ToBool,
		},
		Addr: 0x1139,
	}

	SOC30Flag = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "SOC30Flag",
				Name:      "SOC<=30%标志",
				ValueType: c_enum.EBool,
			},
			DataAccess: c_default.VDataAccessUInt16ToBool,
		},
		Addr: 0x113A,
	}

	SOE = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "SOE",
				Name:      "SOE",
				ValueType: c_enum.EUint16,
				Unit:      "%",
			},
			DataAccess: c_default.VDataAccessUInt16,
		},
		Addr: 0x113B,
	}

	HeartbeatSignal = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "HeartbeatSignal",
				Name:      "心跳信号值",
				ValueType: c_enum.EInt16,
				Desc:      ",范围为 0x0000~0x00FF，每次读取该值都会增加 1",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x113C,
	}

	Switching = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Switching",
				Name:      "开关量指示",
				ValueType: c_enum.EInt16,
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x110F,
	}

	SystemErrorProtection = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "SystemErrorProtection",
				Name:      "系统故障保护",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerErrorBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   3,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	CurrentProtection = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "CurrentProtection",
				Name:      "电流保护",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerErrorBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   4,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	VoltageProtection = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "VoltageProtection",
				Name:      "电压保护",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerErrorBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   5,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	TemperatureProtection = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TemperatureProtection",
				Name:      "温度保护",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerErrorBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   6,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	VoltageAlarm = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "VoltageAlarm",
				Name:      "电压警告",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerAlertBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   7,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	CurrentAlarm = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "CurrentAlarm",
				Name:      "电流警告",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerAlertBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   8,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	TemperatureAlarm = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TemperatureAlarm",
				Name:      "温度警告",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerAlertBool,
				Desc:      "0-正常；1-保护",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   9,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr:          0x1100,
		StatusExplain: c_default.StatusExplainProtectFunc,
	}

	PileSystemIdleStatus = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileSystemIdleStatus",
				Name:      "电池组搁置状态",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Desc:      "0-否，1-搁置",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   10,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
			return c_default.StatusExplainBool(value, "搁置", "否")
		},
	}

	PileSystemChargeStatus = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileSystemChargeStatus",
				Name:      "电池组充电状态",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Desc:      "0-否，1-充电",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   11,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
			return c_default.StatusExplainBool(value, "充电", "否")
		},
	}

	PileSystemDischargeStatus = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileSystemDischargeStatus",
				Name:      "电池组放电状态",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Desc:      "0-否，1-放电",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   12,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
			return c_default.StatusExplainBool(value, "放电", "否")
		},
	}

	PileSystemSleepStatus = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PileSystemSleepStatus",
				Name:      "电池组休眠状态",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Desc:      "0-否，1-休眠",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   13,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
			return c_default.StatusExplainBool(value, "休眠", "否")
		},
	}

	FanWarn = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "FanWarn",
				Name:      "风扇告警",
				ValueType: c_enum.EBool,
				Group:     StatusGroup,
				Trigger:   c_default.FAlarmTriggerAlertBool,
				Desc:      "0-无异常，1-有异常",
			},
			DataAccess: &c_base.SDataAccess{
				BitIndex:   14,
				BitLength:  1,
				DataFormat: c_enum.DataFormatBitRange,
				ByteEndian: c_enum.ByteEndianBig,
				WordOrder:  c_enum.WordOrderHighLow,
			},
		},
		Addr: 0x1100,
		StatusExplain: func(value any) (string, error) {
			return c_default.StatusExplainBool(value, "有异常", "无异常")
		},
	}

	Protection = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Protection",
				Name:      "保护状态",
				ValueType: c_enum.EInt16,
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x1101,
	}

	AlarmStatus1 = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "AlarmStatus1",
				Name:      "告警状态1",
				ValueType: c_enum.EInt16,
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr:    0x1102,
		Trigger: c_default.FAlarmTriggerAlertNotZero,
	}
)
