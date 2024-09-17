package influxdb_1

import (
	"common/c_base"
	"context"
	"influxdb_1/internal"
)

func NewStorageInstance(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	return internal.NewInfluxdb1(ctx, storageConfig)
}
