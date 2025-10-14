package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"fmt"
)

// GetDevicePointsDefinition 获取设备全部点位定义
func (c *ControllerV1) GetDevicePointsDefinition(ctx context.Context, req *v1.GetDevicePointsDefinitionReq) (res *v1.GetDevicePointsDefinitionRes, err error) {
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统初始化中，返回空数据
		return &v1.GetDevicePointsDefinitionRes{Fields: []*c_base.SFieldDefinition{}}, nil
	}

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return &v1.GetDevicePointsDefinitionRes{Fields: []*c_base.SFieldDefinition{}}, nil
	}

	// 检查是否为虚拟设备
	if device.IsVirtualDevice() {
		return c.getVirtualDevicePointsDefinition(device)
	}

	// 分别获取设备点位和遥测点位
	devicePoints := device.GetDevicePoints()
	telemetryPoints := device.GetTelemetryPoints()

	if len(devicePoints) == 0 && len(telemetryPoints) == 0 {
		return &v1.GetDevicePointsDefinitionRes{Fields: []*c_base.SFieldDefinition{}}, nil
	}

	fields := make([]*c_base.SFieldDefinition, 0, len(devicePoints)+len(telemetryPoints))

	// 处理设备点位（保持原有group设置）
	for _, p := range devicePoints {
		if fd := c_base.ConvertIPointToFieldDefinition(p); fd != nil {
			fields = append(fields, fd)
		}
	}

	// 处理遥测点位（强制设置group为"i18n:common.summary"）
	for _, p := range telemetryPoints {
		if fd := c_base.ConvertIPointToFieldDefinition(p); fd != nil {
			// 强制设置group为"i18n:common.summary"
			fd.Group = "i18n:common.summary"
			fields = append(fields, fd)
		}
	}

	return &v1.GetDevicePointsDefinitionRes{Fields: fields}, nil
}

// getVirtualDevicePointsDefinition 获取虚拟设备点位定义
func (c *ControllerV1) getVirtualDevicePointsDefinition(device c_base.IDevice) (*v1.GetDevicePointsDefinitionRes, error) {
	deviceConfig := device.GetConfig()
	if deviceConfig == nil {
		return &v1.GetDevicePointsDefinitionRes{Fields: []*c_base.SFieldDefinition{}}, nil
	}

	var allFields []*c_base.SFieldDefinition

	// 遍历所有子设备
	for _, childConfig := range deviceConfig.ChildDeviceConfig {
		childDevice := common.GetDeviceManager().GetDeviceById(childConfig.Id)
		if childDevice == nil {
			continue
		}

		// 获取子设备的遥测点位
		telemetryPoints := childDevice.GetTelemetryPoints()
		deviceName := childConfig.Name

		for _, point := range telemetryPoints {
			if point == nil || point.IsHidden() {
				continue
			}

			// 转换为FieldDefinition
			if fd := c_base.ConvertIPointToFieldDefinition(point); fd != nil {
				// 设置分组信息：设备名:汇总
				fd.Group = fmt.Sprintf("%s:汇总", deviceName)
				allFields = append(allFields, fd)
			}
		}
	}

	return &v1.GetDevicePointsDefinitionRes{Fields: allFields, IsVirtualDevice: true}, nil
}
