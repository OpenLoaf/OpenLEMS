# 能源系统配置说明

## 总线配置

BusNode组件的位置和长度现在可以通过JSON配置文件来控制，而不是仅依赖自动计算。

### 配置方法

在JSON配置文件中可以配置总线的位置和宽度：

```json
{
  "busNode": {
    "id": "bus",
    "type": "bus",
    "position": { 
      "x": 0,
      "y": 200  // 总线Y坐标（像素）
    },
    "data": { 
      "label": "总线",
      "busWidth": 800  // 总线宽度（像素）
    }
  }
}
```

### 配置示例

- `nodes.json` - 默认配置，总线宽度800px，Y坐标200px
- `nodes-wide.json` - 宽总线配置，总线宽度1200px，Y坐标250px
- `nodes-narrow.json` - 窄总线配置，总线宽度400px，Y坐标150px

### 工作原理

1. **优先使用配置值**：
   - 如果JSON中设置了`busWidth`，BusNode将直接使用该值
   - 如果JSON中设置了`position.y`，BusNode将使用该Y坐标值
2. **后备计算逻辑**：如果没有配置相应值，则使用原有的自动计算逻辑
3. **兼容性**：现有配置文件无需修改，会自动使用计算逻辑

### 使用建议

- 对于简单的布局，可以使用自动计算
- 对于需要精确控制总线位置和长度的场景，建议在JSON中明确设置`position.y`和`busWidth`
- `busWidth`值应该大于等于所有连接节点之间的最大距离
- `position.y`值应该考虑容器高度和节点布局
