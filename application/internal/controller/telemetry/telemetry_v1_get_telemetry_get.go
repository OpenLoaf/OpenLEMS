package telemetry

import (
	"context"
	common "ems-plan"
	"ems-plan/util"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"application/api/telemetry/v1"
)

func (c *ControllerV1) GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error) {
	instance := common.GetDeviceById(req.DeviceId)
	if instance == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}

	// 反射执行方法
	value, err := util.ExecuteFunction(instance, req.TelemetryKey)
	if err != nil {
		return nil, err
	}

	return &v1.GetTelemetryGetRes{
		TestJoinKey:      req.DeviceId + ":" + req.TelemetryKey,
		DeviceId:         req.DeviceId,
		TelemetryKey:     req.TelemetryKey,
		TelemetryKeyName: g.I18n().T(ctx, req.TelemetryKey),
		Value:            value,
	}, nil
}
