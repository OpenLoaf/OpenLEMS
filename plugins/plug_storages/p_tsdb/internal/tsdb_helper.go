package internal

import (
	"common/c_base"
	"common/c_enum"

	"github.com/prometheus/prometheus/pkg/labels"
)

// appender 定义本地的Appender接口，用于appendSampleWithCount函数
type appender interface {
	Add(l labels.Labels, t int64, v float64) (uint64, error)
	Commit() error
	Rollback() error
}

// buildDeviceLabels 构建设备数据标签
func buildDeviceLabels(deviceId string, pointKey string) labels.Labels {
	return labels.FromMap(map[string]string{
		LabelNameMetric: MetricNameEmsMetric,
		LabelNameType:   string(c_base.StorageTypeDevice),
		LabelNameID:     deviceId,
		LabelNameField:  pointKey,
	})
}

// buildProtocolLabels 构建协议数据标签
func buildProtocolLabels(protocolId string, deviceId string, field string) labels.Labels {
	return labels.FromMap(map[string]string{
		LabelNameMetric:   MetricNameEmsMetric,
		LabelNameType:     string(c_base.StorageTypeProtocol),
		LabelNameID:       protocolId,
		LabelNameField:    field,
		LabelNameDeviceID: deviceId,
	})
}

// buildSystemLabels 构建系统数据标签
func buildSystemLabels(measurement string, field string) labels.Labels {
	return labels.FromMap(map[string]string{
		LabelNameMetric: MetricNameEmsMetric,
		LabelNameType:   string(c_base.StorageTypeSystem),
		LabelNameID:     measurement,
		LabelNameField:  field,
	})
}

// buildStatusCacheKey 构建状态缓存键
func buildStatusCacheKey(deviceId string, pointKey string) string {
	return deviceId + StatusCacheKeySeparator + pointKey
}

// appendSampleWithCount 添加样本并计数
// 返回值：是否添加成功（用于统计）
func appendSampleWithCount(app appender, lbls labels.Labels, timestamp int64, value float64) (bool, error) {
	_, err := app.Add(lbls, timestamp, value)
	if err != nil {
		return false, err
	}
	return true, nil
}

// getPriceTypeValue 将价格类型字符串转换为数值
func getPriceTypeValue(priceType string) float64 {
	priceTypeEnum := c_enum.ParsePriceType(priceType)
	return float64(priceTypeEnum)
}

// matchesDeviceIdPrefix 检查缓存键是否以指定设备ID开头
func matchesDeviceIdPrefix(cacheKey string, deviceId string) bool {
	prefixLen := len(deviceId) + len(StatusCacheKeySeparator)
	return len(cacheKey) > prefixLen && cacheKey[:prefixLen] == deviceId+StatusCacheKeySeparator
}
