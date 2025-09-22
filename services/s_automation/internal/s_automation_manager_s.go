package internal

import (
	"common"
	"common/c_base"
	"common/c_log"
	"context"
	"encoding/json"
	"s_db"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/expr-lang/expr"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SAutomationManager 自动化管理器实现
type SAutomationManager struct {
	automations map[int]*SAutomationTask // 内存中的任务缓存（预解析规则）
	mu          sync.RWMutex             // 读写锁
	ctx         context.Context          // 上下文
	cancel      context.CancelFunc       // 取消函数
	ticker      *time.Ticker             // 定时器
	isRunning   bool                     // 运行状态
}

var (
	automationManagerInstance IAutomationManager
	automationManagerOnce     sync.Once
)

// GetAutomationManager 获取自动化管理器单例
func GetAutomationManager() IAutomationManager {
	automationManagerOnce.Do(func() {
		automationManagerInstance = &SAutomationManager{
			automations: make(map[int]*SAutomationTask),
		}
	})
	return automationManagerInstance
}

// Start 启动自动化管理器
func (m *SAutomationManager) Start(ctx context.Context, interval time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.isRunning {
		g.Log().Warning(ctx, "自动化管理器已经在运行中")
		return nil
	}

	// 创建可取消的上下文
	m.ctx, m.cancel = context.WithCancel(ctx)

	// 加载所有自动化任务
	err := m.loadAllAutomations(m.ctx)
	if err != nil {
		g.Log().Errorf(m.ctx, "加载自动化任务失败: %+v", err)
		return err
	}

	// 启动定时器，使用自定义间隔
	m.ticker = time.NewTicker(interval)
	m.isRunning = true

	// 启动执行协程
	go m.executionLoop()

	g.Log().Infof(m.ctx, "自动化管理器启动成功，执行间隔: %v", interval)
	return nil
}

// Stop 停止自动化管理器
func (m *SAutomationManager) Stop(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		g.Log().Warning(ctx, "自动化管理器未在运行")
		return nil
	}

	// 停止定时器
	if m.ticker != nil {
		m.ticker.Stop()
	}

	// 取消上下文
	if m.cancel != nil {
		m.cancel()
	}

	m.isRunning = false
	g.Log().Info(ctx, "自动化管理器已停止")
	return nil
}

// AddAutomation 添加自动化任务
func (m *SAutomationManager) AddAutomation(ctx context.Context, automation *s_db_model.SAutomationModel) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 设置创建和更新时间
	now := gtime.Now()
	automation.SetCreatedAt(now)
	automation.SetUpdatedAt(now)

	// 保存到数据库
	var startTime, endTime *time.Time
	if automation.GetStartTime() != nil {
		t := automation.GetStartTime().Time
		startTime = &t
	}
	if automation.GetEndTime() != nil {
		t := automation.GetEndTime().Time
		endTime = &t
	}

	id, err := s_db.GetAutomationService().CreateAutomation(
		ctx,
		automation.GetName(),
		startTime,
		endTime,
		automation.GetTimeRangeType(),
		automation.GetTimeRangeValue(),
		automation.GetTriggerRule(),
		automation.GetExecuteRule(),
		automation.GetExecutionInterval(),
	)
	if err != nil {
		g.Log().Errorf(ctx, "创建自动化任务失败: %+v", err)
		return err
	}

	// 设置ID并添加到内存缓存
	automation.SetId(id)

	// 创建预解析的任务
	task, err := m.createAutomationTask(automation)
	if err != nil {
		g.Log().Errorf(ctx, "创建自动化任务失败，ID: %d, 错误: %+v", id, err)
		return err
	}

	m.automations[id] = task

	g.Log().Infof(ctx, "成功添加自动化任务，ID: %d", id)
	return nil
}

// RemoveAutomation 删除自动化任务
func (m *SAutomationManager) RemoveAutomation(ctx context.Context, id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 从数据库删除
	err := s_db.GetAutomationService().DeleteAutomation(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "删除自动化任务失败，ID: %d, 错误: %+v", id, err)
		return err
	}

	// 从内存缓存删除
	delete(m.automations, id)

	g.Log().Infof(ctx, "成功删除自动化任务，ID: %d", id)
	return nil
}

