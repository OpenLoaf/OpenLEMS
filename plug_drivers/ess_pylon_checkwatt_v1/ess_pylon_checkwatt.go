//go:build !dev && !windows

//go:generate ../build.sh
package main

import (
	"context"
	"ems-plan/c_base"
	"pylon_checkwatt_v1/internal"
)

func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewPlugin(ctx)
}
