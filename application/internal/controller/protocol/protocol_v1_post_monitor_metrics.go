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

// PostProtocolMonitorMetrics 获取协议详细指标数据
func (c *ControllerV1) PostProtocolMonitorMetrics(ctx context.Context, req *v1.PostProtocolMonitorMetricsReq) (res *v1.PostProtocolMonitorMetricsRes, err error) {
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

	// 处理时间参数 - 如果时间为空，默认查询24小时内的数据
	startTime := req.StartTime
	endTime := req.EndTime
	if startTime == nil || endTime == nil {
		now := time.Now().UnixMilli()
		if startTime == nil {
			// 默认查询24小时前的数据
			startTime = &now
			*startTime = now - int64(24*60*60*1000) // 24小时前
		}
		if endTime == nil {
			// 默认查询到当前时间
			endTime = &now
		}
	}

	// 添加调试日志
	c_log.BizInfof(ctx, "协议监控查询: 筛选后协议数量=%d, 时间范围=%d-%d", len(filteredProtocols), *startTime, *endTime)

	// 创建图表数据
	chartData := c_chart.NewChartData(len(filteredProtocols))

	// 获取所有协议的数据
	for _, protocol := range filteredProtocols {
		// 先尝试查询所有时间范围的数据，看看是否有数据
		allTimeChart, err := common.GetStorageInstance().GetStorageData(
			c_base.StorageTypeProtocol,
			protocol.Id,
			[]string{s_storage.ProtocolMetricTotal},
			nil, // 不限制开始时间
			nil, // 不限制结束时间
			0,   // 不限制步长
		)

		if err != nil {
			c_log.BizErrorf(ctx, "查询协议[%s]所有时间数据失败: %v", protocol.Id, err)
		} else if allTimeChart != nil && len(allTimeChart.Series) > 0 {
			c_log.BizInfof(ctx, "协议[%s]有历史数据，系列数: %d", protocol.Id, len(allTimeChart.Series))
		} else {
			c_log.BizInfof(ctx, "协议[%s]无历史数据", protocol.Id)
		}

		// 获取该协议的详细指标数据
		protocolChart, err := common.GetStorageInstance().GetStorageData(
			c_base.StorageTypeProtocol,
			protocol.Id,
			// 协议详细指标字段列表
			[]string{
				s_storage.ProtocolMetricSuccessRate,
				s_storage.ProtocolMetricAvgResponseMs,
				s_storage.ProtocolMetricTotal,
				s_storage.ProtocolMetricFailed,
			},
			startTime,
			endTime,
			req.Step,
		)
		if err != nil {
			// 添加调试日志
			c_log.BizErrorf(ctx, "获取协议[%s]数据失败: %v, 时间范围: %d-%d", protocol.Id, err, *startTime, *endTime)
			continue // 跳过获取失败的数据
		}

		// 添加调试日志
		if protocolChart == nil {
			c_log.BizInfof(ctx, "协议[%s]返回空数据", protocol.Id)
			continue
		}

		c_log.BizInfof(ctx, "协议[%s]查询到 %d 个系列", protocol.Id, len(protocolChart.Series))

		if len(protocolChart.Series) > 0 {
			// 设置系列名称和类型
			for _, series := range protocolChart.Series {
				series.Name = protocol.Name + " - " + series.Name
				// 根据指标类型设置图表类型
				if series.Name == s_storage.ProtocolMetricSuccessRate || series.Name == s_storage.ProtocolMetricAvgResponseMs {
					series.Type = "heatmap" // 热力图
				} else {
					series.Type = "scatter" // 散点图
				}
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

		if startTime != nil && now < *startTime+int64(60*60*1000) {
			inTimeRange = false
		}
		if endTime != nil && now > *endTime {
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

	return &v1.PostProtocolMonitorMetricsRes{
		ChartData: chartData,
	}, nil
}
