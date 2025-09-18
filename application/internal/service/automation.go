package service

import (
	"context"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/os/gtime"
)

type IAutomation interface {
	// 基础 CRUD 方法
	CreateAutomation(ctx context.Context, startTime, endTime *gtime.Time, timeRangeType, timeRangeValue, triggerRule, executeRule string) (int, error)
	GetAutomationById(ctx context.Context, id int) (*s_db_model.SAutomationModel, error)
	UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error
	DeleteAutomation(ctx context.Context, id int) error

	// 查询方法
	GetAllAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error)
	GetAutomationsByTimeRangeType(ctx context.Context, timeRangeType string) ([]*s_db_model.SAutomationModel, error)
	GetEnabledAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error)
	GetAutomationsByFilters(ctx context.Context, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, error)
	GetAutomationPage(ctx context.Context, page, pageSize int, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, int, error)

	// 统计方法
	ClearAllAutomations(ctx context.Context) error
	GetAutomationCount(ctx context.Context) (int, error)
}

var localAutomation IAutomation

// Automation 获取自动化服务实例
func Automation() IAutomation {
	if localAutomation == nil {
		panic("automation service is not initialized")
	}
	return localAutomation
}

// RegisterAutomation 注册自动化服务
func RegisterAutomation(automation IAutomation) {
	localAutomation = automation
}
