package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetDeviceHistory(ctx context.Context, req *v1.PostDeviceHistoryReq) (res *v1.PostDeviceHistoryRes, err error) {

	chart, err := common.GetStorageInstance().GetStorageData(c_base.StorageTypeDevice, req.DeviceId, req.TelemetryKeys, req.StartTime, req.EndTime, req.Step)

	return &v1.PostDeviceHistoryRes{
		ChartData: chart,
	}, err
}
