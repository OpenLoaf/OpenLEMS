package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
)

// GetDeviceTelemetryService 获取指定设备的所有 Telemetry 和 Service
func (c *ControllerV1) GetDeviceTelemetryService(ctx context.Context, req *v1.GetDeviceTelemetryServiceReq) (res *v1.GetDeviceTelemetryServiceRes, err error) {
	// 参数验证
	if req.DeviceId == "" {
		return nil, errors.New("设备ID不能为空")
	}

	// 获取设备配置
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		c_log.Warningf(ctx, "设备不存在: %s", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "设备不存在")
	}

	// 获取驱动信息
	driverInfo := deviceConfig.DriverInfo
	if driverInfo == nil {
		c_log.Warningf(ctx, "设备驱动信息不存在: %s", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "设备驱动信息不存在")
	}

	// 构建响应数据
	res = &v1.GetDeviceTelemetryServiceRes{
		DeviceId:   req.DeviceId,
		DeviceName: deviceConfig.Name,
		Telemetry:  driverInfo.Telemetry,
		Service:    driverInfo.Service,
	}

	return res, nil
}
