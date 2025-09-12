package internal

import "time"

// IWindowCounter 滑动窗口计数器接口
// 用于统计指定时间窗口内的请求数量，计算QPS等指标
type IWindowCounter interface {
	// Increment 增加计数
	Increment()

	// IncrementBy 增加指定数量的计数
	IncrementBy(count int64)

	// GetCount 获取当前窗口内的总计数
	GetCount() int64

	// GetQPS 获取当前QPS (每秒请求数)
	GetQPS() float64

	// GetWindowSize 获取窗口大小
	GetWindowSize() time.Duration

	// GetBucketCount 获取桶的数量
	GetBucketCount() int

	// Reset 重置计数器
	Reset()

	// GetStats 获取统计信息
	GetStats() *SWindowCounterStats
}

// SWindowCounterStats 计数器统计信息
type SWindowCounterStats struct {
	TotalCount     int64         `json:"total_count"`      // 总计数
	CurrentQPS     float64       `json:"current_qps"`      // 当前QPS
	WindowSize     time.Duration `json:"window_size"`      // 窗口大小
	BucketCount    int           `json:"bucket_count"`     // 桶数量
	BucketSize     time.Duration `json:"bucket_size"`      // 桶大小
	LastUpdateTime time.Time     `json:"last_update_time"` // 最后更新时间
}
