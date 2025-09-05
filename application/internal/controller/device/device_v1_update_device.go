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

func (c *ControllerV1) UpdateDevice(ctx context.Context, req *v1.UpdateDeviceReq) (res *v1.UpdateDeviceRes, err error) {
	// 构建更新数据
	data := make(map[string]interface{})
	data["name"] = req.Name
	data["protocolId"] = req.ProtocolId
	data["driver"] = req.Driver
	data["logLevel"] = req.LogLevel
	data["enabled"] = req.Enabled
	data["sort"] = req.Sort

	// 调用数据库服务更新设备
	err = s_db.GetDeviceService().UpdateDevice(ctx, req.DeviceId, data)
	if err != nil {
		c_log.Error(ctx, "Update Device Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed)
	}

	// 重启设备管理器以应用更改
	common.GetDeviceManager().Shutdown()
	common.GetDeviceManager().Start()

	return &v1.UpdateDeviceRes{
		DeviceId: req.DeviceId,
	}, nil
}
