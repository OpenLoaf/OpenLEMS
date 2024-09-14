package driver

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/os/gtimer"
	"influxdb_1"
	"sync"
	"time"
)

type SStorageCmd struct {
	once        sync.Once
	ctx         context.Context
	cancelFunc  context.CancelFunc
	intervalSec int32
	c_base.IStorage
}

func NewStorageCmd(ctx context.Context) *SStorageCmd {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Driver")
	ctx, cancelFunc := context.WithCancel(ctx)

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

	return &SStorageCmd{
		ctx:         ctx,
		cancelFunc:  cancelFunc,
		intervalSec: storageConfig.IntervalSec,
		IStorage:    influxdb_1.NewStorageInstance(ctx, storageConfig),
	}
}

func (s *SStorageCmd) Start() {
	s.once.Do(func() {
		gtimer.SetInterval(s.ctx, time.Duration(s.intervalSec)*time.Second, func(ctx context.Context) {
			// 保存数据
			//s.Save()
			//common.

		})
	})
}

func (s *SStorageCmd) Stop() {
	s.Close()
	s.cancelFunc()
}
