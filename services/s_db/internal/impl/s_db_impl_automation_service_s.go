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
func (s *sAutomationServiceImpl) CreateAutomation(ctx context.Context, name string, startTime, endTime *time.Time, timeRangeType, timeRangeValue, triggerRule, executeRule string, executionInterval int) (int, error) {
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
		Name:              name,
		StartTime:         startGTime,
		EndTime:           endGTime,
		TimeRangeType:     timeRangeType,
		TimeRangeValue:    timeRangeValue,
		TriggerRule:       triggerRule,
		ExecuteRule:       executeRule,
		ExecutionInterval: executionInterval,
		Enabled:           true,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	err := automationModel.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建自动化规则失败 - 错误: %+v", err)
		return 0, err
	}

	// 获取刚插入的记录ID
	lastIdValue, err := g.Model(s_db_model.TableAutomation).Ctx(ctx).Fields(s_db_model.FieldId).Order("id DESC").Limit(1).Value()
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则ID失败 - 错误: %+v", err)
		return 0, err
	}

	lastId := lastIdValue.Int()

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
	g.Log().Infof(ctx, "开始更新自动化规则，ID: %d, 数据: %+v", id, data)

	automation := &s_db_model.SAutomationModel{}
	err := automation.GetById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则失败，ID: %d - 错误: %+v", id, err)
		return err
	}

	g.Log().Infof(ctx, "获取到自动化规则，当前状态: Enabled=%t", automation.Enabled)

	// 更新字段
	if value, ok := data[s_db_model.FieldAutomationStartTime].(*gtime.Time); ok && value != nil {
		automation.StartTime = value
		g.Log().Infof(ctx, "更新 startTime: %v", value)
	}
	if value, ok := data[s_db_model.FieldAutomationEndTime].(*gtime.Time); ok {
		automation.EndTime = value
		g.Log().Infof(ctx, "更新 endTime: %v", value)
	}
	if value, ok := data[s_db_model.FieldAutomationTimeRangeType].(string); ok {
		automation.TimeRangeType = value
		g.Log().Infof(ctx, "更新 timeRangeType: %s", value)
	}
	if value, ok := data[s_db_model.FieldAutomationTimeRangeValue].(string); ok {
		automation.TimeRangeValue = value
		g.Log().Infof(ctx, "更新 timeRangeValue: %s", value)
	}
	if value, ok := data[s_db_model.FieldAutomationTriggerRule].(string); ok {
		automation.TriggerRule = value
		g.Log().Infof(ctx, "更新 triggerRule: %s", value)
	}
	if value, ok := data[s_db_model.FieldAutomationExecuteRule].(string); ok {
		automation.ExecuteRule = value
		g.Log().Infof(ctx, "更新 executeRule: %s", value)
	}
	if value, ok := data[s_db_model.FieldEnabled].(bool); ok {
		g.Log().Infof(ctx, "找到 enabled 字段，值: %t, 当前值: %t", value, automation.Enabled)
		automation.Enabled = value
		g.Log().Infof(ctx, "更新 enabled: %t", value)
	} else {
		g.Log().Warningf(ctx, "未找到 enabled 字段，数据中的键: %v", getMapKeys(data))
	}
	if value, ok := data[s_db_model.FieldAutomationExecutionInterval].(int); ok {
		automation.ExecutionInterval = value
		g.Log().Infof(ctx, "更新 executionInterval: %d", value)
	}

	// 更新 updated_at
	automation.UpdatedAt = gtime.Now()

	g.Log().Infof(ctx, "准备更新数据库，最终状态: Enabled=%t", automation.Enabled)

	// 使用 UpdateFields 方法，只更新需要的字段
	updateFields := g.Map{
		s_db_model.FieldEnabled:   automation.Enabled,
		s_db_model.FieldUpdatedAt: automation.UpdatedAt,
	}

	// 如果有其他字段需要更新，也添加到 updateFields 中
	if automation.StartTime != nil {
		updateFields[s_db_model.FieldAutomationStartTime] = automation.StartTime
	}
	if automation.EndTime != nil {
		updateFields[s_db_model.FieldAutomationEndTime] = automation.EndTime
	}
	if automation.TimeRangeType != "" {
		updateFields[s_db_model.FieldAutomationTimeRangeType] = automation.TimeRangeType
	}
	if automation.TimeRangeValue != "" {
		updateFields[s_db_model.FieldAutomationTimeRangeValue] = automation.TimeRangeValue
	}
	if automation.TriggerRule != "" {
		updateFields[s_db_model.FieldAutomationTriggerRule] = automation.TriggerRule
	}
	if automation.ExecuteRule != "" {
		updateFields[s_db_model.FieldAutomationExecuteRule] = automation.ExecuteRule
	}
	if automation.ExecutionInterval > 0 {
		updateFields[s_db_model.FieldAutomationExecutionInterval] = automation.ExecutionInterval
	}

	g.Log().Infof(ctx, "更新字段: %+v", updateFields)
	err = automation.UpdateFields(ctx, updateFields)
	if err != nil {
		g.Log().Errorf(ctx, "更新自动化规则失败，ID: %d - 错误: %+v", id, err)
		return err
	}

	g.Log().Infof(ctx, "成功更新自动化规则，ID: %d, 最终状态: Enabled=%t", id, automation.Enabled)
	return nil
}

