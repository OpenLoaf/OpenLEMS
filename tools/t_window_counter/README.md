# Core 包 - 功能模块化架构

这个包提供了各种核心功能模块，采用功能归类的方式组织代码结构。

## 功能模块

### 滑动窗口计数器 (c_window_counter)

高效的滑动窗口计数器，用于统计最近的QPS（每秒请求数）和其他时间窗口内的指标。

#### 功能特性

- **滑动窗口计数**: 使用环形缓冲区实现高效的滑动窗口计数
- **QPS计算**: 自动计算每秒请求数
- **并发安全**: 使用读写锁保证并发访问安全
- **多种预设配置**: 提供多种预设的计数器配置
- **灵活配置**: 支持自定义窗口大小和桶数量
- **统计信息**: 提供详细的统计信息

## 项目结构

```
core/
├── c_window_counter/              # 滑动窗口计数器功能模块
│   ├── internal/                  # 内部实现（不对外暴露）
│   │   ├── c_window_counter_i.go  # 内部接口定义
│   │   ├── c_window_counter_s.go  # 实现结构体
│   │   └── c_window_counter_f.go  # 内部工厂函数
│   └── public/                    # 对外暴露的接口
│       ├── c_window_counter_i.go  # 接口定义
│       ├── c_window_counter_f.go  # 工厂函数
│       └── c_window_counter_test.go # 测试文件
├── examples/                      # 使用示例
│   └── window_counter_example.go  # 滑动窗口计数器示例
└── README.md                      # 说明文档
```

## 使用方法

### 滑动窗口计数器

#### 基本使用

```go
package main

import (
    "fmt"
    "time"
    "core/c_window_counter/public"
)

func main() {
    // 创建默认计数器
    counter := public.NewDefaultWindowCounter()
    
    // 增加计数
    counter.Increment()
    counter.IncrementBy(5)
    
    // 获取统计信息
    count := counter.GetCount()
    qps := counter.GetQPS()
    
    fmt.Printf("当前计数: %d, QPS: %.2f\n", count, qps)
}
```

#### 预设计数器类型

```go
// 默认计数器 (1分钟窗口，10个桶)
counter1 := public.NewDefaultWindowCounter()

// QPS专用计数器 (1分钟窗口，60个桶)
counter2 := public.NewQPSWindowCounter()

// 高精度计数器 (10秒窗口，100个桶)
counter3 := public.NewHighPrecisionWindowCounter()

// 长期统计计数器 (1小时窗口，60个桶)
counter4 := public.NewLongTermWindowCounter()
```

#### 自定义配置

```go
// 创建自定义配置的计数器
counter := public.CreateWindowCounterWithConfig(30*time.Second, 30)

// 获取推荐的桶数量
optimalBuckets := public.GetOptimalBucketCount(time.Minute)
```

#### 获取统计信息

```go
stats := counter.GetStats()
fmt.Printf("总计数: %d\n", stats.TotalCount)
fmt.Printf("当前QPS: %.2f\n", stats.CurrentQPS)
fmt.Printf("窗口大小: %v\n", stats.WindowSize)
fmt.Printf("桶数量: %d\n", stats.BucketCount)
fmt.Printf("桶大小: %v\n", stats.BucketSize)
fmt.Printf("最后更新时间: %v\n", stats.LastUpdateTime)
```

#### 并发使用

```go
// 计数器是并发安全的，可以在多个goroutine中使用
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := 0; j < 100; j++ {
            counter.Increment()
        }
    }()
}
wg.Wait()
```

## API 参考

### 滑动窗口计数器接口

```go
type IWindowCounter interface {
    Increment()                    // 增加计数
    IncrementBy(count int64)       // 增加指定数量的计数
    GetCount() int64              // 获取当前窗口内的总计数
    GetQPS() float64              // 获取当前QPS
    GetWindowSize() time.Duration // 获取窗口大小
    GetBucketCount() int          // 获取桶的数量
    Reset()                       // 重置计数器
    GetStats() *SWindowCounterStats // 获取统计信息
}
```

### 工厂函数

```go
// 创建默认计数器
func NewDefaultWindowCounter() IWindowCounter

// 创建QPS专用计数器
func NewQPSWindowCounter() IWindowCounter

// 创建高精度计数器
func NewHighPrecisionWindowCounter() IWindowCounter

// 创建长期统计计数器
func NewLongTermWindowCounter() IWindowCounter

// 根据配置创建计数器
func CreateWindowCounterWithConfig(windowSize time.Duration, bucketCount int) IWindowCounter

// 计算QPS
func CalculateQPS(count int64, windowSize time.Duration) float64

// 获取推荐的桶数量
func GetOptimalBucketCount(windowSize time.Duration) int
```

## 性能特性

- **内存效率**: 使用固定大小的环形缓冲区，内存占用恒定
- **时间复杂度**: 增量和查询操作都是O(1)时间复杂度
- **并发性能**: 使用读写锁，支持高并发读写
- **精度控制**: 通过调整桶数量可以平衡精度和内存使用

## 测试

运行测试：

```bash
go test ./c_window_counter/public -v
```

运行示例：

```bash
go run examples/window_counter_example.go
```

## 设计原理

滑动窗口计数器使用环形缓冲区实现：

1. **时间分片**: 将时间窗口分成多个桶（bucket）
2. **环形缓冲区**: 使用固定大小的数组存储每个桶的计数
3. **滑动机制**: 根据当前时间自动清理过期的桶
4. **统计计算**: 实时计算窗口内所有桶的总计数和QPS

这种设计既保证了高效性，又提供了良好的精度控制。
