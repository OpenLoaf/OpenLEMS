package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"
	"time"
)

type sPcsStarCharge100E struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率
}

var _ c_type.IPcs = (*sPcsStarCharge100E)(nil)

func (s *sPcsStarCharge100E) Init() error {
	// 注册
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(
			GroupCommand,
			GroupPowerInfo,
			//GroupPhase,
			//GroupSerial, GroupGridModeSetting, GroupTemperature,
			GroupStatus,
		)
		return nil
	})

	return nil
}

func (s *sPcsStarCharge100E) GetFrequency() (*float64, error) {
	return s.GetFromPointFloat64(AverageFrequency)
}

func (s *sPcsStarCharge100E) GetIGBTTemperature() (*float32, error) {
	//return s.GetFromPointFloat32(c_default.VPointIGBTTemp)
	return nil, nil
}

func (s *sPcsStarCharge100E) Shutdown() {
	_ = s.SetPower(0)
	_ = s.SetStatus(c_enum.EPcsStatusOff)
	c_log.Infof(s.DeviceCtx, "销毁成功,设置PCS状态为Off!")
}

func (s *sPcsStarCharge100E) SetReset() error {
	c_log.Warningf(s.DeviceCtx, "sPcsStarCharge100E SetReset() not support!")
	return nil
}

func (s *sPcsStarCharge100E) SetStatus(status c_enum.EEnergyStoreStatus) error {
	if status == c_enum.EPcsStatusOff {
		_ = s.SetPower(0)

		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(OnOffCommand, 0)
		})
	}
	if status == c_enum.EPcsStatusStandby {
		// 这里文档是 On/off command: 0- Shutdown, 1- Startup, 2- Standby
		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(OnOffCommand, 1)
		})
	}
	return c_base.NotSupport
}

func (s *sPcsStarCharge100E) SetGridMode(mode c_enum.EGridMode) error {
	return c_base.NotSupport
}

func (s *sPcsStarCharge100E) GetStatus() (*c_enum.EEnergyStoreStatus, error) {
	v, err := s.GetFromProtocol(func(protocol c_proto.IModbusProtocol) (any, error) {
		value, err := protocol.GetUintValue(InverterOperationStatus)
		if err != nil {
			return c_enum.EPcsStatusUnknown, err
		}
		if value == nil {
			return c_enum.EPcsStatusUnknown, nil // 没有采集到数据
		}
		switch *value {
		// 0 - Waiting for the machine to start, 1 - Power on self check, 2 - Grid connected operation, 3 - Off grid operation, 4 - Reserved, 5 - General error
		case 0, 1:
			// 等待设备启动算是关机的状态
			return c_enum.EPcsStatusOff, nil
		case 2, 3:
			// 离网并网运行中时，说明设备正常。获取功率，如果获取功率失败，说明设备故障，获取成功后正为放电，负为充电
			power, err := s.GetPower()
			if err != nil {
				return c_enum.EPcsStatusFault, err
			}
			if power == nil {
				return c_enum.EPcsStatusStandby, nil
			}
			if *power > 0 {
				return c_enum.EPcsStatusDischarge, nil
			} else if *power < 0 {
				return c_enum.EPcsStatusCharge, nil
			} else {
				return c_enum.EPcsStatusStandby, nil
			}
		case 5:
			return c_enum.EPcsStatusFault, err
		}
		return c_enum.EPcsStatusFault, err
	})

	if err != nil {
		status := c_enum.EPcsStatusUnknown
		return &status, err
	}
	status := v.(c_enum.EEnergyStoreStatus)
	return &status, err
}

func (s *sPcsStarCharge100E) GetGridMode() (*c_enum.EGridMode, error) {
	mode := c_enum.EGridOn
	return &mode, nil
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

func (s *sPcsStarCharge100E) GetTargetPower() (*int32, error) {
	return &s.targetPower, nil
}

func (s *sPcsStarCharge100E) GetTargetReactivePower() (*int32, error) {
	return &s.targetReactivePower, nil
}

func (s *sPcsStarCharge100E) GetTargetPowerFactor() (*float32, error) {
	factor := float32(-1)
	return &factor, nil
}

func (s *sPcsStarCharge100E) GetPower() (*float64, error) {
	return s.GetFromPointFloat64(TotalActivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetApparentPower() (*float64, error) {
	return s.GetFromPointFloat64(TotalApparentPowerInverterSide)
}

func (s *sPcsStarCharge100E) GetReactivePower() (*float64, error) {
	return s.GetFromPointFloat64(TotalReactivePowerInverterSide)
}

func (s *sPcsStarCharge100E) GetRatedPower() (*uint32, error) {
	power := uint32(100)
	return &power, nil
}

func (s *sPcsStarCharge100E) GetMaxInputPower() (*float32, error) {
	power := float32(100)
	return &power, nil
}

func (s *sPcsStarCharge100E) GetMaxOutputPower() (*float32, error) {
	power := float32(100)
	return &power, nil
}

func (s *sPcsStarCharge100E) GetTodayIncomingQuantity() (*float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(DailyBatteryDischargeEnergy, c_enum.EMqHoldingRegisters, time.Minute, true)
	})
}

func (s *sPcsStarCharge100E) GetHistoryIncomingQuantity() (*float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(TotalBatteryDischargeEnergy, c_enum.EMqHoldingRegisters, time.Minute, true)
	})
}

func (s *sPcsStarCharge100E) GetTodayOutgoingQuantity() (*float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(DailyBatteryChargeEnergy, c_enum.EMqHoldingRegisters, time.Minute, true)
	})
}

func (s *sPcsStarCharge100E) GetHistoryOutgoingQuantity() (*float64, error) {
	return s.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.ReadSingleSync(TotalBatteryChargeEnergy, c_enum.EMqHoldingRegisters, time.Minute, true)
	})
}

// 实现新的IDevice接口方法
func (s *sPcsStarCharge100E) GetDevicePoints() []c_base.IPoint {
	// 返回Modbus协议点位
	return []c_base.IPoint{
		InverterOperationStatus,
		TotalActivePowerInverterSide,
		TotalApparentPowerInverterSide,
		TotalReactivePowerInverterSide,
		AverageFrequency,
		ActivePowerSetting,
		ReactivePowerSetting,
		OnOffCommand,
		// 可以继续添加其他协议点位
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sPcsStarCharge100E) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		TotalActivePowerInverterSide, // 总有功功率 - 核心运行参数
		InverterOperationStatus,      // 运行状态 - 设备状态信息
		AverageFrequency,             // 频率 - 电网质量指标
	}
}
