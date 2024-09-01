package telemetry

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"application/api/telemetry/v1"
)

func (c *ControllerV1) GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
