package c_base

import (
	"context"
	"time"
)

type IDriver interface {
	GetId() string
	GetType() EDeviceType

	GetIsMaster() bool
	GetName() string

	GetDescription() SDescription

	GetLastUpdateTime() *time.Time

	Init(ctx context.Context, client IProtocol, cfg any) error

	IsActivate() bool // 是否有效，无效一般是连接断了

	IAlarmHandler
}
