package c_base

import (
	"time"
)

type IDriver interface {
	IAlarm
	Init(client IProtocol, cfg any) error

	GetId() string
	GetType() EDeviceType
	GetName() string
	GetIsMaster() bool
	GetDescription() SDescription
	GetLastUpdateTime() *time.Time
	GetMetaValueList() []*MetaValueWrapper

	IsActivate() bool // 是否有效，无效一般是连接断了

}
