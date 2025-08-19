package internal

import (
	"common/c_base"
	"context"
	"s_db"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

// 使用 Manager 模式管理存储实例，参考 DriverManager 的实现风格

var (
	rwMutex                sync.RWMutex
	storageManagerInstance *SStorageManager
	storageManagerOnce     sync.Once
)

// SStorageManager 统一的存储管理器，实现 IManager + IStorage
type SStorageManager struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	state      c_base.EServerState

	c_base.IStorage
}

// GetInstance 保持对外函数名不变，返回 StorageManager 单例
func GetInstance() *SStorageManager {
	storageManagerOnce.Do(func() {
		storageManagerInstance = &SStorageManager{}
	})
	return storageManagerInstance
}

// Start 启动管理器（此处预留扩展，当前主要依赖 RegisterInstance 完成实例注册）
func (s *SStorageManager) Start(parentCtx context.Context) {
	s.ctx, s.cancelFunc = context.WithCancel(parentCtx)
	s.state = c_base.EStateRunning
	// 预热配置
	_ = s_db.GetSettingService().GetSettingValueByKey(s.ctx, "")
}

// Shutdown 关闭管理器与存储实例
func (s *SStorageManager) Shutdown() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
	if s.IStorage != nil {
		s.IStorage.Close()
		s.IStorage = nil
	}
	s.state = c_base.EStateStopped
}

// Cleanup 清理资源（占位）
func (s *SStorageManager) Cleanup() error {
	return nil
}

// Status 返回当前状态
func (s *SStorageManager) Status() c_base.EServerState {
	return s.state
}

// RegisterInstance 注册/重建存储实例
func RegisterInstance(storage c_base.IStorage) {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	manager := GetInstance()

	// 若已存在，则先注销
	if manager.cancelFunc != nil {
		g.Log().Infof(manager.ctx, "StorageInstance已经注册！准备注销并重新注册！")
		manager.cancelFunc()
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Storage")
	ctx, cancelFunc := context.WithCancel(ctx)
	manager.ctx = ctx
	manager.cancelFunc = cancelFunc
	manager.IStorage = storage

	if manager.IStorage == nil {
		g.Log().Infof(manager.ctx, "未启动长时存储！")
		manager.state = c_base.EStateStopped
		return
	}

	// 保存当前的系统信息
	saveSystemMetrics()

	// 启动系统监测数据保存
	gtimer.SetInterval(ctx, 1*time.Minute, func(ctx context.Context) {
		// 保存数据
		saveSystemMetrics()
	})

	manager.state = c_base.EStateRunning

	go func(localCtx context.Context) {
		_ = <-localCtx.Done()
		if manager.IStorage != nil {
			manager.IStorage.Close()
		}
		manager.IStorage = nil
		g.Log().Infof(localCtx, "存储服务已关闭！")
	}(ctx)
}

func Shutdown() {
	GetInstance().Shutdown()
}

func saveSystemMetrics() {
	manager := GetInstance()
	if manager == nil || manager.IStorage == nil {
		return
	}
	systemInfo := GetSystemInfo()
	_ = manager.IStorage.SaveSystemMetrics(c_base.ConstSystem, systemInfo, GetSystemMetrics())
	_ = manager.IStorage.SaveSystemMetrics(c_base.ConstProcess, systemInfo, GetProcessInfo())
}

// RegisterDriver 注册设备数据的周期存储任务
func (s *SStorageManager) RegisterDriver(storageIntervalSec int32, driver c_base.IDevice) {
	if s == nil || s.IStorage == nil {
		return
	}
	if storageIntervalSec >= 0 {
		var dur time.Duration
		if storageIntervalSec == 0 {
			dur = 1 * time.Minute
		} else {
			dur = time.Duration(storageIntervalSec) * time.Second
		}
		// TODO: 同时监测设备关闭或者存储关闭的情况，自动销毁定时任务
		gtimer.SetInterval(s.ctx, dur, func(ctx context.Context) {
			// 保存数据
			des := driver.GetDriverDescription()
			_ = s.IStorage.SaveDevices(driver.GetDeviceConfig().Id, driver.GetDriverType(), des.GetAllTelemetry(driver))
		})
		g.Log().Infof(s.ctx, "设备[%s]存储间隔：%v", driver.GetDeviceConfig().Name, dur)
	} else {
		g.Log().Infof(s.ctx, "设备[%s] 数据不存储！", driver.GetDeviceConfig().Name)
	}
}
