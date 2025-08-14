package fire_control_v1

import (
	"common/c_base"
)

var (
	DetectorNumber          = &c_base.Meta{Name: "DetectorNumber", Cn: "探测器编号", Addr: 0, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "01: 1号探测器"}
	TemperatureAlarm        = &c_base.Meta{Name: "TemperatureAlarm", Cn: "温度报警", Addr: 1, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	SmokeAlarm              = &c_base.Meta{Name: "SmokeAlarm", Cn: "烟雾报警", Addr: 1, ReadType: c_base.RBit1, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	COAlarm                 = &c_base.Meta{Name: "COAlarm", Cn: "CO报警", Addr: 1, ReadType: c_base.RBit2, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	H2Alarm                 = &c_base.Meta{Name: "H2Alarm", Cn: "H2报警", Addr: 1, ReadType: c_base.RBit3, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	VOCAlarm                = &c_base.Meta{Name: "VOCAlarm", Cn: "VOC报警", Addr: 1, ReadType: c_base.RBit4, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	Level1Alarm             = &c_base.Meta{Name: "Level1Alarm", Cn: "1级报警", Addr: 1, ReadType: c_base.RBit5, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	Level2Alarm             = &c_base.Meta{Name: "Level2Alarm", Cn: "2级报警", Addr: 1, ReadType: c_base.RBit6, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	DetectorFault           = &c_base.Meta{Name: "DetectorFault", Cn: "探测器故障", Addr: 2, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	GasCapsuleHardwareFault = &c_base.Meta{Name: "GasCapsuleHardwareFault", Cn: "气溶胶硬件故障", Addr: 2, ReadType: c_base.RBit1, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	MainCircuitVoltageFault = &c_base.Meta{Name: "MainCircuitVoltageFault", Cn: "主电欠压故障", Addr: 2, ReadType: c_base.RBit2, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	ReportNumber            = &c_base.Meta{Name: "ReportNumber", Cn: "报文编号", Addr: 7, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "0-255自增"}
)

var (
	DetectorNumber_V2    = &c_base.Meta{Name: "DetectorNumber_V2", Cn: "探测器编号", Addr: 0, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "01: 1号探测器"}
	AlarmLevel           = &c_base.Meta{Name: "AlarmLevel", Cn: "报警等级", Addr: 1, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "0x00 正常 0x01 一级报警 0x02 二级报警", Trigger: c_base.IsNotZero}
	COConcentration      = &c_base.Meta{Name: "COConcentration", Cn: "CO浓度", Addr: 2, ReadType: c_base.RUint16, SystemType: c_base.SUint16, Unit: "ppm", Factor: 10, Desc: "一氧化碳气体浓度数据, 数据输出范围0~1000ppm, 实际输出值缩小10倍 例如: 数据为10时: 100ppm"}
	TemperatureData      = &c_base.Meta{Name: "TemperatureData", Cn: "温度数据", Addr: 3, ReadType: c_base.RInt16, SystemType: c_base.SInt16, Unit: "°C", Offset: -40, Desc: "温度数据, 测量温度范围: -40°C ~ +150°C 数据偏移40, 计算方式: 实际温度值 = 温度数据-40 例如: 温度数据为0时: -40°C 温度数据为40时: 0°C 温度数据为165时: 125°C"}
	SmokeAlarm_V2        = &c_base.Meta{Name: "SmokeAlarm_V2", Cn: "烟雾报警", Addr: 4, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	DetectorOfflineFault = &c_base.Meta{Name: "DetectorOfflineFault", Cn: "探测器掉线故障", Addr: 5, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	SensorFault          = &c_base.Meta{Name: "SensorFault", Cn: "传感器故障", Addr: 5, ReadType: c_base.RBit1, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	VOCAlarm_V2          = &c_base.Meta{Name: "VOCAlarm_V2", Cn: "VOC报警", Addr: 6, ReadType: c_base.RBit0, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	H2Alarm_V2           = &c_base.Meta{Name: "H2Alarm_V2", Cn: "H2报警", Addr: 6, ReadType: c_base.RBit4, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	ReportNumber_V2      = &c_base.Meta{Name: "ReportNumber_V2", Cn: "报文编号", Addr: 7, ReadType: c_base.RUint8, SystemType: c_base.SUint8, Unit: "", Desc: "0-255自增"}
)
