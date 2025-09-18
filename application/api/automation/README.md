# Automation API 接口文档

## 概述

Automation API 提供了完整的自动化任务管理功能，包括创建、查询、更新、删除和状态控制等操作。

## 接口列表

### 1. 获取自动化分页列表

**接口路径**: `GET /automation/page`

**功能**: 获取所有自动化任务的分页列表，支持按设备ID和启用状态过滤

**请求参数**:
- `page`: 页码，默认为1
- `pageSize`: 每页数量，默认为10，最大100
- `deviceId`: 设备ID（可选），用于过滤特定设备的自动化任务
- `enabled`: 是否启用（可选），用于过滤启用/停用的任务

**响应数据**:
```json
{
  "list": [
    {
      "id": 1,
      "startTime": "2024-01-01T00:00:00Z",
      "endTime": "2024-12-31T23:59:59Z",
      "timeRangeType": "daily",
      "timeRangeValue": "08:00-18:00",
      "triggerRule": "{\"type\":\"time\",\"value\":\"08:00\"}",
      "executeRule": "{\"deviceId\":\"device001\",\"action\":\"start\"}",
      "enabled": true,
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    }
  ],
  "total": 1,
  "page": 1,
  "pageSize": 10
}
```

### 2. 获取设备自动化列表

**接口路径**: `GET /automation/device/{deviceId}`

**功能**: 获取指定设备的所有自动化任务列表（包括已停用的任务）

**路径参数**:
- `deviceId`: 设备ID

**响应数据**:
```json
{
  "automations": [
    {
      "id": 1,
      "startTime": "2024-01-01T00:00:00Z",
      "endTime": "2024-12-31T23:59:59Z",
      "timeRangeType": "daily",
      "timeRangeValue": "08:00-18:00",
      "triggerRule": "{\"type\":\"time\",\"value\":\"08:00\"}",
      "executeRule": "{\"deviceId\":\"device001\",\"action\":\"start\"}",
      "enabled": true,
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    }
  ],
  "count": 1
}
```

### 3. 创建自动化任务

**接口路径**: `POST /automation`

**功能**: 创建新的自动化任务

**请求参数**:
- `startTime`: 开始时间（可选）
- `endTime`: 结束时间（可选）
- `timeRangeType`: 时间范围类型（必填）
- `timeRangeValue`: 时间范围值（必填）
- `triggerRule`: 触发规则，JSON格式（必填）
- `executeRule`: 执行规则，JSON格式（必填）
- `enabled`: 是否启用，默认为true

**响应数据**:
```json
{
  "id": 1
}
```

### 4. 更新自动化任务

**接口路径**: `PUT /automation/{id}`

**功能**: 更新指定的自动化任务

**路径参数**:
- `id`: 自动化任务ID

**请求参数**:
- `startTime`: 开始时间（可选）
- `endTime`: 结束时间（可选）
- `timeRangeType`: 时间范围类型（可选）
- `timeRangeValue`: 时间范围值（可选）
- `triggerRule`: 触发规则，JSON格式（可选）
- `executeRule`: 执行规则，JSON格式（可选）
- `enabled`: 是否启用（可选）

**响应数据**:
```json
{}
```

### 5. 删除自动化任务

**接口路径**: `DELETE /automation/{id}`

**功能**: 删除指定的自动化任务

**路径参数**:
- `id`: 自动化任务ID

**响应数据**:
```json
{}
```

### 6. 开启/停用自动化任务

**接口路径**: `POST /automation/{id}/toggle`

**功能**: 开启或停用指定的自动化任务

**路径参数**:
- `id`: 自动化任务ID

**请求参数**:
- `enable`: 是否启用（true=开启，false=停用）

**响应数据**:
```json
{
  "enabled": true
}
```

## 数据模型说明

### 自动化任务字段说明

- `id`: 任务唯一标识
- `startTime`: 任务开始时间
- `endTime`: 任务结束时间
- `timeRangeType`: 时间范围类型（如：daily、weekly、monthly等）
- `timeRangeValue`: 时间范围值（如：08:00-18:00）
- `triggerRule`: 触发规则，JSON格式存储
- `executeRule`: 执行规则，JSON格式存储
- `enabled`: 是否启用
- `createdAt`: 创建时间
- `updatedAt`: 更新时间

### 规则格式说明

#### 触发规则 (triggerRule)
```json
{
  "type": "time",
  "value": "08:00",
  "conditions": [
    {
      "deviceId": "device001",
      "point": "temperature",
      "operator": ">",
      "value": 25
    }
  ]
}
```

#### 执行规则 (executeRule)
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

## 使用示例

### 创建定时任务
```bash
curl -X POST http://localhost:8000/automation \
  -H "Content-Type: application/json" \
  -d '{
    "timeRangeType": "daily",
    "timeRangeValue": "08:00-18:00",
    "triggerRule": "{\"type\":\"time\",\"value\":\"08:00\"}",
    "executeRule": "{\"deviceId\":\"device001\",\"action\":\"start\"}",
    "enabled": true
  }'
```

### 获取分页列表
```bash
curl "http://localhost:8000/automation/page?page=1&pageSize=10&enabled=true"
```

### 停用任务
```bash
curl -X POST http://localhost:8000/automation/1/toggle \
  -H "Content-Type: application/json" \
  -d '{"enable": false}'
```

## 注意事项

1. 所有时间字段使用 ISO 8601 格式
2. 触发规则和执行规则必须为有效的JSON格式
3. 删除操作不可逆，请谨慎使用
4. 停用的任务不会执行，但仍会保留在系统中
5. 分页查询的pageSize最大值为100
