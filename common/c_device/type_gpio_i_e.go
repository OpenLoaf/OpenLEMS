package c_device

import (
	"common/c_base"
	"context"
)

type IGpio interface {
	c_base.IDevice

	RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) // 状态变化处理

	GetStatus() *bool // 是否是高电平

	SetHigh() error // 设置为高电平
	SetLow() error  // 设置为低电平
}
