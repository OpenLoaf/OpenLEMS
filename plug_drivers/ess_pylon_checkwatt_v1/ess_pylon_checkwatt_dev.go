//go:build dev || windows

//go:generate ../build.sh
package ess_pylon_checkwatt_v1

import (
	"context"
	"ems-plan/c_base"
	"pylon_checkwatt_v1/internal"
)

func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewPlugin(ctx)
}
