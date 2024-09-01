package telemetry

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"application/api/telemetry/v1"
)

func (c *ControllerV1) GetTelemetryDescription(ctx context.Context, req *v1.GetTelemetryDescriptionReq) (res *v1.GetTelemetryDescriptionRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
