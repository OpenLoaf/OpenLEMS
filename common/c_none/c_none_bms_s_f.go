package c_none

import (
	"common/c_base"
	"common/c_device"
)

type sNoneBms struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneBms) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneBms) Shutdown() {

}

func (s *sNoneBms) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (s *sNoneBms) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneBms) SetReset() error {
	return NoneErr
}

func (s *sNoneBms) SetBmsStatus(status c_device.EBmsStatus) error {
	return NoneErr
}

func (s *sNoneBms) GetCellMinTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCellMaxTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCellAvgTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCellMinVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCellMaxVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCellAvgVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	return c_device.EBmsStatusUnknown, NoneErr
}

func (s *sNoneBms) GetSoc() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneBms) GetSoh() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneBms) GetCapacity() (uint32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetCycleCount() (uint, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetRatedPower() int32 {
	return 0
}

func (s *sNoneBms) GetMaxInputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetMaxOutputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetDcPower() (float64, error) {
	return 0.0, NoneErr
}

func (s *sNoneBms) GetDcVoltage() (float64, error) {
	return 0.0, NoneErr
}

func (s *sNoneBms) GetDcCurrent() (float64, error) {
	return 0.0, NoneErr
}

func (s *sNoneBms) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneBms) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}
