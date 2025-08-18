package c_log

import "context"

// ExampleUsage 展示如何在common包中使用日志代理
func ExampleUsage(ctx context.Context) {
	// 使用方式与g.Log()完全一样
	Log().Infof(ctx, "这是一条信息日志: %s", "示例")
	Log().Debugf(ctx, "这是一条调试日志: %d", 123)
	Log().Warningf(ctx, "这是一条警告日志")
	Log().Errorf(ctx, "这是一条错误日志: %v", "错误信息")
}

// GetLogger 提供一个便捷的函数来获取logger实例
// 可以在common包的其他地方使用
func GetLogger() ILogger {
	return Log()
}
