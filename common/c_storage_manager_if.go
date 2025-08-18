package common

import "common/c_base"

type IStorageManager interface {
	IManager
	c_base.IStorage
}

var storageManager IStorageManager

func RegisterStorageManager(cmd IStorageManager) {
	storageManager = cmd
}

func GetStorageManager() IStorageManager {
	if storageManager == nil {
		panic("storage manager is not registered")
	}
	return storageManager
}
