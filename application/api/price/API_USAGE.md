# 电价管理API使用说明

## 概述

电价管理API提供完整的电价配置管理功能，支持电价的创建、更新、删除、查询、启用/停用等操作，以及获取当前激活电价等功能。

## API接口列表

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 创建电价 | POST | `/api/price` | 创建新的电价配置 |
| 更新电价 | PUT | `/api/price/{id}` | 更新指定电价配置 |
| 删除电价 | DELETE | `/api/price/{id}` | 删除指定电价配置 |
| 电价列表 | GET | `/api/price` | 获取电价列表（分页） |
| 电价详情 | GET | `/api/price/{id}` | 获取指定电价详情 |
| 启用/停用 | PUT | `/api/price/{id}/toggle` | 启用或停用电价 |
| 当前电价 | GET | `/api/price/current` | 获取当前激活的电价 |

## 详细接口说明

### 1. 创建电价

**接口**: `POST /api/price`

**请求参数**:
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
    }
  ],
  "remoteId": null
}
```

**响应**: 空响应体

**参数说明**:
- `description`: 电价描述，必填，长度2-100字符
- `priority`: 优先级，必填，范围1-5，数值越小优先级越高
- `status`: 状态，必填，可选值：Enable/Disable
- `dateRange`: 日期范围配置，必填
- `timeRange`: 时间范围配置，必填
- `priceSegments`: 电价时段配置，必填
- `remoteId`: 远程电价ID，可选，用于远程电价

### 2. 更新电价

**接口**: `PUT /api/price/{id}`

**请求参数**: 与创建电价相同，额外包含 `id` 字段

**响应**: 空响应体

### 3. 删除电价

**接口**: `DELETE /api/price/{id}`

**路径参数**:
- `id`: 电价ID，必填

**响应**: 空响应体

### 4. 电价列表

**接口**: `GET /api/price`

**查询参数**:
- `page`: 页码，必填，最小值1
- `pageSize`: 每页数量，必填，范围1-100
- `status`: 状态筛选，可选，值：Enable/Disable
- `keyword`: 关键词搜索，可选，最大长度50
- `priority`: 优先级筛选，可选，范围1-5

**响应**:
```json
{
  "list": [
    {
      "id": 1,
      "description": "工作日峰谷电价",
      "priority": 1,
      "status": "Enable",
      "isActive": true,
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
        }
      ],
      "remoteId": null,
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z",
      "createdBy": "admin"
    }
  ],
  "total": 1
}
```

### 5. 电价详情

**接口**: `GET /api/price/{id}`

**路径参数**:
- `id`: 电价ID，必填

**响应**:
```json
{
  "price": {
    "id": 1,
    "description": "工作日峰谷电价",
    "priority": 1,
    "status": "Enable",
    "isActive": true,
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
      }
    ],
    "remoteId": null,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "createdBy": "admin"
  }
}
```

### 6. 启用/停用电价

**接口**: `PUT /api/price/{id}/toggle`

**路径参数**:
- `id`: 电价ID，必填

**请求参数**:
```json
{
  "status": "Enable"
}
```

**响应**: 空响应体

### 7. 获取当前激活电价

**接口**: `GET /api/price/current`

**响应**:
```json
{
  "price": {
    "id": 1,
    "description": "工作日峰谷电价",
    "priority": 1,
    "status": "Enable",
    "isActive": true,
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
      }
    ],
    "remoteId": null,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "createdBy": "admin"
  }
}
```

## 数据结构说明

### 电价时段 (PriceSegment)

| 字段 | 类型 | 描述 | 示例 |
|------|------|------|------|
| startTime | string | 开始时间 | "08:30" |
| endTime | string | 结束时间 | "12:45" |
| priceType | string | 电价类型 | "valley" |
| price | number | 电价值 | 0.3 |

### 日期范围 (DateRange)

| 字段 | 类型 | 描述 | 示例 |
|------|------|------|------|
| isLongTerm | boolean | 是否长期有效 | true |
| startDate | string | 开始日期 | "2024-01-01" |
| endDate | string | 结束日期 | "2024-12-31" |

### 时间范围 (TimeRange)

| 字段 | 类型 | 描述 | 示例 |
|------|------|------|------|
| type | string | 时间范围类型 | "weekday" |
| weekdayType | string | 工作日类型 | "workday" |
| customDays | array | 自定义日期 | [1, 15, 30] |
| customMonths | array | 自定义月份 | [1, 6, 12] |

## 电价类型说明

| 类型 | 值 | 描述 |
|------|----|----|
| 谷电 | valley | 低峰时段电价 |
| 峰电 | peak | 高峰时段电价 |
| 平电 | flat | 平时段电价 |
| 尖峰 | sharp | 尖峰时段电价 |
| 深谷 | deep_valley | 深谷时段电价 |

## 时间范围类型说明

### 1. 工作日类型 (weekday)
- `workday`: 工作日（周一到周五）
- `weekend`: 周末（周六到周日）
- `all`: 全部时间

### 2. 自定义类型 (custom)
- 支持自定义日期和月份
- `customDays`: 自定义日期（1-31）
- `customMonths`: 自定义月份（1-12）

### 3. 月度类型 (monthly)
- 每月1日生效

## 使用示例

### 创建工作日峰谷电价

```bash
curl -X POST "http://127.0.0.1:15880/api/price" \
  -H "Content-Type: application/json" \
  -H "Cookie: ems_session_id=your_session_id" \
  -d '{
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
  }'
