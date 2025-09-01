package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_proto"
	"time"

	"github.com/pkg/errors"
)

type sPcsStarCharge100E struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率
}

func (s *sPcsStarCharge100E) Init() error {
	// 注册
	s.RegisterTask(
		GroupCommand,
		GroupPowerInfo,
		//GroupPhase,
		//GroupSerial, GroupGridModeSetting, GroupTemperature,
		GroupStatus,
	)

	return nil
}

func (s *sPcsStarCharge100E) Shutdown() {
	_ = s.SetPower(0)
	_ = s.SetStatus(c_base.EPcsStatusOff)
	c_log.Infof(s.DeviceCtx, "销毁成功,设置PCS状态为Off!")
}

func (s *sPcsStarCharge100E) GetFunctionList() []*c_base.STelemetry {
	return nil
}

func (s *sPcsStarCharge100E) SetReset() error {
	c_log.Warningf(s.DeviceCtx, "sPcsStarCharge100E SetReset() not support!")
	return nil
}

func (s *sPcsStarCharge100E) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusOff {
		_ = s.SetPower(0)

		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(OnOffCommand, 0)
		})
	}
	if status == c_base.EPcsStatusStandby {
		// 这里文档是 On/off command: 0- Shutdown, 1- Startup, 2- Standby
		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(OnOffCommand, 1)
		})
	}
	return errors.Errorf("sPcsStarCharge100E SetStatus status not support!")
}

func (s *sPcsStarCharge100E) SetGridMode(mode c_base.EGridMode) error {
	return errors.Errorf("sPcsStarCharge100E SetGridMode status not support!")
}

func (s *sPcsStarCharge100E) GetStatus() (c_base.EEnergyStoreStatus, error) {
	v, err := s.GetFromProtocol(func(protocol c_proto.IModbusProtocol) (any, error) {
		value, err := protocol.GetUintValue(InverterOperationStatus)
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
	})

	if err != nil {
		return c_base.EPcsStatusUnknown, err
	}
	return v.(c_base.EEnergyStoreStatus), err
}

func (s *sPcsStarCharge100E) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridOn, nil
}

func (s *sPcsStarCharge100E) SetPower(power int32) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		s.targetPower = power
		return protocol.WriteSingleRegister(ActivePowerSetting, power)
	})

}

func (s *sPcsStarCharge100E) SetReactivePower(power int32) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		s.targetReactivePower = power
		return protocol.WriteSingleRegister(ReactivePowerSetting, power)
	})
}

func (s *sPcsStarCharge100E) SetPowerFactor(factor float32) error {
	c_log.Warningf(s.DeviceCtx, "sPcsStarCharge100E SetPowerFactor() not support!")
	return nil
}

func (s *sPcsStarCharge100E) GetTargetPower() (int32, error) {
	return s.targetPower, nil
}

func (s *sPcsStarCharge100E) GetTargetReactivePower() (int32, error) {
	return s.targetReactivePower, nil
}

func (s *sPcsStarCharge100E) GetTargetPowerFactor() (float32, error) {
	return -1, nil
}

func (s *sPcsStarCharge100E) GetPower() (float64, error) {
	return s.GetFromPointFloat64(TotalActivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetApparentPower() (float64, error) {
	return s.GetFromPointFloat64(TotalApparentPowerInverterSide)
}

func (s *sPcsStarCharge100E) GetReactivePower() (float64, error) {
	return s.GetFromPointFloat64(TotalReactivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetRatedPower() (uint32, error) {
	return 100, nil
}

func (s *sPcsStarCharge100E) GetMaxInputPower() (float32, error) {
	return 100, nil
}

func (s *sPcsStarCharge100E) GetMaxOutputPower() (float32, error) {
	return 100, nil
}

func (s *sPcsStarCharge100E) GetTodayIncomingQuantity() (float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(DailyBatteryDischargeEnergy, c_proto.EMqHoldingRegisters, time.Minute, false)
	})
}

func (s *sPcsStarCharge100E) GetHistoryIncomingQuantity() (float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(TotalBatteryDischargeEnergy, c_proto.EMqHoldingRegisters, time.Minute, false)
	})
}

func (s *sPcsStarCharge100E) GetTodayOutgoingQuantity() (float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(DailyBatteryChargeEnergy, c_proto.EMqHoldingRegisters, time.Minute, false)
	})
}

func (s *sPcsStarCharge100E) GetHistoryOutgoingQuantity() (float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(TotalBatteryChargeEnergy, c_proto.EMqHoldingRegisters, time.Minute, false)
	})
}
