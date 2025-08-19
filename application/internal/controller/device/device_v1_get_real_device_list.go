package device

import (
	"application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error) {
	var devices = make([]*entity.SDevice, 0)

	common.GetDeviceManager().IteratorAssAllDevicesWrapper(func(deviceWrapper c_base.IDeviceWrapper) {

		device := &entity.SDevice{
			DeviceId:   deviceWrapper.GetDeviceConfig().Id,
			DeviceType: string(deviceWrapper.GetDriverInfo().Type),
			DeviceName: deviceWrapper.GetDriverInfo().Name,
			//
			AlarmLevel:     0,
			LastUpdateTime: "",
		}

		devices = append(devices, device)
	})

	return &v1.GetRealDeviceListRes{Devices: devices}, nil
}
