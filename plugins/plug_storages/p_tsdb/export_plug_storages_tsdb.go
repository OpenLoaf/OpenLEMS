package p_tsdb

import (
	"common/c_base"
	"context"
	"p_tsdb/internal"
)

func NewStorageInstance(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	return internal.NewPromTSDB(ctx, storageConfig)
}
