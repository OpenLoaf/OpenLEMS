package internal

// 系统资源指标常量定义
const (
	// 系统资源使用指标
	MetricUptimeMinute    = "uptime_minute"     // 系统在线时长（分钟）
	MetricCpu             = "cpu"               // CPU 总使用率
	MetricMemTotalMB      = "mem_total_mb"      // 内存总量（MB）
	MetricMemAvailableMB  = "mem_available_mb"  // 可用内存（MB）
	MetricMemUsedMB       = "mem_used_mb"       // 已用内存（MB）
	MetricMemUsedPercent  = "mem_used_percent"  // 内存使用百分比
	MetricLoad1Min        = "load_1min"         // 1分钟负载
	MetricLoad5Min        = "load_5min"         // 5分钟负载
	MetricLoad15Min       = "load_15min"        // 15分钟负载
	MetricDiskTotalMB     = "disk_total_mb"     // 磁盘总量（MB）
	MetricDiskFreeMB      = "disk_free_mb"      // 磁盘可用空间（MB）
	MetricDiskUsedMB      = "disk_used_mb"      // 磁盘已用空间（MB）
	MetricDiskUsedPercent = "disk_used_percent" // 磁盘使用百分比

	// 网络使用量指标（增量）
	MetricNetAllSentMB = "net_all_sent_mb" // 网络发送量增量（MB）
	MetricNetAllRecvMB = "net_all_recv_mb" // 网络接收量增量（MB）

	// 进程资源使用指标
	MetricProcessCpuPercent    = "cpu_percent"    // 进程 CPU 使用率
	MetricProcessMemoryPercent = "memory_percent" // 进程内存使用率

	// 服务情况指标
	MetricGoroutineCount = "goroutine_count" // Goroutine 数量
	MetricHeapAllocMB    = "heap_alloc_mb"   // 堆分配内存（MB）
	MetricHeapSysMB      = "heap_sys_mb"     // 堆系统内存（MB）
	MetricGCCount        = "gc_count"        // GC 次数

	// 存储统计指标
	MetricSamplesPerSecond = "samples_per_second" // 每秒样本数
	MetricTotalSeries      = "total_series"       // 总时间序列数
	MetricTotalSamples     = "total_samples"      // 总样本数增量
	MetricStorageSizeMB    = "storage_size_mb"    // 存储大小（MB）
)
