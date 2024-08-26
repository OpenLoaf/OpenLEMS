package pylon_tech_us108

import (
	"ems-plan/c_base"
	"plug_protocol_modbus/p_modbus"
	"time"
)

var (
	GroupHeart = &p_modbus.ModbusGroup{
		Name:      "GroupHeart",
		Desc:      "心跳,每读一次心跳自动加一",
		Addr:      HeartbeatSignal.Addr,
		Quantity:  1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{HeartbeatSignal},
	}

	GroupInfo = &p_modbus.ModbusGroup{
		Name:      "GroupInfo",
		Desc:      "查询基本运行信息",
		Addr:      BasicStatus.Addr,
		Quantity:  CycleCount.Addr - BasicStatus.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{BasicStatus, SystemErrorProtection, CurrentProtection, VoltageProtection,
			TemperatureProtection, VoltageProtection, VoltageAlarm, CurrentAlarm, TemperatureAlarm,
			PileSystemIdleStatus, PileSystemChargeStatus, PileSystemDischargeStatus, PileSystemSleepStatus, FanWarn,
			Protection, AlarmStatus1, DCVoltage, DCCurrent, Temperature, SOC, CycleCount},
	}

	GroupTime = &p_modbus.ModbusGroup{
		Name:      "GroupTime",
		Desc:      "查询年月日时分秒",
		Addr:      Year.Addr,
		Quantity:  Second.Addr - Year.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000, //TODO 0
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Year, Month, Day, Hour, Minute, Second},
	}

	GroupStatistics = &p_modbus.ModbusGroup{
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
