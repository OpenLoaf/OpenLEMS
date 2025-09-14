package internal

import (
	"context"
	"s_db/s_db_model"
	"time"
)

// IAutomationManager 自动化管理器接口
type IAutomationManager interface {
	// 启动管理器
	Start(ctx context.Context, interval time.Duration) error
	// 停止管理器
	Stop(ctx context.Context) error
	// 添加自动化任务
	AddAutomation(ctx context.Context, automation *s_db_model.SAutomationModel) error
	// 删除自动化任务
	RemoveAutomation(ctx context.Context, id int) error
	// 更新自动化任务
	UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error
	// 启用自动化任务
	EnableAutomation(ctx context.Context, id int) error
	// 停用自动化任务
	DisableAutomation(ctx context.Context, id int) error
	// 重新加载所有任务
	ReloadAutomations(ctx context.Context) error
}
