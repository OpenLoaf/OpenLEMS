package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_error"
	"common/c_log"
	"common/c_modbus"
	"common/c_util"
	"context"
)

type sPcsStarCharge100E struct {
	c_modbus.IModbusProtocol
	ctx                 context.Context
	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率
	deviceConfig        *c_base.SDeviceConfig
	*c_base.SDriverDescription
}

func (s *sPcsStarCharge100E) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.IModbusProtocol = protocol.(c_modbus.IModbusProtocol)
	s.deviceConfig = deviceConfig

	// 注册
	s.RegisterRead(s.ctx,
		GroupCommand,
		GroupPowerInfo,
		//GroupPhase,
		//GroupSerial, GroupGridModeSetting, GroupTemperature,
		GroupStatus,
	)
}

func (s *sPcsStarCharge100E) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsStarCharge100E) Shutdown() {
	_ = s.SetPower(0)
	_ = s.SetStatus(c_base.EPcsStatusOff)
	c_log.Noticef(s.ctx, "[%s]%s销毁成功,设置PCS状态为Off!", s.deviceConfig.Id, s.deviceConfig.Name)
}

func (s *sPcsStarCharge100E) GetFunctionList() []*c_base.STelemetry {
	return nil
}

func (s *sPcsStarCharge100E) SetReset() error {
	c_log.Warningf(s.ctx, "sPcsStarCharge100E SetReset() not support!")
	return nil
}

func (s *sPcsStarCharge100E) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusOff {
		_ = s.SetPower(0)
		return s.WriteSingleRegister(OnOffCommand, 0)
	}
	if status == c_base.EPcsStatusStandby {
		// 这里文档是 On/off command: 0- Shutdown, 1- Startup, 2- Standby
		return s.WriteSingleRegister(OnOffCommand, 1)
	}

	return c_error.ErrorParam
}

func (s *sPcsStarCharge100E) SetGridMode(mode c_base.EGridMode) error {
	return nil
}

func (s *sPcsStarCharge100E) GetStatus() (c_base.EEnergyStoreStatus, error) {
	value, err := s.GetUintValue(InverterOperationStatus)
	if err != nil {
		return c_base.EPcsStatusUnknown, err
	}

	switch value {
	// 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error
	case 0, 1:
		// 等待设备启动算是关机的状态
		return c_base.EPcsStatusOff, nil
	case 2, 3:
		// 离网并网运行中时，说明设备正常。获取功率，如果获取功率失败，说明设备故障，获取成功后正为放电，负为充电
		power, err := s.GetPower()
		if err != nil {
			return c_base.EPcsStatusFault, err
		}
		if power > 0 {
			return c_base.EPcsStatusDischarge, nil
		} else if power < 0 {
			return c_base.EPcsStatusCharge, nil
		} else {
			return c_base.EPcsStatusStandby, nil
		}
	case 5:
		return c_base.EPcsStatusFault, err
	}
	return c_base.EPcsStatusFault, err
}

func (s *sPcsStarCharge100E) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridOn, nil
}

func (s *sPcsStarCharge100E) SetPower(power int32) error {
	s.targetPower = power
	return s.WriteSingleRegister(ActivePowerSetting, power)
}

func (s *sPcsStarCharge100E) SetReactivePower(power int32) error {
	s.targetReactivePower = power
	return s.WriteSingleRegister(ReactivePowerSetting, power)
}

func (s *sPcsStarCharge100E) SetPowerFactor(factor float32) error {
	c_log.Warningf(s.ctx, "sPcsStarCharge100E SetPowerFactor() not support!")
	return nil
}

func (s *sPcsStarCharge100E) GetTargetPower() int32 {
	return s.targetPower
}

func (s *sPcsStarCharge100E) GetTargetReactivePower() int32 {
	return s.targetReactivePower
}

func (s *sPcsStarCharge100E) GetTargetPowerFactor() float32 {
	return -1
}

func (s *sPcsStarCharge100E) GetPower() (float64, error) {
	return s.GetFloat64Value(TotalActivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetApparentPower() (float64, error) {
	return s.GetFloat64Value(TotalApparentPowerInverterSide)
}

func (s *sPcsStarCharge100E) GetReactivePower() (float64, error) {
	return s.GetFloat64Value(TotalReactivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetRatedPower() int32 {
	return 100
}

func (s *sPcsStarCharge100E) GetMaxInputPower() (float32, error) {
	return 100, nil
}

func (s *sPcsStarCharge100E) GetMaxOutputPower() (float32, error) {
	return 100, nil
}

func (s *sPcsStarCharge100E) GetTodayIncomingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, DailyBatteryDischargeEnergy)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (s *sPcsStarCharge100E) GetHistoryIncomingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, TotalBatteryDischargeEnergy)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (s *sPcsStarCharge100E) GetTodayOutgoingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, DailyBatteryChargeEnergy)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (s *sPcsStarCharge100E) GetHistoryOutgoingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, TotalBatteryChargeEnergy)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}
