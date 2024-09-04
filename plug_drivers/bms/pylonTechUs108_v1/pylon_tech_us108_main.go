//go:generate ../build.sh
package main

import (
	"context"
	"ems-plan/c_device"
	"pylonTechUs108_v1/pylon_tech_us108"
)

// NewPlugin 必须的方法，不能取消
func NewPlugin(ctx context.Context) c_device.IBms {
	return pylon_tech_us108.NewPlugin(ctx)
}
