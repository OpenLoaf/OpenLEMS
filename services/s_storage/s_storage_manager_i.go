package s_storage

import (
	"common/c_base"
	"context"
)

// IStorageManager 存储管理器接口
type IStorageManager interface {
	// Start 启动存储管理器
	Start(ctx context.Context) error

	// Shutdown 关闭存储管理器
	Shutdown(ctx context.Context) error

	// RegisterDriver 注册设备数据存储任务
	RegisterDriver(deviceConfig *c_base.SDeviceConfig) error

	// GetStorageInstance 获取存储实例
	GetStorageInstance() c_base.IStorage
}
