package c_base

import "common/c_util"

var IsNotZero = func(value any) bool {
	return value.(int) != 0
}

var IsZero = func(value any) bool {
	return value.(int) == 0
}

var StatusExplainEnableFunc = func(value any) string {
	if v, err := c_util.ToBool(value); err == nil {
		if v {
			// 1 或者 trues
			return "status_enable" // 启动
		} else {
			// 0 关闭
			return "status_disable" // 关闭
		}
	}
	return "status_undefined"
}

var StatusExplainProtectFunc = func(value any) string {
	if v, err := c_util.ToBool(value); err == nil {
		if v {
			// 1 或者 trues
			return "status_protect" // 保护
		} else {
			// 0 正常
			return "status_normal" // 正常
		}
	}
	return "status_undefined"
}

var StatusExplainIsNotFunc = func(value any) string {
	if v, err := c_util.ToBool(value); err == nil {
		if v {
			// 1 或者 trues
			return "status_is" // 是
		} else {
			// 0 否
			return "status_not" // 否
		}
	}
	return "status_undefined"
}

var StatusExplainErrorFunc = func(value any) string {
	if v, err := c_util.ToBool(value); err == nil {
		if v {
			// 1 或者 trues
			return "status_error" // 异常
		} else {
			// 0 正常
			return "status_normal" // 正常
		}
	}
	return "status_undefined"
}
