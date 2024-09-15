package driver

import (
	"context"
	"ems-plan/c_base"
	"influxdb_1"
	"sync"
)

type SStorageCmd struct {
	once        sync.Once
	ctx         context.Context
	cancelFunc  context.CancelFunc
	intervalSec int32
	c_base.IStorage
}

var (
	storageOnce sync.Once
	storageCmd  *SStorageCmd
)

func GetStorageCmd() *SStorageCmd {
	storageOnce.Do(func() {

		storageConfig := &c_base.SStorageConfig{
			Enable:      true,
			Type:        c_base.EStorageTypeInfluxDB1,
			Url:         "http://192.168.0.248:8086",
			Database:    "mydb",
			Username:    "test",
			Password:    "123456",
			IntervalSec: 1,
			KeepDays:    1,
			Params:      nil,
		}
		ctx := context.Background()
		storageCmd = &SStorageCmd{
			ctx: ctx,
			//cancelFunc:  cancelFunc,
			intervalSec: storageConfig.IntervalSec,
			IStorage:    influxdb_1.NewStorageInstance(ctx, storageConfig),
		}
	})
	return storageCmd
}

func (s *SStorageCmd) Stop() {
	s.Close()
	//s.cancelFunc()
}
