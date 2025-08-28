package c_log

import "context"

// LogLine 结构化日志行
type LogLine struct {
	Timestamp string `json:"timestamp" dc:"时间戳"`
	Id        string `json:"id"`
	Type      string `json:"type" dc:"日志类型：ems、device、protocol、policy"`
	Level     string `json:"level"     dc:"日志等级"`
	Content   string `json:"content"   dc:"日志内容"`
}

// LogQueryParams 日志查询参数
type LogQueryParams struct {
	Type     string `json:"type"`     // 业务类型(ems,device,protocol,policy,all)
	Id       string `json:"id"`       // 相关ID
	Date     string `json:"date"`     // 日期，格式：20060102
	Page     int    `json:"page"`     // 页码，从1开始
	PageSize int    `json:"pageSize"` // 每页条数
	Level    string `json:"level"`    // 日志等级(DEBUG,INFO,WARN,ERROR,ALL)
}

// LogQueryResult 日志查询结果
type LogQueryResult struct {
	Total int       `json:"total"` // 总行数
	Lines []LogLine `json:"lines"` // 日志行列表(倒序)
}

// ILogger 日志接口，用于代理GoFrame的g.Log()功能
type ILogger interface {
	// Debug 调试级别日志
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})

	// Info 信息级别日志
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})

	// Warning 警告级别日志
	Warning(ctx context.Context, v ...interface{})
	Warningf(ctx context.Context, format string, v ...interface{})

	// Error 错误级别日志
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})

	// QueryLogs 查询日志
	QueryLogs(ctx context.Context, params LogQueryParams) (*LogQueryResult, error)
}
