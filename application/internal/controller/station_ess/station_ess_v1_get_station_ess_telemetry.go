package station_ess

import (
	"application/api/station_ess/v1"
	"application/internal/service"
	"context"
)

func (c *ControllerV1) GetStationEssTelemetry(ctx context.Context, req *v1.GetStationEssTelemetryReq) (res *v1.GetStationEssTelemetryRes, err error) {

	return &v1.GetStationEssTelemetryRes{EssStatus: service.Station().GetEssStatus()}, nil
}
