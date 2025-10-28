package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_chart"
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

// translateMetricName 翻译测点名称
func translateMetricName(ctx context.Context, metricName string) string {
	// 尝试翻译测点名称，如果翻译键不存在则返回原始名称
	key := "device.metrics." + metricName
	translated := gi18n.T(ctx, key)

	// 调试日志：输出翻译过程
	g.Log().Debugf(ctx, "翻译测点名称: 原始名称=%s, 翻译键=%s, 翻译结果=%s", metricName, key, translated)

	if translated == "" || translated == key {
		// 翻译失败，使用原始名称
		g.Log().Debugf(ctx, "翻译失败，使用原始名称: %s", metricName)
		return metricName
	}
	// 翻译成功，返回翻译后的名称
	g.Log().Debugf(ctx, "翻译成功: %s -> %s", metricName, translated)
	return translated
}

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
		device := common.GetDeviceManager().GetDeviceById(deviceId)
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
					// 对测点名称进行国际化处理
					translatedName := translateMetricName(ctx, series.Name)
					series.Name = deviceId + " - " + translatedName
					series.Type = "line"
					// 设置单位：优先从设备点位定义读取
					if device != nil {
						// series.Name 此时是 "deviceId - translatedName"，原始字段名为 lineKey
						// 通过点位Key精确获取单位
						unit := getPointUnit(device, lineKey)
						if unit != "" {
							series.Unit = unit
						}
					}
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
					// 对测点名称进行国际化处理
					translatedName := translateMetricName(ctx, series.Name)
					series.Name = deviceId + " - " + translatedName
					series.Type = "bar"
					// 设置单位
					if device != nil {
						unit := getPointUnit(device, barKey)
						if unit != "" {
							series.Unit = unit
						}
					}
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

// getPointUnit 根据点位key从设备定义中获取单位
func getPointUnit(device c_base.IDevice, pointKey string) string {
	if device == nil || pointKey == "" {
		return ""
	}
	// 优先在设备点位中查找
	for _, p := range device.GetDevicePoints() {
		if p != nil && p.GetKey() == pointKey {
			return p.GetUnit()
		}
	}
	// 其次在遥测点位中查找
	for _, p := range device.GetTelemetryPoints() {
		if p != nil && p.GetKey() == pointKey {
			return p.GetUnit()
		}
	}
	return ""
}
