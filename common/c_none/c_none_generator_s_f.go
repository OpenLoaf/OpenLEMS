package c_none

import (
	"common/c_base"
)

type sNoneGenerator struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneGenerator) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneGenerator) Shutdown() {

}

func (s *sNoneGenerator) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceGenerator
}

func (s *sNoneGenerator) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneGenerator) GetGridFrequency() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetUa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetUb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetUc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetIa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetIb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetIc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetPa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetPb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetPc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) SetPower(power float64) error {
	return NoneErr
}

func (s *sNoneGenerator) SetReactivePower(power float64) error {
	return NoneErr
}

func (s *sNoneGenerator) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneGenerator) GetTargetPower() float64 {
	return 0
}

func (s *sNoneGenerator) GetTargetReactivePower() float64 {
	return 0
}

func (s *sNoneGenerator) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneGenerator) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetRatedPower() uint32 {
	return 0
}

func (s *sNoneGenerator) GetMaxInputPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneGenerator) GetMaxOutputPower() (float64, error) {
	return 0, NoneErr
}
