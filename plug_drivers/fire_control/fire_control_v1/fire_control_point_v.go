package fire_control_v1

import (
	"common/c_base"
)

var (
	DetectorId = &c_base.Meta{Name: "DetectorId", Cn: "探测器编号", Addr: 1, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "01：1 号探测器"}

	TemperatureAlarm = &c_base.Meta{Name: "TemperatureAlarm", Cn: "温度报警", Addr: 2, ReadType: c_base.RBit1, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	SmokeAlarm       = &c_base.Meta{Name: "SmokeAlarm", Cn: "烟雾报警", Addr: 2, ReadType: c_base.RBit2, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	COAlarm          = &c_base.Meta{Name: "COAlarm", Cn: "CO报警", Addr: 2, ReadType: c_base.RBit3, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	H2Alarm          = &c_base.Meta{Name: "H2Alarm", Cn: "H2报警", Addr: 2, ReadType: c_base.RBit4, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	VOCAlarm         = &c_base.Meta{Name: "VOCAlarm", Cn: "VOC报警", Addr: 2, ReadType: c_base.RBit5, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	Level1Alarm      = &c_base.Meta{Name: "Level1Alarm", Cn: "1级报警", Addr: 2, ReadType: c_base.RBit6, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	Level2Alarm      = &c_base.Meta{Name: "Level2Alarm", Cn: "2级报警", Addr: 2, ReadType: c_base.RBit7, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}

	DetectorFault           = &c_base.Meta{Name: "DetectorFault", Cn: "探测器故障", Addr: 3, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	GasCapsuleHardwareFault = &c_base.Meta{Name: "GasCapsuleHardwareFault", Cn: "气溶胶硬件故障", Addr: 3, ReadType: c_base.RBit2, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	MainCircuitVoltageFault = &c_base.Meta{Name: "MainCircuitVoltageFault", Cn: "主电欠压故障", Addr: 3, ReadType: c_base.RBit3, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}

	Reserved4   = &c_base.Meta{Name: "Reserved4", Cn: "保留", Addr: 4, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "00"}
	Reserved5   = &c_base.Meta{Name: "Reserved5", Cn: "保留", Addr: 5, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "00"}
	Reserved6   = &c_base.Meta{Name: "Reserved6", Cn: "保留", Addr: 6, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "00"}
	Reserved7   = &c_base.Meta{Name: "Reserved7", Cn: "保留", Addr: 7, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "00"}
	AlarmNumber = &c_base.Meta{Name: "AlarmNumber", Cn: "报警编号", Addr: 8, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "0-255自增"}
)
