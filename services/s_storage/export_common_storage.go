package s_storage

import (
	"common/c_base"
	"s_storage/internal"

	"golang.org/x/net/context"
)

func NewSingleStorageManager(parentCtx context.Context, storage c_base.IStorage) {
	internal.NewSingleInstance(parentCtx, storage)
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal.StorageManagerInstance.RegisterDriver(storageIntervalSec, driver)
}
