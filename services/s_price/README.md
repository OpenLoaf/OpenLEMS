# 电价管理服务 (s_price)

## 概述

电价管理服务是EMS系统的核心组件之一，负责管理本地和远程电价配置，提供电价查询、激活判断、缓存管理和定时存储等功能。

## 核心功能

### 1. 电价配置管理
- 支持本地电价和远程电价配置
- 支持多种电价类型：谷电、峰电、平电、尖峰、深谷
- 支持分钟级精度的时段配置
- 支持日期范围和时间范围配置

### 2. 内存缓存机制
- 启动时加载所有电价到内存
- 增删改操作后自动刷新缓存
- 提供快速查询接口

### 3. 激活判断逻辑
- 根据当前时间、日期范围、时间范围判断激活电价
- 按优先级排序，返回优先级最高的激活电价
- 支持工作日/周末、自定义日期、月度等时间范围类型

### 4. 定时存储
- 每小时自动保存当前电价到Storage
- 支持电价历史记录查询
- 使用专门的 `SavePriceData` 方法保存电价数据
- 存储格式：priceId, price, priceType, timestamp

## 架构设计

### 服务层结构
```
services/s_price/
├── internal/                    # 内部实现
│   ├── s_price_types_s.go      # 类型定义
│   ├── s_price_cache_s.go      # 内存缓存实现
│   ├── s_price_active_f.go     # 激活判断逻辑
│   ├── s_price_storage_f.go    # Storage保存逻辑
│   └── s_price_manager_impl_s.go # 电价管理器实现
├── s_price_export_f.go         # 服务导出
└── go.mod                      # 模块依赖
```

### 核心接口
- `IPriceManager` - 电价管理器接口
- `GetCurrentPrice()` - 获取当前激活电价
- `RefreshCache()` - 刷新缓存

## 数据模型

### 电价时段结构 (SPriceSegment)
```go
type SPriceSegment struct {
    StartTime string           // 开始时间 "HH:MM" 格式
    EndTime   string           // 结束时间 "HH:MM" 格式
    PriceType EPriceType       // 电价类型
    Price     float64          // 电价值
}
```

### 日期范围配置 (SDateRange)
```go
type SDateRange struct {
    StartDate  string // 开始日期 "YYYY-MM-DD" 格式
    EndDate    string // 结束日期 "YYYY-MM-DD" 格式
    IsLongTerm bool   // 是否长期有效
}
```

### 时间范围配置 (STimeRange)
```go
type STimeRange struct {
    Type         string // 时间范围类型：weekday/custom/monthly
    WeekdayType  string // 工作日类型：workday/weekend/all
    CustomDays   []int  // 自定义日期
    CustomMonths []int  // 自定义月份
}
```

## 电价类型

支持以下5种电价类型：

| 类型 | 值 | 描述 |
|------|----|----|
| 谷电 | valley | 低峰时段电价 |
| 峰电 | peak | 高峰时段电价 |
| 平电 | flat | 平时段电价 |
| 尖峰 | sharp | 尖峰时段电价 |
| 深谷 | deep_valley | 深谷时段电价 |

## 时间范围类型

### 1. 工作日类型 (weekday)
- `workday` - 工作日（周一到周五）
- `weekend` - 周末（周六到周日）
- `all` - 全部时间

### 2. 自定义类型 (custom)
- `CustomDays` - 自定义日期（1-31）
- `CustomMonths` - 自定义月份（1-12）

### 3. 月度类型 (monthly)
- 每月1日生效

## 激活判断算法

1. **过滤条件**：status = Enable
2. **时间匹配**：dateRange 和 timeRange 匹配当前时间
3. **排序规则**：按 priority 升序（数值越小优先级越高）
4. **返回结果**：取第一条匹配的电价

## 使用示例

### 获取当前激活电价
```go
// 通过Common层获取
currentPrice, err := common.GetPriceManager().GetCurrentPrice(ctx)

// 通过服务层获取
activePrice, err := s_price.GetCurrentActivePrice(ctx)
```

### 获取当前电价时段
```go
segment, err := s_price.GetCurrentPriceSegment(ctx)
if segment != nil {
    fmt.Printf("当前电价: %.4f, 类型: %s", segment.Price, segment.PriceType)
}
```

### 刷新缓存
```go
err := s_price.RefreshPriceCache(ctx)
```

## 配置示例

### 工作日峰谷电价配置
```json
{
  "description": "工作日峰谷电价",
  "priority": 1,
  "status": "Enable",
  "dateRange": {
    "isLongTerm": true,
    "startDate": "2024-01-01",
    "endDate": null
  },
  "timeRange": {
    "type": "weekday",
    "weekdayType": "workday"
  },
  "priceSegments": [
    {
      "startTime": "00:00",
      "endTime": "08:00",
      "priceType": "valley",
      "price": 0.3
    },
    {
      "startTime": "08:00",
      "endTime": "12:00",
      "priceType": "peak",
      "price": 0.8
    },
    {
      "startTime": "12:00",
      "endTime": "18:00",
      "priceType": "flat",
      "price": 0.5
    },
    {
      "startTime": "18:00",
      "endTime": "22:00",
      "priceType": "peak",
      "price": 0.8
    },
    {
      "startTime": "22:00",
      "endTime": "24:00",
      "priceType": "valley",
      "price": 0.3
    }
  ]
}
```

## 远程电价支持

### 远程电价标识
- `remote_id` 为 NULL：本地电价
- `remote_id` 不为 NULL：远程电价

### MQTT下发电价流程（预留）
1. MQTT接收到电价配置消息
2. 解析消息中的 `remote_id`
3. 查询数据库：WHERE remote_id = ? AND status != 'Deleted'
4. 如果存在：更新电价配置
5. 如果不存在：创建新电价（设置remote_id）
6. 刷新内存缓存

## 性能优化

### 内存缓存
- 所有电价数据加载到内存，避免频繁数据库查询
- 使用读写锁保护并发访问
- 增删改操作后自动刷新缓存

### 定时存储
- 每小时保存一次当前电价，避免频繁存储操作
- 异步执行，不阻塞主流程

## 错误处理

### 常见错误
- 电价配置格式错误
- 时间范围配置无效
- 缓存刷新失败
- Storage保存失败

### 错误恢复
- 配置错误时记录日志但不中断服务
- 缓存刷新失败时使用旧缓存数据
- Storage保存失败时记录错误日志

## 监控和日志

### 日志记录
- 使用 `c_log` 模块记录关键操作
- 记录缓存加载、激活判断、存储保存等操作
- 错误日志包含详细上下文信息

### 性能监控
- 缓存加载时间
- 激活判断耗时
- Storage保存成功率

## 扩展性

### 新增电价类型
1. 在 `common/c_enum/c_enum_price_type_e.go` 中添加新类型
2. 更新相关验证逻辑
3. 更新前端显示逻辑

### 新增时间范围类型
1. 在 `STimeRange` 中添加新字段
2. 在 `isActiveAtTime` 函数中添加判断逻辑
3. 更新API文档

## 注意事项

1. **时间格式**：使用 "HH:MM" 格式，如 "08:30"
2. **优先级**：数值越小优先级越高
3. **缓存一致性**：增删改操作后必须刷新缓存
4. **线程安全**：使用读写锁保护并发访问
5. **错误处理**：配置错误不应中断服务运行
