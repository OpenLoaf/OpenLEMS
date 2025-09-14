package s_automation

import (
	"context"
	"s_automation/internal"
	"s_db/s_db_model"
	"time"
)

// Init 初始化自动化服务
func Init() {
	// 初始化自动化管理器
	manager := internal.GetAutomationManager()
	// 这里可以根据需要添加其他初始化逻辑
	_ = manager
}

// GetAutomationManager 获取自动化管理器
func GetAutomationManager() internal.IAutomationManager {
	return internal.GetAutomationManager()
}

// StartAutomationManager 启动自动化管理器
func StartAutomationManager(ctx context.Context, interval time.Duration) error {
	manager := internal.GetAutomationManager()
	return manager.Start(ctx, interval)
}

// StopAutomationManager 停止自动化管理器
func StopAutomationManager(ctx context.Context) error {
	manager := internal.GetAutomationManager()
	return manager.Stop(ctx)
}

// AddAutomation 添加自动化任务
func AddAutomation(ctx context.Context, automation *s_db_model.SAutomationModel) error {
	manager := internal.GetAutomationManager()
	return manager.AddAutomation(ctx, automation)
}

// RemoveAutomation 删除自动化任务
func RemoveAutomation(ctx context.Context, id int) error {
	manager := internal.GetAutomationManager()
	return manager.RemoveAutomation(ctx, id)
}

// UpdateAutomation 更新自动化任务
func UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error {
	manager := internal.GetAutomationManager()
	return manager.UpdateAutomation(ctx, id, data)
}

// EnableAutomation 启用自动化任务
func EnableAutomation(ctx context.Context, id int) error {
	manager := internal.GetAutomationManager()
	return manager.EnableAutomation(ctx, id)
}

// DisableAutomation 停用自动化任务
func DisableAutomation(ctx context.Context, id int) error {
	manager := internal.GetAutomationManager()
	return manager.DisableAutomation(ctx, id)
}

// ReloadAutomations 重新加载所有任务
func ReloadAutomations(ctx context.Context) error {
	manager := internal.GetAutomationManager()
	return manager.ReloadAutomations(ctx)
}
