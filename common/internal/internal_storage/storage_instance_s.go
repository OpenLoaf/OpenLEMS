package internal_storage

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

var (
	StorageInstance *SStorageInstance
)

type SStorageInstance struct {
	ctx context.Context
	c_base.IStorage
}

func InitStorage(ctx context.Context, storage c_base.IStorage) {
	if StorageInstance != nil {
		panic(gerror.Newf("StorageInstance已经注册！请勿重复！"))
	}
	StorageInstance = &SStorageInstance{}
	//StorageInstance.ctx, StorageInstance.cancelFunc = context.WithCancel(ctx)
	StorageInstance.IStorage = storage

	// 监听取消信号
	go func() {
		_ = <-ctx.Done()
		if StorageInstance.IStorage != nil {
			StorageInstance.IStorage.Close()
		}
		StorageInstance = nil
		g.Log().Infof(ctx, "存储服务已关闭！")
	}()
}

func (s *SStorageInstance) TimerSaveDeviceMetrics(storageIntervalSec int32, driver c_base.IDriver) {
	if storageIntervalSec >= 0 {
		var dur time.Duration
		if storageIntervalSec == 0 {
			dur = 1 * time.Minute
		} else {
			dur = time.Duration(storageIntervalSec) * time.Second
		}
		gtimer.SetInterval(s.ctx, dur, func(ctx context.Context) {
			// 保存数据
			_ = s.IStorage.Save(driver.GetDeviceConfig().Id, driver.GetDriverType(), driver.GetAllTelemetry(driver))
		})
		g.Log().Infof(s.ctx, "设备[%s]存储间隔：%v", driver.GetDeviceConfig().Name, dur)
	} else {
		g.Log().Infof(s.ctx, "设备[%s] 数据不存储！", driver.GetDeviceConfig().Name)
	}
}