// UpdateAutomation 更新自动化任务
func (m *SAutomationManager) UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 更新数据库
	err := s_db.GetAutomationService().UpdateAutomation(ctx, id, data)
	if err != nil {
		g.Log().Errorf(ctx, "更新自动化任务失败，ID: %d, 错误: %+v", id, err)
		return err
	}

	// 重新加载该任务到内存缓存
	automation, err := s_db.GetAutomationService().GetAutomationById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "重新加载自动化任务失败，ID: %d, 错误: %+v", id, err)
		return err
	}

	// 创建预解析的任务
	task, err := m.createAutomationTask(automation)
	if err != nil {
		g.Log().Errorf(ctx, "重新创建自动化任务失败，ID: %d, 错误: %+v", id, err)
		return err
	}

	m.automations[id] = task

	g.Log().Infof(ctx, "成功更新自动化任务，ID: %d", id)
	return nil
}

// EnableAutomation 启用自动化任务
func (m *SAutomationManager) EnableAutomation(ctx context.Context, id int) error {
	return m.UpdateAutomation(ctx, id, map[string]interface{}{
		s_db_model.FieldEnabled:   true,
		s_db_model.FieldUpdatedAt: gtime.Now(),
	})
}

// DisableAutomation 停用自动化任务
func (m *SAutomationManager) DisableAutomation(ctx context.Context, id int) error {
	return m.UpdateAutomation(ctx, id, map[string]interface{}{
		s_db_model.FieldEnabled:   false,
		s_db_model.FieldUpdatedAt: gtime.Now(),
	})
}

// ReloadAutomations 重新加载所有任务
func (m *SAutomationManager) ReloadAutomations(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.loadAllAutomations(ctx)
}

// loadAllAutomations 从数据库加载所有自动化任务
func (m *SAutomationManager) loadAllAutomations(ctx context.Context) error {
	automations, err := s_db.GetAutomationService().GetAllAutomations(ctx)
	if err != nil {
		return err
	}

	// 清空现有缓存
	m.automations = make(map[int]*SAutomationTask)

	// 加载到内存缓存
	for _, automation := range automations {
		task, err := m.createAutomationTask(automation)
		if err != nil {
			g.Log().Errorf(ctx, "创建自动化任务失败，ID: %d, 错误: %+v", automation.GetId(), err)
			continue
		}
		m.automations[automation.GetId()] = task
	}

	g.Log().Infof(ctx, "成功加载 %d 个自动化任务", len(automations))
	return nil
}

// createAutomationTask 创建预解析的自动化任务
func (m *SAutomationManager) createAutomationTask(automation *s_db_model.SAutomationModel) (*SAutomationTask, error) {
	task := &SAutomationTask{
		SAutomationModel: automation,
		TriggerConfig:    nil,
		ExecuteConfig:    nil,
	}

	// 预解析触发配置
	if automation.GetTriggerRule() != "" {
		var triggerConfig SAutomationTriggerConfig
		err := json.Unmarshal([]byte(automation.GetTriggerRule()), &triggerConfig)
		if err != nil {
			return nil, err
		}
		task.SetTriggerConfig(&triggerConfig)
	}

	// 预解析执行配置
	if automation.GetExecuteRule() != "" {
		var executeConfig SAutomationExecuteConfig
		err := json.Unmarshal([]byte(automation.GetExecuteRule()), &executeConfig)
		if err != nil {
			return nil, err
		}
		task.SetExecuteConfig(&executeConfig)
	}

	return task, nil
}

// executionLoop 执行循环
func (m *SAutomationManager) executionLoop() {
	for {
		select {
		case <-m.ctx.Done():
			g.Log().Info(m.ctx, "自动化执行循环已停止")
			return
		case <-m.ticker.C:
			m.executeAutomations()
		}
	}
}

// executeAutomations 执行所有启用的自动化任务
func (m *SAutomationManager) executeAutomations() {
	m.mu.RLock()
	enabledTasks := make([]*SAutomationTask, 0)
	for _, task := range m.automations {
		if task.GetEnabled() {
			enabledTasks = append(enabledTasks, task)
		}
	}
	m.mu.RUnlock()

	for _, task := range enabledTasks {
		go m.executeAutomation(task)
	}
}

