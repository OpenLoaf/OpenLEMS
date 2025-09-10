package c_base

import "time"

type IPointTask interface {
	GetName() string
	GetDescription() string
	GetPoints() []IPoint // 获取点位
	GetLifeTime() time.Duration
}
