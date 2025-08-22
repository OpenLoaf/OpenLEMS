package log

import (
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

// BizRouterLoggerAdapter 业务日志适配器：统一写入EMS单文件
// 分类信息通过 JSON 字段(type/id)保留

type BizRouterLoggerAdapter struct{}

func NewBizRouterLoggerAdapter() c_log.ILogger { return &BizRouterLoggerAdapter{} }

func (b *BizRouterLoggerAdapter) pick(ctx context.Context) *glog.Logger {
	return BizEMS()
}

func (b *BizRouterLoggerAdapter) Debug(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Debug(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Debugf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Debugf(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Info(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Info(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Infof(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Infof(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Notice(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Notice(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Noticef(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Noticef(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Warning(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Warning(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Warningf(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Error(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Error(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Errorf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Errorf(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Critical(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Critical(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Criticalf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Criticalf(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Panic(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Panic(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Panicf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Panicf(ctx, format, v...)
}
func (b *BizRouterLoggerAdapter) Fatal(ctx context.Context, v ...interface{}) {
	b.pick(ctx).Fatal(ctx, v...)
}
func (b *BizRouterLoggerAdapter) Fatalf(ctx context.Context, format string, v ...interface{}) {
	b.pick(ctx).Fatalf(ctx, format, v...)
}
