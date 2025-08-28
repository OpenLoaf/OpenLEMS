package log

import (
	"common/c_base"
	"common/c_log"
	"context"
	"fmt"
	"s_db"
	"s_db/s_db_basic"
)

// DbLoggerAdapter 数据库日志适配器：将日志保存到数据库中
type DbLoggerAdapter struct {
	logService s_db_basic.ILogService
}

// NewDbLoggerAdapter 创建数据库日志适配器
func NewDatabaseAdapter() c_log.ILogger {
	return &DbLoggerAdapter{
		logService: s_db.GetLogService(),
	}
}

// getLogInfo 从上下文中提取日志信息
func (d *DbLoggerAdapter) getLogInfo(ctx context.Context) (logType, deviceId string) {
	logType = "ems" // 默认类型
	deviceId = ""   // 默认设备ID

	if ctx != nil {
		// 优先检查设备ID
		if v := ctx.Value(c_base.ConstCtxKeyDeviceId); v != nil {
			if s, ok := v.(string); ok && s != "" {
				logType = "device"
				deviceId = s
				return
			}
		}

		// 检查协议ID
		if v := ctx.Value(c_base.ConstCtxKeyProtocolId); v != nil {
			if s, ok := v.(string); ok && s != "" {
				logType = "protocol"
				deviceId = s
				return
			}
		}

		// 检查策略ID
		if v := ctx.Value("PolicyId"); v != nil {
			if s, ok := v.(string); ok && s != "" {
				logType = "policy"
				deviceId = s
				return
			}
		}
	}

	return
}

// saveToDb 保存日志到数据库
func (d *DbLoggerAdapter) saveToDb(ctx context.Context, level, content string) {
	logType, deviceId := d.getLogInfo(ctx)

	// 异步保存到数据库，避免阻塞主流程
	go func() {
		// 创建新的上下文用于数据库操作
		dbCtx := context.Background()
		err := d.logService.CreateLog(dbCtx, logType, deviceId, level, content)
		if err != nil {
			// 如果数据库保存失败，这里可以选择记录到系统日志
			// 但为了避免循环调用，我们暂时不处理
			fmt.Printf("保存日志到数据库失败: %v\n", err)
		}
	}()
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
	d.saveToDb(ctx, "WARNING", content)
}

// Warningf 警告级别格式化日志
func (d *DbLoggerAdapter) Warningf(ctx context.Context, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	d.saveToDb(ctx, "WARNING", content)
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
	// 构建查询过滤条件
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

	// 调用数据库服务查询
	logs, total, err := d.logService.GetLogPage(ctx, params.Page, params.PageSize, filters)
	if err != nil {
		return nil, fmt.Errorf("查询数据库日志失败: %w", err)
	}

	// 转换为统一的日志行格式
	lines := make([]c_log.LogLine, 0, len(logs))
	for _, log := range logs {
		lines = append(lines, c_log.LogLine{
			Timestamp: log.CreatedAt,
			Id:        log.DeviceId,
			Type:      log.Type,
			Level:     log.Level,
			Content:   log.Content,
		})
	}

	return &c_log.LogQueryResult{
		Total: total,
		Lines: lines,
	}, nil
}
