package public

import (
	"c_window_counter/internal"
	"time"
)

// NewDefaultWindowCounter 创建默认配置的滑动窗口计数器
// 窗口大小: 1分钟，桶数量: 10个
func NewDefaultWindowCounter() IWindowCounter {
	return internal.NewDefaultWindowCounter()
}

// NewQPSWindowCounter 创建专门用于QPS统计的计数器
// 窗口大小: 1分钟，桶数量: 60个 (每秒一个桶)
func NewQPSWindowCounter() IWindowCounter {
	return internal.NewQPSWindowCounter()
}

// NewHighPrecisionWindowCounter 创建高精度QPS计数器
// 窗口大小: 10秒，桶数量: 100个 (每100毫秒一个桶)
func NewHighPrecisionWindowCounter() IWindowCounter {
	return internal.NewHighPrecisionWindowCounter()
}

// NewLongTermWindowCounter 创建长期统计计数器
// 窗口大小: 1小时，桶数量: 60个 (每分钟一个桶)
func NewLongTermWindowCounter() IWindowCounter {
	return internal.NewLongTermWindowCounter()
}

// CreateWindowCounterWithConfig 根据配置创建计数器
func CreateWindowCounterWithConfig(windowSize time.Duration, bucketCount int) IWindowCounter {
	return internal.CreateWindowCounterWithConfig(windowSize, bucketCount)
}

// CalculateQPS 计算指定时间窗口内的QPS
// 这是一个静态函数，用于计算任意时间窗口的QPS
func CalculateQPS(count int64, windowSize time.Duration) float64 {
	return internal.CalculateQPS(count, windowSize)
}

// GetOptimalBucketCount 根据窗口大小获取推荐的桶数量
// 返回推荐的桶数量，确保每个桶的时间大小在合理范围内
func GetOptimalBucketCount(windowSize time.Duration) int {
	return internal.GetOptimalBucketCount(windowSize)
}
