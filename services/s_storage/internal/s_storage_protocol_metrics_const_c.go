package internal

// 协议性能指标常量
const (
	// 基础指标
	ProtocolMetricTotal       = "total"
	ProtocolMetricSuccess     = "success"
	ProtocolMetricFailed      = "failed"
	ProtocolMetricResultSize  = "result_size"
	ProtocolMetricMaxWaitMs   = "max_wait_ms"
	ProtocolMetricMaxTaskName = "max_task_name"

	// 性能指标
	ProtocolMetricAvgResponseMs = "avg_response_ms"
	ProtocolMetricMinWaitMs     = "min_wait_ms"

	// 可靠性指标
	ProtocolMetricSuccessRate         = "success_rate"
	ProtocolMetricConsecutiveFailures = "consecutive_failures"
	ProtocolMetricTimeoutCount        = "timeout_count"
	ProtocolMetricReconnectCount      = "reconnect_count"

	// 吞吐量指标
	ProtocolMetricAvgBytesPerRequest = "avg_bytes_per_request"
)

// 协议指标字段列表
var ProtocolMetricsFields = []string{
	ProtocolMetricTotal,
	ProtocolMetricSuccess,
	ProtocolMetricFailed,
	ProtocolMetricResultSize,
	ProtocolMetricMaxWaitMs,
	ProtocolMetricMaxTaskName,
	ProtocolMetricAvgResponseMs,
	ProtocolMetricMinWaitMs,
	ProtocolMetricSuccessRate,
	ProtocolMetricConsecutiveFailures,
	ProtocolMetricTimeoutCount,
	ProtocolMetricReconnectCount,
	ProtocolMetricAvgBytesPerRequest,
}

// 协议性能趋势指标字段列表
var ProtocolTrendMetricsFields = []string{
	ProtocolMetricSuccessRate,
	ProtocolMetricAvgResponseMs,
	ProtocolMetricTotal,
	ProtocolMetricFailed,
}
