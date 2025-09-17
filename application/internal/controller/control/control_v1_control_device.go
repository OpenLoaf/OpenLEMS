package control

import (
	v1 "application/api/control/v1"
	"common"
	"common/c_base"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ControlDevice(ctx context.Context, req *v1.ControlDeviceReq) (res *v1.ControlDeviceRes, err error) {

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	config := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if config == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	err = c_base.ExecuteCustomService(req.CommandName, device, req.Parameters)

	return &v1.ControlDeviceRes{}, err
}
