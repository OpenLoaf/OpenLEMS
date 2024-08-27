package common

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/internal/internal_device"
	"ems-plan/internal/internal_meta"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// DeviceInstance 所有设备实例
var DeviceInstance = internal_device.Instances

func MetaTransformAndCache(ctx context.Context, deviceId string, meta *c_base.Meta, value any, alarmProvider any, cache *gcache.Cache, lifetime time.Duration) (*gvar.Var, error) {
	return internal_meta.MetaProcess(ctx, deviceId, meta, value, alarmProvider, cache, lifetime)
}
