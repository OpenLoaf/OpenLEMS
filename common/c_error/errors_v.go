package c_error

import (
	"github.com/pkg/errors"
)

var (
	NoData          = errors.New("数据不存在")
	NonSupportError = errors.New("不支持的操作") // 不支持的操作不会显示到页面上
	ErrorParam      = errors.New("参数错误")
	OverLimitError  = errors.New("数值超过限制")
	TypeError       = errors.New("类型不匹配")
)
