//go:build dev || windows

//go:generate ../build.sh
package bms_pylon_tech_us108_v1

import (
	"context"
	"ems-plan/c_base"
	"pylonTechUs108_v1/internal"
)

// NewPlugin 必须的方法，不能取消
func NewPlugin(ctx context.Context) c_base.IDriver {
	return internal.NewPlugin(ctx)
}
