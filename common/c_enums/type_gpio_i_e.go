package c_type

import (
	"common/c_base"
	"context"
)

type IGpio interface {
	c_base.IDriver

	RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) // 状态变化处理

	GetGpioStatus() *bool // 是否是高电平

	SetHigh() error // 设置为高电平
	SetLow() error  // 设置为低电平
}
