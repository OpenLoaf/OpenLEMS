package p_gpio_sysfs

import "common/c_enum"

type SProtocolGpioConfig struct {
	Period     uint               // 读取周期，单位毫秒. 0  表示默认1秒
	Level      c_enum.EAlarmLevel //  告警级别
	LowTrigger bool               // false：高电平触发 true：低电平触发
}
