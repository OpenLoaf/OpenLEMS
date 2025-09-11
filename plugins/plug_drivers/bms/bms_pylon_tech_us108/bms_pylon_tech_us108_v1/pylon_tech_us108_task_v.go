package bms_pylon_tech_us108_v1

import (
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	GroupHeart = &c_proto.SModbusPointTask{
		Name:      "GroupHeart",
		Desc:      "心跳,等点位",
		Addr:      ChargeForbiddenMark.Addr,
		Quantity:  5,
		Function:  c_enum.EMqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_proto.DefaultCacheLifeTime,
		Points:    []*c_proto.SModbusPoint{ChargeForbiddenMark, DischargeForbiddenMark, SOC30Flag, SOE, HeartbeatSignal},
	}

	GroupInfo = &c_proto.SModbusPointTask{
		Name:      "GroupInfo",
		Desc:      "查询基本运行信息",
		Addr:      BasicStatus.Addr,
		Quantity:  Switching.Addr - BasicStatus.Addr + 1,
		Function:  c_enum.EMqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_proto.DefaultCacheLifeTime,
		Points: []*c_proto.SModbusPoint{BasicStatus, SystemErrorProtection, CurrentProtection, VoltageProtection,
			TemperatureProtection, VoltageAlarm, CurrentAlarm, TemperatureAlarm,
			PileSystemIdleStatus, PileSystemChargeStatus, PileSystemDischargeStatus, PileSystemSleepStatus, FanWarn,
			Protection, AlarmStatus1, DCVoltage, DCCurrent, Temperature, SOC, CycleCount,
			PileMaxV, PileMaxI, PileMinV, PileMaxDI, Switching,
		},
	}

	GroupTime = &c_proto.SModbusPointTask{
		Name:      "GroupTime",
		Desc:      "查询年月日时分秒",
		Addr:      Year.Addr,
		Quantity:  Second.Addr - Year.Addr + 1,
		Function:  c_enum.EMqHoldingRegisters,
		CycleMill: 1000, //TODO 0
		Lifetime:  c_proto.DefaultCacheLifeTime,
		Points:    []*c_proto.SModbusPoint{Year, Month, Day, Hour, Minute, Second},
	}

	GroupStatistics = &c_proto.SModbusPointTask{
		Name:      "GroupStatistics",
		Desc:      "查询统计信息",
		Addr:      SOH.Addr,
		Quantity:  HistoryDischarge.Addr - SOH.Addr + 2,
		Function:  c_enum.EMqHoldingRegisters,
		CycleMill: 1000,             //TODO 0
		Lifetime:  30 * time.Second, // 30s后过期
		Points:    []*c_proto.SModbusPoint{SOH, RemainCapacity, ChargeCapacity, DischargeCapacity, TodayCharge, TodayDischarge, HistoryCharge, HistoryDischarge},
	}
)
