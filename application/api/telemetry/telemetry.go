// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package telemetry

import (
	"context"

	"application/api/telemetry/v1"
)

type ITelemetryV1 interface {
	GetTelemetryDescription(ctx context.Context, req *v1.GetTelemetryDescriptionReq) (res *v1.GetTelemetryDescriptionRes, err error)
	GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error)
}
