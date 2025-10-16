// //go:generate mockgen -source=storage_i.go -package=mock_c_base -destination=mock_c_base/storage_i.mock.go
package c_base

import (
	"common/c_chart"
	"time"
)

// StorageStats 存储统计信息
type StorageStats struct {
	TotalSeries      int64      `json:"total_series"`       // 总时间序列数量
	TotalSamples     int64      `json:"total_samples"`      // 总样本数量
	StorageSize      int64      `json:"storage_size"`       // 存储大小（字节）
	OldestTimestamp  *time.Time `json:"oldest_timestamp"`   // 最老数据时间戳
	NewestTimestamp  *time.Time `json:"newest_timestamp"`   // 最新数据时间戳
	RetentionTime    int64      `json:"retention_time"`     // 数据保留时间（秒）
	AvgSeriesSize    float64    `json:"avg_series_size"`    // 平均每个序列占用数据大小（字节）
	SamplesPerSecond float64    `json:"samples_per_second"` // 每秒存储样本数
	StorageSizeMB    float64    `json:"storage_size_mb"`    // 存储大小（MB）
	RetentionHours   float64    `json:"retention_hours"`    // 数据保留时间（小时）
}

type StorageType string

const (
	StorageTypeDevice   StorageType = "device"
	StorageTypeProtocol StorageType = "protocol"
	StorageTypeSystem   StorageType = "system"
)

type IStorage interface {

	// SaveDevices 保存设备数据
	SaveDevices(deviceId string, fields map[string]any) error

	// SaveProtocolMetrics 保存协议指标数据
	SaveProtocolMetrics(protocolConfig *SProtocolConfig, deviceConfig *SDeviceConfig, metrics map[string]any) error

	// SaveSystemMetrics 保存系统指标数据
	SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error

	// GetStorageData 获取存储数据
	GetStorageData(storageType StorageType, id string, pointKey []string, startTime, endTime *int64, step int) (*c_chart.ChartData, error)

	// GetStorageStats 获取存储统计信息
	GetStorageStats() (*StorageStats, error)

	// Close 关闭数据库
	Close()
}
