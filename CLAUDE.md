# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## 中文开发指南

> 📝 **始终使用中文回复** 按照 AGENTS.md 中的三阶段工作流进行。

本项目是一个分布式能源管理系统 (EMS)，采用微服务架构和插件化设计。项目的核心规则都在 `.cursor/rules/` 目录下，请参考对应的规范文件。

## 常用构建命令

### 基础构建
```bash
# 进入项目根目录
cd /path/to/ems-plan

# 构建所有模块（开发环境，需要 -tags=dev）
go build -tags=dev -o ems ./application

# 生产环境构建
go build -o ems ./application

# 构建驱动插件
cd plugins/plug_drivers
make project=gpio_basic              # 构建单个驱动
make all                             # 构建所有驱动
make list                            # 列出所有驱动
make status                          # 检查构建状态
```

### C++ 库构建
```bash
cd @cpp/hexlib
make build                           # 编译 C++ 库（必须先执行）
make clean                           # 清理编译产物
make info                            # 显示编译信息
```

### 应用程序启动
```bash
cd application

# 基本启动（强制启动）
go run main.go --force=true

# 启动 Web 服务
go run main.go --web=true --force=true

# 开发环境启动
go run main.go --web=true --profile=dev --force=true

# 测试模式（3秒后自动关闭）
go run main.go --test=true --web=true --force=true

# 使用 GoFrame 开发模式（热重载）
gf run main.go -p ./bin -a "--web=true --force=true" -w "api/*.go" -w "internal/*.go"
```

### 测试
```bash
# 运行所有测试
go test ./...

# 运行特定模块测试
go test ./common/c_func/...

# 带覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# 带 CGO 的测试
CGO_ENABLED=1 go test ./... -v
```

## 项目架构

