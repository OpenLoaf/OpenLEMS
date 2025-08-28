package device

import (
	"application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error) {
	var devices = make([]*c_base.SDeviceConfig, 0)

	common.GetDeviceManager().IteratorAllDevices(func(config *c_base.SDeviceConfig, instance c_base.IDevice) bool {
		devices = append(devices, config)
		return true
	})

	return &v1.GetRealDeviceListRes{Devices: devices}, nil
}
