// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package station_ess

import (
	"context"

	"application/api/station_ess/v1"
)

type IStationEssV1 interface {
	GetStationEssChildrenTelemetry(ctx context.Context, req *v1.GetStationEssChildrenTelemetryReq) (res *v1.GetStationEssChildrenTelemetryRes, err error)
	GetStationEssTelemetry(ctx context.Context, req *v1.GetStationEssTelemetryReq) (res *v1.GetStationEssTelemetryRes, err error)
}
