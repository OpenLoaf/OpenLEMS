package p_gpio_sysfs

import (
	"context"
	"ems-plan/c_base"
)

type IGpioSysfsProtocol interface {
	c_base.IProtocol

	RegisterHighHandler(fc func(ctx context.Context))
	RegisterLowHandler(fc func(ctx context.Context))

	IsHigh() bool // 是否是高电平
	IsLow() bool  // 是否是低电平

	SetHigh() error // 设置为高电平
	SetLow() error  // 设置为低电平
}
