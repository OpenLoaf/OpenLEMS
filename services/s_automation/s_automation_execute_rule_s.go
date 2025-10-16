package s_automation

import "s_automation/types"

// 公开别名，供外部直接通过 s_automation 包引用
type (
	SAutomationExecuteRule   = types.SAutomationExecuteRule
	SAutomationExecuteConfig = types.SAutomationExecuteConfig
)

var (
	NewAutomationExecuteRule   = types.NewAutomationExecuteRule
	NewAutomationExecuteConfig = types.NewAutomationExecuteConfig
)
