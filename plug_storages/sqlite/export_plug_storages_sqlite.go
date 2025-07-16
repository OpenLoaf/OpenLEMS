package sqlite

import (
	"context"
	"sqlite/internal"
	"sqlite/service"
)

// Init 初始化数据库表
func Init() {
	internal.Init()
}

func NewConfigManage(ctx context.Context) service.IConfigManage {
	return service.NewConfigManage(ctx)
}

func NewProtocolManage(ctx context.Context) service.IProtocolManage {
	return service.NewProtocolManage(ctx)
}

func NewDeviceManage(ctx context.Context) service.IDeviceManage {
	return service.NewDeviceManage(ctx)
}
