package control

import (
	v1 "application/api/control/v1"
	"common"
	"common/c_base"
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ControlDevice(ctx context.Context, req *v1.ControlDeviceReq) (res *v1.ControlDeviceRes, err error) {

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		c_log.BizErrorf(ctx, "手动控制设备失败！设备[%s]不存在", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	config := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if config == nil {
		c_log.BizErrorf(ctx, "手动控制设备失败！设备[%s]配置不存在", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "device not found")
	}

	// 记录开始执行服务
	c_log.BizInfof(ctx, "手动控制设备开始执行服务！设备[%s]，服务[%s]，参数[%v]",
		config.Name, req.CommandName, req.Parameters)

	err = c_base.ExecuteCustomService(req.CommandName, device, req.Parameters)
	if err != nil {
		c_log.BizErrorf(ctx, "手动控制设备执行服务失败！设备[%s]，服务[%s]，原因：%v",
			config.Name, req.CommandName, err)
		return nil, err
	}

	c_log.BizInfof(ctx, "手动控制设备执行服务成功！设备[%s]，服务[%s]",
		config.Name, req.CommandName)

	return &v1.ControlDeviceRes{}, nil
}
