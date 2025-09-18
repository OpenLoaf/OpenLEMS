package impl

import (
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sAutomationServiceImpl struct {
}

var (
	automationServiceInstance s_db_basic.IAutomationService
	automationServiceOnce     sync.Once
)

func GetAutomationService() s_db_basic.IAutomationService {
	automationServiceOnce.Do(func() {
		automationServiceInstance = &sAutomationServiceImpl{}
	})
	return automationServiceInstance
}

// CreateAutomation 创建自动化规则
func (s *sAutomationServiceImpl) CreateAutomation(ctx context.Context, startTime, endTime *time.Time, timeRangeType, timeRangeValue, triggerRule, executeRule string) (int, error) {
	now := gtime.Now()

	// 转换 time.Time 为 gtime.Time
	var startGTime, endGTime *gtime.Time
	if startTime != nil {
		startGTime = gtime.New(*startTime)
	}
	if endTime != nil {
		endGTime = gtime.New(*endTime)
	}

	automationModel := &s_db_model.SAutomationModel{
		StartTime:      startGTime,
		EndTime:        endGTime,
		TimeRangeType:  timeRangeType,
		TimeRangeValue: timeRangeValue,
		TriggerRule:    triggerRule,
		ExecuteRule:    executeRule,
		Enabled:        true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err := automationModel.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建自动化规则失败 - 错误: %+v", err)
		return 0, err
	}

	// 获取刚插入的记录ID
	var lastId int
	err = g.Model(s_db_model.TableAutomation).Ctx(ctx).Order("id DESC").Limit(1).Scan(&lastId)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则ID失败 - 错误: %+v", err)
		return 0, err
	}

	g.Log().Infof(ctx, "成功创建自动化规则，ID: %d", lastId)
	return lastId, nil
}

// GetAutomationById 根据ID获取自动化规则
func (s *sAutomationServiceImpl) GetAutomationById(ctx context.Context, id int) (*s_db_model.SAutomationModel, error) {
	automation := &s_db_model.SAutomationModel{}
	err := automation.GetById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则失败，ID: %d - 错误: %+v", id, err)
		return nil, err
	}
	return automation, nil
}

// UpdateAutomation 更新自动化规则
func (s *sAutomationServiceImpl) UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error {
	automation := &s_db_model.SAutomationModel{}
	err := automation.GetById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则失败，ID: %d - 错误: %+v", id, err)
		return err
	}

	// 更新字段
	if value, ok := data["startTime"].(*time.Time); ok && value != nil {
		automation.StartTime = gtime.New(*value)
	}
	if value, ok := data["endTime"].(*time.Time); ok {
		if value != nil {
			automation.EndTime = gtime.New(*value)
		} else {
			automation.EndTime = nil
		}
	}
	if value, ok := data["timeRangeType"].(string); ok {
		automation.TimeRangeType = value
	}
	if value, ok := data["timeRangeValue"].(string); ok {
		automation.TimeRangeValue = value
	}
	if value, ok := data["triggerRule"].(string); ok {
		automation.TriggerRule = value
	}
	if value, ok := data["executeRule"].(string); ok {
		automation.ExecuteRule = value
	}
	if value, ok := data["enable"].(bool); ok {
		automation.Enabled = value
	}

	// 更新 updated_at
	automation.UpdatedAt = gtime.Now()

	err = automation.Update(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "更新自动化规则失败，ID: %d - 错误: %+v", id, err)
		return err
	}

	g.Log().Infof(ctx, "成功更新自动化规则，ID: %d", id)
	return nil
}

// DeleteAutomation 删除自动化规则
func (s *sAutomationServiceImpl) DeleteAutomation(ctx context.Context, id int) error {
	err := s_db_model.DeleteAutomationById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "删除自动化规则失败，ID: %d - 错误: %+v", id, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除自动化规则，ID: %d", id)
	return nil
}

// GetAllAutomations 获取所有自动化规则
func (s *sAutomationServiceImpl) GetAllAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error) {
	automations, err := s_db_model.GetAllAutomations(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取所有自动化规则失败 - 错误: %+v", err)
		return nil, err
	}
	return automations, nil
}

