package s_driver

import (
	"common"
	"context"
	"s_driver/internal"
	"s_driver/s_driver_interface"
)

func NewDeviceCmd(ctx context.Context) common.IDeviceCmd {
	return internal.NewDeviceCmd(ctx)
}

func GetDriverManager() s_driver_interface.IDriverManager {
	return internal.GetDriverManager()
}
