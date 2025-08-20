package control

import (
	v1 "application/api/control/v1"
	"common"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetCustomServices(ctx context.Context, req *v1.GetCustomServicesReq) (res *v1.GetCustomServicesRes, err error) {

	deviceWrapper := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if deviceWrapper == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	driverDescription := deviceWrapper.GetDeviceDetail().DriverDescription
	if driverDescription == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "driver description not found")
	}

	return &v1.GetCustomServicesRes{
		Services: driverDescription.CustomService,
	}, nil
}