### 目录结构
- **common/**: 共享库（接口、类型定义、基础功能）
- **services/**: 核心服务层（s_db、s_driver、s_storage 等）
- **plugins/**: 插件系统（驱动、协议、存储、策略等）
- **application/**: 主应用程序
- **@cpp/**: C++ 扩展库（通过 CGO 集成）
- **tools/**: 工具模块

### 核心模块关系
```
application (主程序)
    ↓
services (核心服务)
    ├── s_db        (数据库服务)
    ├── s_driver    (驱动管理)
    ├── s_storage   (存储管理)
    └── s_policy    (策略管理)
    ↓
plugins (可插拔模块)
    ├── plug_drivers    (设备驱动)
    ├── plug_protocols  (通信协议)
    ├── plug_storages   (数据存储)
    └── plug_policy     (策略实现)
    ↓
common (共享库)
    ├── c_base      (基础接口)
    ├── c_device    (设备接口)
    └── c_enum      (枚举定义)
```

## 关键规则和约定

### 命名规范
- **接口**: `I` 前缀 (e.g., `IDevice`, `IDriver`)
- **结构体**: `S` 前缀 (e.g., `SDeviceConfig`)
- **枚举**: `E` 前缀 (e.g., `EDeviceType`)
- **常量**: `C` 前缀 (e.g., `CDeviceTypeAmmeter`)

### 文件命名 (common/ 目录)
- `c_*_i.go`: 接口定义
- `c_*_s.go`: 结构体定义
- `c_*_e.go`: 枚举定义
- `c_*_f.go`: 函数定义
- `c_*_c.go`: 常量定义

### 重要约定
1. **开发环境构建**: 必须使用 `-tags=dev` 参数
2. **CGO 支持**: Application 中必须 `CGO_ENABLED=1`
3. **错误处理**: 统一使用 `github.com/pkg/errors`
4. **日志系统**: 使用 `c_log` 包进行统一日志管理
5. **上下文传递**: 所有服务方法应接受 `context.Context`
6. **注解规范**: 所有公共方法、字段、接口都必须有中文注解

## 核心规则文件索引

这些规则文件定义了项目开发的关键规范，根据任务类型参考。表格列说明：
- **alwaysApply**: ✅ 表示自动应用，❌ 表示按需应用
- **globs**: 文件匹配模式，用于自动激活规则

| 规则文件 | 自动应用 | 适用场景 | 文件匹配模式 | 路径 |
|---------|--------|--------|-----------|------|
| **code-common-library.mdc** | ✅ | Common库开发、接口定义、命名规范、错误处理、代码注解 | `common/**/*.go` | `.cursor/rules/code-common-library.mdc` |
| **arch-project-structure.mdc** | ❌ | 项目结构、开发规范、技术栈、启动方式 | `**/*.go,**/*.md,**/*.yaml,**/*.sh` | `.cursor/rules/arch-project-structure.mdc` |
| **tech-startup-commands.mdc** | ❌ | 启动命令、参数配置、测试方法、调试指南、PID管理 | `application/**/*.go,application/**/*.sh,application/**/*.json,application/**/*.yaml` | `.cursor/rules/tech-startup-commands.mdc` |
| **arch-services-architecture.mdc** | ❌ | 服务层架构、模块设计、技术栈、设计模式 | `services/**/*.go` | `.cursor/rules/arch-services-architecture.mdc` |
| **tech-hexlib-cgo.mdc** | ✅ | CGO 集成、C++ 库开发、构建流程、接口设计 | `@cpp/hexlib/**/*.{go,cpp,h,make}` | `.cursor/rules/tech-hexlib-cgo.mdc` |
| **api-generation-standards.mdc** | ❌ | API 开发、代码生成、控制器设计、接口拆分 | `application/api/**/*.go,application/internal/controller/**/*.go` | `.cursor/rules/api-generation-standards.mdc` |
| **code-enum-creation-standards.mdc** | ❌ | 枚举定义、String 方法生成、JSON序列化 | `common/c_enum/*.go,common/c_enums/*.go` | `.cursor/rules/code-enum-creation-standards.mdc` |
| **code-config-struct-tags.mdc** | ❌ | 配置结构体、标签规范、动态标签解析 | `*config*.go,*Config*.go,common/c_base/**/*.go,plugins/**/*config*.go` | `.cursor/rules/code-config-struct-tags.mdc` |
| **arch-unified-point-system.mdc** | ✅ | 点位系统设计、数据映射、构造函数模式 | `common/c_base/*point*.go,common/c_proto/**/*.go,common/c_default/**/*.go,plugins/plug_drivers/**/*.go` | `.cursor/rules/arch-unified-point-system.mdc` |
| **tech-tsdb-storage.mdc** | ❌ | 时序数据库存储、TSDB 插件开发 | `plugins/plug_storages/p_tsdb/**/*.go,plugins/plug_storages/**/*.go,common/c_base/*storage*.go` | `.cursor/rules/tech-tsdb-storage.mdc` |
| **biz-policy-management.mdc** | ❌ | 策略管理、策略引擎设计、策略插件开发 | `plugins/plug_policy/**/*.go,services/s_policy/**/*.go,common/c_base/*policy*.go` | `.cursor/rules/biz-policy-management.mdc` |
| **biz-setting-system.mdc** | ❌ | 系统设置、配置管理、SSystemSettingDefine 使用 | `services/s_db/**/*setting*.go,services/s_db/**/*Setting*.go` | `.cursor/rules/biz-setting-system.mdc` |
| **biz-price-management.mdc** | ❌ | 电价管理、价格策略、电价配置 | `services/s_price/**/*.go,application/api/price/**/*.go` | `.cursor/rules/biz-price-management.mdc` |
| **api-test-login.mdc** | ❌ | API 测试、登录验证、认证流程 | `application/api/**/*.go,application/internal/controller/**/*.go,*.md,*.yaml,*.json,*.sh` | `.cursor/rules/api-test-login.mdc` |
| **mgt-cursor-rules.mdc** | ✅ | Cursor规则文件创建、管理和使用规范 | `.cursor/rules/**/*.mdc` | `.cursor/rules/mgt-cursor-rules.mdc` |
| **mgt-cursor-rules-naming.mdc** | ✅ | Cursor规则文件命名规范、分类体系、文件组织 | `.cursor/rules/**/*.mdc` | `.cursor/rules/mgt-cursor-rules-naming.mdc` |

## 三阶段工作流 (AGENTS.md)

项目遵循严格的三阶段工作流程。启动任何工作前，参考 AGENTS.md 文件中的规则：

1. **【分析问题】**: 理解需求、搜索代码、识别根因
2. **【制定方案】**: 列出变更、消除重复、确保架构一致
3. **【执行方案】**: 严格实现、运行类型检查、不提交代码（除非明确要求）

## Web 服务访问

启动 Web 服务后访问以下地址：
- **主界面**: http://localhost:8000
- **API 文档**: http://localhost:8000/api.json

## 环境变量

```bash
# Go 模块配置
GO111MODULE=on
GOPROXY=https://goproxy.cn,direct
GOSUMDB=sum.golang.google.cn

# CGO 配置
CGO_ENABLED=1
LD_LIBRARY_PATH=/app/lib

# EMS 特定配置
EMS_PROFILE=dev
EMS_DRIVER_PATH=resources/drivers
EMS_WEB_ENABLED=true
```

## Docker 开发

```bash
# 构建开发镜像
docker-compose -f script/docker-compose.dev.yml build

# 运行开发容器
docker-compose -f script/docker-compose.dev.yml up -d ems-dev

# 进入容器
docker exec -it ems-plan-dev /bin/bash
```

## 常见问题排查

### 构建失败
```bash
# 检查 CGO 是否启用
go env CGO_ENABLED

# 检查编译器
clang++ --version  # macOS
g++ --version      # Linux

# 清理并重新构建 C++ 库
cd @cpp/hexlib && make clean && make build
```

### 驱动构建失败
```bash
# 检查驱动列表
cd plugins/plug_drivers && make list

# 查看驱动状态
make status

# 查看完整输出
make project=gpio_basic 2>&1 | head -50
```

### 程序无法启动
```bash
# 检查 PID 文件
cat out/ems.pid

# 检查端口占用
lsof -i :8000

# 查看日志
tail -f out/logs/app/*.log
```

## 相关文档

- **README.md**: 项目完整说明文档
- **AGENTS.md**: 三阶段工作流规则
- **script/docker-compose.dev.yml**: Docker 开发配置
- **script/Dockerfile.dev**: Docker 镜像定义
