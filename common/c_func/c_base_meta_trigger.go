package c_func

import (
	"fmt"

	"github.com/shockerli/cvt"
)

var IsNotZero = func(value any) (bool, error) {
	return value.(int) != 0, nil
}

var IsZero = func(value any) (bool, error) {
	return value.(int) == 0, nil
}

var StatusExplainEnableFunc = func(value any) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_enable", nil // 启动
		} else {
			// 0 关闭
			return "g18n:status_disable", nil // 关闭
		}
	}
	return "g18n:status_undefined", nil
}

var StatusExplainProtectFunc = func(value any) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_protect", nil // 保护
		} else {
			// 0 正常
			return "g18n:status_normal", nil // 正常
		}
	}
	return "g18n:status_undefined", nil
}

var StatusExplainIsNotFunc = func(value any) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_is", nil // 是
		} else {
			// 0 否
			return "g18n:status_not", nil // 否
		}
	}
	return "g18n:status_undefined", nil
}

var StatusExplainErrorFunc = func(value any) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_error", nil // 异常
		} else {
			// 0 正常
			return "g18n:status_normal", nil // 正常
		}
	}
	return "g18n:status_undefined", nil
}

// BoolExplain 布尔值解释
func BoolExplain(value any, trueExplain, falseExplain string) (string, error) {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			return trueExplain, nil
		} else {
			return falseExplain, nil
		}
	}
	return fmt.Sprintf("未知值: %v", value), nil
}
