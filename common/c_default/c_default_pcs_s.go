package c_default

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
)

type SDevicePcs[P c_base.IProtocol] struct {
	*c_device.SRealDeviceImpl[P]
}

func (s *SDevicePcs[P]) GetStatus() (*c_enum.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetGridMode() (*c_enum.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetTargetPower() (*int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetTargetReactivePower() (*int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetTargetPowerFactor() (*float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetPower() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetApparentPower() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetReactivePower() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetRatedPower() (*uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetMaxInputPower() (*float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetMaxOutputPower() (*float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetTodayIncomingQuantity() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetHistoryIncomingQuantity() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetTodayOutgoingQuantity() (*float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SDevicePcs[P]) GetHistoryOutgoingQuantity() (*float64, error) {
	//TODO implement me
	panic("implement me")
}
