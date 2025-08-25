package s_storage

import (
	"common/c_base"
	"s_storage/internal"

	"golang.org/x/net/context"
)

func NewSingleStorageManager(parentCtx context.Context, storage c_base.IStorage) {
	internal.NewSingleInstance(parentCtx, storage)
}

func RegisterStorageDriver(deviceConfig *c_base.SDeviceConfig) {
	internal.StorageManagerInstance.RegisterDriver(deviceConfig)
}
