package device

import (
	v2 "application/api/device/v1"
	"common"
	"common/c_log"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"s_db"
)

func (c *ControllerV2) EnableDevice(ctx context.Context, req *v2.EnableDeviceReq) (res *v2.EnableDeviceRes, err error) {
	data := make(map[string]interface{})
	data["enabled"] = true
	err = s_db.GetDeviceService().UpdateDevice(ctx, req.DeviceId, data)
	if err != nil {
		c_log.Error(ctx, "Update Device Error", err)
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed)
	}

	common.GetDeviceManager().Shutdown()
	common.GetDeviceManager().Start()

	return &v2.EnableDeviceRes{}, err
}
