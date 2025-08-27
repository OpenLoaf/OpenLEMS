package c_func

import (
	"fmt"
	"github.com/shockerli/cvt"
)

var IsNotZero = func(value any) bool {
	return value.(int) != 0
}

var IsZero = func(value any) bool {
	return value.(int) == 0
}

var StatusExplainEnableFunc = func(value any) string {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_enable" // 启动
		} else {
			// 0 关闭
			return "g18n:status_disable" // 关闭
		}
	}
	return "g18n:status_undefined"
}

var StatusExplainProtectFunc = func(value any) string {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_protect" // 保护
		} else {
			// 0 正常
			return "g18n:status_normal" // 正常
		}
	}
	return "g18n:status_undefined"
}

var StatusExplainIsNotFunc = func(value any) string {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_is" // 是
		} else {
			// 0 否
			return "g18n:status_not" // 否
		}
	}
	return "g18n:status_undefined"
}

var StatusExplainErrorFunc = func(value any) string {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			// 1 或者 trues
			return "g18n:status_error" // 异常
		} else {
			// 0 正常
			return "g18n:status_normal" // 正常
		}
	}
	return "g18n:status_undefined"
}

// BoolExplain 布尔值解释
func BoolExplain(value any, trueExplain, falseExplain string) string {
	if v, err := cvt.BoolE(value); err == nil {
		if v {
			return trueExplain
		} else {
			return falseExplain
		}
	}
	return fmt.Sprintf("未知值: %v", value)
}