// executeAutomation 执行单个自动化任务
func (m *SAutomationManager) executeAutomation(task *SAutomationTask) {
	if task.ExecuteConfig == nil || task.TriggerConfig == nil {
		return
	}
	// 检查时间范围
	if !m.isInTimeRange(task.SAutomationModel) {
		return
	}

	// 检查触发配置（使用预解析的配置）
	if !m.checkTriggerConfig(task) {
		return
	}

	// 执行规则（暂时留空，等待后续实现）
	m.executeAutomationRules(task)
}

// isInTimeRange 检查是否在时间范围内
func (m *SAutomationManager) isInTimeRange(automation *s_db_model.SAutomationModel) bool {
	now := time.Now()

	// 检查开始时间
	if automation.GetStartTime() != nil && now.Before(automation.GetStartTime().Time) {
		return false
	}

	// 检查结束时间
	if automation.GetEndTime() != nil && now.After(automation.GetEndTime().Time) {
		return false
	}

	// TODO: 根据 timeRangeType 和 timeRangeValue 进行更复杂的时间范围检查
	// 例如：每天特定时间、每周特定日期等

	return true
}

// checkTriggerConfig 检查触发配置（使用预解析的配置）
func (m *SAutomationManager) checkTriggerConfig(task *SAutomationTask) bool {
	config := task.GetTriggerConfig()
	if config == nil {
		return true // 没有触发配置时默认触发
	}

	// 检查 anyMatch 条件（OR 逻辑）
	if len(config.AnyMatch) > 0 {
		anyMatchResult := false
		for _, condition := range config.AnyMatch {
			if m.checkTriggerCondition(condition) {
				anyMatchResult = true
				break // 任意一个满足即可
			}
		}
		if !anyMatchResult {
			return false // anyMatch 都不满足
		}
	}

	// 检查 subMatch 条件
	if len(config.SubMatch) > 0 {
		if config.SubMatchAll {
			// 全部满足
			for _, condition := range config.SubMatch {
				if !m.checkTriggerCondition(condition) {
					return false
				}
			}
		} else {
			// 任意一个满足
			subMatchResult := false
			for _, condition := range config.SubMatch {
				if m.checkTriggerCondition(condition) {
					subMatchResult = true
					break
				}
			}
			if !subMatchResult {
				return false
			}
		}
	}

	return true
}

// checkTriggerCondition 检查触发条件
func (m *SAutomationManager) checkTriggerCondition(condition *SAutomationTriggerCondition) bool {
	// 验证触发条件
	if err := condition.Validate(); err != nil {
		g.Log().Errorf(m.ctx, "触发条件验证失败: %+v", err)
		return false
	}

	// 检查设备条件
	deviceResult := true
	if condition.IsDeviceCondition() {
		deviceResult = m.checkDeviceCondition(condition.DeviceCondition)
	}

	// 检查时间条件
	timeResult := true
	if condition.IsTimeCondition() {
		timeResult = m.checkTimeCondition(condition.TimeCondition)
	}

	// 如果同时设置了设备条件和时间条件，需要都满足
	// 如果只设置了一种条件，只需要该条件满足
	finalResult := deviceResult && timeResult

	g.Log().Debugf(m.ctx, "触发条件检查结果 - 设备条件: %v, 时间条件: %v, 最终结果: %v",
		deviceResult, timeResult, finalResult)

	return finalResult
}

// checkDeviceCondition 检查设备条件
func (m *SAutomationManager) checkDeviceCondition(deviceCondition *SAutomationDeviceCondition) bool {
	// 获取设备实例
	deviceInstance := common.GetDeviceManager().GetDeviceById(deviceCondition.DeviceId)
	if deviceInstance == nil {
		g.Log().Warningf(m.ctx, "设备不存在: %s", deviceCondition.DeviceId)
		return false
	}

	// 获取设备的遥测数据
	telemetryMap := c_base.GetAllTelemetry(deviceInstance)
	if len(telemetryMap) == 0 {
		g.Log().Debugf(m.ctx, "设备遥测数据为空: %s", deviceCondition.DeviceId)
		return false
	}

	g.Log().Infof(m.ctx, "检查设备条件 - 设备: %s, 规则: %s, 遥测数据: %+v",
		deviceCondition.DeviceId, deviceCondition.Rule, telemetryMap)

	// 使用 expr 包解析和验证规则表达式
	result, err := m.evaluateRuleExpression(deviceCondition.Rule, telemetryMap)
	if err != nil {
		g.Log().Errorf(m.ctx, "规则表达式解析失败 - 设备: %s, 规则: %s, 错误: %+v",
			deviceCondition.DeviceId, deviceCondition.Rule, err)
		return false
	}

	g.Log().Infof(m.ctx, "设备条件检查结果 - 设备: %s, 规则: %s, 结果: %v",
		deviceCondition.DeviceId, deviceCondition.Rule, result)

	return result
}

