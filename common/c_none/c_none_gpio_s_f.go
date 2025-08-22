package c_none

import (
	"common/c_base"
	"context"
)

type sNoneGpio struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneGpio) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneGpio) Shutdown() {

}

func (s *sNoneGpio) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceGpio
}

func (s *sNoneGpio) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneGpio) RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) {

}

func (s *sNoneGpio) GetStatus() *bool {
	return nil
}

func (s *sNoneGpio) SetHigh() error {
	return NoneErr
}

func (s *sNoneGpio) SetLow() error {
	return NoneErr
}
