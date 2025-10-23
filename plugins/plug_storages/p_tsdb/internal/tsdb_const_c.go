package internal

// TSDB 标签名称常量定义
const (
	// LabelNameMetric 指标名称标签
	LabelNameMetric = "__name__"

	// LabelNameType 类型标签 (device/protocol/system)
	LabelNameType = "type"

	// LabelNameID ID标签 (deviceId/protocolId/measurement)
	LabelNameID = "id"

	// LabelNameField 字段标签
	LabelNameField = "field"

	// LabelNameDeviceID 设备ID标签
	LabelNameDeviceID = "device_id"
)

// TSDB 指标名称常量定义
const (
	// MetricNameEmsMetric 数值指标名称
	MetricNameEmsMetric = "ems_metric"
)

// 时间转换常量
const (
	// MillisecondsPerSecond 毫秒到秒的转换系数
	MillisecondsPerSecond = 1000
)

// 状态缓存常量
const (
	// StatusCacheKeySeparator 状态缓存键分隔符
	StatusCacheKeySeparator = ":"
)

// 统计缓存常量
const (
	// StatsCacheDuration 统计缓存有效期（秒）
	StatsCacheDuration = 30
)
