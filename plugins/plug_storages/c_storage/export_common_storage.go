package c_storage

import (
	"c_storage/internal"
	"common/c_base"
	"context"
)

func RegisterStorageInstance(builder func(ctx context.Context) c_base.IStorage) {
	internal.RegisterInstance(builder)
}

func CloseStorage() {
	internal.Close()
}

func GetStorageInstance() c_base.IStorage {
	return internal.GetInstance()
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal.GetInstance().RegisterDriver(storageIntervalSec, driver)
}
