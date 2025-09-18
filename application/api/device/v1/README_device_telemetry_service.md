# 设备 Telemetry 和 Service 接口文档

## 接口概述

新增的 `GetDeviceTelemetryService` 接口用于获取指定设备的所有 Telemetry（遥测）和 Service（自定义服务）信息。

## 接口详情

### 请求

- **路径**: `GET /api/device/telemetry-service/{deviceId}`
- **标签**: 设备相关
- **摘要**: 获取指定设备的所有 Telemetry 和 Service

#### 请求参数

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| deviceId | string | 是 | 设备ID（路径参数） |

#### 请求示例

```bash
curl -X GET "http://localhost:15880/api/device/telemetry-service/device-001"
```

### 响应

#### 响应结构

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "deviceId": "device-001",
    "deviceName": "测试设备",
    "telemetry": [
      {
        "key": "voltage",
        "name": "电压",
        "precise": 2,
        "unit": "V",
        "desc": "设备电压值",
        "valueExplain": {
          "0": "正常",
          "1": "异常"
        },
        "paramExplain": null
      }
    ],
    "service": [
      {
        "key": "restart",
        "name": "重启设备",
        "description": "重启设备服务",
        "params": [
          {
            "key": "delay",
            "name": "延迟时间",
            "type": "int",
            "required": false,
            "default": 0,
            "desc": "重启延迟时间（秒）"
          }
        ]
      }
    ]
  }
}
```

#### 响应字段说明

| 字段名 | 类型 | 描述 |
|--------|------|------|
| deviceId | string | 设备ID |
| deviceName | string | 设备名称 |
| telemetry | array | 遥测信息列表 |
| service | array | 自定义服务列表 |

#### Telemetry 字段说明

| 字段名 | 类型 | 描述 |
|--------|------|------|
| key | string | 遥测键名 |
| name | string | 遥测名称（支持国际化） |
| precise | number | 浮点数精度 |
| unit | string | 单位 |
| desc | string | 描述 |
| valueExplain | object | 值解释映射 |
| paramExplain | object | 参数解释映射 |

#### Service 字段说明

| 字段名 | 类型 | 描述 |
|--------|------|------|
| key | string | 服务方法名 |
| name | string | 服务名称（支持国际化） |
| description | string | 服务描述 |
| params | array | 参数定义列表 |

## 错误处理

### 常见错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 400 | 设备ID不能为空 | 检查请求参数 |
| 404 | 设备不存在 | 确认设备ID是否正确 |
| 404 | 设备实例不存在 | 确认设备是否已启动 |
| 404 | 设备驱动信息不存在 | 确认设备驱动是否已加载 |
| 500 | 系统正在初始化中 | 等待系统初始化完成 |

### 错误响应示例

```json
{
  "code": 404,
  "message": "设备不存在",
  "data": null
}
```

## 使用场景

1. **设备配置管理**: 获取设备的遥测点和服务列表，用于前端配置界面
2. **设备监控**: 了解设备支持的遥测数据类型
3. **设备控制**: 获取设备支持的自定义服务，用于设备控制操作
4. **系统集成**: 第三方系统集成时获取设备能力信息

## 注意事项

1. 该接口返回的是设备驱动定义的静态信息，不包含实时数据
2. 如果设备未启动或驱动未加载，将返回相应的错误信息
3. 遥测和服务信息来源于设备的驱动配置文件
4. 接口支持国际化，返回的名称字段会根据请求的语言头进行本地化

## 相关接口

- `GET /api/device/telemetry` - 获取设备实时遥测数据
- `GET /api/device/tree/{deviceId}` - 获取设备树信息
- `GET /api/device/real-list` - 获取真实设备列表
