# GPIO协议插件构建说明

## 构建标签

本插件使用构建标签来控制Linux provider和Mock provider的编译选择：

### 标签说明

- `linux`: 系统平台标签，表示Linux系统
- `gpio_enable`: 功能标签，表示启用GPIO功能

### 编译规则

1. **Linux Provider** (`*_linux_provider_s_f.go`):
   - 构建标签: `//go:build linux && gpio_enable`
   - 编译条件: 只有在Linux系统且启用gpio_enable标签时才编译

2. **Mock Provider** (`*_mock_provider_s_f.go`):
   - 构建标签: `//go:build !linux || !gpio_enable`
   - 编译条件: 非Linux系统，或Linux系统但未启用gpio_enable标签时编译

### 使用示例

#### 启用GPIO功能（Linux系统）
```bash
# 编译时启用GPIO功能
go build -tags="gpio_enable" ./...

# 或者同时启用多个标签
go build -tags="gpio_enable,dev" ./...
```

#### 禁用GPIO功能（使用Mock）
```bash
# 不指定gpio_enable标签，将使用Mock provider
go build ./...

# 或者明确禁用
go build -tags="!gpio_enable" ./...
```

### 编译结果

| 系统平台 | 构建标签 | 编译结果 |
|---------|---------|---------|
| Linux | `gpio_enable` | Linux Provider |
| Linux | 无标签 | Mock Provider |
| 非Linux | 任意 | Mock Provider |

### 注意事项

1. 在Linux系统上，如果不指定`gpio_enable`标签，将使用Mock provider
2. Mock provider不依赖Linux特定的GPIO库，可以在任何平台上运行
3. 构建标签是编译时决定的，运行时无法动态切换
