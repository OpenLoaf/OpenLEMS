package driver

import (
	"common/c_base"
	"context"
	"driver/internal"
)

func NewDeviceCmd(ctx context.Context) c_base.IService {
	return internal.NewDeviceCmd(ctx)
}

func NewDriverManager() IDriverManager {
	return internal.NewDriverManager()
}
