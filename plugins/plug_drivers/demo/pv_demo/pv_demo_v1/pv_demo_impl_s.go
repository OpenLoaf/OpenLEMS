package pv_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
)

type sPvDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sPvDemo) Init() error {
	s.RegisterTask(ReadTask)
	return nil
}

func (s *sPvDemo) Shutdown() {

}

func (s *sPvDemo) GetStatus() (c_enum.EPvStatus, error) {
	power, err := s.GetPower()
	if err != nil {
		return c_enum.EPvUnknown, err
	}
	if power == 0.0 {
		return c_enum.EPvStandby, nil
	}
	return c_enum.EPvGenerate, nil
}

func (s *sPvDemo) SetPower(power float64) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(PowerLimit, int32(power))
	})
}

func (s *sPvDemo) SetReactivePower(power float64) error {
	return c_base.NotSupport
}

func (s *sPvDemo) SetPowerFactor(factor float32) error {
	return c_base.NotSupport
}

func (s *sPvDemo) GetTargetPower() (float64, error) {
	return s.GetFromPointFloat64(PowerLimit)
}

func (s *sPvDemo) GetTargetReactivePower() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetTargetPowerFactor() (float32, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetPower() (float64, error) {
	return s.GetFromPointFloat64(Power)
}

func (s *sPvDemo) GetApparentPower() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetReactivePower() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetDcPower() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetDcVoltage() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetDcCurrent() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetTodayIncomingQuantity() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetHistoryIncomingQuantity() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetTodayOutgoingQuantity() (float64, error) {
	return 0, c_base.NotSupport
}

func (s *sPvDemo) GetHistoryOutgoingQuantity() (float64, error) {
	return s.GetFromPointFloat64(GeneratedEnergy)
}

func (s *sPvDemo) GetCapacity() (uint32, error) {
	return s.GetFromPointUint32(InstalledCapacity)
}

func (s *sPvDemo) GetTemperature() (uint32, error) {
	return s.GetFromPointUint32(Temperature)
}

func (s *sPvDemo) GetIrradiance() (uint32, error) {
	return s.GetFromPointUint32(InstalledCapacity)
}
