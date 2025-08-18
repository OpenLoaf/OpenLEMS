package s_storage

import (
	"common/c_base"
	"context"
	"s_storage/internal"
)

func RegisterStorageInstance(builder func(ctx context.Context) c_base.IStorage) {
	internal.RegisterInstance(builder)
}

func CloseStorage() {
	internal.Shutdown()
}

func GetStorageInstance() c_base.IStorage {
	return internal.GetInstance()
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal.GetInstance().RegisterDriver(storageIntervalSec, driver)
}
