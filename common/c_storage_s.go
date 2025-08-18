package common

import "common/c_base"

var storageInstance c_base.IStorage

func RegisterStorageInstance(storage c_base.IStorage) {
	storageInstance = storage
}

func GetStorageInstance() c_base.IStorage {
	if storageInstance == nil {
		panic("storage instance is nil !")
	}
	return storageInstance
}
