package c_log

import (
	"context"
	"fmt"
	"log"
)

var (
	// 系统日志实例
	systemLogger ILogger = &defaultLogger{}
	// 业务日志实例（可对接数据库）
	businessLogger ILogger = &defaultLogger{}
)

// SetLogger 便捷方法：系统/业务同时设置
func SetLogger(logger ILogger) {
	if logger != nil {
		systemLogger = logger
		businessLogger = logger
	}
}

// SetSystemLogger 设置系统日志实例
func SetSystemLogger(logger ILogger) {
	if logger != nil {
		systemLogger = logger
	}
}

// SetBusinessLogger 设置业务日志实例
func SetBusinessLogger(logger ILogger) {
	if logger != nil {
		businessLogger = logger
	}
}

// Log 为兼容保留，等价于系统日志实例
func Log() ILogger { return systemLogger }

// ---- 顶层方法（系统日志）----
func Debug(ctx context.Context, v ...interface{}) { systemLogger.Debug(ctx, v...) }
func Debugf(ctx context.Context, format string, v ...interface{}) {
	systemLogger.Debugf(ctx, format, v...)
}
func Info(ctx context.Context, v ...interface{}) { systemLogger.Info(ctx, v...) }
func Infof(ctx context.Context, format string, v ...interface{}) {
	systemLogger.Infof(ctx, format, v...)
}
func Warning(ctx context.Context, v ...interface{}) { systemLogger.Warning(ctx, v...) }
func Warningf(ctx context.Context, format string, v ...interface{}) {
	systemLogger.Warningf(ctx, format, v...)
}
func Error(ctx context.Context, v ...interface{}) { systemLogger.Error(ctx, v...) }
func Errorf(ctx context.Context, format string, v ...interface{}) {
	systemLogger.Errorf(ctx, format, v...)
}

// ---- 顶层方法（业务日志，仅四种级别）----
func BizDebug(ctx context.Context, v ...interface{}) { businessLogger.Debug(ctx, v...) }
func BizDebugf(ctx context.Context, format string, v ...interface{}) {
	businessLogger.Debugf(ctx, format, v...)
}
func BizInfo(ctx context.Context, v ...interface{}) { businessLogger.Info(ctx, v...) }
func BizInfof(ctx context.Context, format string, v ...interface{}) {
	businessLogger.Infof(ctx, format, v...)
}
func BizWarning(ctx context.Context, v ...interface{}) { businessLogger.Warning(ctx, v...) }
func BizWarningf(ctx context.Context, format string, v ...interface{}) {
	businessLogger.Warningf(ctx, format, v...)
}
func BizError(ctx context.Context, v ...interface{}) { businessLogger.Error(ctx, v...) }
func BizErrorf(ctx context.Context, format string, v ...interface{}) {
	businessLogger.Errorf(ctx, format, v...)
}

// ---- 业务日志查询方法 ----
func BizQueryLogs(ctx context.Context, params LogQueryParams) (*LogQueryResult, error) {
	return businessLogger.QueryLogs(ctx, params)
}

// 默认日志实现（标准库）
type defaultLogger struct{}

func (l *defaultLogger) Debug(ctx context.Context, v ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}
func (l *defaultLogger) Info(ctx context.Context, v ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Infof(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}
func (l *defaultLogger) Notice(ctx context.Context, v ...interface{}) {
	log.Printf("[NOTICE] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Noticef(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[NOTICE] "+format, v...)
}
func (l *defaultLogger) Warning(ctx context.Context, v ...interface{}) {
	log.Printf("[WARNING] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Warningf(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[WARNING] "+format, v...)
}
func (l *defaultLogger) Error(ctx context.Context, v ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
func (l *defaultLogger) Critical(ctx context.Context, v ...interface{}) {
	log.Printf("[CRITICAL] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Criticalf(ctx context.Context, format string, v ...interface{}) {
	log.Printf("[CRITICAL] "+format, v...)
}
func (l *defaultLogger) Panic(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Printf("[PANIC] %s", msg)
	panic(msg)
}
func (l *defaultLogger) Panicf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("[PANIC] %s", msg)
	panic(msg)
}
func (l *defaultLogger) Fatal(ctx context.Context, v ...interface{}) {
	log.Fatalf("[FATAL] %s", fmt.Sprint(v...))
}
func (l *defaultLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	log.Fatalf("[FATAL] "+format, v...)
}

func (l *defaultLogger) QueryLogs(ctx context.Context, params LogQueryParams) (*LogQueryResult, error) {
	log.Printf("[WARNING] 默认日志实现不支持查询功能")
	return &LogQueryResult{Total: 0, Lines: []LogLine{}}, nil
}
