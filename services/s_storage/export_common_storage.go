package s_storage

import (
	"common/c_base"
	"golang.org/x/net/context"
	"s_storage/internal"
)

func NewSingleStorageManager(parentCtx context.Context, storage c_base.IStorage) {
	internal.NewSingleInstance(parentCtx, storage)
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal.StorageManagerInstance.RegisterDriver(storageIntervalSec, driver)
}
