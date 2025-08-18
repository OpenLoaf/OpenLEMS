package log

import (
	"common/c_log"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

// BizLogSaver 业务日志持久化接口（后续对接数据库）
type BizLogSaver interface {
	Save(ctx context.Context, level string, message string) error
}

// BizLoggerAdapter 业务日志适配器：打印到GoFrame，并调用Saver保存
// 注意：实现 c_log.ILogger 接口
// 所有 *f 方法会先格式化，再调用非格式化方法以减少重复代码
// Saver 可为 nil，为空时不进行保存
// 仅 Debug/Info/Warning/Error 会触发保存，其它级别仅打印
type BizLoggerAdapter struct {
	logger *glog.Logger
	saver  BizLogSaver
}

// NewBizLoggerAdapter 创建业务日志适配器
func NewBizLoggerAdapter(logger *glog.Logger, saver BizLogSaver) c_log.ILogger {
	return &BizLoggerAdapter{logger: logger, saver: saver}
}

// 辅助：保存占位
func (b *BizLoggerAdapter) save(ctx context.Context, level, msg string) {
	if b.saver == nil {
		return
	}
	_ = b.saver.Save(ctx, level, msg)
}

// 仅以下四种级别触发保存
func (b *BizLoggerAdapter) Debug(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Debug(ctx, msg)
	b.save(ctx, "DEBUG", msg)
}
func (b *BizLoggerAdapter) Debugf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Debug(ctx, msg)
}
func (b *BizLoggerAdapter) Info(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Info(ctx, msg)
	b.save(ctx, "INFO", msg)
}
func (b *BizLoggerAdapter) Infof(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Info(ctx, msg)
}
func (b *BizLoggerAdapter) Warning(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Warning(ctx, msg)
	b.save(ctx, "WARNING", msg)
}
func (b *BizLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Warning(ctx, msg)
}
func (b *BizLoggerAdapter) Error(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Error(ctx, msg)
	b.save(ctx, "ERROR", msg)
}
func (b *BizLoggerAdapter) Errorf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Error(ctx, msg)
}

// 其余级别：仅打印，不保存
func (b *BizLoggerAdapter) Notice(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Notice(ctx, msg)
}
func (b *BizLoggerAdapter) Noticef(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Notice(ctx, msg)
}
func (b *BizLoggerAdapter) Critical(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Critical(ctx, msg)
}
func (b *BizLoggerAdapter) Criticalf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Critical(ctx, msg)
}
func (b *BizLoggerAdapter) Panic(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Panic(ctx, msg)
}
func (b *BizLoggerAdapter) Panicf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Panic(ctx, msg)
}
func (b *BizLoggerAdapter) Fatal(ctx context.Context, v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Fatal(ctx, msg)
}
func (b *BizLoggerAdapter) Fatalf(ctx context.Context, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.Fatal(ctx, msg)
}

// BizLogNoopSaver 默认占位实现：仅用GoFrame打印提示（不真正入库）
type BizLogNoopSaver struct{ logger *glog.Logger }

func NewBizLogNoopSaver(logger *glog.Logger) BizLogSaver { return &BizLogNoopSaver{logger: logger} }

func (s *BizLogNoopSaver) Save(ctx context.Context, level string, message string) error {
	s.logger.Debugf(ctx, "[BIZ-DB][placeholder] level=%s msg=%s", level, message)
	return nil
}
