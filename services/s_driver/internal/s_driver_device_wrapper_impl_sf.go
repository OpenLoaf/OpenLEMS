package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
)

type SDeviceWrapper struct {
	ctx context.Context

	deviceConfig   *c_base.SDeviceConfig
	driverInfo     *c_base.SDriverInfo
	protocolConfig *c_base.SProtocolConfig
	instance       c_base.IDevice
	deviceState    c_base.EServerState
}

func (s *SDeviceWrapper) GetDeviceConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *SDeviceWrapper) GetDriverInfo() *c_base.SDriverInfo {
	return s.driverInfo
}

func (s *SDeviceWrapper) GetProtocolConfig() *c_base.SProtocolConfig {
	return s.protocolConfig
}

func (s *SDeviceWrapper) GetDeviceDetail() *c_base.SDeviceDetail {
	return c_base.NewDeviceDetail(s.deviceConfig, s.driverInfo, s.protocolConfig, s.instance, s.deviceState)
}

func (s *SDeviceWrapper) GetDeviceInstance() c_base.IDevice {
	return s.instance
}

func (s *SDeviceWrapper) GetDeviceState() c_base.EServerState {
	return s.deviceState
}

func (s *SDeviceWrapper) Shutdown() {
	if s.instance != nil {
		s.instance.Shutdown()
		c_log.BizInfof(s.ctx, "设备关闭成功！")
	}
}

func (s *SDeviceWrapper) UpdateState(deviceState c_base.EServerState) {
	s.deviceState = deviceState
}
