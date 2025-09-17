# 自动化执行规则使用示例

## 结构体说明

### SAutomationExecuteRule
单个执行规则，包含：
- `DeviceId`: 设备ID
- `Service`: 服务名称  
- `Params`: 参数列表

### SAutomationExecuteConfig
执行配置，包含多个执行规则的数组。

## 使用示例

### 1. 创建执行规则

```go
// 创建单个执行规则
rule1 := NewAutomationExecuteRule("device001", "startService", []string{"param1", "param2"})
rule2 := NewAutomationExecuteRule("device002", "stopService", []string{"param3"})

// 创建执行配置
config := NewAutomationExecuteConfig()
config.AddRule(rule1)
config.AddRule(rule2)
```

### 2. JSON 序列化

```go
// 序列化为 JSON 字符串（用于存储到数据库）
jsonData, err := json.Marshal(config)
if err != nil {
    log.Fatal(err)
}
executeRuleString := string(jsonData)
```

### 3. JSON 反序列化

```go
// 从 JSON 字符串反序列化（从数据库读取）
var config SAutomationExecuteConfig
err := json.Unmarshal([]byte(executeRuleString), &config)
if err != nil {
    log.Fatal(err)
}
```

### 4. 在自动化任务中使用

```go
// 创建自动化任务时，执行规则会自动解析
automation := &s_db_model.SAutomationModel{
    ExecuteRule: `{"rules":[{"deviceId":"device001","service":"startService","params":["param1","param2"]}]}`,
}

task, err := NewAutomationTask(automation)
if err != nil {
    log.Fatal(err)
}

// 获取执行配置
executeConfig := task.GetExecuteConfig()
if executeConfig != nil {
    for _, rule := range executeConfig.GetRules() {
        fmt.Printf("设备: %s, 服务: %s, 参数: %v\n", 
            rule.GetDeviceId(), rule.GetService(), rule.GetParams())
    }
}
```

## JSON 格式示例

```json
{
  "rules": [
    {
      "deviceId": "device001",
      "service": "startService", 
      "params": ["param1", "param2"]
    },
    {
      "deviceId": "device002",
      "service": "stopService",
      "params": ["param3"]
    }
  ]
}
```