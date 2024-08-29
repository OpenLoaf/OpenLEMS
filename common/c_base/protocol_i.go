package c_base

import (
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type IProtocol interface {
	GetId() string                 // 获取ID
	GetType() EDeviceType          // 获取类型
	Init(deviceType EDeviceType)   // 初始化
	Close() error                  // 关闭
	IsActivate() bool              // 是否有效，无效一般是连接断了
	PrintCacheValues()             // 打印缓存值
	GetCache() *gcache.Cache       // 获取缓存
	GetLastUpdateTime() *time.Time // 获取最后更新时间

	IAlarmHandler
}
