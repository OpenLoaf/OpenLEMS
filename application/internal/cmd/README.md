# CMD 目录重构说明

## 重构概述

本次重构将原本分散在 `cmd.go` 和 `cmd_web.go` 中的代码重新组织，使其结构更清晰、更易于维护。

## 新的目录结构

```
application/internal/
├── cmd/                    # 主程序入口
│   ├── main.go            # 主程序逻辑（简化）
│   ├── web.go             # Web服务启动逻辑
│   ├── ems.go             # EMS系统核心功能
│   └── cmd_test.go        # 测试文件（保留）
└── utils/                 # 工具函数
    ├── pid.go            # PID文件管理
    ├── network.go        # 网络工具函数
    ├── middleware.go     # Web中间件
    └── web.go            # Web服务工具函数
```

## 文件功能说明

### cmd/main.go
- 主程序入口点
- 命令行参数定义
- 系统初始化和服务启动的协调
- 信号处理和优雅关闭

### cmd/web.go
- Web服务启动逻辑
- API路由设置
- 静态文件服务配置
- 中间件配置

### cmd/ems.go
- EMS系统核心功能
- 系统初始化逻辑
- 服务启动逻辑
- 关闭信号处理

### utils/pid.go
- PID文件的创建和删除
- 进程管理相关工具函数

### utils/network.go
- 获取本地IPv4地址
- 网络相关工具函数

### utils/middleware.go
- Web中间件实现
- 请求ID生成
- 错误处理
- 访问日志记录

### utils/web.go
- Web服务信息打印
- 服务器地址显示


## 重构优势

1. **职责分离**：每个文件都有明确的职责，代码更易理解
2. **模块化**：工具函数独立成包，便于复用
3. **可维护性**：代码结构清晰，修改时影响范围小
4. **可测试性**：各个模块可以独立测试
5. **可扩展性**：新增功能时结构清晰，易于扩展

## 保持不变的功能

- 所有原有的业务逻辑保持不变
- 命令行参数和配置保持不变
- Web服务功能保持不变
- 系统启动和关闭流程保持不变

## 使用方式

重构后的使用方式与之前完全相同：

```bash
# 启动服务（不启动Web）
go run main.go

# 启动服务（启动Web）
go run main.go --web

# 其他参数保持不变
go run main.go --web --profile=dev --device-name=mydevice
```

## 注意事项

1. 所有导入路径已更新，确保编译时路径正确
2. 原有的测试文件 `cmd_test.go` 已保留
3. 如果需要在其他地方使用工具函数，请从 `application/internal/utils` 包导入
4. 系统初始化逻辑现在在 `cmd/ems.go` 中，便于其他模块调用
