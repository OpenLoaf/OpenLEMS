package fire_control_v1

import (
	"canbus/p_canbus"
	"common/c_base"
)

var (
	AlarmAndFault = p_canbus.SCanbusTask{
		Name:        "AlarmAndFault",
		GetCanbusID: func(map[string]any) uint32 { return 0x1C000109 },
		Metas:       []*c_base.Meta{DetectorNumber, TemperatureAlarm, SmokeAlarm, COAlarm, H2Alarm, VOCAlarm, Level1Alarm, Level2Alarm, DetectorFault, GasCapsuleHardwareFault, MainCircuitVoltageFault, ReportNumber},
	}
	Detail = p_canbus.SCanbusTask{
		Name:        "Detail",
		GetCanbusID: func(map[string]any) uint32 { return 0x1C00010A },
		Metas:       []*c_base.Meta{DetectorNumber_V2, AlarmLevel, COConcentration, TemperatureData, SmokeAlarm_V2, DetectorOfflineFault, SensorFault, VOCAlarm_V2, H2Alarm_V2, ReportNumber_V2},
	}
)
