package c_none

import (
	"common/c_base"
)

type sNoneAmmeter struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneAmmeter) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneAmmeter) Shutdown() {

}

func (s *sNoneAmmeter) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceAmmeter
}

func (s *sNoneAmmeter) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneAmmeter) GetUa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetUb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetUc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetIa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetIb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetIc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPTotal() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetQa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetQb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetQc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetQTotal() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetSa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetSb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetSc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetSTotal() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPfa() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPfb() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPfc() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPfTotal() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetPtCt() (float32, float32, error) {
	return 0, 0, NoneErr
}

func (s *sNoneAmmeter) GetFrequency() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneAmmeter) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}
