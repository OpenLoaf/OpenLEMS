package c_base

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type IDriver interface {
	GetId() string
	GetType() EDeviceType
	// GetName() string

	GetDescription() SDescription

	GetCache() *gcache.Cache // 获取缓存
	GetLastUpdateTime() *time.Time

	Init(ctx context.Context, client IProtocol, config any) error

	HasAlarm() (bool, error) // 是否存在故障
	IsActivate() bool        // 是否有效，无效一般是连接断了
}
