# 日志系统说明

## 概述

本日志系统支持多种日志输出方式，包括文件日志和数据库日志。

## 文件结构

```
application/internal/log/
├── system_adapter.go      # 系统日志适配器
├── file_adapter.go        # 文件日志适配器
├── database_adapter.go    # 数据库日志适配器
├── file_helpers.go        # 文件日志帮助函数
├── adapters_test.go       # 统一测试文件
└── README.md             # 文档说明
```

## 日志适配器

### 1. SystemAdapter (system_adapter.go)
- **功能**: 系统日志适配器，直接转发给GoFrame的g.Log()
- **用途**: 系统级别的日志输出
- **创建函数**: `NewSystemAdapter(logger)`

### 2. FileAdapter (file_adapter.go)
- **功能**: 文件日志适配器，输出到EMS单文件
- **特点**: 支持JSON格式，包含type/id分类字段
- **用途**: 业务级别的日志输出到文件
- **创建函数**: `NewFileAdapter()`

### 3. DatabaseAdapter (database_adapter.go)
- **功能**: 数据库日志适配器，将日志保存到数据库中
- **特点**: 
  - 异步保存，不阻塞主流程
  - 支持从上下文提取设备ID、协议ID、策略ID
  - 自动分类日志类型（ems/device/protocol/policy）
- **用途**: 业务级别的日志持久化存储
- **创建函数**: `NewDatabaseAdapter()`



## 日志级别

支持以下日志级别：
- `DEBUG`: 调试信息
- `INFO`: 一般信息
- `WARNING`: 警告信息  
- `ERROR`: 错误信息

## 上下文支持

日志系统支持从上下文中提取以下信息：
- `c_base.ConstCtxKeyDeviceId`: 设备ID
- `c_base.ConstCtxKeyProtocolId`: 协议ID
- `PolicyId`: 策略ID

## 数据库日志表结构

日志保存在`log`表中，包含以下字段：
- `id`: 主键
- `type`: 日志类型（ems/device/protocol/policy）
- `device_id`: 设备ID
- `level`: 日志级别
- `content`: 日志内容
- `created_at`: 创建时间

## 使用示例

```go
// 在cmd.go中注册数据库日志适配器
c_log.SetSystemLogger(applog.NewSystemAdapter(g.Log()))
c_log.SetBusinessLogger(applog.NewDatabaseAdapter())

// 使用业务日志
ctx := context.Background()
c_log.BizInfo(ctx, "这是一条业务日志")

// 带设备ID的日志
deviceCtx := context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, "device_001")
c_log.BizInfo(deviceCtx, "设备操作日志")

// 查询业务日志
params := c_log.LogQueryParams{
    Type:     "device",
    Id:       "device_001",
    Date:     "20241201",
    Page:     1,
    PageSize: 100,
    Level:    "INFO",
}
result, err := c_log.BizQueryLogs(ctx, params)
if err != nil {
    // 处理错误
}
// 使用查询结果
fmt.Printf("总记录数: %d\n", result.Total)
for _, line := range result.Lines {
    fmt.Printf("时间: %s, 级别: %s, 内容: %s\n", line.Timestamp, line.Level, line.Content)
}
```

## 配置说明

在`cmd.go`中，系统会自动注册适配器：

```go
// 注入系统日志（GoFrame）
c_log.SetSystemLogger(applog.NewSystemAdapter(g.Log()))
// 注入业务日志（输出到数据库）
c_log.SetBusinessLogger(applog.NewDatabaseAdapter())
```

这样配置后，所有的业务日志都会保存到数据库中（结构化存储，便于查询和管理）。

## 查询功能

### 查询接口
所有日志适配器都实现了统一的查询接口：
```go
QueryLogs(ctx context.Context, params LogQueryParams) (*LogQueryResult, error)
```

### 查询参数
- `Type`: 业务类型（ems/device/protocol/policy/all）
- `Id`: 相关ID（设备ID、协议ID、策略ID）
- `Date`: 日期，格式：20060102
- `Page`: 页码，从1开始
- `PageSize`: 每页条数
- `Level`: 日志等级（DEBUG/INFO/WARN/ERROR/ALL）

### 适配器查询实现
1. **DatabaseAdapter**: 从数据库查询日志
2. **FileAdapter**: 从文件查询日志
3. **SystemAdapter**: 返回空结果，提示不支持查询
4. **defaultLogger**: 返回空结果，提示不支持查询

### API接口
通过`/log/biz`接口可以查询业务日志，系统会根据当前注册的业务日志适配器自动选择查询方式。
