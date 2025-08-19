package s_storage

import (
	"common/c_base"
	"s_storage/internal"
)

func RegisterStorageInstance(storage c_base.IStorage) {
	internal.RegisterInstance(storage)
}

func CloseStorage() {
	internal.Shutdown()
}

func GetStorageInstance() c_base.IStorage {
	mgr := internal.GetInstance()
	if mgr == nil {
		return nil
	}
	return mgr.IStorage
}

func RegisterStorageDriver(storageIntervalSec int32, driver c_base.IDevice) {
	internal.GetInstance().RegisterDriver(storageIntervalSec, driver)
}
