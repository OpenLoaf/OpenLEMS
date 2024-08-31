package c_base

import (
	"time"
)

type IProtocol interface {
	IAlarm
	Init(deviceType EDeviceType) // 初始化
	Close() error                // 关闭
	IsActivate() bool            // 是否有效，无效一般是连接断了
	PrintCacheValues()           // 打印缓存值

	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time         // 获取最后更新时间

}
