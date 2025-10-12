package s_export_mqtt

import (
	"context"

	"common/c_log"
	"s_mqtt/internal"
)

// Init 初始化MQTT服务
func Init() error {
	// 初始化MQTT管理器
	manager := internal.GetMqttManager()
	_ = manager // 这里可以根据需要添加其他初始化逻辑

	c_log.Info(context.Background(), "MQTT服务初始化完成")
	return nil
}

// StartMqtt 启动MQTT服务
func StartMqtt(ctx context.Context) error {
	manager := internal.GetMqttManager()
	return manager.Start(ctx)
}

// StopMqtt 停止MQTT服务
func StopMqtt(ctx context.Context) error {
	manager := internal.GetMqttManager()
	return manager.Stop(ctx)
}

// ReloadMqtt 重新加载MQTT配置
func ReloadMqtt(ctx context.Context) error {
	manager := internal.GetMqttManager()
	return manager.Reload(ctx)
}

// GetMqttStatus 获取MQTT服务状态
func GetMqttStatus() (bool, int, []*internal.SMqttClientStatus) {
	manager := internal.GetMqttManager()
	clientStatusList := manager.GetAllClientStatus()
	return manager.IsRunning(), manager.GetClientCount(), clientStatusList
}
