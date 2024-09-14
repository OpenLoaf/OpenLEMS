package influxdb_1

import (
	"context"
	"ems-plan/c_base"
	"influxdb_1/internal"
)

func NewStorageInstance(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	return internal.NewInfluxdb1(ctx, storageConfig)
}
