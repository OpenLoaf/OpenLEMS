# MQTT数据推送服务 (s_export_mqtt)

## 概述

MQTT数据推送服务负责从数据库配置中读取MQTT连接列表，定时获取指定设备的数据并按照standard格式推送到MQTT服务器。

## 功能特性

- **配置驱动**: 从数据库设置中读取MQTT配置列表
- **多连接支持**: 支持同时管理多个MQTT连接
- **定时推送**: 根据配置的`uploadPeriod`定时推送设备数据
- **数据格式化**: 支持standard格式的数据格式化，可扩展其他格式
- **动态重载**: 支持运行时重新加载配置
- **错误处理**: 单个连接失败不影响其他连接继续工作
- **自动重连**: MQTT连接断开时自动重连

## 配置格式

MQTT配置存储在数据库的`mqtt_config_list`设置中，JSON格式如下：

```json
[
  {
    "serverAddress": "127.0.0.1",
    "serverPort": 1883,
    "serviceStandard": "standard",
    "allowControl": false,
    "enabled": true,
    "deviceIds": ["pylon_bms"],
    "rewriteChannel": false,
    "pushChannel": "111",
    "subscribeChannel": "222",
    "uploadPeriod": 60
  }
]
```

### 配置字段说明

- `serverAddress`: MQTT服务器地址
- `serverPort`: MQTT服务器端口
- `serviceStandard`: 服务标准（目前支持"standard"）
- `allowControl`: 是否允许控制（本期不实现）
- `enabled`: 是否启用该配置
- `deviceIds`: 要推送数据的设备ID列表
- `rewriteChannel`: 是否重写通道
- `pushChannel`: 自定义推送通道（topic）
- `subscribeChannel`: 订阅通道（本期不实现）
- `uploadPeriod`: 上传周期（秒）

## 数据格式

### Standard格式

推送的数据格式为：
```json
{
  "sn": "设备ID",
  "time": 时间戳（毫秒）,
  "data": {
    "pointA": 数值,
    "pointB": 数值
  }
}
```

### Topic格式

默认topic格式：`lems/{system_number}/info`

其中`{system_number}`会被替换为系统序列号（从数据库设置中获取）。

如果配置了`rewriteChannel=true`且`pushChannel`不为空，则使用`pushChannel`作为topic。

## 使用方法

### 1. 初始化服务

```go
import "s_export_mqtt"

// 初始化MQTT导出服务
s_export_mqtt.Init()
```

### 2. 启动服务

```go
// 启动MQTT导出服务
err := s_export_mqtt.StartMqttExporter(ctx)
if err != nil {
    log.Fatalf("启动MQTT导出服务失败: %v", err)
}
```

### 3. 停止服务

```go
// 停止MQTT导出服务
err := s_export_mqtt.StopMqttExporter(ctx)
if err != nil {
    log.Errorf("停止MQTT导出服务失败: %v", err)
}
```

### 4. 重新加载配置

```go
// 重新加载MQTT配置
err := s_export_mqtt.ReloadMqttExporter(ctx)
if err != nil {
    log.Errorf("重新加载MQTT配置失败: %v", err)
}
```

### 5. 获取服务状态

```go
// 获取MQTT导出服务状态
isRunning, clientCount := s_export_mqtt.GetMqttExporterStatus()
fmt.Printf("服务运行状态: %v, 客户端数量: %d\n", isRunning, clientCount)
```

## 架构设计

### 核心组件

1. **SMqttExportConfig**: MQTT配置结构体
2. **IDataFormatter**: 数据格式化器接口
3. **SStandardFormatter**: Standard格式实现
4. **SMqttClient**: 单个MQTT连接管理
5. **SMqttExportManager**: MQTT导出管理器

### 扩展性

- **数据格式扩展**: 实现`IDataFormatter`接口可支持新的数据格式
- **服务标准扩展**: 在`createFormatter`方法中添加新的格式化器

## 错误处理

- 单个MQTT连接失败不影响其他连接
- 连接失败时自动重连（配置AutoReconnect）
- 发布失败记录日志但不中断定时器
- 设备无数据时跳过推送

## 日志记录

服务会记录以下日志：
- MQTT连接成功/失败
- 数据推送成功/失败
- 配置加载和重载
- 服务启动和停止

## 依赖

- `github.com/eclipse/paho.mqtt.golang`: MQTT客户端库
- `common`: 通用库（设备管理、日志等）
- `s_db`: 数据库服务（配置读取）
