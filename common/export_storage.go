package common

import (
	"common/c_base"
	"common/internal/internal_storage"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

func RegisterStorageInstance(builder func(ctx context.Context) c_base.IStorage) {
	internal_storage.RegisterInstance(builder)
}

func CloseStorage() {
	internal_storage.Close()
}

func GetStorageInstance() c_base.IStorage {
	return internal_storage.GetInstance()
}

func StorageTimerSaveDeviceMetrics(ctx context.Context, storageIntervalSec int32, driver c_base.IDriver) {
	if storageIntervalSec >= 0 {
		var dur time.Duration
		if storageIntervalSec == 0 {
			dur = 1 * time.Minute
		} else {
			dur = time.Duration(storageIntervalSec) * time.Second
		}
		gtimer.SetInterval(ctx, dur, func(ctx context.Context) {
			// 保存数据
			_ = GetStorageInstance().SaveDevices(driver.GetDeviceConfig().Id, driver.GetDriverType(), driver.GetAllTelemetry(driver))
		})
		g.Log().Infof(ctx, "设备[%s]存储间隔：%v", driver.GetDeviceConfig().Name, dur)
	} else {
		g.Log().Infof(ctx, "设备[%s] 数据不存储！", driver.GetDeviceConfig().Name)
	}

}
