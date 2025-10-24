package device

import (
	v2 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {
	configTree := common.GetDeviceManager().GetDeviceConfigTree()

	// 应用过滤条件
	filteredTree := c.filterDeviceTree(configTree, req)

	var deviceTree = make([]*entity.SDeviceTree, 0)
	_ = gconv.Scan(filteredTree, &deviceTree)
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}

// filterDeviceTree 根据请求参数过滤设备树
func (c *ControllerV1) filterDeviceTree(configTree []*c_base.SDeviceConfig, req *v2.GetDeviceTreeReq) []*c_base.SDeviceConfig {
	var filteredTree []*c_base.SDeviceConfig

	for _, device := range configTree {
		// 递归过滤设备及其子设备
		filteredDevice := c.filterDeviceRecursive(device, req)
		if filteredDevice != nil {
			filteredTree = append(filteredTree, filteredDevice)
		}
	}

	return filteredTree
}

// filterDeviceRecursive 递归过滤设备
func (c *ControllerV1) filterDeviceRecursive(device *c_base.SDeviceConfig, req *v2.GetDeviceTreeReq) *c_base.SDeviceConfig {
	// 检查 RunningStatus 条件
	if req.RunningStatus != nil {
		deviceInstance := common.GetDeviceManager().GetDeviceById(device.Id)
		isRunning := deviceInstance != nil

		// 根据状态过滤
		shouldInclude := false
		switch *req.RunningStatus {
		case c_enum.EDeviceRunningStatusRunning:
			shouldInclude = isRunning
		case c_enum.EDeviceRunningStatusStopped:
			shouldInclude = !isRunning
		}

		if !shouldInclude {
			// 如果设备本身不满足条件，检查是否有子设备满足条件
			var filteredChildren []*c_base.SDeviceConfig
			for _, child := range device.ChildDeviceConfig {
				filteredChild := c.filterDeviceRecursive(child, req)
				if filteredChild != nil {
					filteredChildren = append(filteredChildren, filteredChild)
				}
			}

			// 如果有子设备满足条件，保留该设备但只显示满足条件的子设备
			if len(filteredChildren) > 0 {
				deviceCopy := *device
				deviceCopy.ChildDeviceConfig = filteredChildren
				return &deviceCopy
			}

			// 没有满足条件的子设备，过滤掉该设备
			return nil
		}
	}

	// 检查 Enabled 条件
	if req.Enabled != nil {
		if device.Enabled != *req.Enabled {
			// 如果设备本身不满足条件，检查是否有子设备满足条件
			var filteredChildren []*c_base.SDeviceConfig
			for _, child := range device.ChildDeviceConfig {
				filteredChild := c.filterDeviceRecursive(child, req)
				if filteredChild != nil {
					filteredChildren = append(filteredChildren, filteredChild)
				}
			}

			// 如果有子设备满足条件，保留该设备但只显示满足条件的子设备
			if len(filteredChildren) > 0 {
				deviceCopy := *device
				deviceCopy.ChildDeviceConfig = filteredChildren
				return &deviceCopy
			}

			// 没有满足条件的子设备，过滤掉该设备
			return nil
		}
	}

	// 递归过滤子设备
	var filteredChildren []*c_base.SDeviceConfig
	for _, child := range device.ChildDeviceConfig {
		filteredChild := c.filterDeviceRecursive(child, req)
		if filteredChild != nil {
			filteredChildren = append(filteredChildren, filteredChild)
		}
	}

	// 创建设备副本并更新子设备列表
	deviceCopy := *device
	deviceCopy.ChildDeviceConfig = filteredChildren
	return &deviceCopy
}
