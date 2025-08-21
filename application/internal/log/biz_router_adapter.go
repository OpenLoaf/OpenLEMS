package log

import (
	"common/c_base"
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

// BizRouterLoggerAdapter 业务日志适配器：按上下文中的ID进行分类落盘
// 优先级：DeviceId > ProtocolId > PolicyId > EMS
type BizRouterLoggerAdapter struct{}

func NewBizRouterLoggerAdapter() c_log.ILogger { return &BizRouterLoggerAdapter{} }

func (b *BizRouterLoggerAdapter) pick(ctx context.Context) *glog.Logger {
	if ctx == nil {
		return BizEMS()
	}
	if v := ctx.Value(c_base.ConstCtxKeyDeviceId); v != nil {
		if s, ok := v.(string); ok && s != "" {
			return BizDevice(s)
		}
	}
	if v := ctx.Value(c_base.ConstCtxKeyProtocolId); v != nil {
		if s, ok := v.(string); ok && s != "" {
			return BizProtocol(s)
		}
	}
	// PolicyId 未在 c_base 中定义上下文常量，这里用约定键名做兜底
	if v := ctx.Value("PolicyId"); v != nil {
		if s, ok := v.(string); ok && s != "" {
			return BizPolicy(s)
		}
	}
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
