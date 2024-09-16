package common

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/internal/internal_storage"
)

func InitStorage(ctx context.Context, storage c_base.IStorage) {
	internal_storage.InitStorage(ctx, storage)
}

func GetStorageInstance() *internal_storage.SStorageInstance {
	return internal_storage.StorageInstance
}
