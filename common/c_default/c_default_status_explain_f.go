package c_default

import (
	"github.com/shockerli/cvt"
)

// StatusExplainBool 布尔值解释
func StatusExplainBool(value any, trueExplain, falseExplain string) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			return trueExplain, nil
		} else {
			return falseExplain, nil
		}
	}
	return "g18n:status_undefined", nil
}

// StatusExplainEnableFunc 解释启用/禁用状态
func StatusExplainEnableFunc(value any) (string, error) {
	return StatusExplainBool(value, "g18n:status_enable", "g18n:status_disable")
}

// StatusExplainProtectFunc 解释保护/正常状态
func StatusExplainProtectFunc(value any) (string, error) {
	return StatusExplainBool(value, "保护", "正常")
}

// StatusExplainIsNotFunc 解释是/否状态
func StatusExplainIsNotFunc(value any) (string, error) {
	return StatusExplainBool(value, "g18n:status_is", "g18n:status_not")
}

// StatusExplainErrorFunc 解释异常/正常状态
func StatusExplainErrorFunc(value any) (string, error) {
	return StatusExplainBool(value, "g18n:status_error", "g18n:status_normal")
}
