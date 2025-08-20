package internal

import (
	"common/c_base"
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

// 使用 Manager 模式管理存储实例，参考 DriverManager 的实现风格

var (
	rwMutex                sync.RWMutex
	StorageManagerInstance *SStorageManager
	storageManagerOnce     sync.Once
)

// SStorageManager 统一的存储管理器，实现 IManager + IStorage
type SStorageManager struct {
	parentCtx  context.Context
	ctx        context.Context
	cancelFunc context.CancelFunc
	state      c_base.EServerState

	c_base.IStorage
}

func NewSingleInstance(parentCtx context.Context, storage c_base.IStorage) *SStorageManager {
	storageManagerOnce.Do(func() {
		StorageManagerInstance = &SStorageManager{
			parentCtx: parentCtx,
			IStorage:  storage,
		}
	})
	return StorageManagerInstance
}

// Start 启动管理器（此处预留扩展，当前主要依赖 RegisterInstance 完成实例注册）
func (s *SStorageManager) Start() {
	s.ctx, s.cancelFunc = context.WithCancel(s.parentCtx)
	s.state = c_base.EStateInit
	if s.IStorage != nil {
		s.state = c_base.EStateRunning
	}

	// 保存当前的系统信息
	StorageManagerInstance.saveSystemMetrics()

	// 启动系统监测数据保存
	gtimer.SetInterval(s.ctx, 1*time.Minute, func(ctx context.Context) {
		// 保存数据
		StorageManagerInstance.saveSystemMetrics()
	})
}

// Shutdown 关闭管理器与存储实例
func (s *SStorageManager) Shutdown() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
	//if s.IStorage != nil {
	//	s.IStorage.Close()
	//	s.IStorage = nil
	//}
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

func Shutdown() {
	StorageManagerInstance.Shutdown()
}

func (s *SStorageManager) saveSystemMetrics() {
	if s.state != c_base.EStateRunning {
		return
	}
	systemInfo := GetSystemInfo()
	_ = StorageManagerInstance.IStorage.SaveSystemMetrics(c_base.ConstSystem, systemInfo, GetSystemMetrics())
	_ = StorageManagerInstance.IStorage.SaveSystemMetrics(c_base.ConstProcess, systemInfo, GetProcessInfo())
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
