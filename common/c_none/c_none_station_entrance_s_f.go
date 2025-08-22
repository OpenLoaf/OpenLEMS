package c_none

import (
	"common/c_base"
)

type sNoneStationEntrance struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneStationEntrance) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneStationEntrance) Shutdown() {

}

func (s *sNoneStationEntrance) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceNone
}

func (s *sNoneStationEntrance) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneStationEntrance) GetGridFrequency() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetUa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetUb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetUc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetIa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetIb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetIc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetPa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetPb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetPc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) SetPower(power float64) error {
	return NoneErr
}

func (s *sNoneStationEntrance) SetReactivePower(power float64) error {
	return NoneErr
}

func (s *sNoneStationEntrance) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneStationEntrance) GetTargetPower() float64 {
	return 0
}

func (s *sNoneStationEntrance) GetTargetReactivePower() float64 {
	return 0
}

func (s *sNoneStationEntrance) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneStationEntrance) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetPtCt() (float32, float32, error) {
	return 0, 0, NoneErr
}

func (s *sNoneStationEntrance) GetFrequency() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetPowerFactor() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEntrance) GetVoltage() (float32, float32, float32, error) {
	return 0, 0, 0, NoneErr
}

func (s *sNoneStationEntrance) GetCurrent() (float32, float32, float32, error) {
	return 0, 0, 0, NoneErr
}

func (s *sNoneStationEntrance) GetAllowControl() bool {
	return false
}

func (s *sNoneStationEntrance) GetChildren() []c_base.IDevice {
	return nil
}
