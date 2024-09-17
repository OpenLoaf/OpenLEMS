package common

import (
	"common/c_base"
	"common/internal/internal_storage"
	"context"
)

func InitStorage(ctx context.Context, storage c_base.IStorage) {
	internal_storage.InitStorage(ctx, storage)
}

func GetStorageInstance() *internal_storage.SStorageInstance {
	return internal_storage.StorageInstance
}
