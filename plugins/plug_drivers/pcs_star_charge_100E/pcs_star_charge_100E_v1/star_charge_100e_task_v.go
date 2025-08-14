package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"modbus/p_modbus"
	"time"
)

// 非定时读取的任务
var (
	SyGroupStatistics = &p_modbus.SModbusTask{
		Name:      "GroupStatistics",
		Desc:      "查询统计信息",
		Addr:      DailyBatteryChargeEnergy.Addr,
		Quantity:  TotalBatteryDischargeEnergy.Addr - DailyBatteryChargeEnergy.Addr + 2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 0,                // 不需要定时读取，需要的时候读取
		Lifetime:  30 * time.Second, // 30s后过期
		Metas:     []*c_base.Meta{DailyBatteryChargeEnergy, DailyBatteryDischargeEnergy, TotalBatteryChargeEnergy, TotalBatteryDischargeEnergy},
	}
	SyGroupTime = &p_modbus.SModbusTask{
		Name:      "TimeGroup",
		Desc:      "查询年月日时分秒",
		Addr:      Year.Addr,
		Quantity:  Second.Addr - Year.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 0,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Year, Month, Day, Hour, Minute, Second},
	}

	SyGroupSoftwareVersion = &p_modbus.SModbusTask{
		Name:      "GroupSoftwareVersion",
		Desc:      "查询软件版本",
		Addr:      SoftwareVersion.Addr,
		Quantity:  2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 0,
		Lifetime:  0, // 永不过期
		Metas:     []*c_base.Meta{SoftwareVersion},
	}

	GroupPhase = &p_modbus.SModbusTask{
		Name:      "GroupPhase",
		Desc:      "查询相位信息",
		Addr:      PhaseAVoltageGridSide.Addr,
		Quantity:  PhaseCPowerInverterSide.Addr - PhaseAVoltageGridSide.Addr + 2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 0,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{PhaseAVoltageGridSide, PhaseACurrentInverterSide, PhaseACurrentGridSide, PhaseAPowerInverterSide,
			PhaseBVoltageGridSide, PhaseBCurrentInverterSide, PhaseBCurrentGridSide, PhaseBPowerInverterSide, PhaseCVoltageGridSide,
			PhaseCCurrentInverterSide, PhaseCCurrentGridSide, PhaseCPowerInverterSide},
	}
)

// 定时读取的任务
var (
	GroupStatus = &p_modbus.SModbusTask{
		Name:      "GroupStatus",
		Desc:      "查询PCS状态信息",
		Addr:      InverterOperationStatus.Addr,
		Quantity:  1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{InverterOperationStatus},
	}

	GroupCommand = &p_modbus.SModbusTask{
		Name:      "GroupCommand",
		Desc:      "控制命令",
		Addr:      OnOffCommand.Addr,
		Quantity:  5,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 200,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{OnOffCommand, ActivePowerSetting, ReactivePowerSetting},
	}

	GroupPowerInfo = &p_modbus.SModbusTask{
		Name:      "GroupPowerInfo",
		Desc:      "查询功率信息",
		Addr:      AverageFrequency.Addr,
		Quantity:  TotalApparentPowerInverterSide.Addr - AverageFrequency.Addr + 2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 200,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas: []*c_base.Meta{AverageFrequency, AveragePowerFactor, AverageVoltageBus, AverageVoltagePositive, AverageVoltageNegative,
			TotalActivePowerInverterSide, TotalReactivePowerInverterSide, TotalApparentPowerInverterSide},
	}

	GroupSerial = &p_modbus.SModbusTask{
		Name:      "GroupSerial",
		Desc:      "查询Pcs序列信息",
		Addr:      SerialNumber1.Addr,
		Quantity:  SerialNumber5.Addr - SerialNumber1.Addr + 2,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 2000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{SerialNumber1, SerialNumber2, SerialNumber3, SerialNumber4, SerialNumber5},
	}

	GroupGridModeSetting = &p_modbus.SModbusTask{
		Name:      "GroupGridModeSetting",
		Desc:      "获取PCS并离网状态",
		Addr:      GridModeSetting.Addr,
		Quantity:  1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{GridModeSetting},
	}

	GroupTemperature = &p_modbus.SModbusTask{
		Name:      "GroupTemperature",
		Desc:      "查询温度信息",
		Addr:      Module1Temperature.Addr,
		Quantity:  InternalAmbientTemperature.Addr - Module1Temperature.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 2000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{Module1Temperature, Module2Temperature, Module3Temperature, RadiatorTemperature, InternalAmbientTemperature},
	}

	GroupOtherStatus = &p_modbus.SModbusTask{
		Name:      "GroupStatus",
		Desc:      "查询PCS状态信息",
		Addr:      DCPositiveRelayStatus.Addr,
		Quantity:  GridOutageStatus.Addr - DCPositiveRelayStatus.Addr + 1,
		Function:  p_modbus.MqHoldingRegisters,
		CycleMill: 1000,
		Lifetime:  c_base.DefaultCacheLifeTime,
		Metas:     []*c_base.Meta{DCPositiveRelayStatus, DCNegativeRelayStatus, ACRelayStatus, GridOutageStatus},
	}
)
