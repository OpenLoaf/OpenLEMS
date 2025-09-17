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

	// 调用数据库服务更新设备
	err = s_db.GetDeviceService().UpdateDevice(ctx, req.DeviceId, data)
	if err != nil {
		c_log.Error(ctx, "Update Device Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed)
	}

	if needRestart {
		// 重启设备管理器以应用更改
		common.GetDeviceManager().Shutdown()
		common.GetDeviceManager().Start()
	}

	return &v1.UpdateDeviceRes{
		DeviceId: req.DeviceId,
	}, nil
}
