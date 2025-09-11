package main

import (
	"core/c_window_counter/public"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 滑动窗口计数器使用示例 ===")

	// 1. 创建默认计数器 (1分钟窗口，10个桶)
	counter1 := public.NewDefaultWindowCounter()
	fmt.Printf("默认计数器: 窗口大小=%v, 桶数量=%d\n",
		counter1.GetWindowSize(), counter1.GetBucketCount())

	// 2. 创建QPS专用计数器 (1分钟窗口，60个桶)
	counter2 := public.NewQPSWindowCounter()
	fmt.Printf("QPS计数器: 窗口大小=%v, 桶数量=%d\n",
		counter2.GetWindowSize(), counter2.GetBucketCount())

	// 3. 创建高精度计数器 (10秒窗口，100个桶)
	counter3 := public.NewHighPrecisionWindowCounter()
	fmt.Printf("高精度计数器: 窗口大小=%v, 桶数量=%d\n",
		counter3.GetWindowSize(), counter3.GetBucketCount())

	// 4. 模拟请求处理
	fmt.Println("\n=== 模拟请求处理 ===")

	// 使用QPS计数器进行演示
	qpsCounter := public.NewQPSWindowCounter()

	// 模拟100个请求
	for i := 0; i < 100; i++ {
		qpsCounter.Increment()
		if i%20 == 0 {
			fmt.Printf("处理了 %d 个请求, 当前计数: %d, QPS: %.2f\n",
				i+1, qpsCounter.GetCount(), qpsCounter.GetQPS())
		}
	}

	// 5. 获取详细统计信息
	fmt.Println("\n=== 统计信息 ===")
	stats := qpsCounter.GetStats()
	fmt.Printf("总计数: %d\n", stats.TotalCount)
	fmt.Printf("当前QPS: %.2f\n", stats.CurrentQPS)
	fmt.Printf("窗口大小: %v\n", stats.WindowSize)
	fmt.Printf("桶数量: %d\n", stats.BucketCount)
	fmt.Printf("桶大小: %v\n", stats.BucketSize)
	fmt.Printf("最后更新时间: %v\n", stats.LastUpdateTime.Format("15:04:05"))

	// 6. 演示时间窗口滑动
	fmt.Println("\n=== 时间窗口滑动演示 ===")
	shortCounter := public.CreateWindowCounterWithConfig(2*time.Second, 4)

	// 增加一些计数
	shortCounter.IncrementBy(10)
	fmt.Printf("增加10个计数后: 计数=%d, QPS=%.2f\n",
		shortCounter.GetCount(), shortCounter.GetQPS())

	// 等待窗口滑动
	fmt.Println("等待3秒让窗口滑动...")
	time.Sleep(3 * time.Second)

	fmt.Printf("窗口滑动后: 计数=%d, QPS=%.2f\n",
		shortCounter.GetCount(), shortCounter.GetQPS())

	// 7. 演示并发安全
	fmt.Println("\n=== 并发安全演示 ===")
	concurrentCounter := public.CreateWindowCounterWithConfig(time.Second, 10)

	// 模拟并发请求
	done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 20; j++ {
				concurrentCounter.Increment()
			}
			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 5; i++ {
		<-done
	}

	fmt.Printf("并发处理完成: 总计数=%d, QPS=%.2f\n",
		concurrentCounter.GetCount(), concurrentCounter.GetQPS())

	// 8. 演示重置功能
	fmt.Println("\n=== 重置功能演示 ===")
	fmt.Printf("重置前: 计数=%d\n", concurrentCounter.GetCount())
	concurrentCounter.Reset()
	fmt.Printf("重置后: 计数=%d, QPS=%.2f\n",
		concurrentCounter.GetCount(), concurrentCounter.GetQPS())

	// 9. 演示QPS监控
	fmt.Println("\n=== QPS监控示例 ===")
	qpsMonitor := public.NewQPSWindowCounter()

	// 模拟API请求处理
	processRequest := func() {
		// 模拟请求处理时间
		time.Sleep(10 * time.Millisecond)
		qpsMonitor.Increment()
	}

	// 模拟1秒内的请求处理
	start := time.Now()
	requestCount := 0

	for time.Since(start) < time.Second {
		processRequest()
		requestCount++

		// 每100个请求打印一次统计
		if requestCount%100 == 0 {
			fmt.Printf("已处理 %d 个请求, 当前QPS: %.2f\n",
				requestCount, qpsMonitor.GetQPS())
		}
	}

	// 最终统计
	finalStats := qpsMonitor.GetStats()
	fmt.Printf("\n最终统计:\n")
	fmt.Printf("总请求数: %d\n", finalStats.TotalCount)
	fmt.Printf("平均QPS: %.2f\n", finalStats.CurrentQPS)
	fmt.Printf("监控窗口: %v\n", finalStats.WindowSize)

	fmt.Println("\n=== 示例完成 ===")
}
