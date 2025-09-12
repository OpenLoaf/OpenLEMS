package public

import (
	"sync"
	"testing"
	"time"
)

func TestNewWindowCounter(t *testing.T) {
	// 测试默认参数
	counter := CreateWindowCounterWithConfig(time.Minute, 10)
	if counter == nil {
		t.Fatal("NewWindowCounter returned nil")
	}

	if counter.GetWindowSize() != time.Minute {
		t.Errorf("Expected window size %v, got %v", time.Minute, counter.GetWindowSize())
	}

	if counter.GetBucketCount() != 10 {
		t.Errorf("Expected bucket count 10, got %d", counter.GetBucketCount())
	}
}

func TestIncrementAndGetCount(t *testing.T) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 初始计数应该为0
	if count := counter.GetCount(); count != 0 {
		t.Errorf("Expected initial count 0, got %d", count)
	}

	// 增加计数
	counter.Increment()
	if count := counter.GetCount(); count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}

	// 增加指定数量
	counter.IncrementBy(5)
	if count := counter.GetCount(); count != 6 {
		t.Errorf("Expected count 6, got %d", count)
	}
}

func TestQPS(t *testing.T) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 初始QPS应该为0
	if qps := counter.GetQPS(); qps != 0 {
		t.Errorf("Expected initial QPS 0, got %f", qps)
	}

	// 增加10个请求
	for i := 0; i < 10; i++ {
		counter.Increment()
	}

	// QPS应该接近10
	qps := counter.GetQPS()
	if qps < 9 || qps > 11 {
		t.Errorf("Expected QPS around 10, got %f", qps)
	}
}

func TestReset(t *testing.T) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 增加一些计数
	counter.IncrementBy(5)
	if count := counter.GetCount(); count != 5 {
		t.Errorf("Expected count 5, got %d", count)
	}

	// 重置
	counter.Reset()
	if count := counter.GetCount(); count != 0 {
		t.Errorf("Expected count 0 after reset, got %d", count)
	}

	if qps := counter.GetQPS(); qps != 0 {
		t.Errorf("Expected QPS 0 after reset, got %f", qps)
	}
}

func TestGetStats(t *testing.T) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 增加一些计数
	counter.IncrementBy(3)

	stats := counter.GetStats()
	if stats == nil {
		t.Fatal("GetStats returned nil")
	}

	if stats.TotalCount != 3 {
		t.Errorf("Expected total count 3, got %d", stats.TotalCount)
	}

	if stats.WindowSize != time.Second {
		t.Errorf("Expected window size %v, got %v", time.Second, stats.WindowSize)
	}

	if stats.BucketCount != 10 {
		t.Errorf("Expected bucket count 10, got %d", stats.BucketCount)
	}
}

func TestConcurrentAccess(t *testing.T) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 并发测试
	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expectedCount := int64(numGoroutines * incrementsPerGoroutine)
	if count := counter.GetCount(); count != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, count)
	}
}

func TestTimeWindowSliding(t *testing.T) {
	// 使用较短的窗口进行测试
	counter := CreateWindowCounterWithConfig(100*time.Millisecond, 10)

	// 增加一些计数
	counter.IncrementBy(5)
	if count := counter.GetCount(); count != 5 {
		t.Errorf("Expected count 5, got %d", count)
	}

	// 等待窗口滑动
	time.Sleep(150 * time.Millisecond)

	// 计数应该被清空
	if count := counter.GetCount(); count != 0 {
		t.Errorf("Expected count 0 after window slide, got %d", count)
	}
}

func TestNewDefaultWindowCounter(t *testing.T) {
	counter := NewDefaultWindowCounter()
	if counter == nil {
		t.Fatal("NewDefaultWindowCounter returned nil")
	}

	if counter.GetWindowSize() != time.Minute {
		t.Errorf("Expected window size %v, got %v", time.Minute, counter.GetWindowSize())
	}

	if counter.GetBucketCount() != 10 {
		t.Errorf("Expected bucket count 10, got %d", counter.GetBucketCount())
	}
}

func TestNewQPSWindowCounter(t *testing.T) {
	counter := NewQPSWindowCounter()
	if counter == nil {
		t.Fatal("NewQPSWindowCounter returned nil")
	}

	if counter.GetWindowSize() != time.Minute {
		t.Errorf("Expected window size %v, got %v", time.Minute, counter.GetWindowSize())
	}

	if counter.GetBucketCount() != 60 {
		t.Errorf("Expected bucket count 60, got %d", counter.GetBucketCount())
	}
}

func TestCalculateQPS(t *testing.T) {
	// 测试静态QPS计算函数
	qps := CalculateQPS(100, time.Second)
	if qps != 100 {
		t.Errorf("Expected QPS 100, got %f", qps)
	}

	qps = CalculateQPS(60, time.Minute)
	if qps != 1 {
		t.Errorf("Expected QPS 1, got %f", qps)
	}

	// 测试边界情况
	qps = CalculateQPS(0, time.Second)
	if qps != 0 {
		t.Errorf("Expected QPS 0, got %f", qps)
	}

	qps = CalculateQPS(100, 0)
	if qps != 0 {
		t.Errorf("Expected QPS 0 for zero duration, got %f", qps)
	}
}

func TestGetOptimalBucketCount(t *testing.T) {
	// 测试不同窗口大小的推荐桶数量
	bucketCount := GetOptimalBucketCount(time.Second)
	if bucketCount < 1 || bucketCount > 1000 {
		t.Errorf("Bucket count %d is out of reasonable range", bucketCount)
	}

	bucketCount = GetOptimalBucketCount(time.Minute)
	if bucketCount < 10 || bucketCount > 1000 {
		t.Errorf("Bucket count %d is out of reasonable range", bucketCount)
	}

	bucketCount = GetOptimalBucketCount(time.Hour)
	if bucketCount < 10 || bucketCount > 1000 {
		t.Errorf("Bucket count %d is out of reasonable range", bucketCount)
	}
}

// 基准测试
func BenchmarkIncrement(b *testing.B) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		counter.Increment()
	}
}

func BenchmarkGetCount(b *testing.B) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 预先增加一些计数
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.GetCount()
	}
}

func BenchmarkGetQPS(b *testing.B) {
	counter := CreateWindowCounterWithConfig(time.Second, 10)

	// 预先增加一些计数
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.GetQPS()
	}
}
