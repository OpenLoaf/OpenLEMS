package s_storage

import (
	"common/c_base"
	"context"
	"s_storage/internal"
)

// NewStorageManager 创建存储管理器
func NewStorageManager(ctx context.Context, storage c_base.IStorage) IStorageManager {
	return internal.NewStorageManagerImpl(ctx, storage)
}

// StartStorageManager 启动存储管理器
func StartStorageManager(ctx context.Context) error {
	manager := GetStorageManager()
	return manager.Start(ctx)
}

// RegisterStorageDriver 注册设备存储任务
func RegisterStorageDriver(deviceConfig *c_base.SDeviceConfig) error {
	manager := GetStorageManager()
	return manager.RegisterDriver(deviceConfig)
}

// GetStorageInstance 获取存储实例
func GetStorageInstance() c_base.IStorage {
	manager := GetStorageManager()
	return manager.GetStorageInstance()
}

// GetSystemMetrics 获取系统指标
func GetSystemMetrics() map[string]any {
	return internal.GetSystemMetrics()
}

// GetProcessInfo 获取进程信息
func GetProcessInfo() map[string]any {
	return internal.GetProcessInfo()
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() map[string]string {
	return internal.GetSystemInfo()
}



