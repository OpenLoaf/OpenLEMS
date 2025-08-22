package c_none

import (
	"common/c_base"
	"common/c_type"
)

type sNoneCoolingLiquid struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneCoolingLiquid) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneCoolingLiquid) Shutdown() {

}

func (s *sNoneCoolingLiquid) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceCoolingLiquid
}

func (s *sNoneCoolingLiquid) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneCoolingLiquid) GetLiquidCoolingStatus() (c_type.ECoolingStatus, error) {
	return c_type.ECoolingStatusStop, NoneErr
}

func (s *sNoneCoolingLiquid) GetInputWaterPressure() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneCoolingLiquid) GetInputWaterTemperature() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneCoolingLiquid) GetOutputWaterPressure() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneCoolingLiquid) GetOutputWaterTemperature() (float32, error) {
	return 0, NoneErr
}
