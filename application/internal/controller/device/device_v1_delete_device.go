package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_log"
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) DeleteDevice(ctx context.Context, req *v1.DeleteDeviceReq) (res *v1.DeleteDeviceRes, err error) {
	// 记录业务操作开始
	c_log.BizInfo(ctx, "删除设备操作开始", "deviceId", req.DeviceId)

	// 验证设备是否存在
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		c_log.BizError(ctx, "设备不存在", "deviceId", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "设备不存在")
	}

	// 调用数据库服务删除设备
	err = s_db.GetDeviceService().DeleteDeviceById(ctx, req.DeviceId)
	if err != nil {
		c_log.BizError(ctx, "删除设备失败", "deviceId", req.DeviceId, "error", err)
		c_log.Error(ctx, "Delete Device Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除设备失败")
	}

	// 记录业务操作成功
	c_log.BizInfo(ctx, "删除设备操作成功", "deviceId", req.DeviceId)

	return &v1.DeleteDeviceRes{}, nil
}
