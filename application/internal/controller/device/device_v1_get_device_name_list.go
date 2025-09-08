package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
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
	deviceManager := common.GetDeviceManager()

	// 如果指定了设备类型过滤
	if len(req.DeviceTypes) > 0 {
		// 创建设备类型映射，用于快速查找
		typeFilter := make(map[c_base.EDeviceType]bool)
		for _, deviceType := range req.DeviceTypes {
			typeFilter[deviceType] = true
		}

		for _, device := range devices {
			// 获取设备配置以获取设备类型
			deviceConfig := deviceManager.GetDeviceConfigById(device.Id)
			if deviceConfig == nil {
				continue // 跳过无法获取配置的设备
			}

			// 检查设备类型是否在过滤列表中
			deviceType := deviceConfig.GetType()
			if typeFilter[deviceType] {
				deviceNames[device.Id] = device.Name
			}
		}
	} else {
		// 没有类型过滤，返回所有设备
		for _, device := range devices {
			deviceNames[device.Id] = device.Name
		}
	}

	return &v1.GetDeviceNameListRes{
		DeviceNames: deviceNames,
	}, nil
}
