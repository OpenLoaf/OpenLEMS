package internal

import (
	"sync"
	"time"
)

// SWindowCounter 滑动窗口计数器实现
// 使用环形缓冲区实现高效的滑动窗口计数
type SWindowCounter struct {
	windowSize     time.Duration // 窗口大小
	bucketCount    int           // 桶数量
	bucketSize     time.Duration // 每个桶的时间大小
	buckets        []int64       // 环形缓冲区，存储每个桶的计数
	currentIndex   int           // 当前桶的索引
	lastUpdateTime time.Time     // 最后更新时间
	mutex          sync.RWMutex  // 读写锁，保证并发安全
}

// NewWindowCounter 创建新的滑动窗口计数器
// windowSize: 窗口大小，如 1分钟、5分钟等
// bucketCount: 桶的数量，建议为10-60个，桶越多精度越高但内存占用越大
func NewWindowCounter(windowSize time.Duration, bucketCount int) *SWindowCounter {
	if bucketCount <= 0 {
		bucketCount = 10 // 默认10个桶
	}
	if windowSize <= 0 {
		windowSize = time.Minute // 默认1分钟窗口
	}

	bucketSize := windowSize / time.Duration(bucketCount)

	return &SWindowCounter{
		windowSize:     windowSize,
		bucketCount:    bucketCount,
		bucketSize:     bucketSize,
		buckets:        make([]int64, bucketCount),
		currentIndex:   0,
		lastUpdateTime: time.Now(),
	}
}

// Increment 增加计数
func (c *SWindowCounter) Increment() {
	c.IncrementBy(1)
}

// IncrementBy 增加指定数量的计数
func (c *SWindowCounter) IncrementBy(count int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	c.updateBuckets(now)
	c.buckets[c.currentIndex] += count
}

// GetCount 获取当前窗口内的总计数
func (c *SWindowCounter) GetCount() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	now := time.Now()
	c.updateBuckets(now)

	var total int64
	for _, count := range c.buckets {
		total += count
	}
	return total
}

// GetQPS 获取当前QPS (每秒请求数)
func (c *SWindowCounter) GetQPS() float64 {
	count := c.GetCount()
	windowSeconds := c.windowSize.Seconds()
	if windowSeconds <= 0 {
		return 0
	}
	return float64(count) / windowSeconds
}

// GetWindowSize 获取窗口大小
func (c *SWindowCounter) GetWindowSize() time.Duration {
	return c.windowSize
}

// GetBucketCount 获取桶的数量
func (c *SWindowCounter) GetBucketCount() int {
	return c.bucketCount
}

// Reset 重置计数器
func (c *SWindowCounter) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := range c.buckets {
		c.buckets[i] = 0
	}
	c.currentIndex = 0
	c.lastUpdateTime = time.Now()
}

// GetStats 获取统计信息
func (c *SWindowCounter) GetStats() *SWindowCounterStats {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	now := time.Now()
	c.updateBuckets(now)

	var totalCount int64
	for _, count := range c.buckets {
		totalCount += count
	}

	windowSeconds := c.windowSize.Seconds()
	var currentQPS float64
	if windowSeconds > 0 {
		currentQPS = float64(totalCount) / windowSeconds
	}

	return &SWindowCounterStats{
		TotalCount:     totalCount,
		CurrentQPS:     currentQPS,
		WindowSize:     c.windowSize,
		BucketCount:    c.bucketCount,
		BucketSize:     c.bucketSize,
		LastUpdateTime: now,
	}
}

// updateBuckets 更新桶状态，清理过期数据
func (c *SWindowCounter) updateBuckets(now time.Time) {
	// 计算需要清理的桶数量
	elapsed := now.Sub(c.lastUpdateTime)
	bucketsToClear := int(elapsed / c.bucketSize)

	if bucketsToClear > 0 {
		// 清理过期的桶
		for i := 0; i < bucketsToClear && i < c.bucketCount; i++ {
			c.currentIndex = (c.currentIndex + 1) % c.bucketCount
			c.buckets[c.currentIndex] = 0
		}

		// 如果时间跨度超过整个窗口，清空所有桶
		if bucketsToClear >= c.bucketCount {
			for i := range c.buckets {
				c.buckets[i] = 0
			}
		}

		c.lastUpdateTime = now
	}
}
