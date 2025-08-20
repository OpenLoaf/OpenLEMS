package device

import (
	"application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error) {
	var devices = make([]*c_base.SDeviceDetail, 0)

	common.GetDeviceManager().IteratorAssAllDevicesWrapper(func(deviceWrapper c_base.IDeviceWrapper) {
		deviceDetail := deviceWrapper.GetDeviceDetail()

		switch req.ShowType {
		case 1:
			if deviceDetail.IsPhysics {
				devices = append(devices, deviceDetail)
			}
		case 2:
			if !deviceDetail.IsPhysics {
				devices = append(devices, deviceDetail)
			}
		default:
			devices = append(devices, deviceDetail)
		}
	})

	return &v1.GetRealDeviceListRes{Devices: devices}, nil
}
