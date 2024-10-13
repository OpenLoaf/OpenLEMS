package services

import (
	"common/c_base"
	"context"
	"services/internal/internal_cmd_device"
)

func NewDeviceCmd(ctx context.Context) c_base.IService {
	return internal_cmd_device.NewDeviceCmd(ctx)
}
