package s_storage

import (
	"common/c_base"
	"context"
	"s_storage/internal"
	"sync"
)

var (
	storageManagerInstance IStorageManager
	storageManagerOnce     sync.Once
)

// RegisterStorageManager 注册存储管理器实例
func RegisterStorageManager(manager IStorageManager) {
	storageManagerOnce.Do(func() {
		storageManagerInstance = manager
	})
}

// GetStorageManager 获取存储管理器实例
func GetStorageManager() IStorageManager {
	if storageManagerInstance == nil {
		panic("storage manager is nil !")
	}
	return storageManagerInstance
}

// NewStorageManagerWithConfig 根据配置创建存储管理器
func NewStorageManagerWithConfig(ctx context.Context, storageConfig *c_base.SStorageConfig) IStorageManager {
	// 这里可以根据配置创建不同的存储实例
	// 目前使用TSDB作为默认存储
	storage := internal.NewTSDBStorageInstance(ctx, storageConfig)
	managerImpl := internal.NewStorageManagerImpl(ctx, storage)
	RegisterStorageManager(managerImpl)
	return managerImpl
}
