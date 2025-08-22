package c_none

import (
	"common/c_base"
)

type sNonePcs struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNonePcs) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNonePcs) Shutdown() {
}

func (s *sNonePcs) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sNonePcs) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNonePcs) SetReset() error {
	return NoneErr
}

func (s *sNonePcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	return NoneErr
}

func (s *sNonePcs) SetGridMode(mode c_base.EGridMode) error {
	return NoneErr
}

func (s *sNonePcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return c_base.EPcsStatusUnknown, NoneErr
}

func (s *sNonePcs) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridUnknown, NoneErr
}

func (s *sNonePcs) SetPower(power int32) error {
	return NoneErr
}

func (s *sNonePcs) SetReactivePower(power int32) error {
	return NoneErr
}

func (s *sNonePcs) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNonePcs) GetTargetPower() int32 {
	return 0
}

func (s *sNonePcs) GetTargetReactivePower() int32 {
	return 0
}

func (s *sNonePcs) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNonePcs) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetRatedPower() int32 {
	return 0
}

func (s *sNonePcs) GetMaxInputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetMaxOutputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNonePcs) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}
