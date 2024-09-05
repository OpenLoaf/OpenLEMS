//go:build !dev && !windows
// +build !dev,!windows

//go:generate ../build.sh
package main

import (
	"basic_v1/internal"
	"context"
	"ems-plan/c_base"
)

func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewPlugin(ctx)
}
