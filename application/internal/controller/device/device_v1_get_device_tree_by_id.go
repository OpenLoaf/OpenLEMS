package device

import (
	v2 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetDeviceTreeById(ctx context.Context, req *v2.GetDeviceTreeByIdReq) (res *v2.GetDeviceTreeByIdRes, err error) {
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)

	var deviceTree = &entity.SDeviceTree{}
	_ = gconv.Scan(deviceConfig, &deviceTree)
	return &v2.GetDeviceTreeByIdRes{
		DeviceTree: deviceTree,
	}, nil
}
