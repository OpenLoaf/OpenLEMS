package common

import (
	"common/c_base"
	"common/internal/internal_storage"
	"context"
)

func RegisterStorageInstance(builder func(ctx context.Context) c_base.IStorage) {
	internal_storage.RegisterInstance(builder)
}

func CloseStorage() {
	internal_storage.Close()
}

func GetStorageInstance() c_base.IStorage {
	return internal_storage.GetInstance()
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal_storage.GetInstance().RegisterDriver(storageIntervalSec, driver)
}
