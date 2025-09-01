package ess_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
	"github.com/shockerli/cvt"
)

type sEssDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sEssDemo) Init() error {
	s.RegisterTask(ReadTask)
	return nil
}

func (s *sEssDemo) Shutdown() {

}

func (s *sEssDemo) GetCellMinTemp() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetCellMaxTemp() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetCellAvgTemp() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetCellMinVoltage() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetCellMaxVoltage() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetCellAvgVoltage() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetSoc() (float32, error) {
	return s.GetFromPointFloat32(SOC)
}

func (s *sEssDemo) GetSoh() (float32, error) {
	return 100, nil
}

func (s *sEssDemo) GetCapacity() (uint32, error) {
	return s.GetFromPointUint32(EnergyCapacity)
}

func (s *sEssDemo) GetCycleCount() (uint, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) GetDcPower() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) SetReset() error {
	return s.SetPower(0)
}

func (s *sEssDemo) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusOff {
		_ = s.SetPower(0)

		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(DeviceControl, 0)
		})
	}
	return nil
}

func (s *sEssDemo) SetGridMode(mode c_base.EGridMode) error {
	return c_base.NotSupport
}

func (s *sEssDemo) GetStatus() (c_base.EEnergyStoreStatus, error) {
	v, err := s.GetFromProtocol(func(protocol c_proto.IModbusProtocol) (any, error) {
		value, err := protocol.GetUintValue(Status)
		if err != nil {
			return c_base.EPcsStatusUnknown, err
		}

		if v, err := cvt.Uint8E(value); err == nil {
			switch v {
			case 0:
				return c_base.EPcsStatusOff, nil
			case 1:
				return c_base.EPcsStatusStandby, nil
			case 2:
				return c_base.EPcsStatusCharge, nil
			case 3:
				return c_base.EPcsStatusDischarge, nil
			case 4:
				return c_base.EPcsStatusFault, nil
			}
		}
		return c_base.EPcsStatusUnknown, err
	})
	return v.(c_base.EEnergyStoreStatus), err
}

func (s *sEssDemo) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridOn, nil
}

func (s *sEssDemo) SetPower(power int32) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(TargetPower, power)
	})
}

func (s *sEssDemo) SetReactivePower(power int32) error {
	return c_base.NotSupport
}

func (s *sEssDemo) SetPowerFactor(factor float32) error {
	return c_base.NotSupport
}

func (s *sEssDemo) GetTargetPower() int32 {
	targetPower, err := s.GetFromProtocolInt32(func(protocol c_proto.IModbusProtocol) (any, error) {
		return protocol.GetValue(TargetPower)
	})
	if err != nil {
		return 0
	}
	return targetPower
}

func (s *sEssDemo) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetRatedPower() (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssDemo) GetFireEnvTemperature() (float64, error) {
	return 0.0, c_base.NotSupport
}

func (s *sEssDemo) GetCarbonMonoxideConcentration() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sEssDemo) HasSmoke() (bool, error) {
	return false, nil
}
