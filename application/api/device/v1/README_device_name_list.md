# 设备名称列表接口

## 接口描述

获取所有设备的名称列表，返回格式为 `{设备ID: 设备名称}` 的映射。

## 接口信息

- **路径**: `/api/device/name-list`
- **方法**: `GET`
- **标签**: 设备相关
- **摘要**: 获取所有设备的名称列表

## 请求参数

无参数

## 响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "deviceNames": {
      "1": "星星充电PCS",
      "2": "派能BMS", 
      "3": "亿兰科MAC",
      "4": "储能柜"
    }
  }
}
```

## 响应字段说明

| 字段名 | 类型 | 说明 |
|--------|------|------|
| deviceNames | map[string]string | 设备名称映射，key为设备ID，value为设备名称 |

## 使用示例

### cURL 示例

```bash
curl -X GET "http://localhost:8000/api/device/name-list" \
  -H "Accept: application/json"
```

### JavaScript 示例

```javascript
fetch('/api/device/name-list')
  .then(response => response.json())
  .then(data => {
    console.log('设备名称列表:', data.data.deviceNames);
  });
```

## 实现说明

该接口使用 `s_db.GetDeviceService().GetAllDevices()` 方法获取所有设备信息，然后提取设备ID和名称构造成映射返回。

## 错误处理

- 如果数据库连接失败，会返回相应的错误信息
- 如果设备列表为空，会返回空的映射对象
