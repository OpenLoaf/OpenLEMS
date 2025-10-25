package internal

import (
	"common/c_base"
	"context"
	"p_tsdb"
)

// NewTSDBStorageInstance 创建TSDB存储实例
func NewTSDBStorageInstance(ctx context.Context, config *c_base.SStorageConfig) c_base.IStorage {
	return p_tsdb.NewStorageInstance(ctx, config)
}
