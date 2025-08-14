package device

import (
	"context"
	"s_db"

	v2 "application/api/device/v2"
)

func (c *ControllerV2) CreateDevice(ctx context.Context, req *v2.CreateDeviceReq) (res *v2.CreateDeviceRes, err error) {
	deviceManage := s_db.GetDeviceService()
	deviceId, err := deviceManage.CreateDevice(ctx, req.DeviceName, req.DevicePId)
	if err != nil {
		return nil, err
	}
	return &v2.CreateDeviceRes{
		DeviceId: deviceId,
	}, nil
}
