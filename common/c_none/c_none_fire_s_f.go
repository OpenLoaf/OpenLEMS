package c_none

import (
	"common/c_base"
)

type sNoneFire struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneFire) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneFire) Shutdown() {

}

func (s *sNoneFire) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceFire
}

func (s *sNoneFire) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneFire) GetFireEnvTemperature() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneFire) GetCarbonMonoxideConcentration() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneFire) HasSmoke() (bool, error) {
	return false, NoneErr
}
