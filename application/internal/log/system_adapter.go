package log

import (
	"common/c_log"
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/os/glog"
)

// GoFrameLoggerAdapter 系统日志适配器：直接把调用转发给GoFrame
type GoFrameLoggerAdapter struct {
	logger *glog.Logger
	mu     sync.RWMutex // 保护并发访问
}

func NewSystemAdapter(logger *glog.Logger) c_log.ILogger {
	fmt.Printf("日志等级：%v\n", logger.GetLevel())
	return &GoFrameLoggerAdapter{logger: logger}
}

// Debug 调试级别日志
func (g *GoFrameLoggerAdapter) Debug(ctx context.Context, v ...interface{}) {
	if g.logger != nil {
		g.logger.Debug(ctx, v...)
	}
}

// Debugf 调试级别格式化日志
func (g *GoFrameLoggerAdapter) Debugf(ctx context.Context, format string, v ...interface{}) {
	if g.logger != nil {
		g.logger.Debugf(ctx, format, v...)
	}
}

// Info 信息级别日志
func (g *GoFrameLoggerAdapter) Info(ctx context.Context, v ...interface{}) {
	if g.logger != nil {
		g.logger.Info(ctx, v...)
	}
}

// Infof 信息级别格式化日志
func (g *GoFrameLoggerAdapter) Infof(ctx context.Context, format string, v ...interface{}) {
	if g.logger != nil {
		g.logger.Infof(ctx, format, v...)
	}
}

// Warning 警告级别日志
func (g *GoFrameLoggerAdapter) Warning(ctx context.Context, v ...interface{}) {
	if g.logger != nil {
		g.logger.Warning(ctx, v...)
	}
}

// Warningf 警告级别格式化日志
func (g *GoFrameLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	if g.logger != nil {
		g.logger.Warningf(ctx, format, v...)
	}
}

// Error 错误级别日志
func (g *GoFrameLoggerAdapter) Error(ctx context.Context, v ...interface{}) {
	if g.logger != nil {
		g.logger.Error(ctx, v...)
	}
}

// Errorf 错误级别格式化日志
func (g *GoFrameLoggerAdapter) Errorf(ctx context.Context, format string, v ...interface{}) {
	if g.logger != nil {
		g.logger.Errorf(ctx, format, v...)
	}
}

// QueryLogs 查询日志（系统适配器不支持查询功能）
func (g *GoFrameLoggerAdapter) QueryLogs(ctx context.Context, params c_log.LogQueryParams) (*c_log.LogQueryResult, error) {
	if g.logger != nil {
		g.logger.Warningf(ctx, "GoFrame日志适配器不支持查询功能")
	}
	return &c_log.LogQueryResult{Total: 0, Lines: []c_log.LogLine{}}, nil
}
