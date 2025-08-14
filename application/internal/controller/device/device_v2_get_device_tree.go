package device

import (
	"context"
	"s_db"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/frame/g"

	v2 "application/api/device/v2"
	"application/internal/model/entity"
)

func (c *ControllerV2) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {
	deviceList, err := s_db.GetDeviceService().GetDeviceList(ctx)
	if err != nil {
		return nil, err
	}
	deviceTree := BuildDeviceTree(ctx, deviceList, "0")
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}

// BuildDeviceTree 递归构建设备树结构
func BuildDeviceTree(ctx context.Context, devices []*s_db_model.Device, parentId string) []*entity.SDeviceTree {
	var tree []*entity.SDeviceTree

	for _, device := range devices {
		if device.Pid == parentId {
			// 获取设备参数
			_, err := device.GetParamsMap()
			if err != nil {
				g.Log().Errorf(context.Background(), "获取设备参数失败 - 设备ID: %d, 设备名称: %s, 参数原始值: %s, 错误: %v",
					device.Id, device.Name, device.Params, err)
				continue
			}

			// 创建驱动配置
			driverConfig := &entity.SDeviceTree{
				DeviceId:     device.Id,
				ProtocolId:   device.ProtocolId,
				DeviceName:   device.Name,
				DeviceDriver: device.Driver,
				LogLevel:     device.LogLevel,
				Enable:       device.Enable,
				Sort:         device.Sort,
			}

			// 递归获取子设备
			children := BuildDeviceTree(ctx, devices, device.Id)
			if children != nil {
				driverConfig.DeviceChildren = children
			}

			tree = append(tree, driverConfig)
		}
	}

	// 修复潜在的数组越界问题
	if len(tree) == 0 {
		return nil
	}
	return tree
}
