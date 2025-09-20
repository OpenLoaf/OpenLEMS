package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_enum"
	"context"
	"s_db"
)

func (c *ControllerV1) GetDeviceDetailedList(ctx context.Context, req *v1.GetDeviceDetailedListReq) (res *v1.GetDeviceDetailedListRes, err error) {
	// 获取所有设备
	devices, err := s_db.GetDeviceService().GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}

	// 构建设备详细信息列表
	var deviceList []*v1.DeviceDetailedInfo
	deviceManager := common.GetDeviceManager()

	// 创建设备类型过滤映射
	typeFilter := make(map[c_enum.EDeviceType]bool)
	if len(req.DeviceTypes) > 0 {
		for _, deviceType := range req.DeviceTypes {
			typeFilter[deviceType] = true
		}
	}

	for _, device := range devices {
		// 获取设备配置以获取详细信息
		deviceConfig := deviceManager.GetDeviceConfigById(device.Id)
		if deviceConfig == nil {
			continue // 跳过无法获取配置的设备
		}

		// 设备类型过滤
		if len(req.DeviceTypes) > 0 {
			deviceType := deviceConfig.GetType()
			if !typeFilter[deviceType] {
				continue
			}
		}

		// 设备启用状态过滤
		if req.Enabled != nil {
			if device.Enabled != *req.Enabled {
				continue
			}
		}

		// 构建设备详细信息
		deviceInfo := &v1.DeviceDetailedInfo{
			Id:          device.Id,
			Pid:         device.Pid,
			Name:        device.Name,
			ProtocolId:  device.ProtocolId,
			Driver:      device.Driver,
			EnableDebug: device.EnableDebug,
			//Strategy:           device.Strategy,
			StorageEnable:      device.StorageEnable,
			StorageIntervalSec: device.StorageIntervalSec,
			Sort:               device.Sort,
			Enabled:            device.Enabled,
			CreatedAt:          &device.CreatedAt.Time,
			UpdatedAt:          &device.UpdatedAt.Time,
			DeviceType:         deviceConfig.GetType(),
			IsVirtualDevice:    deviceConfig.ProtocolId == "", // 没有协议ID的是虚拟设备
			FailedMessage:      deviceConfig.FailedMessage,
		}

		// 获取设备参数
		if paramsMap, err := device.GetParamsMap(); err == nil {
			deviceInfo.Params = make(map[string]any)
			for k, v := range paramsMap {
				deviceInfo.Params[k] = v
			}
		}

		// 获取驱动信息
		if deviceConfig.DriverInfo != nil {
			deviceInfo.DriverType = string(deviceConfig.DriverInfo.Type)
			deviceInfo.DriverBrand = deviceConfig.DriverInfo.Brand
			deviceInfo.DriverModel = deviceConfig.DriverInfo.Model
			deviceInfo.DriverVersion = deviceConfig.DriverInfo.Version
		}

		// 获取协议信息
		if deviceConfig.ProtocolConfig != nil {
			deviceInfo.ProtocolName = deviceConfig.ProtocolConfig.Name
			deviceInfo.ProtocolType = string(deviceConfig.ProtocolConfig.Type)
			deviceInfo.ProtocolAddress = deviceConfig.ProtocolConfig.Address
		}

		deviceList = append(deviceList, deviceInfo)
	}

	return &v1.GetDeviceDetailedListRes{
		Devices: deviceList,
		Total:   len(deviceList),
	}, nil
}
