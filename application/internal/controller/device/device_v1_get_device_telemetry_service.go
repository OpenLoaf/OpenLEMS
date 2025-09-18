package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_enum"
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

	// 检查设备管理器状态
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		c_log.Warning(ctx, "设备管理器正在初始化中")
		return nil, gerror.NewCode(gcode.CodeInternalError, "系统正在初始化中，请稍后重试")
	}

	// 获取设备配置
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		c_log.Warningf(ctx, "设备不存在: %s", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "设备不存在")
	}

	// 获取设备实例
	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		c_log.Warningf(ctx, "设备实例不存在: %s", req.DeviceId)
		return nil, gerror.NewCode(gcode.CodeNotFound, "设备实例不存在")
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

	c_log.Infof(ctx, "成功获取设备 Telemetry 和 Service: 设备ID=%s, 遥测数量=%d, 服务数量=%d",
		req.DeviceId, len(res.Telemetry), len(res.Service))

	return res, nil
}
