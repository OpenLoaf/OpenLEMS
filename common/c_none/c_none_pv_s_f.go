package c_none

import (
	"common/c_base"
)

type sNonePv struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNonePv) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNonePv) Shutdown() {

}

func (s *sNonePv) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePv
}

func (s *sNonePv) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNonePv) GetGridFrequency() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetUa() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetUb() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetUc() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetIa() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetIb() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetIc() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetPa() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetPb() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetPc() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePv) SetPower(power float64) error {
	return NoneErr
}

func (s *sNonePv) SetReactivePower(power float64) error {
	return NoneErr
}

func (s *sNonePv) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNonePv) GetTargetPower() float64 {
	return 0
}

func (s *sNonePv) GetTargetReactivePower() float64 {
	return 0
}

func (s *sNonePv) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNonePv) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetDcPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetDcVoltage() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetDcCurrent() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePv) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}
