package internal

import (
	"time"
)

// NewDefaultWindowCounter 创建默认配置的滑动窗口计数器
// 窗口大小: 1分钟，桶数量: 10个
func NewDefaultWindowCounter() *SWindowCounter {
	return NewWindowCounter(time.Minute, 10)
}

// NewQPSWindowCounter 创建专门用于QPS统计的计数器
// 窗口大小: 1分钟，桶数量: 60个 (每秒一个桶)
func NewQPSWindowCounter() *SWindowCounter {
	return NewWindowCounter(time.Minute, 60)
}

// NewHighPrecisionWindowCounter 创建高精度QPS计数器
// 窗口大小: 10秒，桶数量: 100个 (每100毫秒一个桶)
func NewHighPrecisionWindowCounter() *SWindowCounter {
	return NewWindowCounter(10*time.Second, 100)
}

// NewLongTermWindowCounter 创建长期统计计数器
// 窗口大小: 1小时，桶数量: 60个 (每分钟一个桶)
func NewLongTermWindowCounter() *SWindowCounter {
	return NewWindowCounter(time.Hour, 60)
}

// CreateWindowCounterWithConfig 根据配置创建计数器
func CreateWindowCounterWithConfig(windowSize time.Duration, bucketCount int) IWindowCounter {
	return NewWindowCounter(windowSize, bucketCount)
}

// CalculateQPS 计算指定时间窗口内的QPS
// 这是一个静态函数，用于计算任意时间窗口的QPS
func CalculateQPS(count int64, windowSize time.Duration) float64 {
	if windowSize <= 0 {
		return 0
	}
	windowSeconds := windowSize.Seconds()
	if windowSeconds <= 0 {
		return 0
	}
	return float64(count) / windowSeconds
}

// GetOptimalBucketCount 根据窗口大小获取推荐的桶数量
// 返回推荐的桶数量，确保每个桶的时间大小在合理范围内
func GetOptimalBucketCount(windowSize time.Duration) int {
	// 推荐每个桶的时间大小为1s
	maxBucketSize := time.Second

	optimalBucketCount := int(windowSize / maxBucketSize)
	if optimalBucketCount < 1 {
		optimalBucketCount = 1
	}

	// 确保桶数量不会太多，避免内存浪费
	maxBuckets := 1000
	if optimalBucketCount > maxBuckets {
		optimalBucketCount = maxBuckets
	}

	// 确保桶数量不会太少，保证精度
	minBuckets := 10
	if optimalBucketCount < minBuckets {
		optimalBucketCount = minBuckets
	}

	return optimalBucketCount
}
