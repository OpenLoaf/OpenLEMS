package c_type

import (
	"common/c_base"
)

type IGpioOut interface {
	c_base.IDriver

	RegisterHandler(handler func(status bool, isChange bool)) // 状态变化处理
	GetStatus() *bool                                         // 是否是高电平
	SetHigh() error                                           // 设置为高电平
	SetLow() error                                            // 设置为低电平
}

type IGpioIn interface {
	c_base.IDriver

	RegisterHandler(handler func(status bool, isChange bool)) // 状态变化处理
	GetStatus() *bool                                         // 是否是高电平
}
