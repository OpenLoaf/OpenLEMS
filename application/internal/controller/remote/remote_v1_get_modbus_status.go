package remote

import (
	v1 "application/api/remote/v1"
	"common/c_default"
	"common/c_enum"
	"common/c_log"
	"context"
	"s_export_modbus"
)

// GetModbusStatus 获取Modbus服务状态
func (c *ControllerV1) GetModbusStatus(ctx context.Context, req *v1.GetModbusStatusReq) (res *v1.GetModbusStatusRes, err error) {
	c_log.Info(ctx, "获取Modbus服务状态")

	// 获取Modbus服务状态
	isRunning, port, deviceCount := s_export_modbus.GetModbusStatus()

	// 获取所有设备状态
	deviceStatusList := s_export_modbus.GetModbusDeviceStatus()

	// 转换为API响应格式
	apiDeviceStatus := make([]*v1.ModbusDeviceStatus, 0, len(deviceStatusList))

	for _, deviceStatus := range deviceStatusList {
		apiStatus := &v1.ModbusDeviceStatus{
			DeviceId:       deviceStatus.DeviceId,
			ModbusId:       deviceStatus.ModbusId,
			StartAddr:      deviceStatus.StartAddr,
			IsOnline:       deviceStatus.IsOnline,
			LastUpdateTime: deviceStatus.LastUpdateTime,
			Error:          deviceStatus.Error,
		}

		// 获取设备的详细映射信息
		deviceMappings := c.getDevicePointMappings(ctx, deviceStatus.DeviceId)
		apiStatus.PointMappings = deviceMappings
		apiStatus.TotalRegisters = c.calculateTotalRegisters(deviceMappings)

		apiDeviceStatus = append(apiDeviceStatus, apiStatus)
	}

	// 获取连接数（暂时设为0，因为当前实现中连接数信息不可用）
	connectionCount := 0

	res = &v1.GetModbusStatusRes{
		IsRunning:       isRunning,
		ListenPort:      port,
		DeviceCount:     deviceCount,
		ConnectionCount: connectionCount,
		DeviceStatus:    apiDeviceStatus,
	}

	c_log.Infof(ctx, "Modbus服务状态获取成功: 运行状态=%v, 端口=%d, 设备数量=%d", isRunning, port, deviceCount)
	return res, nil
}

// getDevicePointMappings 获取设备的点位映射信息
func (c *ControllerV1) getDevicePointMappings(ctx context.Context, deviceId string) []*v1.ModbusPointMapping {
	// 获取设备映射信息
	deviceMaps := s_export_modbus.GetModbusDeviceMaps()
	deviceMap, exists := deviceMaps[deviceId]
	if !exists {
		c_log.Warningf(ctx, "设备 %s 的映射信息不存在", deviceId)
		return []*v1.ModbusPointMapping{}
	}

	mappings := make([]*v1.ModbusPointMapping, 0, len(deviceMap.Mappings))

	for _, mapping := range deviceMap.Mappings {
		pointMapping := &v1.ModbusPointMapping{
			StartOffset:   mapping.StartOffset,
			RegisterCount: mapping.RegisterCount,
		}

		// 处理系统固定点位
		if mapping.Point == nil {
			switch mapping.StartOffset {
			case 0:
				pointMapping.PointKey = c_default.VPointSystemOnlineStatus.Key
				pointMapping.PointName = c_default.VPointSystemOnlineStatus.Name
				pointMapping.ValueType = c_default.VPointSystemOnlineStatus.ValueType.String()
				pointMapping.Unit = c_default.VPointSystemOnlineStatus.Unit
				pointMapping.Description = c_default.VPointSystemOnlineStatus.Desc
			case 1:
				pointMapping.PointKey = c_default.VPointSystemTimestamp.Key
				pointMapping.PointName = c_default.VPointSystemTimestamp.Name
				pointMapping.ValueType = c_default.VPointSystemTimestamp.ValueType.String()
				pointMapping.Unit = c_default.VPointSystemTimestamp.Unit
				pointMapping.Description = c_default.VPointSystemTimestamp.Desc
			default:
				pointMapping.PointKey = "system.unknown"
				pointMapping.PointName = "未知系统点位"
				pointMapping.ValueType = c_enum.EUint16.String()
				pointMapping.Unit = ""
				pointMapping.Description = "未知系统点位"
			}
		} else {
			// 处理设备实际点位
			pointMapping.PointKey = mapping.Point.GetKey()
			pointMapping.PointName = mapping.Point.GetName()
			pointMapping.ValueType = mapping.Point.GetValueType().String()
			pointMapping.Unit = mapping.Point.GetUnit()
			pointMapping.Description = mapping.Point.GetName() // 使用名称作为描述
		}

		mappings = append(mappings, pointMapping)
	}

	return mappings
}

// calculateTotalRegisters 计算总寄存器数量
func (c *ControllerV1) calculateTotalRegisters(mappings []*v1.ModbusPointMapping) uint16 {
	var total uint16
	for _, mapping := range mappings {
		total += mapping.RegisterCount
	}
	return total
}
