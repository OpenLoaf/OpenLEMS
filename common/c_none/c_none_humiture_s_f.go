package c_none

import (
	"common/c_base"
)

type sNoneHumiture struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneHumiture) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneHumiture) Shutdown() {

}

func (s *sNoneHumiture) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceHumiture
}

func (s *sNoneHumiture) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneHumiture) GetTemperature() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneHumiture) GetHumidity() (float64, error) {
	return 0, NoneErr
}
