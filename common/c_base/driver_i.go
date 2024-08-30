package c_base

import (
	"time"
)

type IDriver interface {
	GetId() string
	GetType() EDeviceType

	GetIsMaster() bool
	GetName() string

	GetDescription() SDescription

	GetLastUpdateTime() *time.Time

	Init(client IProtocol, cfg any) error

	IsActivate() bool // 是否有效，无效一般是连接断了

	IAlarm
}
