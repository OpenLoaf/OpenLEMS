package protocol

import (
	v1 "application/api/protocol/v1"
	"common"
	"common/c_base"
	"common/c_chart"
	"common/c_log"
	"context"
	"errors"
	"s_db"
	"s_storage"
	"time"
)

// PostProtocolMonitorTrend 获取协议性能趋势数据
func (c *ControllerV1) PostProtocolMonitorTrend(ctx context.Context, req *v1.PostProtocolMonitorTrendReq) (res *v1.PostProtocolMonitorTrendRes, err error) {
	// 参数验证
	if len(req.MetricKeys) == 0 {
		return nil, errors.New("指标key列表不能为空")
	}

	// 获取协议列表
	protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, "")
	if err != nil {
		return nil, err
	}

	// 应用筛选条件
	filteredProtocols := protocols
	if len(req.ProtocolIds) > 0 {
		filteredProtocols = filterProtocolsByIds(filteredProtocols, req.ProtocolIds)
	}
	if len(req.ProtocolTypes) > 0 {
		filteredProtocols = filterProtocolsByTypes(filteredProtocols, req.ProtocolTypes)
	}

	if len(filteredProtocols) == 0 {
		return nil, errors.New("没有找到符合条件的协议")
	}

	// 映射指标key到实际字段名
	metricKeyMap := map[string]string{
		"ProtocolMetricTotal":         s_storage.ProtocolMetricTotal,
		"ProtocolMetricSuccess":       s_storage.ProtocolMetricSuccess,
		"ProtocolMetricFailed":        s_storage.ProtocolMetricFailed,
		"ProtocolMetricSuccessRate":   s_storage.ProtocolMetricSuccessRate,
		"ProtocolMetricAvgResponseMs": s_storage.ProtocolMetricAvgResponseMs,
	}

	// 转换用户传入的指标key为实际字段名
	actualMetricKeys := make([]string, 0, len(req.MetricKeys))
	for _, key := range req.MetricKeys {
		if actualKey, ok := metricKeyMap[key]; ok {
			actualMetricKeys = append(actualMetricKeys, actualKey)
		} else {
			// 如果没有映射，直接使用原始key（可能已经是正确的字段名）
			actualMetricKeys = append(actualMetricKeys, key)
		}
	}

	c_log.BizInfof(ctx, "协议监控趋势查询: 协议数=%d, 原始指标=%v, 实际指标=%v, 时间范围=%v-%v",
		len(filteredProtocols), req.MetricKeys, actualMetricKeys, req.StartTime, req.EndTime)

	// 创建图表数据
	chartData := c_chart.NewChartData(len(filteredProtocols) * len(actualMetricKeys))

	// 获取所有协议的数据
	for _, protocol := range filteredProtocols {
		// 获取该协议的趋势数据
		protocolChart, err := common.GetStorageInstance().GetStorageData(
			c_base.StorageTypeProtocol,
			protocol.Id,
			actualMetricKeys,
			req.StartTime,
			req.EndTime,
			req.Step,
		)
		if err != nil {
			c_log.BizErrorf(ctx, "获取协议[%s]趋势数据失败: %v", protocol.Id, err)
			continue // 跳过获取失败的数据
		}

		if protocolChart != nil && len(protocolChart.Series) > 0 {
			// 设置系列名称和类型
			for _, series := range protocolChart.Series {
				series.Name = protocol.Name + " - " + series.Name
				series.Type = "line" // 趋势图使用线图
				chartData.AddSeries(series)
			}

			// 合并X轴数据
			if len(chartData.XAxis.Data) == 0 {
				chartData.XAxis = protocolChart.XAxis
			}
		}
	}

	// 添加最新数据防止断层
	if chartData != nil && len(chartData.Series) != 0 {
		now := time.Now().UnixMilli()
		inTimeRange := true

		if req.StartTime != nil && now < *req.StartTime+int64(60*60*1000) {
			inTimeRange = false
		}
		if req.EndTime != nil && now > *req.EndTime {
			inTimeRange = false
		}

		// 只有在时间范围内才添加最新数据
		if inTimeRange {
			chartData.AddTimestamp(now)
			for _, s := range chartData.Series {
				s.AppendData("")
			}
		}
	}

	return &v1.PostProtocolMonitorTrendRes{
		ChartData: chartData,
	}, nil
}
