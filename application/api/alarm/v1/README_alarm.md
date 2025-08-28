# 告警接口文档

## 接口概述

本模块提供了完整的告警管理功能，包括当前告警查询、历史告警查询、告警忽略管理等功能。

## 接口列表

### 1. 当前告警分页查询

**接口地址：** `GET /alarm/current`

**功能描述：** 获取当前告警的分页列表

**请求参数：**
```json
{
  "deviceId": "string",    // 可选，指定设备ID；为空则查询全部
  "level": "string",       // 可选，告警级别过滤(LOW/MEDIUM/HIGH/CRITICAL/ALL)
  "point": "string",       // 可选，告警点位名称过滤
  "page": 1,              // 页码，从1开始，默认1
  "pageSize": 20          // 每页条数，最大100，默认20
}
```

**响应示例：**
```json
{
  "total": 100,
  "items": [
    {
      "id": 1,
      "deviceId": "device001",
      "point": "temperature",
      "level": "HIGH",
      "title": "温度过高",
      "detail": "设备温度超过阈值",
      "createdAt": "2024-01-01 12:00:00"
    }
  ]
}
```

### 2. 历史告警分页查询

**接口地址：** `GET /alarm/history`

**功能描述：** 获取历史告警的分页列表

**请求参数：**
```json
{
  "deviceId": "string",    // 可选，指定设备ID；为空则查询全部
  "level": "string",       // 可选，告警级别过滤(LOW/MEDIUM/HIGH/CRITICAL/ALL)
  "point": "string",       // 可选，告警点位名称过滤
  "title": "string",       // 可选，告警标题模糊搜索
  "date": "string",        // 可选，日期过滤，格式：2006-01-02
  "page": 1,              // 页码，从1开始，默认1
  "pageSize": 20          // 每页条数，最大100，默认20
}
```

**响应示例：**
```json
{
  "total": 1000,
  "items": [
    {
      "id": 1,
      "deviceId": "device001",
      "point": "voltage",
      "level": "CRITICAL",
      "title": "电压异常",
      "detail": "设备电压超出安全范围",
      "createdAt": "2024-01-01 10:00:00"
    }
  ]
}
```

### 3. 创建忽略告警

**接口地址：** `POST /alarm/ignore`

**功能描述：** 创建告警忽略记录，指定设备ID和告警点位名称

**请求参数：**
```json
{
  "deviceId": "device001",  // 必填，设备ID
  "point": "temperature"    // 必填，告警点位名称
}
```

**响应示例：**
```json
{
  "success": true,
  "message": "成功创建忽略告警"
}
```

### 4. 删除忽略告警

**接口地址：** `DELETE /alarm/ignore`

**功能描述：** 删除指定的告警忽略记录

**请求参数：**
```json
{
  "deviceId": "device001",  // 必填，设备ID
  "point": "temperature"    // 必填，告警点位名称
}
```

**响应示例：**
```json
{
  "success": true,
  "message": "成功删除忽略告警"
}
```

### 5. 忽略告警分页查询

**接口地址：** `GET /alarm/ignore`

**功能描述：** 获取忽略告警的分页列表

**请求参数：**
```json
{
  "deviceId": "string",    // 可选，指定设备ID；为空则查询全部
  "point": "string",       // 可选，告警点位名称过滤
  "date": "string",        // 可选，日期过滤，格式：2006-01-02
  "page": 1,              // 页码，从1开始，默认1
  "pageSize": 20          // 每页条数，最大100，默认20
}
```

**响应示例：**
```json
{
  "total": 50,
  "items": [
    {
      "id": 1,
      "deviceId": "device001",
      "point": "temperature",
      "createdAt": "2024-01-01 09:00:00"
    }
  ]
}
```

## 告警级别说明

- `LOW`: 低级别告警，不影响系统运行
- `MEDIUM`: 中级别告警，需要注意
- `HIGH`: 高级别告警，可能影响功能
- `CRITICAL`: 严重告警，需要立即处理

## 使用示例

### 查询当前告警
```bash
curl -X GET "http://localhost:8000/alarm/current?deviceId=device001&level=HIGH&page=1&pageSize=10"
```

### 创建忽略告警
```bash
curl -X POST "http://localhost:8000/alarm/ignore" \
  -H "Content-Type: application/json" \
  -d '{"deviceId": "device001", "point": "temperature"}'
```

### 删除忽略告警
```bash
curl -X DELETE "http://localhost:8000/alarm/ignore" \
  -H "Content-Type: application/json" \
  -d '{"deviceId": "device001", "point": "temperature"}'
```

## 注意事项

1. 所有分页查询的页码从1开始
2. 每页条数最大限制为100
3. 日期格式统一使用 `2006-01-02` 格式
4. 告警忽略功能会影响告警的触发，请谨慎使用
5. 接口返回的时间格式为 `YYYY-MM-DD HH:mm:ss`
