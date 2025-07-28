package fire_control_v1

import (
	"canbus/p_canbus"
	"common/c_base"
)

var (
	AlarmAndFault = p_canbus.SCanbusTask{
		Name:     "AlarmAndFault",
		CanbusID: 0x1C000109,
		Metas:    []*c_base.Meta{DetectorId, TemperatureAlarm, SmokeAlarm, COAlarm, H2Alarm, VOCAlarm, Level1Alarm, Level2Alarm, DetectorFault, GasCapsuleHardwareFault, MainCircuitVoltageFault, Reserved4, Reserved5, Reserved6, Reserved7, AlarmNumber},
	}
)
