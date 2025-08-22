package c_none

import (
	"common/c_base"
)

type sNoneCharge struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneCharge) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneCharge) Shutdown() {

}

func (s *sNoneCharge) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceChargePile
}

func (s *sNoneCharge) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneCharge) SetPower(power float64) error {
	return NoneErr
}

func (s *sNoneCharge) SetReactivePower(power float64) error {
	return NoneErr
}

func (s *sNoneCharge) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneCharge) GetTargetPower() float64 {
	return 0
}

func (s *sNoneCharge) GetTargetReactivePower() float64 {
	return 0
}

func (s *sNoneCharge) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneCharge) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetRatedPower() uint32 {
	return 0
}

func (s *sNoneCharge) GetMaxInputPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetMaxOutputPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneCharge) GetCarSoc() (float64, error) {
	return -1, NoneErr
}
