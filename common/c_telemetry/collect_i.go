package c_telemetry

import (
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type ICollection interface {
	GetCache() *gcache.Cache
	GetLastUpdateTime() *time.Time
}
