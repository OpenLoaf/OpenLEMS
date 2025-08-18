package internal

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

var (
	rwMutex      sync.RWMutex
	pushInstance *sPushInstance
)

type sPushInstance struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	c_base.IPush
}

// GetPushInstance 获取推送服务实例
func GetPushInstance() c_base.IPush {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return pushInstance
}

// Close 关闭推送服务
func Close() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if pushInstance != nil {
		g.Log().Infof(pushInstance.ctx, "PushInstance已经注册！准备注销并重新注册！")
		pushInstance.cancelFunc()
	}
}

// BuildInstance 构建推送服务实例
func BuildInstance(builder func(ctx context.Context) c_base.IPush) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if builder == nil {
		pushInstance = nil
		return
	}
	if pushInstance != nil {
		g.Log().Infof(pushInstance.ctx, "PushInstance已经注册！准备注销并重新注册！")
		pushInstance.cancelFunc()
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Push")
	ctx, cancelFunc := context.WithCancel(ctx)
	pushInstance = &sPushInstance{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		IPush:      builder(ctx),
	}

	go func() {
		_ = <-ctx.Done()
		if pushInstance.IPush != nil {
			pushInstance.IPush.Close()
		}
		pushInstance = nil
		g.Log().Infof(ctx, "推送服务已关闭！")
	}()
}
