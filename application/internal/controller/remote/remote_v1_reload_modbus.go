package remote

import (
	"common/c_log"
	"context"
	"s_export_modbus"

	v1 "application/api/remote/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ReloadModbus 重新加载Modbus服务配置
func (c *ControllerV1) ReloadModbus(ctx context.Context, req *v1.ReloadModbusReq) (res *v1.ReloadModbusRes, err error) {
	c_log.Info(ctx, "开始重新加载Modbus服务配置")

	// 调用Modbus服务重新加载配置
	err = s_export_modbus.ReloadModbus(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载Modbus服务配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "重新加载Modbus服务配置失败")
	}

	c_log.Info(ctx, "成功重新加载Modbus服务配置")
	return &v1.ReloadModbusRes{}, nil
}
