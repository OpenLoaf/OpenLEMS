package common

import (
	"common/c_base"
	"common/internal/internal_push"
	"context"
)

func BuildPushInstance(builder func(ctx context.Context) c_base.IPush) {
	internal_push.BuildInstance(builder)
}

func GetPushInstance() c_base.IPush {
	return internal_push.GetPushInstance()
}

func ClosePush() {
	internal_push.Close()
}
