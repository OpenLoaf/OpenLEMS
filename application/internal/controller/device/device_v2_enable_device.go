package device

import (
	v2 "application/api/device/v2"
	"context"
)

func (c *ControllerV2) EnableDevice(ctx context.Context, req *v2.EnableDeviceReq) (res *v2.EnableDeviceRes, err error) {
	// TODO: 实现启用设备的业务逻辑
	// 1. 验证设备是否存在
	// 2. 检查设备当前状态
	// 3. 执行启用操作
	// 4. 更新设备状态

	return &v2.EnableDeviceRes{}, nil
}
