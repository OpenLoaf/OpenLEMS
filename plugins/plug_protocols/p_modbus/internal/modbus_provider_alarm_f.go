package internal

import "common/c_base"

func (p *ModbusProtocolProvider) GetAlarmList() []*c_base.SAlarmDetail {
	//TODO implement me
	panic("implement me")
}

func (p *ModbusProtocolProvider) RegisterAlarmTriggerFunc(handler func(maxAlarm c_base.EAlarmLevel, nowAlarm *c_base.SAlarmDetail)) {
	//TODO implement me
	panic("implement me")
}

func (p *ModbusProtocolProvider) RegisterAlarmRemoveFunc(handler func(maxAlarm c_base.EAlarmLevel, nowAlarm *c_base.SAlarmDetail)) {
	//TODO implement me
	panic("implement me")
}
