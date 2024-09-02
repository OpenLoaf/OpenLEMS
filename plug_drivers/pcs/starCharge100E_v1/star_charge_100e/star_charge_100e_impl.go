package star_charge_100e

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"plug_protocol_modbus/p_modbus"
)

type StarCharge100EPcs struct {
	p_modbus.IModbusProtocol
	Ctx                 context.Context
	log                 *glog.Logger
	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率
	description         *c_base.SDescription
}

func (s *StarCharge100EPcs) GetDescription() *c_base.SDescription {
	return s.description
}

func (s *StarCharge100EPcs) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *StarCharge100EPcs) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.IModbusProtocol = protocol.(p_modbus.IModbusProtocol)
	s.description = &c_base.SDescription{
		Brand:  "Star",
		Model:  "100E",
		Remark: "星星充电100E CS",
	}
	// 注册
	s.RegisterRead(s.Ctx,
		GroupCommand,
		GroupPowerInfo,
		//GroupPhase,
		//GroupSerial, GroupGridModeSetting, GroupTemperature,
		GroupStatus,
	)

	g.Log().Noticef(s.Ctx, "StarCharge100EPcs 初始化完毕！")
}

func (s *StarCharge100EPcs) GetFunctionList() []*c_base.STelemetry {
	return nil
}

func (s *StarCharge100EPcs) SetReset() error {
	g.Log().Warningf(s.Ctx, "StarCharge100EPcs SetReset() not support!")
	return nil
}

func (s *StarCharge100EPcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	g.Log().Warningf(s.Ctx, "StarCharge100EPcs SetStatus() not support!")
	return nil
}

func (s *StarCharge100EPcs) SetGridMode(mode c_base.EGridMode) error {
	return nil
}

func (s *StarCharge100EPcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	value, err := s.GetUintValue(InverterOperationStatus)
	if err != nil {
		return c_base.EPcsStatusUnknown, err
	}

	switch value {
	// 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error
	case 0, 1:
		// 等待设备启动算是同步中状态
		return c_base.EPcsStatusSync, nil
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

func (s *StarCharge100EPcs) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridOn, nil
}

func (s *StarCharge100EPcs) SetPower(power int32) error {
	s.targetPower = power
	return s.WriteSingleRegister(ActivePowerSetting, power)
}

func (s *StarCharge100EPcs) SetReactivePower(power int32) error {
	s.targetReactivePower = power
	return s.WriteSingleRegister(ReactivePowerSetting, power)
}

func (s *StarCharge100EPcs) SetPowerFactor(factor float32) error {
	g.Log().Warningf(s.Ctx, "StarCharge100EPcs SetPowerFactor() not support!")
	return nil
}

func (s *StarCharge100EPcs) GetTargetPower() int32 {
	return s.targetPower
}

func (s *StarCharge100EPcs) GetTargetReactivePower() int32 {
	return s.targetReactivePower
}

func (s *StarCharge100EPcs) GetTargetPowerFactor() float32 {
	return -1
}

func (s *StarCharge100EPcs) GetPower() (float64, error) {
	return s.GetFloat64Value(TotalActivePowerInverterSide)
}

func (s *StarCharge100EPcs) GetApparentPower() (float64, error) {
	return s.GetFloat64Value(TotalApparentPowerInverterSide)
}

func (s *StarCharge100EPcs) GetReactivePower() (float64, error) {
	return s.GetFloat64Value(TotalReactivePowerInverterSide)
}

func (s *StarCharge100EPcs) GetRatedPower() uint32 {
	return 100
}

func (s *StarCharge100EPcs) GetMaxInputPower() (float32, error) {
	return 100, nil
}

func (s *StarCharge100EPcs) GetMaxOutputPower() (float32, error) {
	return 100, nil
}

func (s *StarCharge100EPcs) GetTodayIncomingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, DailyBatteryDischargeEnergy)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (s *StarCharge100EPcs) GetHistoryIncomingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, TotalBatteryDischargeEnergy)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (s *StarCharge100EPcs) GetTodayOutgoingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, DailyBatteryChargeEnergy)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (s *StarCharge100EPcs) GetHistoryOutgoingQuantity() (float64, error) {
	read, err := s.ReadGroupSync(SyGroupStatistics, true, TotalBatteryChargeEnergy)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}
