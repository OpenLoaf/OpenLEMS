package influxdb_2

import (
	"common/c_base"
	"context"
	"influxdb_2/internal"
)

func NewStorageInstance(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	return internal.NewInfluxdb2(ctx, storageConfig)
}
