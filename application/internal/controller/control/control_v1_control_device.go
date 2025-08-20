package control

import (
	v1 "application/api/control/v1"
	"common"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ControlDevice(ctx context.Context, req *v1.ControlDeviceReq) (res *v1.ControlDeviceRes, err error) {
	deviceWrapper := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if deviceWrapper == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	//instance:=deviceWrapper.GetDeviceInstance()
	driverDescription := deviceWrapper.GetDeviceDetail().DriverDescription
	if driverDescription == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "driver description is nil")
	}

	err = driverDescription.ExecuteCustomService(req.CommandName, deviceWrapper.GetDeviceInstance(), nil)

	return &v1.ControlDeviceRes{}, err
}
