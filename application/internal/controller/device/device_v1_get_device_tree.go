package device

import (
	v2 "application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {
	var (
		pid        string
		deviceTree = make([]*c_base.SDeviceConfig, 0)
	)

	common.GetDeviceManager().IteratorAssAllDevicesWrapper(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if pid == "" {
			pid = config.Pid
		} else if pid != config.Pid {
			return false
		}
		deviceTree = append(deviceTree, config)
		return true
	})
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}
