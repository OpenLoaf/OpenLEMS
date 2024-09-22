package services

import (
	"context"
	"services/internal/internal_cmd_device"
)

func NewDeviceCmd(ctx context.Context) *internal_cmd_device.SDeviceCmd {
	return internal_cmd_device.NewDeviceCmd(ctx)
}
