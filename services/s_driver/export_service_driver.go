package s_driver

import (
	"common/c_base"
	"context"
	"s_driver/internal"
	"s_driver/s_driver_interface"
)

func NewDeviceCmd(ctx context.Context) c_base.IService {
	return internal.NewDeviceCmd(ctx)
}

func GetDriverManager() s_driver_interface.IDriverManager {
	return internal.GetDriverManager()
}
