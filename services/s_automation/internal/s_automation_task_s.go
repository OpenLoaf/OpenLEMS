package internal

import (
	"s_db/s_db_model"
)

// SAutomationTask 自动化任务结构体，包含预解析的规则
type SAutomationTask struct {
	*s_db_model.SAutomationModel
	TriggerConfig *SAutomationTriggerConfig // 预解析的触发配置
	ExecuteConfig *SAutomationExecuteConfig // 预解析的执行配置
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
func (t *SAutomationTask) SetTriggerConfig(config *SAutomationTriggerConfig) {
	t.TriggerConfig = config
}

// GetTriggerConfig 获取触发配置
func (t *SAutomationTask) GetTriggerConfig() *SAutomationTriggerConfig {
	return t.TriggerConfig
}

// parseExecuteConfig 解析执行配置
func (t *SAutomationTask) parseExecuteConfig() error {
	// 这里需要导入 json 包，但为了避免循环导入，我们在调用处处理
	// 实际解析逻辑在 SAutomationManager 中实现
	return nil
}

// SetExecuteConfig 设置执行配置
func (t *SAutomationTask) SetExecuteConfig(config *SAutomationExecuteConfig) {
	t.ExecuteConfig = config
}

// GetExecuteConfig 获取执行配置
func (t *SAutomationTask) GetExecuteConfig() *SAutomationExecuteConfig {
	return t.ExecuteConfig
}
