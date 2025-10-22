package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ClearDeviceHistory(ctx context.Context, req *v1.ClearDeviceHistoryReq) (res *v1.ClearDeviceHistoryRes, err error) {
	// 记录业务操作开始
	c_log.BizInfo(ctx, "清除设备历史数据操作开始", "deviceId", req.DeviceId)

	// 验证设备是否存在
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		c_log.BizError(ctx, "设备不存在", "deviceId", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "设备不存在")
	}

	// 获取存储实例
	storageInstance := common.GetStorageInstance()
	if storageInstance == nil {
		c_log.BizError(ctx, "存储实例未初始化", "deviceId", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeInternalError, "存储服务不可用")
	}

	// 调用存储服务清除设备历史数据
	err = storageInstance.ClearDeviceHistoryAll(req.DeviceId)
	if err != nil {
		c_log.BizError(ctx, "清除设备历史数据失败", "deviceId", req.DeviceId, "error", err)
		c_log.Error(ctx, "Clear Device History Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "清除设备历史数据失败")
	}

	// 记录业务操作成功
	c_log.BizInfo(ctx, "清除设备历史数据操作成功", "deviceId", req.DeviceId)

	return &v1.ClearDeviceHistoryRes{}, nil
}
