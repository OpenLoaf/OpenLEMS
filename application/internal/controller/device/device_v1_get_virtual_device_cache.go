package device

import (
	v1 "application/api/device/v1"
	"common"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetVirtualDeviceCache(ctx context.Context, req *v1.GetVirtualDeviceCacheReq) (res *v1.GetVirtualDeviceCacheRes, err error) {
	// TODO: 实现虚拟设备缓存获取业务逻辑

	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}

	//for i, config := range deviceConfig.ChildDeviceConfig {
	//
	//}

	return &v1.GetVirtualDeviceCacheRes{
		DeviceServerState: "",
		AlarmLevel:        "",
		LastUpdateTime:    nil,
		Groups:            nil,
	}, nil
}
