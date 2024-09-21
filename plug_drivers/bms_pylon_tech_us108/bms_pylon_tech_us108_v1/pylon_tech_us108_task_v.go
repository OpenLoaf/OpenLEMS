package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"modbus/p_modbus"
	"time"
)

var (
	GroupHeart = &p_modbus.SModbusTask{
		Name:      "GroupHeart",
		Desc:      "心跳,等点位",
		Addr:      ChargeForbiddenMark.Addr,
		Quantity:  5,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{ChargeForbiddenMark, DischargeForbiddenMark, SOC30Flag, SOE, HeartbeatSignal},
	}

	GroupInfo = &p_modbus.SModbusTask{
		Name:      "GroupInfo",
		Desc:      "查询基本运行信息",
		Addr:      BasicStatus.Addr,
		Quantity:  Switching.Addr - BasicStatus.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{BasicStatus, SystemErrorProtection, CurrentProtection, VoltageProtection,
			TemperatureProtection, VoltageProtection, VoltageAlarm, CurrentAlarm, TemperatureAlarm,
			PileSystemIdleStatus, PileSystemChargeStatus, PileSystemDischargeStatus, PileSystemSleepStatus, FanWarn,
			Protection, AlarmStatus1, DCVoltage, DCCurrent, Temperature, SOC, CycleCount,
			PileMaxV, PileMaxI, PileMinV, PileMaxDI, Switching,
		},
	}

	GroupTime = &p_modbus.SModbusTask{
		Name:      "GroupTime",
		Desc:      "查询年月日时分秒",
		Addr:      Year.Addr,
		Quantity:  Second.Addr - Year.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000, //TODO 0
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Year, Month, Day, Hour, Minute, Second},
	}

	GroupStatistics = &p_modbus.SModbusTask{
		Name:      "GroupStatistics",
		Desc:      "查询统计信息",
		Addr:      SOH.Addr,
		Quantity:  HistoryDischarge.Addr - SOH.Addr + 2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,             //TODO 0
		Lifetime:  30 * time.Second, // 30s后过期
		Metas:     []*c_base.Meta{SOH, RemainCapacity, ChargeCapacity, DischargeCapacity, TodayCharge, TodayDischarge, HistoryCharge, HistoryDischarge},
	}
)
