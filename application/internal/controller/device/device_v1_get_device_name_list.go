package device

import (
	v1 "application/api/device/v1"
	"context"
	"s_db"
)

func (c *ControllerV1) GetDeviceNameList(ctx context.Context, req *v1.GetDeviceNameListReq) (res *v1.GetDeviceNameListRes, err error) {
	// 获取所有设备
	devices, err := s_db.GetDeviceService().GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}

	// 构建设备名称映射
	deviceNames := make(map[string]string)
	for _, device := range devices {
		deviceNames[device.Id] = device.Name
	}

	return &v1.GetDeviceNameListRes{
		DeviceNames: deviceNames,
	}, nil
}
