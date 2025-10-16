package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
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
			Min:       0,
			Max:       100,
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
	SleepControl = c_proto.NewModbusPointWithDesc(0x1090, "SleepControl", "SleepControl", c_enum.EInt16, "", "休眠控制, 写入0xAA进入休眠状态，写入0x55唤醒", c_default.VDataAccessInt16)

	AllowCharge = c_proto.NewModbusPointWithDesc(0x1091, "AllowCharge", "AllowCharge", c_enum.EInt16, "", "允许充电控制, 写入0xAA允许充电，其他无效", c_default.VDataAccessInt16)

	AllowDischarge = c_proto.NewModbusPointWithDesc(0x1092, "AllowDischarge", "AllowDischarge", c_enum.EInt16, "", "允许放电控制, 写入0xAA允许放电，其他无效", c_default.VDataAccessInt16)

	TempShield = c_proto.NewModbusPointWithDesc(0x1093, "TempShield", "TempShield", c_enum.EInt16, "", "临时屏蔽\"无通信时切离继电器功能\"的请求", c_default.VDataAccessInt16)

	AllowRun = c_proto.NewModbusPointWithDesc(0x1094, "AllowRun", "AllowRun", c_enum.EInt16, "", "运行指令，通知系统开始并柜, 写入0xAA允许运行，其他无效", c_default.VDataAccessInt16)
)

// 时间
var (
	Year = c_proto.NewModbusPointWithDesc(0x10E0, "Year", "Year", c_enum.EUint16, "", "年 00~99（表示 2000~2099）", c_default.VDataAccessUInt16)

	Month = c_proto.NewModbusPointWithDesc(0x10E1, "Month", "Month", c_enum.EUint16, "", "月 1~12", c_default.VDataAccessUInt16)

	Day = c_proto.NewModbusPointWithDesc(0x10E2, "Day", "Day", c_enum.EUint16, "", "日 1~31", c_default.VDataAccessUInt16)

	Hour = c_proto.NewModbusPointWithDesc(0x10E3, "Hour", "Hour", c_enum.EUint16, "", "时 0~23", c_default.VDataAccessUInt16)

	Minute = c_proto.NewModbusPointWithDesc(0x10E4, "Minute", "Minute", c_enum.EUint16, "", "分 0~59", c_default.VDataAccessUInt16)

	Second = c_proto.NewModbusPointWithDesc(0x10E5, "Second", "Second", c_enum.EUint16, "", "秒 0~59", c_default.VDataAccessUInt16)
)

