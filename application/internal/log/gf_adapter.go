package log

import (
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

// GoFrameLoggerAdapter 系统日志适配器：直接把调用转发给GoFrame
type GoFrameLoggerAdapter struct{ logger *glog.Logger }

func NewGoFrameLoggerAdapter(logger *glog.Logger) c_log.ILogger {
	return &GoFrameLoggerAdapter{logger: logger}
}

func (g *GoFrameLoggerAdapter) Debug(ctx context.Context, v ...interface{}) {
	g.logger.Debug(ctx, v...)
}
func (g *GoFrameLoggerAdapter) Debugf(ctx context.Context, format string, v ...interface{}) {
	g.logger.Debugf(ctx, format, v...)
}
func (g *GoFrameLoggerAdapter) Info(ctx context.Context, v ...interface{}) { g.logger.Info(ctx, v...) }
func (g *GoFrameLoggerAdapter) Infof(ctx context.Context, format string, v ...interface{}) {
	g.logger.Infof(ctx, format, v...)
}

func (g *GoFrameLoggerAdapter) Warning(ctx context.Context, v ...interface{}) {
	g.logger.Warning(ctx, v...)
}
func (g *GoFrameLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	g.logger.Warningf(ctx, format, v...)
}
func (g *GoFrameLoggerAdapter) Error(ctx context.Context, v ...interface{}) {
	g.logger.Error(ctx, v...)
}
func (g *GoFrameLoggerAdapter) Errorf(ctx context.Context, format string, v ...interface{}) {
	g.logger.Errorf(ctx, format, v...)
}
