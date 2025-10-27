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
	"github.com/shirou/gopsutil/v4/net"
)

// SStorageManager 存储管理器实现
type SStorageManager struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	storage    c_base.IStorage
	started    bool
	mutex      sync.RWMutex

	// 增量计算缓存字段
	lastNetSentKB    *float64 // 上次网络发送量（KB）
	lastNetRecvKB    *float64 // 上次网络接收量（KB）
	lastTotalSamples *float64 // 上次总样本数
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

	// 保存系统指标（包含增量网络指标）
	if err := s.storage.SaveSystemMetrics(c_base.ConstSystem, systemInfo, s.getSystemMetricsWithIncrement()); err != nil {
		c_log.BizError(s.ctx, "保存系统指标失败", "error", err)
	}

	// 保存进程指标
	if err := s.storage.SaveSystemMetrics(c_base.ConstProcess, systemInfo, GetProcessInfo()); err != nil {
		c_log.BizError(s.ctx, "保存进程指标失败", "error", err)
	}

	// 保存TSDB统计信息（包含增量总样本数）
	if stats, err := s.storage.GetStorageStats(); err == nil {
		tsdbMetrics := s.getTSDBMetricsWithIncrement(stats)
		if err := s.storage.SaveSystemMetrics("tsdb_stats", systemInfo, tsdbMetrics); err != nil {
			c_log.BizError(s.ctx, "保存TSDB统计信息失败", "error", err)
		}
	}
}

// getSystemMetricsWithIncrement 获取包含增量网络指标的系统指标
func (s *SStorageManager) getSystemMetricsWithIncrement() map[string]any {
	result := GetSystemMetrics()

	// 计算网络指标增量
	if counters, err := net.IOCounters(false); err == nil {
		var totalSentKB, totalRecvKB float64
		for _, counter := range counters {
			totalSentKB += float64(counter.BytesSent / 1024)
			totalRecvKB += float64(counter.BytesRecv / 1024)
		}

		// 检查是否为首次运行
		if s.lastNetSentKB == nil || s.lastNetRecvKB == nil {
			// 首次运行，不计算增量，直接保存当前值
			s.lastNetSentKB = &totalSentKB
			s.lastNetRecvKB = &totalRecvKB
			// 首次运行不设置增量指标，避免错误的大数值
		} else {
			// 计算增量
			sentIncrement := totalSentKB - *s.lastNetSentKB
			recvIncrement := totalRecvKB - *s.lastNetRecvKB

			// 如果为负数，说明重置了，设为0
			if sentIncrement < 0 {
				sentIncrement = 0
			}
			if recvIncrement < 0 {
				recvIncrement = 0
			}

			// 更新缓存值
			*s.lastNetSentKB = totalSentKB
			*s.lastNetRecvKB = totalRecvKB

			// 设置增量指标
			result[MetricNetAllSentKB] = sentIncrement
			result[MetricNetAllRecvKB] = recvIncrement
		}
	}

	return result
}

// getTSDBMetricsWithIncrement 获取包含增量总样本数的TSDB指标
func (s *SStorageManager) getTSDBMetricsWithIncrement(stats *c_base.StorageStats) map[string]any {
	currentSamples := float64(stats.TotalSamples)

	result := map[string]any{
		MetricSamplesPerSecond: stats.SamplesPerSecond,
		MetricTotalSeries:      float64(stats.TotalSeries),
		MetricStorageSizeMB:    stats.StorageSizeMB,
	}

	// 检查是否为首次运行
	if s.lastTotalSamples == nil {
		// 首次运行，不计算增量，直接保存当前值
		s.lastTotalSamples = &currentSamples
		// 首次运行不设置增量指标，避免错误的大数值
	} else {
		// 计算总样本数增量
		samplesIncrement := currentSamples - *s.lastTotalSamples

		// 如果为负数，说明重置了，设为0
		if samplesIncrement < 0 {
			samplesIncrement = 0
		}

		// 更新缓存值
		*s.lastTotalSamples = currentSamples

		// 设置增量指标
		result[MetricTotalSamples] = samplesIncrement
	}

	return result
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
