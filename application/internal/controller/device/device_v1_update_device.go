package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_log"
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/shockerli/cvt"
)

func (c *ControllerV1) UpdateDevice(ctx context.Context, req *v1.UpdateDeviceReq) (res *v1.UpdateDeviceRes, err error) {
	// 记录业务操作开始
	c_log.BizInfo(ctx, "设备更新操作开始", "deviceId", req.DeviceId)

	// 构建更新数据 - 只更新有值的字段
	data := make(map[string]interface{})

	config := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)

	needRestart := false
	// 字符串字段：只有非空字符串才更新
	if req.Name != "" {
		data["name"] = req.Name
		if config != nil {
			config.Name = req.Name
		}
	}
	if req.ProtocolId != "" {
		data["protocolId"] = req.ProtocolId
		needRestart = true
	}
	if req.Driver != "" {
		data["driver"] = req.Driver
		needRestart = true
	}
	if req.LogLevel != "" {
		data["logLevel"] = req.LogLevel
		if config != nil {
			config.LogLevel = req.LogLevel
		}
	}
	if req.Params != "" {
		data["params"] = req.Params
		needRestart = true
	}

	// 指针字段：只有非nil才更新
	if manualMode, er := cvt.BoolE(req.ManualMode); er == nil {
		data["manualMode"] = manualMode
		if config != nil {
			config.ManualMode = manualMode
		}
	}
	if req.Enabled != nil {
		data["enabled"] = *req.Enabled
		needRestart = true
	}
	if req.Sort != nil {
		data["sort"] = *req.Sort
	}
	if sort, er := cvt.IntE(req.Sort); er == nil {
		data["sort"] = sort
		if config != nil {
			config.Sort = sort
		}
	}

	// 记录更新字段信息
	if len(data) > 0 {
		c_log.BizInfo(ctx, "设备更新字段", "deviceId", req.DeviceId, "fields", data)
	} else {
		c_log.BizWarning(ctx, "设备更新无有效字段", "deviceId", req.DeviceId)
	}

	// 调用数据库服务更新设备
	err = s_db.GetDeviceService().UpdateDevice(ctx, req.DeviceId, data)
	if err != nil {
		c_log.BizError(ctx, "设备更新失败", "deviceId", req.DeviceId, "error", err)
		c_log.Error(ctx, "Update Device Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed)
	}

	if needRestart {
		// 记录设备重启操作
		c_log.BizInfo(ctx, "设备配置变更需要重启", "deviceId", req.DeviceId)

		// 重启设备管理器以应用更改
		common.GetDeviceManager().Shutdown()
		common.GetDeviceManager().Start()

		c_log.BizInfo(ctx, "设备重启完成", "deviceId", req.DeviceId)
	}

	// 记录业务操作成功
	c_log.BizInfo(ctx, "设备更新操作成功", "deviceId", req.DeviceId, "updatedFields", len(data))

	return &v1.UpdateDeviceRes{
		DeviceId: req.DeviceId,
	}, nil
}
