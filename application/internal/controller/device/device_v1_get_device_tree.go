package device

import (
	v2 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {
	deviceConfigs := common.GetDeviceManager().GetTopDeviceConfigs()

	var deviceTree = make([]*entity.SDeviceTree, 0)
	_ = gconv.Scan(deviceConfigs, &deviceTree)
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}
