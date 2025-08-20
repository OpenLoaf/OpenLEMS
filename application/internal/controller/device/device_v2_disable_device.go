package device

import (
	v2 "application/api/device/v2"
	"context"
	"s_db"
)

func (c *ControllerV2) DisableDevice(ctx context.Context, req *v2.DisableDeviceReq) (res *v2.DisableDeviceRes, err error) {
	// TODO: 实现停用设备的业务逻辑
	// 1. 验证设备是否存在
	// 2. 检查设备当前状态
	// 3. 执行停用操作
	// 4. 更新设备状态

	s_db.GetDeviceService()

	return &v2.DisableDeviceRes{}, nil
}
