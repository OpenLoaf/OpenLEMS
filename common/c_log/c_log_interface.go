package c_log

import "context"

// ILogger 日志接口，用于代理GoFrame的g.Log()功能
type ILogger interface {
	// Debug 调试级别日志
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})

	// Info 信息级别日志
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})

	// Notice 通知级别日志
	Notice(ctx context.Context, v ...interface{})
	Noticef(ctx context.Context, format string, v ...interface{})

	// Warning 警告级别日志
	Warning(ctx context.Context, v ...interface{})
	Warningf(ctx context.Context, format string, v ...interface{})

	// Error 错误级别日志
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})

	// Critical 严重级别日志
	Critical(ctx context.Context, v ...interface{})
	Criticalf(ctx context.Context, format string, v ...interface{})

	// Panic panic级别日志
	Panic(ctx context.Context, v ...interface{})
	Panicf(ctx context.Context, format string, v ...interface{})

	// Fatal 致命级别日志
	Fatal(ctx context.Context, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
}
