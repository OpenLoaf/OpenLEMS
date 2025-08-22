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
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellMinTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellMaxTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellAvgTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellMinVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellMaxVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCellAvgVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetSoc() (float32, error) {
	return 0.0, noneErr
}

func (s *sNoneBms) GetSoh() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCapacity() (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetCycleCount() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sNoneBms) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