// checkTimeCondition 检查时间条件
func (m *SAutomationManager) checkTimeCondition(timeCondition *SAutomationTimeCondition) bool {
	now := time.Now()

	g.Log().Infof(m.ctx, "检查时间条件 - 当前时间: %s, 条件: %+v",
		now.Format("2006-01-02 15:04:05"), timeCondition)

	result := timeCondition.IsTimeMatch(now)

	g.Log().Infof(m.ctx, "时间条件检查结果 - 当前时间: %s, 结果: %v",
		now.Format("2006-01-02 15:04:05"), result)

	return result
}

// evaluateRuleExpression 使用 expr 包评估规则表达式
func (m *SAutomationManager) evaluateRuleExpression(rule string, telemetryMap map[string]any) (bool, error) {
	// 编译表达式
	program, err := expr.Compile(rule, expr.Env(telemetryMap))
	if err != nil {
		return false, err
	}

	// 执行表达式
	result, err := expr.Run(program, telemetryMap)
	if err != nil {
		return false, err
	}

	// 将结果转换为布尔值
	switch v := result.(type) {
	case bool:
		return v, nil
	case string:
		// 处理字符串类型的布尔值
		return v == "true" || v == "1", nil
	case int, int8, int16, int32, int64:
		// 处理整数类型：非零为 true
		return v != 0, nil
	case float32, float64:
		// 处理浮点数类型：非零为 true
		return v != 0.0, nil
	default:
		// 其他类型尝试转换为字符串再判断
		return false, nil
	}
}

// executeAutomationRules 执行自动化规则
func (m *SAutomationManager) executeAutomationRules(task *SAutomationTask) {
	// TODO: 根据 ExecuteRule 字段执行相应的操作
	// 这里需要集成设备管理器和控制接口
	// 示例逻辑：
	// 1. 解析 ExecuteRule JSON 字符串
	// 2. 根据解析结果执行相应的设备操作
	// 3. 处理执行结果

	executeConfig := task.ExecuteConfig
	if executeConfig == nil {
		return
	}

	for _, rule := range executeConfig.Rules {
		deviceInstance := common.GetDeviceManager().GetDeviceById(rule.DeviceId)
		if deviceInstance == nil {
			c_log.BizErrorf(m.ctx, "自动化任务执行失败！设备[%s]不存在", rule.DeviceId)
			continue
		}

		if deviceInstance.GetConfig().ManualMode {
			c_log.BizInfof(m.ctx, "自动化任务执行跳过！设备[%s]处于手动模式，跳过[%s]服务执行", deviceInstance.GetConfig().Name, rule.Service)
			continue
		}

		// 记录开始执行服务
		c_log.BizInfof(m.ctx, "自动化任务开始执行服务！设备[%s]，服务[%s]，参数[%v]",
			deviceInstance.GetConfig().Name, rule.Service, rule.Params)

		err := c_base.ExecuteCustomService(rule.Service, deviceInstance, rule.Params)
		if err != nil {
			c_log.BizErrorf(m.ctx, "自动化任务执行服务失败！设备[%s]，服务[%s]，原因：%v",
				deviceInstance.GetConfig().Name, rule.Service, err)
		} else {
			c_log.BizInfof(m.ctx, "自动化任务执行服务成功！设备[%s]，服务[%s]",
				deviceInstance.GetConfig().Name, rule.Service)
		}
	}

	//c_log.BizInfof(m.ctx, "自动化任务执行完成！任务ID[%d]，执行规则数量[%d]", task.GetId(), len(executeConfig.Rules))

}
