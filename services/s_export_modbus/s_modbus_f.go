package s_export_modbus

import (
	"context"

	"common/c_log"
	"s_export_modbus/internal"
)

// Init 初始化Modbus服务
func Init() error {
	// 初始化Modbus管理器
	manager := internal.GetModbusManager()
	_ = manager // 这里可以根据需要添加其他初始化逻辑

	c_log.Info(context.Background(), "Modbus服务初始化完成")
	return nil
}

// StartModbus 启动Modbus服务
func StartModbus(ctx context.Context) error {
	manager := internal.GetModbusManager()
	return manager.Start(ctx)
}

// StopModbus 停止Modbus服务
func StopModbus(ctx context.Context) error {
	manager := internal.GetModbusManager()
	return manager.Stop(ctx)
}

// ReloadModbus 重新加载Modbus配置
func ReloadModbus(ctx context.Context) error {
	manager := internal.GetModbusManager()
	return manager.Reload(ctx)
}

// GetModbusStatus 获取Modbus服务状态
func GetModbusStatus() (isRunning bool, port int, deviceCount int) {
	manager := internal.GetModbusManager()
	isRunning, connectionCount, deviceCount := manager.GetServerStatus()
	_ = connectionCount                // 暂时不使用连接数
	return isRunning, 502, deviceCount // 端口暂时硬编码为502
}

// GetModbusDeviceStatus 获取所有设备状态
func GetModbusDeviceStatus() []*internal.SModbusDeviceStatus {
	manager := internal.GetModbusManager()
	return manager.GetAllDeviceStatus()
}
