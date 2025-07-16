package device

import (
	"context"

	v2 "application/api/device/v2"
	"sqlite"
)

func (c *ControllerV2) CreateDevice(ctx context.Context, req *v2.CreateDeviceReq) (res *v2.CreateDeviceRes, err error) {
	deviceManage := sqlite.NewDeviceManage(ctx)
	deviceId, err := deviceManage.CreateDevice(ctx, req.DeviceName, req.DevicePId)
	if err != nil {
		return nil, err
	}
	return &v2.CreateDeviceRes{
		DeviceId: deviceId,
	}, nil
}
