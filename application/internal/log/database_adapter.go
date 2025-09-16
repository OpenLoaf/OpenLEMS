package log

import (
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"fmt"
	"s_db"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

// DbLoggerAdapter 数据库日志适配器：将日志保存到数据库中
type DbLoggerAdapter struct {
	logService   s_db_basic.ILogService
	stdout       bool
	stdoutLogger c_log.ILogger // 标准输出日志适配器
	mu           sync.RWMutex  // 保护并发访问
}

// NewDatabaseAdapter 创建数据库日志适配器
func NewDatabaseAdapter() c_log.ILogger {
	stdout, _ := g.Cfg().Get(context.Background(), "logger.biz_ems_stdout", false)
	fmt.Printf("===>  new logger adapter stdout: %s\n", stdout.String())

	adapter := &DbLoggerAdapter{
		logService: s_db.GetLogService(),
		stdout:     stdout.Bool(),
	}

	// 如果启用标准输出，创建GoFrame日志适配器
	if adapter.stdout {
		adapter.stdoutLogger = NewSystemAdapter(g.Log())
	}

	return adapter
}

// getLogInfo 从上下文中提取日志信息
func (d *DbLoggerAdapter) getLogInfo(ctx context.Context) (logType, deviceId string) {
	if ctx == nil {
		return "ems", ""
	}

	// 定义上下文键和对应的日志类型映射
	type contextMapping struct {
		key     interface{}
		logType string
	}

	mappings := []contextMapping{
		{c_base.ConstCtxKeyDeviceId, c_enum.ELogTypeDevice.String()},
		{c_base.ConstCtxKeyProtocolId, c_enum.ELogTypeProtocol.String()},
		{c_base.ConstCtxKeyPolicyId, c_enum.ELogTypePolicy.String()},
	}

	// 按优先级检查上下文键
	for _, mapping := range mappings {
		if v := ctx.Value(mapping.key); v != nil {
			if s, ok := v.(string); ok && s != "" {
				return mapping.logType, s
			}
		}
	}

	return c_enum.ELogTypeEms.String(), ""
}

// saveToDb 保存日志到数据库
func (d *DbLoggerAdapter) saveToDb(ctx context.Context, level, content string) {
	// 参数验证
	if content == "" {
		return
	}

	logType, deviceId := d.getLogInfo(ctx)

	// 如果启用标准输出，先输出到标准输出
	if d.stdout && d.stdoutLogger != nil {
		d.outputToStdout(ctx, level, content)
	}

	// 异步保存到数据库，避免阻塞主流程
	go func() {
		// 创建新的上下文用于数据库操作
		dbCtx := context.Background()
		err := d.logService.CreateLog(dbCtx, logType, deviceId, level, content)
		if err != nil {
			// 如果数据库保存失败，记录错误但不阻塞主流程
			fmt.Printf("保存日志到数据库失败 [类型:%s, ID:%s, 级别:%s]: %v\n",
				logType, deviceId, level, err)
		}
	}()
}

// outputToStdout 输出日志到标准输出
func (d *DbLoggerAdapter) outputToStdout(ctx context.Context, level, content string) {
	if d.stdoutLogger == nil {
		return
	}

	// 使用互斥锁保护并发访问
	d.mu.RLock()
	defer d.mu.RUnlock()
	content = "BIZ ====> " + content
	// 根据日志级别调用对应的输出方法
	switch level {
	case "DEBUG":
		d.stdoutLogger.Debug(ctx, content)
	case "INFO":
		d.stdoutLogger.Info(ctx, content)
	case "WARN":
		d.stdoutLogger.Warning(ctx, content)
	case "ERROR":
		d.stdoutLogger.Error(ctx, content)
	default:
		d.stdoutLogger.Info(ctx, content)
	}
}

// Debug 调试级别日志
func (d *DbLoggerAdapter) Debug(ctx context.Context, v ...interface{}) {
	content := fmt.Sprint(v...)
	d.saveToDb(ctx, "DEBUG", content)
}

// Debugf 调试级别格式化日志
func (d *DbLoggerAdapter) Debugf(ctx context.Context, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	d.saveToDb(ctx, "DEBUG", content)
}

// Info 信息级别日志
func (d *DbLoggerAdapter) Info(ctx context.Context, v ...interface{}) {
	content := fmt.Sprint(v...)
	d.saveToDb(ctx, "INFO", content)
}

// Infof 信息级别格式化日志
func (d *DbLoggerAdapter) Infof(ctx context.Context, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	d.saveToDb(ctx, "INFO", content)
}

// Warning 警告级别日志
func (d *DbLoggerAdapter) Warning(ctx context.Context, v ...interface{}) {
	content := fmt.Sprint(v...)
	d.saveToDb(ctx, "WARN", content)
}

// Warningf 警告级别格式化日志
func (d *DbLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	d.saveToDb(ctx, "WARN", content)
}

// Error 错误级别日志
func (d *DbLoggerAdapter) Error(ctx context.Context, v ...interface{}) {
	content := fmt.Sprint(v...)
	d.saveToDb(ctx, "ERROR", content)
}

// Errorf 错误级别格式化日志
func (d *DbLoggerAdapter) Errorf(ctx context.Context, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	d.saveToDb(ctx, "ERROR", content)
}

// QueryLogs 查询数据库日志
func (d *DbLoggerAdapter) QueryLogs(ctx context.Context, params c_log.LogQueryParams) (*c_log.LogQueryResult, error) {
	// 参数验证
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}
	if params.PageSize > 1000 {
		params.PageSize = 1000 // 限制最大页面大小
	}

	// 构建查询过滤条件
	filters := d.buildQueryFilters(params)

	// 调用数据库服务查询
	logs, total, err := d.logService.GetLogPage(ctx, params.Page, params.PageSize, filters)
	if err != nil {
		return nil, fmt.Errorf("查询数据库日志失败: %w", err)
	}

	// 转换为统一的日志行格式
	lines := d.convertToLogLines(logs)

	return &c_log.LogQueryResult{
		Total: total,
		Lines: lines,
	}, nil
}

// buildQueryFilters 构建查询过滤条件
func (d *DbLoggerAdapter) buildQueryFilters(params c_log.LogQueryParams) map[string]interface{} {
	filters := make(map[string]interface{})

	// 日期过滤
	if params.Date != "" {
		filters["date"] = params.Date
	}

	// 类型过滤
	if params.Type != "" && params.Type != "all" {
		filters["type"] = params.Type
	}

	// 设备ID过滤
	if params.Id != "" {
		filters["device_id"] = params.Id
	}

	// 级别过滤
	if params.Level != "" && params.Level != "ALL" {
		filters["level"] = params.Level
	}

	return filters
}

// convertToLogLines 转换日志记录为统一格式
func (d *DbLoggerAdapter) convertToLogLines(logs []*s_db_model.SLogModel) []c_log.LogLine {
	lines := make([]c_log.LogLine, 0, len(logs))
	for _, log := range logs {
		lines = append(lines, c_log.LogLine{
			Id:        log.DeviceId,
			Type:      log.Type,
			Level:     log.Level,
			Content:   log.Content,
			CreatedAt: log.CreatedAt,
		})
	}
	return lines
}
