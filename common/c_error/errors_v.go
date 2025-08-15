package c_error

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	NoData          = gerror.New("数据不存在")
	NonSupportError = gerror.New("不支持的操作")
	ErrorParam      = gerror.New("参数错误")

	OverLimitError = gerror.New("数值超过限制")
)
