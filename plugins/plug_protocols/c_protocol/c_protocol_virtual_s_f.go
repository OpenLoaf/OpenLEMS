package c_protocol

import (
	"common/c_base"
	"time"
)

type SVirtualProtocol struct {
	deviceConfig   *c_base.SDeviceConfig
	protocolConfig *c_base.SProtocolConfig
}

func NewVirtualProtocol(deviceConfig *c_base.SDeviceConfig, protocolConfig *c_base.SProtocolConfig) c_base.IProtocol {
	return &SVirtualProtocol{deviceConfig: deviceConfig, protocolConfig: protocolConfig}
}

func (s *SVirtualProtocol) ResetAlarm() {

}

func (s *SVirtualProtocol) TriggerAlarm(alarm *c_base.SAlarmDetail) {

}

func (s *SVirtualProtocol) RegisterMonitorChan(details chan<- *c_base.SAlarmDetail) {
}

func (s *SVirtualProtocol) GetAlarmLevel() c_base.EAlarmLevel {
	return c_base.ENone
}

func (s *SVirtualProtocol) GetAlarmDetails() []*c_base.SAlarmDetail {
	return nil
}

func (s *SVirtualProtocol) GetMonitorChan() chan<- *c_base.SAlarmDetail {
	return nil
}

func (s *SVirtualProtocol) IsPhysical() bool {
	return false
}

func (s *SVirtualProtocol) GetDeviceConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *SVirtualProtocol) GetMetaValueList() []*c_base.MetaValueWrapper {
	return nil
}

func (s *SVirtualProtocol) GetLastUpdateTime() *time.Time {
	return nil
}

func (s *SVirtualProtocol) ProtocolListen() {

}

func (s *SVirtualProtocol) IsActivate() bool {
	return true
}

func (s *SVirtualProtocol) GetProtocolConfig() *c_base.SProtocolConfig {
	return s.protocolConfig
}
