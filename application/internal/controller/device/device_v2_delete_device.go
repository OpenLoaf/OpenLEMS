package device

import (
	"context"

	v2 "application/api/device/v2"
	"sqlite"
)

func (c *ControllerV2) DeleteDevice(ctx context.Context, req *v2.DeleteDeviceReq) (res *v2.DeleteDeviceRes, err error) {
	deviceManage := sqlite.NewDeviceManage(ctx)
	err = deviceManage.DeleteDevice(ctx, req.DeviceId)
	if err != nil {
		return nil, err
	}
	return &v2.DeleteDeviceRes{
		DeviceId: req.DeviceId,
	}, nil
}
