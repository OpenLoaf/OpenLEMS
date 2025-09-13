package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"context"
	"time"
)

func (c *ControllerV1) GetDeviceHistory(ctx context.Context, req *v1.PostDeviceHistoryReq) (res *v1.PostDeviceHistoryRes, err error) {

	chart, err := common.GetStorageInstance().GetStorageData(c_base.StorageTypeDevice, req.DeviceId, req.TelemetryKeys, req.StartTime, req.EndTime, req.Step)
	if chart != nil && len(chart.Series) != 0 {
		// 添加最新的一条数据，防止数据右边没有断层
		now := time.Now().UnixMilli()
		inTimeRange := true

		if req.StartTime != nil && now < *req.StartTime+int64(60*60*1000) {
			inTimeRange = false
		}
		if req.EndTime != nil && now > *req.EndTime {
			inTimeRange = false
		}

		// 只有在时间范围内才添加最新数据，防止数据右边没有断层
		if inTimeRange {
			chart.AddTimestamp(now)
			for _, s := range chart.Series {
				s.AppendData("")
			}
		}
	}
	//chart.AddSeries()
	return &v1.PostDeviceHistoryRes{
		ChartData: chart,
	}, err
}
