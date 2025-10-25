package internal

// 系统资源指标分类
const (
	// 进程资源使用指标
	MetricProcessCpuPercent    = "cpu_percent"
	MetricProcessMemoryPercent = "memory_percent"

	// 网络使用量指标
	MetricNetAllSentMB = "net_all_sent_mb"
	MetricNetAllRecvMB = "net_all_recv_mb"

	// 服务情况指标
	MetricGoroutineCount = "goroutine_count"
	MetricHeapAllocMB    = "heap_alloc_mb"
	MetricHeapSysMB      = "heap_sys_mb"
	MetricGCCount        = "gc_count"
)

// 资源指标字段映射
var ResourceMetricsMap = map[string][]string{
	"process": {
		MetricProcessCpuPercent,
		MetricProcessMemoryPercent,
	},
	"network": {
		MetricNetAllSentMB,
		MetricNetAllRecvMB,
	},
	"service": {
		MetricGoroutineCount,
		MetricHeapAllocMB,
		MetricHeapSysMB,
		MetricGCCount,
	},
	"storage": {
		"samples_per_second",
		"total_series",
	},
}
