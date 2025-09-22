# 配置结构体验证功能

## 概述

本模块提供了配置结构体标签验证功能，包括两个主要的验证方法：

1. **`ValidateConfigStructFields(config any) (bool, error)`** - 验证配置结构体字段定义是否正确
2. **`ValidateConfigData(config any, data map[string]any) (bool, error)`** - 验证前端传递的数据是否符合字段定义要求

## 功能特性

### 1. 字段定义验证

验证配置结构体字段定义的正确性，包括：

- ✅ 必需标签检查（`code`、`ct/componentType/component_type`）
- ✅ 组件类型和值类型匹配性检查
- ✅ 数值范围验证（`min/max`）
- ✅ 正则表达式验证
- ✅ 选择组件选项配置验证

### 2. 数据验证

验证前端传递的数据是否符合字段定义要求，包括：

- ✅ 必填字段验证
- ✅ 数据类型验证
- ✅ 数值范围验证
- ✅ 正则表达式验证
- ✅ 选择选项验证（`singleSelect`/`multiSelect`）

## 使用方法

### 基本用法

```go
package main

import (
    "fmt"
    "common/c_base"
)

type DeviceConfig struct {
    Name     string `code:"name" name:"设备名称" ct:"text" vt:"string" req:"true" desc:"设备显示名称"`
    Port     int    `code:"port" name:"端口号" ct:"number" vt:"int" min:"1" max:"65535" def:"8080"`
    Enabled  bool   `code:"enabled" name:"启用状态" ct:"switch" vt:"bool" def:"true"`
    Protocol string `code:"protocol" name:"协议类型" ct:"singleSelect" vt:"string" ve:"tcp:TCP,udp:UDP,modbus:Modbus"`
}

func main() {
    config := &DeviceConfig{}
    
    // 1. 验证配置结构体字段定义
    valid, err := c_base.ValidateConfigStructFields(config)
    if err != nil {
        fmt.Printf("配置结构体验证失败: %v\n", err)
        return
    }
    fmt.Printf("配置结构体验证结果: %v\n", valid)
    
    // 2. 验证前端传递的数据
    data := map[string]any{
        "name":     "测试设备",
        "port":     8080,
        "enabled":  true,
        "protocol": "tcp",
    }
    
    valid, err = c_base.ValidateConfigData(config, data)
    if err != nil {
        fmt.Printf("数据验证失败: %v\n", err)
        return
    }
    fmt.Printf("数据验证结果: %v\n", valid)
}
```

### 高级用法

```go
// 构建字段定义用于前端渲染
fields, err := c_base.BuildConfigStructFields(config)
if err != nil {
    log.Fatal(err)
}

for _, field := range fields {
    fmt.Printf("字段: %s, 类型: %s, 组件: %s, 必填: %v\n", 
        field.Key, field.ValueType, field.ComponentType, field.Required)
}
```

## 支持的标签

### 必需标签

- `code` - JSON序列化字段名（必需）
- `ct/componentType/component_type` - 组件类型（必需）

### 可选标签

- `name` - 字段显示名称
- `vt/valueType/value_type` - 值类型
- `req/required/required` - 是否必填
- `desc/description/description` - 字段描述
- `group/group/group` - 字段分组
- `min/min/min` - 最小值限制
- `max/max/max` - 最大值限制
- `def/default/default` - 默认值
- `unit/unit/unit` - 单位信息
- `step/step/step` - 数值步长
- `regex/regex/regex` - 正则表达式验证
- `rfm/regexFailedMessage/regex_failed_message` - 正则验证失败提示
- `ve/valueExplain/valueExplain` - 值解释配置
- `pe/paramExplain/paramExplain` - 参数解释配置

## 组件类型支持

- `text` - 文本输入框
- `number` - 数字输入框
- `switch` - 开关组件
- `singleSelect` - 单选下拉框
- `multiSelect` - 多选下拉框
- `date` - 日期选择器
- `time` - 时间选择器
- `dateTime` - 日期时间选择器

## 值类型支持

- `string` - 字符串类型
- `int` - 整数类型
- `float` - 浮点数类型
- `bool` - 布尔类型

## 验证规则

### 1. 字段定义验证规则

- 所有字段必须有 `code` 标签
- 所有字段必须有 `ct/componentType/component_type` 标签
- 组件类型和值类型必须匹配
- 数值范围：`min` 不能大于 `max`
- 正则表达式必须有效
- 选择组件必须有选项配置

### 2. 数据验证规则

- 必填字段不能为空
- 数据类型必须匹配
- 数值必须在指定范围内
- 字符串必须匹配正则表达式（如果配置了）
- 选择值必须在有效选项中

## 错误处理

验证方法返回详细的错误信息，包括：

- 字段名和错误类型
- 具体的验证失败原因
- 建议的修复方法

## 性能特性

- 使用反射缓存提高性能
- 支持并发安全
- 内存占用优化

## 测试覆盖

- 单元测试覆盖所有验证逻辑
- 集成测试验证完整流程
- 基准测试确保性能要求

## 示例代码

详细的使用示例请参考 `c_base_config_validation_example.go` 文件。

## 注意事项

1. 正则表达式标签中的逗号不会被错误分割
2. 支持嵌套结构体的平铺展开
3. 向后兼容传统的 `json` 标签
4. 所有验证都是非阻塞的，遇到错误立即返回
