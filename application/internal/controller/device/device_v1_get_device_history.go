package device

import (
	v1 "application/api/device/v1"
	"context"
)

func (c *ControllerV1) GetDeviceHistory(ctx context.Context, req *v1.GetDeviceHistoryReq) (res *v1.GetDeviceHistoryRes, err error) {
	// TODO: 实现获取设备历史数据的业务逻辑
	// 1. 验证设备是否存在
	// 2. 验证遥测点位是否有效
	// 3. 验证日期格式是否正确
	// 4. 从存储系统查询历史数据
	// 5. 格式化返回数据

	// 临时返回空数据，等待业务逻辑实现
	return &v1.GetDeviceHistoryRes{
		DeviceId: req.DeviceId,
		Date:     req.Date,
		Data:     []*v1.TelemetryHistoryData{},
	}, nil
}