// 主要状态和参数点位
var (
	BasicStatus = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("BasicStatus"),
		c_proto.WithName("基本状态"),
		c_proto.WithValueType(c_enum.EUint16),
		c_proto.WithDesc("基本状态:00：休眠 01：充电 02：放电 03：搁置"),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   0,
			BitLength:  2,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "0", Value: "休眠", Color: "#1890ff"},
			{Key: "1", Value: "充电", Color: "#52c41a"},
			{Key: "2", Value: "放电", Color: "#faad14"},
			{Key: "3", Value: "搁置", Color: "#d9d9d9"},
		}),
	)

	DCVoltage = c_proto.NewModbusPointFromPreset(0x1103, c_default.VPointDcVoltage, c_default.VDataAccessUInt16Scale01)

	DCCurrent = c_proto.NewModbusPointFromPreset(0x1104, c_default.VPointDcCurrent, c_default.VDataAccessInt32Scale001)

	Temperature = c_proto.NewModbusPointFromPreset(0x1106, c_default.VPointTemp, c_default.VDataAccessUInt16Scale01)

	SOC = c_proto.NewModbusPointFromPreset(0x1107, c_default.VPointSOC, c_default.VDataAccessUInt16)

	CycleCount = c_proto.NewModbusPointWithDesc(0x1108, "CycleCount", "循环次数", c_enum.EUint16, "", "循环次数", c_default.VDataAccessUInt16)

	PileMaxV = c_proto.NewModbusPointWithDesc(0x1109, "PileMaxV", "电池组最大充电电压值", c_enum.EFloat32, "V", "电池组最大充电电压值", c_default.VDataAccessUInt16Scale01)

	PileMaxI = c_proto.NewModbusPointWithDesc(0x110A, "PileMaxI", "电池组最大充电电流值", c_enum.EFloat64, "", "电池组最大充电电流值", c_default.VDataAccessUInt32Scale001)

	PileMinV = c_proto.NewModbusPointWithDesc(0x110C, "PileMinV", "电池组最小放电电压值", c_enum.EFloat32, "V", "电池组最小放电电压值", c_default.VDataAccessUInt16Scale01)

	PileMaxDI = c_proto.NewModbusPointWithDesc(0x110D, "PileMaxDI", "电池组最大放电电流值", c_enum.EFloat64, "", "电池组最大放电电流值", c_default.VDataAccessUInt32Scale001)

	BatteryCellMaxVoltage = c_proto.NewModbusPointWithDesc(0x1110, "BatteryCellMaxVoltage", "电池最高电压", c_enum.EFloat32, "V", "电池最高电压", c_default.VDataAccessUInt16Scale0001)

	BatteryCellMinVoltage = c_proto.NewModbusPointWithDesc(0x1111, "BatteryCellMinVoltage", "电池最低电压", c_enum.EFloat32, "V", "电池最低电压", c_default.VDataAccessUInt16Scale0001)

	BatteryCellMaxTemp = c_proto.NewModbusPointFromPreset(0x1114, c_default.VPointTempMax, c_default.VDataAccessUInt16Scale01)

	BatteryCellMinTemp = c_proto.NewModbusPointFromPreset(0x1115, c_default.VPointTempMin, c_default.VDataAccessUInt16Scale01)

	SOH = c_proto.NewModbusPointFromPreset(0x1120, c_default.VPointSOH, c_default.VDataAccessUInt16)

	RemainCapacity = c_proto.NewModbusPointWithDesc(0x1121, "RemainCapacity", "电池可放电能量", c_enum.EFloat64, "kWh", "电池可放电能量", c_default.VDataAccessUInt32Scale001)

	ChargeCapacity = c_proto.NewModbusPointWithDesc(0x1123, "ChargeCapacity", "蓄电池充电量", c_enum.EFloat64, "kWh", "蓄电池充电量", c_default.VDataAccessUInt32Scale001)

	DischargeCapacity = c_proto.NewModbusPointWithDesc(0x1125, "DischargeCapacity", "蓄电池放电量", c_enum.EFloat64, "kWh", "蓄电池放电量", c_default.VDataAccessUInt32Scale001)

	TodayCharge = c_proto.NewModbusPointWithDesc(0x1127, "TodayCharge", "当日累积充电量", c_enum.EFloat64, "kWh", "当日累积充电量", c_default.VDataAccessUInt32Scale001)

	TodayDischarge = c_proto.NewModbusPointWithDesc(0x1129, "TodayDischarge", "当日累积放电量", c_enum.EFloat64, "kWh", "当日累积放电量", c_default.VDataAccessUInt32Scale001)

	HistoryCharge = c_proto.NewModbusPointWithDesc(0x112B, "HistoryCharge", "历史累积充电量", c_enum.EUint32, "kWh", "历史累积充电量", c_default.VDataAccessUInt32)

	HistoryDischarge = c_proto.NewModbusPointWithDesc(0x112D, "HistoryDischarge", "历史累积放电量", c_enum.EUint32, "kWh", "历史累积放电量", c_default.VDataAccessUInt32)

	ChargeForbiddenMark = c_proto.NewModbusPointExt(0x1138,
		c_proto.WithKey("ChargeForbiddenMark"),
		c_proto.WithName("禁止充电标志"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDataAccess(c_default.VDataAccessUInt16ToBool),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "允许充电", Color: "#52c41a"},
			{Key: "true", Value: "禁止充电", Color: "#f5222d"},
		}),
	)

	DischargeForbiddenMark = c_proto.NewModbusPointExt(0x1139,
		c_proto.WithKey("DischargeForbiddenMark"),
		c_proto.WithName("禁止放电标志"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDataAccess(c_default.VDataAccessUInt16ToBool),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "允许放电", Color: "#52c41a"},
			{Key: "true", Value: "禁止放电", Color: "#f5222d"},
		}),
	)

	SOC30Flag = c_proto.NewModbusPointExt(0x113A,
		c_proto.WithKey("SOC30Flag"),
		c_proto.WithName("SOC<=30%标志"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDataAccess(c_default.VDataAccessUInt16ToBool),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "否", Color: "#52c41a"},
			{Key: "true", Value: "是", Color: "#faad14"},
		}),
	)

	SOE = c_proto.NewModbusPointWithDesc(0x113B, "SOE", "SOE", c_enum.EUint16, "%", "SOE", c_default.VDataAccessUInt16)

	HeartbeatSignal = c_proto.NewModbusPointWithDesc(0x113C, "HeartbeatSignal", "心跳信号值", c_enum.EInt16, "", ",范围为 0x0000~0x00FF，每次读取该值都会增加 1", c_default.VDataAccessInt16)

	Switching = c_proto.NewModbusPointWithDesc(0x110F, "Switching", "开关量指示", c_enum.EInt16, "", "开关量指示", c_default.VDataAccessInt16)

	SystemErrorProtection = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("SystemErrorProtection"),
		c_proto.WithName("系统故障保护"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerErrorBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   3,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	CurrentProtection = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("CurrentProtection"),
		c_proto.WithName("电流保护"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerErrorBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   4,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	VoltageProtection = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("VoltageProtection"),
		c_proto.WithName("电压保护"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerErrorBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   5,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	TemperatureProtection = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("TemperatureProtection"),
		c_proto.WithName("温度保护"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerErrorBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   6,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	VoltageAlarm = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("VoltageAlarm"),
		c_proto.WithName("电压警告"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerAlertBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   7,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	CurrentAlarm = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("CurrentAlarm"),
		c_proto.WithName("电流警告"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerAlertBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   8,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	TemperatureAlarm = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("TemperatureAlarm"),
		c_proto.WithName("温度警告"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-正常；1-保护"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerAlertBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   9,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "正常", Color: "#52c41a"},
			{Key: "true", Value: "保护", Color: "#f5222d"},
		}),
	)

	PileSystemIdleStatus = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("PileSystemIdleStatus"),
		c_proto.WithName("电池组搁置状态"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-否，1-搁置"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   10,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "否", Color: "#52c41a"},
			{Key: "true", Value: "搁置", Color: "#d9d9d9"},
		}),
	)

	PileSystemChargeStatus = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("PileSystemChargeStatus"),
		c_proto.WithName("电池组充电状态"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-否，1-充电"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   11,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "否", Color: "#52c41a"},
			{Key: "true", Value: "充电", Color: "#52c41a"},
		}),
	)

	PileSystemDischargeStatus = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("PileSystemDischargeStatus"),
		c_proto.WithName("电池组放电状态"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-否，1-放电"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   12,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "否", Color: "#52c41a"},
			{Key: "true", Value: "放电", Color: "#faad14"},
		}),
	)

	PileSystemSleepStatus = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("PileSystemSleepStatus"),
		c_proto.WithName("电池组休眠状态"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-否，1-休眠"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   13,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "否", Color: "#52c41a"},
			{Key: "true", Value: "休眠", Color: "#d9d9d9"},
		}),
	)

	FanWarn = c_proto.NewModbusPointExt(0x1100,
		c_proto.WithKey("FanWarn"),
		c_proto.WithName("风扇告警"),
		c_proto.WithValueType(c_enum.EBool),
		c_proto.WithDesc("0-无异常，1-有异常"),
		c_proto.WithGroup(StatusGroup),
		c_proto.WithTrigger(c_default.FAlarmTriggerAlertBool),
		c_proto.WithDataAccess(&c_base.SDataAccess{
			BitIndex:   14,
			BitLength:  1,
			DataFormat: c_enum.DataFormatBitRange,
			ByteEndian: c_enum.ByteEndianBig,
			WordOrder:  c_enum.WordOrderHighLow,
		}),
		c_proto.WithValueExplain([]*c_base.SFieldExplain{
			{Key: "false", Value: "无异常", Color: "#52c41a"},
			{Key: "true", Value: "有异常", Color: "#f5222d"},
		}),
	)

	Protection = c_proto.NewModbusPointWithDesc(0x1101, "Protection", "保护状态", c_enum.EInt16, "", "保护状态", c_default.VDataAccessInt16)

	AlarmStatus1 = c_proto.NewModbusPointExt(0x1102,
		c_proto.WithKey("AlarmStatus1"),
		c_proto.WithName("告警状态1"),
		c_proto.WithValueType(c_enum.EInt16),
		c_proto.WithDataAccess(c_default.VDataAccessInt16),
		c_proto.WithTrigger(c_default.FAlarmTriggerAlertNotZero),
	)
)
