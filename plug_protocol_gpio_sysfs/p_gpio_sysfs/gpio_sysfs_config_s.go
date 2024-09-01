package p_gpio_sysfs

import "ems-plan/c_base"

type SGpioSysfsDeviceConfig struct {
	Direction  EGpioDirection     `json:"direction" dc:"GPIO方向IN/OUT" name:"direction" brief:"GPIO方向"`
	Path       string             `json:"path" dc:"GPIO路径"`
	ExportPath string             `json:"exportPath" dc:"执行Export的路径"`
	ExportPort int                `json:"exportPort" dc:"执行Export的端口"`
	Level      c_base.EAlarmLevel `json:"level" dc:"告警级别" name:"level" brief:"告警级别"`
	Reverse    bool               `json:"reverse" dc:"是否反转" name:"reverse" brief:"是否反转"`
}
