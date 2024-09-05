//go:build dev || windows

package sess_basic_v1

import (
	"context"
	"ems-plan/c_base"
	"station_energy_store/internal"
)

func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewGroupEnergyStore(ctx)
}
