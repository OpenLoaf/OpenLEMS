package s_storage

import (
	"common/c_base"
	"context"
	"s_storage/internal"
)

// 重新导出系统资源指标常量，供外部包使用
const (
	// 系统资源使用指标
	MetricUptimeMinute    = internal.MetricUptimeMinute
	MetricCpu             = internal.MetricCpu
	MetricMemTotalMB      = internal.MetricMemTotalMB
	MetricMemAvailableMB  = internal.MetricMemAvailableMB
	MetricMemUsedMB       = internal.MetricMemUsedMB
	MetricMemUsedPercent  = internal.MetricMemUsedPercent
	MetricLoad1Min        = internal.MetricLoad1Min
	MetricLoad5Min        = internal.MetricLoad5Min
	MetricLoad15Min       = internal.MetricLoad15Min
	MetricDiskTotalMB     = internal.MetricDiskTotalMB
	MetricDiskFreeMB      = internal.MetricDiskFreeMB
	MetricDiskUsedMB      = internal.MetricDiskUsedMB
	MetricDiskUsedPercent = internal.MetricDiskUsedPercent

	// 网络使用量指标
	MetricNetAllSentKB = internal.MetricNetAllSentKB
	MetricNetAllRecvKB = internal.MetricNetAllRecvKB

	// 进程资源使用指标
	MetricProcessCpuPercent    = internal.MetricProcessCpuPercent
	MetricProcessMemoryPercent = internal.MetricProcessMemoryPercent

	// 服务情况指标
	MetricGoroutineCount = internal.MetricGoroutineCount
	MetricHeapAllocMB    = internal.MetricHeapAllocMB
	MetricHeapSysMB      = internal.MetricHeapSysMB
	MetricGCCount        = internal.MetricGCCount

	// 存储统计指标
	MetricSamplesPerSecond = internal.MetricSamplesPerSecond
	MetricTotalSeries      = internal.MetricTotalSeries
	MetricTotalSamples     = internal.MetricTotalSamples
	MetricStorageSizeMB    = internal.MetricStorageSizeMB
)

// NewStorageManager 创建存储管理器
func NewStorageManager(ctx context.Context, storage c_base.IStorage) IStorageManager {
	return internal.NewStorageManagerImpl(ctx, storage)
}

// StartStorageManager 启动存储管理器
func StartStorageManager(ctx context.Context) error {
	manager := GetStorageManager()
	return manager.Start(ctx)
}

// RegisterStorageDriver 注册设备存储任务
func RegisterStorageDriver(deviceConfig *c_base.SDeviceConfig) error {
	manager := GetStorageManager()
	return manager.RegisterDriver(deviceConfig)
}

// GetStorageInstance 获取存储实例
func GetStorageInstance() c_base.IStorage {
	manager := GetStorageManager()
	return manager.GetStorageInstance()
}

// GetSystemMetrics 获取系统指标
func GetSystemMetrics() map[string]any {
	return internal.GetSystemMetrics()
}

// GetProcessInfo 获取进程信息
func GetProcessInfo() map[string]any {
	return internal.GetProcessInfo()
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() map[string]string {
	return internal.GetSystemInfo()
}

// ResourceMetricsMap 资源类别对应的指标字段映射
var ResourceMetricsMap = map[string][]string{
	"process": {
		MetricProcessCpuPercent,
		MetricProcessMemoryPercent,
	},
	"network": {
		MetricNetAllSentKB,
		MetricNetAllRecvKB,
	},
	"service": {
		MetricGoroutineCount,
		MetricMemUsedMB,

		MetricGCCount,
	},
	"storage": {
		MetricSamplesPerSecond,
		MetricTotalSeries,
	},
}
