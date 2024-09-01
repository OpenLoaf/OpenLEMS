package telemetry

import "example/ems/application/internal/service"

type sTelemetry struct{}

func init() {
	service.RegisterTelemetry(New())
}

func New() service.ITelemetry {
	return &sTelemetry{}
}
