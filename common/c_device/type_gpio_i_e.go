package c_device

import (
	"context"
	"ems-plan/c_base"
)

type IGpio interface {
	c_base.IDriver

	RegisterHandler(handler func(ctx context.Context, status bool)) // 状态变化处理

	IsHigh() bool // 是否是高电平
	IsLow() bool  // 是否是低电平

	SetHigh() error // 设置为高电平
	SetLow() error  // 设置为低电平
}
