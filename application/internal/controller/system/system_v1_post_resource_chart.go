package system

import (
	v1 "application/api/system/v1"
	"common"
	"common/c_base"
	"context"
	"errors"
	"s_storage"
	"time"
)

func (c *ControllerV1) PostSystemResourceChart(ctx context.Context, req *v1.PostSystemResourceChartReq) (res *v1.PostSystemResourceChartRes, err error) {
	// 获取指标字段列表
	metricKeys, ok := s_storage.ResourceMetricsMap[req.Category]
	if !ok {
		return nil, errors.New("不支持的资源类别")
	}

	// 处理时间参数 - 默认查询最近1小时
	startTime := req.StartTime
	endTime := req.EndTime
	if startTime == nil || endTime == nil {
		now := time.Now().UnixMilli()
		if startTime == nil {
			defaultStart := now - int64(60*60*1000) // 1小时前
			startTime = &defaultStart
		}
		if endTime == nil {
			endTime = &now
		}
	}

	// 特殊处理：storage类别查询TSDB统计信息
	if req.Category == "storage" {
		return c.getStorageChart(ctx, startTime, endTime, req.Step)
	}

	// 根据类别选择数据源
	var measurement string
	if req.Category == "process" || req.Category == "service" {
		measurement = c_base.ConstProcess
	} else if req.Category == "network" {
		measurement = c_base.ConstSystem
	}

	// 查询指标数据
	chartData, err := common.GetStorageInstance().GetStorageData(
		c_base.StorageTypeSystem,
		measurement,
		metricKeys,
		startTime,
		endTime,
		req.Step,
	)
	if err != nil {
		return nil, err
	}

	return &v1.PostSystemResourceChartRes{
		ChartData: chartData,
	}, nil
}

// getStorageChart 获取TSDB存储统计图表
func (c *ControllerV1) getStorageChart(ctx context.Context, startTime, endTime *int64, step int) (*v1.PostSystemResourceChartRes, error) {
	// 查询TSDB统计历史数据
	// 注意：需要先确保TSDB统计数据也被定期保存到TSDB中
	metricKeys := []string{
		"samples_per_second",
		"total_series",
	}

	chartData, err := common.GetStorageInstance().GetStorageData(
		c_base.StorageTypeSystem,
		"tsdb_stats",
		metricKeys,
		startTime,
		endTime,
		step,
	)
	if err != nil {
		return nil, err
	}

	return &v1.PostSystemResourceChartRes{
		ChartData: chartData,
	}, nil
}
