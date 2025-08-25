package control

import (
	v1 "application/api/control/v1"
	"common"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ControlDevice(ctx context.Context, req *v1.ControlDeviceReq) (res *v1.ControlDeviceRes, err error) {
	deviceConfig := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if deviceConfig == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	// todo 完善执行命令方法
	////instance:=deviceConfig.GetDeviceInstance()
	//driverDescription := deviceConfig.GetDeviceDetail().DriverDescription
	//if driverDescription == nil {
	//	return nil, gerror.NewCode(gcode.CodeNotFound, "driver description is nil")
	//}
	//
	//err = driverDescription.ExecuteCustomService(req.CommandName, deviceConfig.GetDeviceInstance(), nil)

	return &v1.ControlDeviceRes{}, err
}
