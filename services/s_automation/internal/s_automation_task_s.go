package internal

import (
	stypes "s_automation/types"
	"s_db/s_db_model"
	"time"
)

// SAutomationTask 自动化任务结构体，包含预解析的规则
type SAutomationTask struct {
	*s_db_model.SAutomationModel
	TriggerConfig     *stypes.SAutomationTriggerConfig // 预解析的触发配置
	ExecuteConfig     *stypes.SAutomationExecuteConfig // 预解析的执行配置
	LastExecutionTime time.Time                        // 上次执行时间（内存缓存）
}

// NewAutomationTask 创建新的自动化任务
func NewAutomationTask(automation *s_db_model.SAutomationModel) (*SAutomationTask, error) {
	task := &SAutomationTask{
		SAutomationModel: automation,
		TriggerConfig:    nil,
		ExecuteConfig:    nil,
	}

	// 预解析触发规则
	if automation.GetTriggerRule() != "" {
		err := task.parseTriggerConfig()
		if err != nil {
			return nil, err
		}
	}

	// 预解析执行规则
	if automation.GetExecuteRule() != "" {
		err := task.parseExecuteConfig()
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}

// parseTriggerConfig 解析触发配置
func (t *SAutomationTask) parseTriggerConfig() error {
	// 这里需要导入 json 包，但为了避免循环导入，我们在调用处处理
	// 实际解析逻辑在 SAutomationManager 中实现
	return nil
}

// SetTriggerConfig 设置触发配置
func (t *SAutomationTask) SetTriggerConfig(config *stypes.SAutomationTriggerConfig) {
	t.TriggerConfig = config
}

// GetTriggerConfig 获取触发配置
func (t *SAutomationTask) GetTriggerConfig() *stypes.SAutomationTriggerConfig {
	return t.TriggerConfig
}

// parseExecuteConfig 解析执行配置
func (t *SAutomationTask) parseExecuteConfig() error {
	// 这里需要导入 json 包，但为了避免循环导入，我们在调用处处理
	// 实际解析逻辑在 SAutomationManager 中实现
	return nil
}

// SetExecuteConfig 设置执行配置
func (t *SAutomationTask) SetExecuteConfig(config *stypes.SAutomationExecuteConfig) {
	t.ExecuteConfig = config
}

// GetExecuteConfig 获取执行配置
func (t *SAutomationTask) GetExecuteConfig() *stypes.SAutomationExecuteConfig {
	return t.ExecuteConfig
}

// GetLastExecutionTime 获取上次执行时间
func (t *SAutomationTask) GetLastExecutionTime() time.Time {
	return t.LastExecutionTime
}

// SetLastExecutionTime 设置上次执行时间
func (t *SAutomationTask) SetLastExecutionTime(executionTime time.Time) {
	t.LastExecutionTime = executionTime
}

// ShouldExecute 判断任务是否应该执行
// defaultInterval: 系统默认执行间隔（毫秒）
func (t *SAutomationTask) ShouldExecute(defaultInterval int64) bool {
	// 如果上次执行时间是零值，说明是首次执行，应该立即执行
	if t.LastExecutionTime.IsZero() {
		return true
	}

	// 获取任务的执行间隔（毫秒）
	interval := int64(t.GetExecutionInterval())

	// 如果任务的执行间隔为0，使用系统默认间隔
	if interval <= 0 {
		interval = defaultInterval
	}

	// 计算距离上次执行的时间（毫秒）
	elapsedMilliseconds := time.Since(t.LastExecutionTime).Milliseconds()

	// 判断是否达到执行间隔
	return elapsedMilliseconds >= interval
}
