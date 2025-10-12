package s_export_mqtt

import (
	"context"

	"common/c_log"
	"s_export_mqtt/internal"
)

// Init 初始化MQTT服务
func Init() error {
	// 初始化MQTT管理器
	manager := internal.GetMqttExportManager()
	_ = manager // 这里可以根据需要添加其他初始化逻辑

	c_log.Info(context.Background(), "MQTT服务初始化完成")
	return nil
}

// StartMqttExporter 启动MQTT器
func StartMqttExporter(ctx context.Context) error {
	manager := internal.GetMqttExportManager()
	return manager.Start(ctx)
}

// StopMqttExporter 停止MQTT器
func StopMqttExporter(ctx context.Context) error {
	manager := internal.GetMqttExportManager()
	return manager.Stop(ctx)
}

// ReloadMqttExporter 重新加载MQTT器配置
func ReloadMqttExporter(ctx context.Context) error {
	manager := internal.GetMqttExportManager()
	return manager.Reload(ctx)
}

// GetMqttExporterStatus 获取MQTT器状态
func GetMqttExporterStatus() (bool, int) {
	manager := internal.GetMqttExportManager()
	return manager.IsRunning(), manager.GetClientCount()
}
