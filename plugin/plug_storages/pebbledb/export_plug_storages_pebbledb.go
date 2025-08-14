package pebbledb

import (
	"common/c_base"
	"context"
	"pebbledb/internal"
)

func NewStorageInstance(ctx context.Context) c_base.IStorage {
	return internal.NewPebbledb(ctx)
}
