package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_chart"
	"context"
	"errors"
	"time"
)

// PostDeviceCompare 获取多设备数据对比
func (c *ControllerV1) PostDeviceCompare(ctx context.Context, req *v1.PostDeviceCompareReq) (res *v1.PostDeviceCompareRes, err error) {
	// 参数验证
	if len(req.DeviceIds) == 0 {
		return nil, errors.New("设备ID列表不能为空")
	}

	if len(req.LineKeys) == 0 && len(req.BarKeys) == 0 {
		return nil, errors.New("线图点位和柱状图点位不能同时为空")
	}

	// 检查并调整时间范围
	startTime := req.StartTime
	endTime := req.EndTime

	// 如果结束时间是未来时间，设置为当前时间的整分钟
	if endTime != nil && *endTime > time.Now().UnixMilli() {
		now := time.Now().Truncate(time.Minute)
		adjustedEndTime := now.UnixMilli()
		endTime = &adjustedEndTime
	}

	// 创建图表数据
	chartData := c_chart.NewChartData(len(req.DeviceIds) * (len(req.LineKeys) + len(req.BarKeys)))

	// 获取所有设备的数据
	for _, deviceId := range req.DeviceIds {
		// 获取线图数据
		for _, lineKey := range req.LineKeys {
			lineChart, err := common.GetStorageInstance().GetStorageData(
				c_base.StorageTypeDevice,
				deviceId,
				[]string{lineKey},
				startTime,
				endTime,
				req.Step,
			)
			if err != nil {
				continue // 跳过获取失败的数据
			}

			if lineChart != nil && len(lineChart.Series) > 0 {
				// 设置系列为线图类型
				for _, series := range lineChart.Series {
					series.Name = deviceId + " - " + series.Name
					series.Type = "line"
					chartData.AddSeries(series)
				}

				// 合并X轴数据
				if len(chartData.XAxis.Data) == 0 {
					chartData.XAxis = lineChart.XAxis
				}
			}
		}

		// 获取柱状图数据
		for _, barKey := range req.BarKeys {
			barChart, err := common.GetStorageInstance().GetStorageData(
				c_base.StorageTypeDevice,
				deviceId,
				[]string{barKey},
				startTime,
				endTime,
				req.Step,
			)
			if err != nil {
				continue // 跳过获取失败的数据
			}

			if barChart != nil && len(barChart.Series) > 0 {
				// 设置系列为柱状图类型
				for _, series := range barChart.Series {
					series.Name = deviceId + " - " + series.Name
					series.Type = "bar"
					chartData.AddSeries(series)
				}

				// 合并X轴数据
				if len(chartData.XAxis.Data) == 0 {
					chartData.XAxis = barChart.XAxis
				}
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

	return &v1.PostDeviceCompareRes{
		ChartData: chartData,
	}, nil
}
