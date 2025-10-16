package automation

import (
	"fmt"
	"s_automation"

	"github.com/expr-lang/expr"
)

// validateTriggerRule 验证触发规则中的所有设备条件规则语法
func validateTriggerRule(triggerRule *s_automation.SAutomationTriggerConfig) error {
	if triggerRule == nil {
		return nil
	}

	// 验证 AnyMatch 中的设备条件规则
	for i, condition := range triggerRule.AnyMatch {
		if condition.DeviceCondition != nil && condition.DeviceCondition.Rule != "" {
			_, err := expr.Compile(condition.DeviceCondition.Rule)
			if err != nil {
				return fmt.Errorf("AnyMatch[%d] 设备条件规则语法错误: %s - %v", i, condition.DeviceCondition.Rule, err)
			}
		}
	}

	// 验证 SubMatch 中的设备条件规则
	for i, condition := range triggerRule.SubMatch {
		if condition.DeviceCondition != nil && condition.DeviceCondition.Rule != "" {
			_, err := expr.Compile(condition.DeviceCondition.Rule)
			if err != nil {
				return fmt.Errorf("SubMatch[%d] 设备条件规则语法错误: %s - %v", i, condition.DeviceCondition.Rule, err)
			}
		}
	}

	return nil
}
