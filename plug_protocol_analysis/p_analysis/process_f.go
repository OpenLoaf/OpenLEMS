package p_analysis

import (
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"golang.org/x/net/context"
	"plug_protocol_analysis/internal"
	"time"
)

func Process(ctx context.Context, value any, cache *gcache.Cache, alarmProvider alarm.IProvider, lifetime time.Duration, meta *c_base.Meta) (*gvar.Var, error) {
	return internal.Process(ctx, value, cache, alarmProvider, lifetime, meta)
}
