package telemetry

import (
	"context"
	common "ems-plan"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"application/api/telemetry/v1"
)

func (c *ControllerV1) GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error) {
	instance := common.DeviceInstance.FindById(req.DeviceId)
	if instance == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}

	// 反射执行方法

	return &v1.GetTelemetryGetRes{}, nil
}