```

### 获取电价列表

```bash
curl -X GET "http://127.0.0.1:15880/api/price?page=1&pageSize=10" \
  -H "Cookie: ems_session_id=your_session_id"
```

### 获取当前激活电价

```bash
curl -X GET "http://127.0.0.1:15880/api/price/current" \
  -H "Cookie: ems_session_id=your_session_id"
```

### 启用/停用电价

```bash
curl -X PUT "http://127.0.0.1:15880/api/price/1/toggle" \
  -H "Content-Type: application/json" \
  -H "Cookie: ems_session_id=your_session_id" \
  -d '{
    "status": "Disable"
  }'
```

## 错误处理

### 常见错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 400 | 参数验证失败 | 检查请求参数格式和必填字段 |
| 401 | 未授权访问 | 检查登录状态和权限 |
| 404 | 电价不存在 | 检查电价ID是否正确 |
| 500 | 服务器内部错误 | 联系系统管理员 |

### 错误响应格式

```json
{
  "code": 400,
  "message": "参数验证失败",
  "data": null
}
```

## 权限说明

- 所有电价管理接口都需要管理员权限
- 需要先登录获取 `ems_session_id`
- 详情请参考 [API接口测试登录规则](mdc:.cursor/rules/api-test-login.mdc)

## 注意事项

1. **时间格式**: 使用 "HH:MM" 格式，如 "08:30"
2. **优先级**: 数值越小优先级越高
3. **状态**: Enable表示启用，Disable表示停用
4. **远程电价**: remoteId为null表示本地电价，非null表示远程电价
5. **激活状态**: isActive字段表示当前是否在时间/日期范围内生效
6. **跨天时段**: 支持跨天的时间段，如 "23:00" 到 "07:00"

## 扩展功能

### 远程电价支持
- 通过 `remoteId` 字段区分本地和远程电价
- 支持MQTT下发电价配置（预留功能）
- 远程电价和本地电价使用相同的激活判断逻辑

### 定时存储
- 系统每小时自动保存当前电价到Storage
- 支持电价历史记录查询
- 存储路径：`system/price/`，包含以下字段：
  - `priceId`: 电价ID
  - `price`: 电价值
  - `priceType`: 电价类型（数值：1=谷电, 2=峰电, 3=平电, 4=尖峰, 5=深谷）
  - `timestamp`: 时间戳

## 相关文档

- [电价管理服务说明](mdc:services/s_price/README.md)
- [API接口测试登录规则](mdc:.cursor/rules/api-test-login.mdc)
- [API生成规范](mdc:.cursor/rules/api-generation-standards.mdc)
