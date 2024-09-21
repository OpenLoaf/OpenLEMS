package internal_storage

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

var (
	rwMutex          sync.RWMutex
	sStorageInstance *SStorageInstance
)

type SStorageInstance struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	c_base.IStorage
}

func GetInstance() c_base.IStorage {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return sStorageInstance
}

func Close() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if sStorageInstance != nil {
		g.Log().Infof(sStorageInstance.ctx, "StorageInstance已经注册！准备注销并重新注册！")
		sStorageInstance.cancelFunc()
	}

}

func RegisterInstance(builder func(ctx context.Context) c_base.IStorage) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if builder == nil {
		sStorageInstance = nil
		return
	}
	if sStorageInstance != nil {
		g.Log().Infof(sStorageInstance.ctx, "StorageInstance已经注册！准备注销并重新注册！")
		sStorageInstance.cancelFunc()
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Storage")
	ctx, cancelFunc := context.WithCancel(ctx)
	sStorageInstance = &SStorageInstance{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		IStorage:   builder(ctx),
	}

	go func() {
		_ = <-ctx.Done()
		if sStorageInstance.IStorage != nil {
			sStorageInstance.IStorage.Close()
		}
		sStorageInstance = nil
		g.Log().Infof(ctx, "存储服务已关闭！")
	}()
}
