# 自动化服务 (s_automation)

自动化服务提供任务调度和执行功能，支持基于时间范围和触发条件的自动化任务管理。

## 功能特性

- **任务管理**: 支持添加、删除、更新、启用、停用自动化任务
- **定时执行**: 可自定义间隔检查并执行符合条件的自动化任务
- **时间范围控制**: 支持开始时间、结束时间和复杂的时间范围规则
- **触发规则**: 支持基于设备遥测数据的多条件触发（AND 逻辑）
- **执行规则**: 支持调用多个设备服务执行具体操作
- **性能优化**: 预解析 JSON 规则，避免重复解析提升执行性能
- **内存缓存**: 使用内存缓存提高查询性能
- **线程安全**: 使用读写锁保证并发安全

## 使用方法

### 1. 初始化服务

```go
import "s_automation"

// 初始化自动化服务
s_automation.Init()
```

### 2. 启动自动化管理器

```go
ctx := context.Background()
// 每秒执行一次（默认）
err := s_automation.StartAutomationManager(ctx, time.Second)
if err != nil {
    log.Fatalf("启动自动化管理器失败: %v", err)
}

// 或者自定义执行间隔
// 每5秒执行一次
err = s_automation.StartAutomationManager(ctx, 5*time.Second)

// 每30秒执行一次
err = s_automation.StartAutomationManager(ctx, 30*time.Second)
```

### 3. 添加自动化任务

```go
import (
    "s_db/s_db_model"
    "github.com/gogf/gf/v2/os/gtime"
    "time"
)

// 创建自动化任务
automation := &s_db_model.SAutomationModel{
    StartTime:      gtime.NewFromTime(time.Now().Add(1 * time.Hour)),
    EndTime:        gtime.NewFromTime(time.Now().Add(2 * time.Hour)),
    TimeRangeType:  "daily",
    TimeRangeValue: "09:00-18:00",
    TriggerRule:    `{"anyMatch":[{"deviceId":"device001","rule":"P>30"}],"subMatch":[{"deviceId":"device002","rule":"Ia<100"}],"subMatchAll":true}`,
    ExecuteRule:    `{"deviceId":"device001","action":"setTemperature","value":26.0}`,
    Enable:         true,
}

err := s_automation.AddAutomation(ctx, automation)
if err != nil {
    log.Printf("添加自动化任务失败: %v", err)
}
```

### 4. 管理自动化任务

```go
import "s_db"

// 获取所有自动化任务
automations, err := s_db.GetAutomationService().GetAllAutomations(ctx)

// 获取启用的自动化任务
enabledAutomations, err := s_db.GetAutomationService().GetEnabledAutomations(ctx)

// 启用任务
err = s_automation.EnableAutomation(ctx, taskId)

// 停用任务
err = s_automation.DisableAutomation(ctx, taskId)

// 更新任务
updateData := map[string]interface{}{
    "enable": false,
    "triggerRule": `{"anyMatch":[{"deviceId":"device002","rule":"H>60"}]}`,
}
err = s_automation.UpdateAutomation(ctx, taskId, updateData)

// 删除任务
err = s_automation.RemoveAutomation(ctx, taskId)
```

### 5. 停止服务

```go
err := s_automation.StopAutomationManager(ctx)
if err != nil {
    log.Printf("停止自动化管理器失败: %v", err)
}
```

## 规则格式

### 触发规则 (TriggerRule)

触发规则使用 JSON 对象格式定义，支持复杂的条件组合：

```json
{
    "anyMatch": [
        {
            "deviceId": "设备ID",
            "rule": "规则表达式"
        }
    ],
    "subMatch": [
        {
            "deviceId": "设备ID", 
            "rule": "规则表达式"
        }
    ],
    "subMatchAll": true
}
```

示例：

```json
{
    "anyMatch": [
        {
            "deviceId": "device001",
            "rule": "P>30"
        },
        {
            "deviceId": "device002", 
            "rule": "Ia<100"
        }
    ],
    "subMatch": [
        {
            "deviceId": "device001",
            "rule": "P>30"
        },
        {
            "deviceId": "device002",
            "rule": "Ia<100"
        }
    ],
    "subMatchAll": true
}
```

**字段说明**：
- `anyMatch`: 任意匹配条件（OR 逻辑），任意一个满足即可
- `subMatch`: 子匹配条件，根据 `subMatchAll` 决定逻辑
- `subMatchAll`: `true` 表示全部满足（AND 逻辑），`false` 表示任意满足（OR 逻辑）
- `rule`: 规则表达式，如 "P>30"、"Ia<100"、"T==25.5" 等

### 执行规则 (ExecuteRule)

执行规则使用 JSON 字符串格式定义，具体格式待后续实现：

```json
"执行规则的JSON字符串"
```

示例：

```json
"{\"deviceId\":\"controller001\",\"action\":\"setTemperature\",\"value\":26.0}"
```

## 时间范围类型

- `daily`: 每日时间范围
- `weekly`: 每周时间范围
- `monthly`: 每月时间范围
- `custom`: 自定义时间范围

## 性能优化

### 预解析机制

- **规则预解析**: 在任务加载时预解析 JSON 规则，避免执行时重复解析
- **内存缓存**: 使用 `SAutomationTask` 结构体缓存预解析的规则
- **执行优化**: 执行时直接使用预解析的结构体，大幅提升性能

### 性能对比

- **优化前**: 每次执行都需要 JSON 解析，CPU 开销大
- **优化后**: 规则只在加载/更新时解析一次，执行时零解析开销

## 注意事项

1. **数据库依赖**: 需要先初始化 `s_db` 服务
2. **设备集成**: 触发规则和执行规则需要与设备管理器集成
3. **性能考虑**: 大量任务时建议合理设置执行间隔
4. **错误处理**: 所有操作都会记录详细的日志信息
5. **并发安全**: 管理器使用读写锁保证线程安全
6. **规则格式**: 确保 JSON 格式正确，解析失败的任务会被跳过

## 扩展开发

当前版本中，以下功能标记为 TODO，需要根据具体业务需求进行实现：

1. **时间范围检查**: 根据 `timeRangeType` 和 `timeRangeValue` 进行复杂的时间范围判断
2. **触发条件检查**: 集成设备管理器，实时检查遥测数据是否满足触发条件
3. **执行操作**: 集成设备控制接口，执行具体的设备操作

这些功能的具体实现需要根据项目的设备管理架构和控制接口进行开发。
