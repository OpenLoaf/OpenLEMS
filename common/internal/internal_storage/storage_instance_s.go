package internal_storage

import (
	"common/c_base"
	"common/util"
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

var (
	rwMutex          sync.RWMutex
	sStorageInstance *SStorageInstance
)

type SStorageInstance struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	c_base.IStorage
}

func GetInstance() *SStorageInstance {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	if sStorageInstance == nil || sStorageInstance.IStorage == nil {
		return nil
	}
	return sStorageInstance
}

func Close() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if sStorageInstance != nil {
		g.Log().Infof(sStorageInstance.ctx, "StorageInstance已经注册！准备注销并重新注册！")
		sStorageInstance.cancelFunc()
	}

}

func RegisterInstance(builder func(ctx context.Context) c_base.IStorage) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if builder == nil {
		sStorageInstance = nil
		return
	}
	if sStorageInstance != nil {
		g.Log().Infof(sStorageInstance.ctx, "StorageInstance已经注册！准备注销并重新注册！")
		sStorageInstance.cancelFunc()
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Storage")
	ctx, cancelFunc := context.WithCancel(ctx)
	sStorageInstance = &SStorageInstance{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		IStorage:   builder(ctx),
	}
	if sStorageInstance.IStorage == nil {
		g.Log().Infof(sStorageInstance.ctx, "未启动长时存储！")
		return
	}

	// 保存当前的系统信息
	saveSystemMetrics()

	// 启动系统监测数据保存
	gtimer.SetInterval(ctx, 1*time.Minute, func(ctx context.Context) {
		// 保存数据
		saveSystemMetrics()
	})

	go func() {
		_ = <-ctx.Done()
		if sStorageInstance.IStorage != nil {
			sStorageInstance.IStorage.Close()
		}
		sStorageInstance = nil
		g.Log().Infof(ctx, "存储服务已关闭！")
	}()
}

func saveSystemMetrics() {
	systemInfo := util.GetSystemInfo()
	_ = sStorageInstance.IStorage.SaveSystemMetrics(c_base.ConstSystem, systemInfo, util.GetSystemMetrics())
	_ = sStorageInstance.IStorage.SaveSystemMetrics(c_base.ConstProcess, systemInfo, util.GetProcessInfo())
}

func (s *SStorageInstance) RegisterDriver(storageIntervalSec int32, driver c_base.IDriver) {
	if storageIntervalSec >= 0 {
		var dur time.Duration
		if storageIntervalSec == 0 {
			dur = 1 * time.Minute
		} else {
			dur = time.Duration(storageIntervalSec) * time.Second
		}
		//TODO 此处需要能同时监测设备关闭或者存储关闭的情况，需要能自动销毁定时任务。现在只有storage关闭时会销毁定时任务，device如果关闭了，定时任务不会销毁
		gtimer.SetInterval(s.ctx, dur, func(ctx context.Context) {
			// 保存数据
			_ = s.IStorage.SaveDevices(driver.GetDeviceConfig().Id, driver.GetDriverType(), driver.GetAllTelemetry(driver))
		})
		g.Log().Infof(s.ctx, "设备[%s]存储间隔：%v", driver.GetDeviceConfig().Name, dur)
	} else {
		g.Log().Infof(s.ctx, "设备[%s] 数据不存储！", driver.GetDeviceConfig().Name)
	}
}