// GetAutomationsByTimeRangeType 根据时间范围类型获取自动化规则
func (s *sAutomationServiceImpl) GetAutomationsByTimeRangeType(ctx context.Context, timeRangeType string) ([]*s_db_model.SAutomationModel, error) {
	automations, err := s_db_model.GetAutomationsByTimeRangeType(ctx, timeRangeType)
	if err != nil {
		g.Log().Errorf(ctx, "根据时间范围类型获取自动化规则失败，类型: %s - 错误: %+v", timeRangeType, err)
		return nil, err
	}
	return automations, nil
}

// GetEnabledAutomations 获取所有启用的自动化规则
func (s *sAutomationServiceImpl) GetEnabledAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error) {
	automations, err := s_db_model.GetEnabledAutomations(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取启用的自动化规则失败 - 错误: %+v", err)
		return nil, err
	}
	return automations, nil
}

// buildAutomationQuery 构建自动化规则查询条件
func (s *sAutomationServiceImpl) buildAutomationQuery(ctx context.Context, deviceId string, filters map[string]interface{}) *gdb.Model {
	// 构建查询条件
	model := g.Model(s_db_model.TableAutomation).Ctx(ctx)

	// 如果提供了 deviceId，使用 json_extract 从 trigger_rule 和 execute_rule 中查找
	if deviceId != "" {
		// 使用 OR 条件查找 deviceId 在 trigger_rule 或 execute_rule 中的记录
		// json_extract 用于从 JSON 数组中提取 deviceId
		model = model.WhereOr(
			"json_extract(trigger_rule, '$[*].deviceId') LIKE ?", "%"+deviceId+"%",
		).WhereOr(
			"json_extract(execute_rule, '$[*].deviceId') LIKE ?", "%"+deviceId+"%",
		)
	}

	// 应用其他过滤条件
	if filters != nil {
		if timeRangeType, ok := filters["timeRangeType"].(string); ok && timeRangeType != "" {
			model = model.Where(s_db_model.FieldAutomationTimeRangeType, timeRangeType)
		}
		if enable, ok := filters["enable"].(bool); ok {
			model = model.Where(s_db_model.FieldEnabled, enable)
		}
	}

	return model
}

// GetAutomationsByFilters 根据过滤条件获取自动化规则（不分页）
func (s *sAutomationServiceImpl) GetAutomationsByFilters(ctx context.Context, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, error) {
	// 使用公共的查询构建逻辑
	model := s.buildAutomationQuery(ctx, deviceId, filters)

	// 查询所有符合条件的记录
	var automations []*s_db_model.SAutomationModel
	err := model.Scan(&automations)
	if err != nil {
		g.Log().Errorf(ctx, "根据过滤条件获取自动化规则失败 - 错误: %+v", err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取自动化规则，数量: %d", len(automations))
	return automations, nil
}

// GetAutomationPage 分页获取自动化规则
func (s *sAutomationServiceImpl) GetAutomationPage(ctx context.Context, page, pageSize int, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, int, error) {
	// 使用公共的查询构建逻辑
	model := s.buildAutomationQuery(ctx, deviceId, filters)

	// 获取总数
	total, err := model.Count()
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则总数失败 - 错误: %+v", err)
		return nil, 0, err
	}

	// 分页查询
	var automations []*s_db_model.SAutomationModel
	err = model.Page(page, pageSize).Scan(&automations)
	if err != nil {
		g.Log().Errorf(ctx, "分页获取自动化规则失败 - 错误: %+v", err)
		return nil, 0, err
	}

	return automations, total, nil
}

// ClearAllAutomations 清空所有自动化规则
func (s *sAutomationServiceImpl) ClearAllAutomations(ctx context.Context) error {
	_, err := g.Model(s_db_model.TableAutomation).Ctx(ctx).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "清空所有自动化规则失败 - 错误: %+v", err)
		return err
	}

	g.Log().Info(ctx, "成功清空所有自动化规则")
	return nil
}

// GetAutomationCount 获取自动化规则总数
func (s *sAutomationServiceImpl) GetAutomationCount(ctx context.Context) (int, error) {
	count, err := s_db_model.CountAutomations(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则总数失败 - 错误: %+v", err)
		return 0, err
	}
	return count, nil
}