// getMapKeys 获取 map 的所有键
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

	// 如果提供了 deviceId，使用优化的JSON查询从 trigger_rule 和 execute_rule 中查找
	if deviceId != "" {
		g.Log().Infof(ctx, "构建设备ID过滤条件: %s", deviceId)

		// 使用优化的JSON查询条件
		model = s.buildDeviceIdQuery(model, deviceId)
		g.Log().Infof(ctx, "应用设备ID过滤条件完成")
	}

	// 应用其他过滤条件
	if filters != nil {
		if timeRangeType, ok := filters["timeRangeType"].(string); ok && timeRangeType != "" {
			model = model.Where(s_db_model.FieldAutomationTimeRangeType, timeRangeType)
			g.Log().Infof(ctx, "应用时间范围类型过滤: %s", timeRangeType)
		}
		if enable, ok := filters[s_db_model.FieldEnabled].(bool); ok {
			model = model.Where(s_db_model.FieldEnabled, enable)
			g.Log().Infof(ctx, "应用启用状态过滤: %t", enable)
		}
	}

	return model
}

// buildDeviceIdQuery 构建设备ID查询条件
// 使用优化的JSON查询，支持多种JSON格式的动态查找
func (s *sAutomationServiceImpl) buildDeviceIdQuery(model *gdb.Model, deviceId string) *gdb.Model {
	// 使用SQLite的JSON函数进行动态查询，支持以下格式：
	// 1. trigger_rule: $.anyMatch[*].deviceCondition.deviceId, $.subMatch[*].deviceCondition.deviceId
	// 2. execute_rule: $[*].deviceId (数组格式), $.rules[*].deviceId (对象格式)
	// 3. 字符串匹配: trigger_rule 或 execute_rule 中包含 deviceId

	// 构建优化的OR条件查询
	orCondition := `(
		-- trigger_rule 中的 anyMatch 数组查询
		EXISTS (
			SELECT 1 FROM json_each(trigger_rule, '$.anyMatch') 
			WHERE json_extract(value, '$.deviceCondition.deviceId') = ?
		) OR
		-- trigger_rule 中的 subMatch 数组查询  
		EXISTS (
			SELECT 1 FROM json_each(trigger_rule, '$.subMatch') 
			WHERE json_extract(value, '$.deviceCondition.deviceId') = ?
		) OR
		-- execute_rule 数组格式查询
		EXISTS (
			SELECT 1 FROM json_each(execute_rule) 
			WHERE json_extract(value, '$.deviceId') = ?
		) OR
		-- execute_rule 对象格式查询
		EXISTS (
			SELECT 1 FROM json_each(execute_rule, '$.rules') 
			WHERE json_extract(value, '$.deviceId') = ?
		) OR
		-- 字符串匹配查询（兜底方案）
		trigger_rule LIKE ? OR
		execute_rule LIKE ?
	)`

	// 构建参数数组
	params := []interface{}{
		deviceId,             // anyMatch 查询
		deviceId,             // subMatch 查询
		deviceId,             // execute_rule 数组查询
		deviceId,             // execute_rule 对象查询
		"%" + deviceId + "%", // trigger_rule 字符串匹配
		"%" + deviceId + "%", // execute_rule 字符串匹配
	}

	return model.Where(orCondition, params...)
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
	g.Log().Infof(ctx, "开始分页查询自动化规则 - 页码: %d, 每页数量: %d, 设备ID: %s", page, pageSize, deviceId)

	// 使用公共的查询构建逻辑
	model := s.buildAutomationQuery(ctx, deviceId, filters)

	// 获取总数
	total, err := model.Count()
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化规则总数失败 - 错误: %+v", err)
		return nil, 0, err
	}

	g.Log().Infof(ctx, "查询到自动化规则总数: %d", total)

	// 分页查询
	var automations []*s_db_model.SAutomationModel
	err = model.Page(page, pageSize).Scan(&automations)
	if err != nil {
		g.Log().Errorf(ctx, "分页获取自动化规则失败 - 错误: %+v", err)
		return nil, 0, err
	}

	g.Log().Infof(ctx, "成功获取自动化规则分页数据 - 返回数量: %d", len(automations))
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
