package c_none

import (
	"common/c_base"
	"common/c_type"
)

type sNoneCoolingAc struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneCoolingAc) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneCoolingAc) Shutdown() {

}

func (s *sNoneCoolingAc) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceCoolingAc
}

func (s *sNoneCoolingAc) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneCoolingAc) GetCoolingAcStatus() (c_type.ECoolingStatus, error) {
	return c_type.ECoolingStatusStop, NoneErr
}
