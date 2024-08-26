package c_base

import (
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type IProtocol interface {
	Start()           // 打开
	Close() error     // 关闭
	IsActivate() bool // 是否有效，无效一般是连接断了
	PrintCacheValues()

	GetCache() *gcache.Cache
	GetLastUpdateTime() *time.Time
}
