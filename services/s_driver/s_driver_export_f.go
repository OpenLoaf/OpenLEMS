package s_driver

import (
	"common"
	"context"
	"s_driver/internal"
)

func NewDriverManagerImpl(parentCtx context.Context) common.IDeviceManager {
	return internal.NewSingleDriverManager(parentCtx)
}
