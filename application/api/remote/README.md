# Remote API 文档

## 概述

Remote API 提供了远程管理功能的接口，包括 MQTT 和 Modbus 服务的状态查询和管理。

## API 接口

### Modbus 服务管理

#### 获取 Modbus 服务状态

**接口路径**: `GET /remote/modbus/status`

**功能描述**: 获取 Modbus TCP 服务的运行状态和所有设备的点位定义情况

**请求参数**: 无

**响应数据结构**:

```json
{
  "isRunning": true,
  "listenPort": 502,
  "deviceCount": 2,
  "connectionCount": 0,
  "deviceStatus": [
    {
      "deviceId": "pylon_bms",
      "modbusId": 1,
      "startAddr": 40000,
      "isOnline": true,
      "lastUpdateTime": "2024-01-01T12:00:00Z",
      "error": "",
      "totalRegisters": 10,
      "pointMappings": [
        {
          "pointKey": "system.online_status",
          "pointName": "设备在线状态",
          "valueType": "EBool",
          "startOffset": 0,
          "registerCount": 1,
          "isSystemPoint": true,
          "description": "设备在线状态：0=离线，1=在线"
        },
        {
          "pointKey": "system.timestamp",
          "pointName": "通讯时间戳",
          "valueType": "EUint32",
          "startOffset": 1,
          "registerCount": 2,
          "isSystemPoint": true,
          "description": "通讯时间戳：Unix时间戳（秒）"
        },
        {
          "pointKey": "battery.voltage",
          "pointName": "电池电压",
          "valueType": "EFloat32",
          "startOffset": 3,
          "registerCount": 2,
          "isSystemPoint": false,
          "description": "电池电压"
        }
      ]
    }
  ]
}
```

**响应字段说明**:

- `isRunning`: Modbus 服务是否正在运行
- `listenPort`: 监听端口（默认 502）
- `deviceCount`: 设备数量
- `connectionCount`: 连接数量
- `deviceStatus`: 设备状态列表
  - `deviceId`: 设备ID
  - `modbusId`: Modbus 从站ID
  - `startAddr`: 起始地址
  - `isOnline`: 设备是否在线
  - `lastUpdateTime`: 最后更新时间
  - `error`: 错误信息（如果有）
  - `totalRegisters`: 总寄存器数量
  - `pointMappings`: 点位映射列表
    - `pointKey`: 点位键名
    - `pointName`: 点位名称
    - `valueType`: 数据类型
    - `startOffset`: 相对起始地址的偏移
    - `registerCount`: 占用寄存器数量
    - `isSystemPoint`: 是否为系统固定点位
    - `description`: 点位描述

### MQTT 服务管理

#### 获取 MQTT 服务状态

**接口路径**: `GET /remote/mqtt/status`

**功能描述**: 获取 MQTT 服务的运行状态和客户端连接情况

#### 重新加载 MQTT 配置

**接口路径**: `POST /remote/mqtt/reload`

**功能描述**: 重新加载 MQTT 服务配置

## 使用示例

### 获取 Modbus 服务状态

```bash
curl -X GET "http://localhost:8000/remote/modbus/status"
```

### 获取 MQTT 服务状态

```bash
curl -X GET "http://localhost:8000/remote/mqtt/status"
```

### 重新加载 MQTT 配置

```bash
curl -X POST "http://localhost:8000/remote/mqtt/reload"
```

## 注意事项

1. **Modbus 服务状态**: 需要确保 Modbus 服务已正确配置和启动
2. **设备映射**: 设备映射信息基于设备的 `GetExportModbusPoints()` 方法返回的点位
3. **系统固定点位**: 每个设备的前两个点位为系统保留（在线状态、时间戳）
4. **寄存器映射**: 支持多种数据类型的寄存器映射，字符串类型会被跳过
5. **错误处理**: 单个设备启动失败不会影响其他设备的运行

## 相关服务

- `s_export_modbus`: Modbus TCP 数据导出服务
- `s_export_mqtt`: MQTT 数据导出服务
