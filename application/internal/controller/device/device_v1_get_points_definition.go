package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
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

	points := device.GetDevicePoints()
	if len(points) == 0 {
		return &v1.GetDevicePointsDefinitionRes{Fields: []*c_base.SFieldDefinition{}}, nil
	}

	fields := make([]*c_base.SFieldDefinition, 0, len(points))
	for _, p := range points {
		if fd := c_base.ConvertIPointToFieldDefinition(p); fd != nil {
			fields = append(fields, fd)
		}
	}

	return &v1.GetDevicePointsDefinitionRes{Fields: fields}, nil
}
