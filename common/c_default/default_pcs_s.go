package c_default

import (
	"common/c_base"
	"time"
)

type SDefaultPcs struct {
}

func (S *SDefaultPcs) ResetAlarm() {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) TriggerAlarm(alarm *c_base.SAlarmDetail) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) RegisterMonitorChan(details chan<- *c_base.SAlarmDetail) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetAlarmLevel() c_base.EAlarmLevel {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetAlarmDetails() []*c_base.SAlarmDetail {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetMonitorChan() chan<- *c_base.SAlarmDetail {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) InitDevice(protocol c_base.IProtocol, deviceConfig *c_base.SDeviceConfig) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (S *SDefaultPcs) GetLastUpdateTime() *time.Time {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetMetaValueList() []*c_base.MetaValueWrapper {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetDeviceConfig() *c_base.SDeviceConfig {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetDescription() *c_base.SDriverDescription {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTelemetry(key string, instance any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetAllTelemetry(instance any) map[string]any {
	//TODO implement me
	panic("implement me")
}
