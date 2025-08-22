package c_none

import (
	"common/c_base"
)

type sNoneLoad struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneLoad) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneLoad) Shutdown() {

}

func (s *sNoneLoad) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceLoad
}

func (s *sNoneLoad) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneLoad) SetPower(power float64) error {
	return NoneErr
}

func (s *sNoneLoad) SetReactivePower(power float64) error {
	return NoneErr
}

func (s *sNoneLoad) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneLoad) GetTargetPower() float64 {
	return 0
}

func (s *sNoneLoad) GetTargetReactivePower() float64 {
	return 0
}

func (s *sNoneLoad) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneLoad) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetRatedPower() uint32 {
	return 0
}

func (s *sNoneLoad) GetMaxInputPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneLoad) GetMaxOutputPower() (float64, error) {
	return 0, NoneErr
}
