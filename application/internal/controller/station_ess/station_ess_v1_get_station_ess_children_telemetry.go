package station_ess

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"application/api/station_ess/v1"
)

func (c *ControllerV1) GetStationEssChildrenTelemetry(ctx context.Context, req *v1.GetStationEssChildrenTelemetryReq) (res *v1.GetStationEssChildrenTelemetryRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
