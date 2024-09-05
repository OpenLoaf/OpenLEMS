//go:build !dev && !windows

//go:generate ../build.sh
package main

import (
	"context"
	"ems-plan/c_base"
	"station_energy_store/internal"
)

func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewGroupEnergyStore(ctx)
}

func main() {

}
