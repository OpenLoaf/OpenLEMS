package internal

import (
	"common"
	"common/c_base"
	"common/c_log"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/gtimer"
)

// SStorageManager 存储管理器实现
type SStorageManager struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	storage    c_base.IStorage
	started    bool
	mutex      sync.RWMutex
}

// NewStorageManagerImpl 创建存储管理器实例
func NewStorageManagerImpl(ctx context.Context, storage c_base.IStorage) *SStorageManager {
	return &SStorageManager{
		storage: storage,
	}
}

// Start 启动存储管理器
func (s *SStorageManager) Start(ctx context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.started {
		return nil
	}

	s.ctx, s.cancelFunc = context.WithCancel(ctx)

	// 立即保存一次系统指标
	s.saveSystemMetrics()

	// 启动系统指标保存定时器（每1分钟）
	gtimer.SetInterval(s.ctx, 1*time.Minute, func(ctx context.Context) {
		s.saveSystemMetrics()
	})

	s.started = true
	c_log.BizInfo(ctx, "存储管理器启动成功")

	return nil
}

// Shutdown 关闭存储管理器
func (s *SStorageManager) Shutdown(ctx context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.started {
		return nil
	}

	if s.cancelFunc != nil {
		s.cancelFunc()
	}

	s.started = false
	c_log.BizInfo(ctx, "存储管理器已关闭")

	return nil
}

// RegisterDriver 注册设备数据存储任务
func (s *SStorageManager) RegisterDriver(deviceConfig *c_base.SDeviceConfig) error {
	if s.storage == nil {
		return errors.New("存储实例未初始化")
	}

	if deviceConfig.StorageIntervalSec >= 0 {
		var dur time.Duration
		if deviceConfig.StorageIntervalSec == 0 {
			dur = 1 * time.Minute
		} else {
			dur = time.Duration(deviceConfig.StorageIntervalSec) * time.Second
		}

		deviceId := deviceConfig.Id
		// 启动设备数据存储定时任务
		gtimer.SetInterval(s.ctx, dur, func(ctx context.Context) {
			s.saveDeviceData(deviceConfig.Id)
		})

		c_log.BizInfo(s.ctx, "设备存储任务已注册", "deviceId", deviceId, "interval", dur)
	} else {
		c_log.BizInfo(s.ctx, "设备数据不存储", "deviceId", deviceConfig.Id)
	}

	return nil
}

// GetStorageInstance 获取存储实例
func (s *SStorageManager) GetStorageInstance() c_base.IStorage {
	return s.storage
}

// saveSystemMetrics 保存系统指标
func (s *SStorageManager) saveSystemMetrics() {
	if s.storage == nil {
		return
	}

	systemInfo := GetSystemInfo()

	// 保存系统指标
	if err := s.storage.SaveSystemMetrics(c_base.ConstSystem, systemInfo, GetSystemMetrics()); err != nil {
		c_log.BizError(s.ctx, "保存系统指标失败", "error", err)
	}

	// 保存进程指标
	if err := s.storage.SaveSystemMetrics(c_base.ConstProcess, systemInfo, GetProcessInfo()); err != nil {
		c_log.BizError(s.ctx, "保存进程指标失败", "error", err)
	}

	// 保存TSDB统计信息
	if stats, err := s.storage.GetStorageStats(); err == nil {
		tsdbMetrics := map[string]any{
			"samples_per_second": stats.SamplesPerSecond,
			"total_series":       float64(stats.TotalSeries),
			"total_samples":      float64(stats.TotalSamples),
			"storage_size_mb":    stats.StorageSizeMB,
		}
		if err := s.storage.SaveSystemMetrics("tsdb_stats", systemInfo, tsdbMetrics); err != nil {
			c_log.BizError(s.ctx, "保存TSDB统计信息失败", "error", err)
		}
	}
}

// saveDeviceData 保存设备数据
func (s *SStorageManager) saveDeviceData(deviceId string) {
	if s.storage == nil {
		return
	}

	device := common.GetDeviceManager().GetDeviceById(deviceId)
	if device == nil {
		c_log.BizDebug(s.ctx, "设备实例未找到", "deviceId", deviceId)
		return
	}

	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(deviceId)
	if deviceConfig == nil || deviceConfig.DriverInfo == nil {
		c_log.BizDebug(s.ctx, "设备配置未找到", "deviceId", deviceId)
		return
	}

	// 保存设备遥测数据
	if err := s.storage.SaveDevices(deviceId, c_base.GetAllTelemetryPoint(device)); err != nil {
		c_log.BizError(s.ctx, "保存设备数据失败", "deviceId", deviceId, "error", err)
	}
}
