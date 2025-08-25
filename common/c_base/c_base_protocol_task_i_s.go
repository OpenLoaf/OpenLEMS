package c_base

import "time"

type ITask interface {
	GetName() string
	GetDescription() string
	GetMetas() []*Meta // 获取点位
	GetLifeTime() time.Duration
}
