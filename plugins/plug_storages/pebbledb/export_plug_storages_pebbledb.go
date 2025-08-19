package pebbledb

import (
	"common/c_base"
	"context"
	"pebbledb/internal"
)

func NewStorageInstance(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	return internal.NewPebbledb(ctx, storageConfig)
}
