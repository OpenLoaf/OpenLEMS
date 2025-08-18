package c_error

import (
	"fmt"
)

var (
	NoData          = fmt.Errorf("数据不存在")
	NonSupportError = fmt.Errorf("不支持的操作")
	ErrorParam      = fmt.Errorf("参数错误")

	OverLimitError = fmt.Errorf("数值超过限制")
)
