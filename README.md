

# EMS Plan - 分布式能源管理系统

一个面向分布式能源场景的轻量级能量管理系统（Energy Management System, EMS）。本项目提供设备驱动、通信协议、数据存储与推送的插件化能力，可快速对接 PCS（变流器/逆变器）、BMS（电池管理系统）、电表、消防等各类设备，完成数据采集、状态监控、策略控制与数据持久化。

## 🚀 项目特性

- **🔌 插件化架构**：驱动（Drivers）、协议（Protocols）、存储（Storages）、推送（Push）、策略（Policy）均为可插拔模块，便于扩展与替换
- **🌐 多协议支持**：内置 Modbus(TCP/RTU)、CANBus（含 UDP 透传）、GPIO(Sysfs) 等工业通信协议
- **📱 设备生态丰富**：已包含多类设备驱动（PCS/BMS/电表/消防/GPIO 等），可通过清单/数据库配置启停与参数化
- **💾 多存储方案**：支持 InfluxDB 1.x/2.x、PebbleDB、SQLite 等多种存储方案
- **📡 数据推送**：支持 MQTT 等方式将运行数据/事件上报到上层平台
- **🌍 Web 管理界面**：内置 Web 服务提供设备管理、监控、控制等功能
- **🔧 策略引擎**：支持基础储能策略和微电网策略
- **📊 实时监控**：WebSocket 实时数据推送，支持遥测数据和电站状态监控

## 🏗️ 系统架构

### 项目结构
```
ems-plan/
├── application/           # 应用主程序
│   ├── api/              # API 接口定义
│   ├── internal/         # 内部实现
│   │   ├── cmd/          # 命令行管理
│   │   ├── controller/   # 控制器层
│   │   ├── model/        # 数据模型
│   │   └── ws/           # WebSocket 服务
│   ├── manifest/         # 配置和静态资源
│   └── main.go           # 程序入口
├── common/               # 通用基础能力
│   ├── c_base/          # 基础接口和结构
│   ├── c_device/        # 设备相关接口
│   ├── c_proto/         # 协议相关接口
│   └── c_util/          # 工具函数
├── plugins/              # 插件系统
│   ├── plug_drivers/    # 设备驱动插件
│   ├── plug_protocols/  # 通信协议插件
│   ├── plug_storages/   # 数据存储插件
│   ├── plug_push/       # 数据推送插件
│   └── plug_policy/     # 策略插件
└── services/            # 服务编排层
    ├── s_db/           # 数据库服务
    ├── s_driver/       # 驱动管理服务
    └── s_storage/      # 存储管理服务
```

### 核心组件

#### 设备驱动插件 (plug_drivers/)
- **PCS 驱动**：亿兰科 MAC/MDC、协能、享能、星星充电 100E
- **BMS 驱动**：派能科技 US108、协能
- **电表驱动**：安科瑞 10R
- **EMS 驱动**：高特、协能、派能科技 CheckWatt
- **其他设备**：消防控制、GPIO 基础驱动、储能站

#### 通信协议插件 (plug_protocols/)
- **Modbus**：支持 TCP/RTU 模式
- **CANBus**：支持 UDP 透传
- **GPIO Sysfs**：Linux GPIO 系统文件接口
- **通用协议框架**：c_protocol 提供协议抽象层

#### 数据存储插件 (plug_storages/)
- **InfluxDB 1.x/2.x**：时序数据库存储
- **PebbleDB**：轻量级键值存储
- **TSDB**：时间序列数据库
- **SQLite**：关系型数据库

#### 数据推送插件 (plug_push/)
- **MQTT**：消息队列遥测传输
- **通用推送框架**：c_push 提供推送抽象层

#### 策略插件 (plug_policy/)
- **基础储能策略**：basic_ess_policy
- **微电网策略**：basic_microgrid_policy

## 🚀 快速开始

### 环境要求
- **Go 1.24+**（推荐使用最新稳定版）
- **可选依赖**：
  - InfluxDB 1.x/2.x（时序数据存储）
  - MQTT Broker（数据推送）
  - Docker（容器化部署）

### 本地开发运行

1. **克隆项目**
```bash
git clone <repository-url>
cd ems-plan
```

2. **安装依赖**
```bash
go mod download
```

3. **开发模式运行**
```bash
cd application
gf run main.go -p ./bin -a "--web=true" -w "api/*.go" -w "internal/*.go"
```

4. **生产模式运行**
```bash
cd application
go run main.go --web=true --driver-path=resources/driver
```

### 常用启动参数

| 参数 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--web` | `-w` | `false` | 是否启用内置 Web 服务 |
| `--driver-path` | `-dp` | `./driver` | 驱动文件存放路径 |
| `--device-name` | `-d` | `device` | 设备配置文件名称 |
| `--runtime-path` | `-rp` | `./out/runtime` | 实时数据库路径 |
| `--db-path` | `-cp` | `./out/db.sqlite3` | 配置数据库路径 |
| `--language` | `-l` | `zh-CN` | 全局语言设置 |
| `--profile` | `-p` | `prod` | 配置环境 (default/dev/prod) |

### Docker 部署

1. **构建镜像**
```bash
docker build -t ems-plan .
```

2. **运行容器**
```bash
# 开发模式（挂载本地目录）
docker run -v .:/root/work -ti ems-plan /bin/bash

# 生产模式
docker run -d -p 8000:8000 -v ./config:/root/work/config ems-plan
```

## 🌐 Web 管理界面

启动 Web 服务后，可通过以下地址访问：

- **主界面**：http://localhost:8000
- **API 文档**：http://localhost:8000/api.json
- **Swagger UI**：http://localhost:8000/swagger/

### 主要功能模块

- **设备管理**：设备树查看、启用/停用设备、设备配置
- **遥测监控**：实时数据查看、历史数据查询
- **系统监控**：CPU、内存、磁盘、网络状态
- **网络配置**：网络接口管理、DNS 配置
- **设备控制**：设备指令下发、自定义服务调用
- **日志查看**：系统日志、业务日志查询

## 🔧 开发指南

### 项目构建

使用 Go Workspace 管理多模块项目：

```bash
# 查看工作空间状态
go work use .

# 构建所有模块
go build ./...

# 运行测试
go test ./...
```

### 驱动开发

1. **创建新驱动**
```bash
cd plugins/plug_drivers
mkdir my_device
cd my_device
go mod init my_device
```

2. **实现驱动接口**
```go
package main

import (
    "common/c_device"
    "common/c_base"
)

type MyDeviceDriver struct {
    // 实现设备驱动接口
}

func (d *MyDeviceDriver) Init(config *c_base.SDeviceConfig) error {
    // 初始化逻辑
    return nil
}

func (d *MyDeviceDriver) Start() error {
    // 启动逻辑
    return nil
}

// 实现其他必要接口...
```

3. **构建驱动**
```bash
# 使用 Makefile 构建
make project=my_device

# 或手动构建
go build -buildmode=plugin -o my_device_v1.driver
```

### API 开发

项目使用 GoFrame 框架，支持自动生成 API 代码：

```bash
# 生成控制器接口
gf gen ctrl

# 生成数据访问层
gf gen dao

# 生成服务层
gf gen service
```

## 📊 数据存储

### 配置数据库 (SQLite)
- **用途**：存储设备配置、协议配置、系统设置
- **默认路径**：`./out/db.sqlite3`
- **管理工具**：
```bash
# 导出数据
sqlite3 db.sqlite3 .dump > init.sql

# 查看表结构
sqlite3 db.sqlite3 ".schema"
```

### 时序数据库 (InfluxDB)
- **用途**：存储设备遥测数据、历史记录
- **版本支持**：InfluxDB 1.x 和 2.x
- **安装**：
```bash
# Ubuntu/Debian
sudo apt install influxdb-client

# 或使用 Docker
docker run -d -p 8086:8086 influxdb:1.8
```

## 🔍 监控与调试

### 日志系统
- **系统日志**：基于 GoFrame 的日志框架
- **业务日志**：按设备类型和 ID 分文件存储
- **日志级别**：DEBUG、INFO、WARNING、ERROR

### 性能监控
- **系统资源**：CPU、内存、磁盘使用率
- **网络流量**：实时网络 I/O 统计
- **设备状态**：设备连接状态、数据采集频率

### 故障排查
```bash
# 查看系统日志
tail -f logs/system.log

# 查看设备日志
tail -f logs/device_*.log

# 检查数据库连接
sqlite3 ./out/db.sqlite3 "SELECT * FROM devices;"
```

## 📝 命名规范

### Golang 类型缩写
- **枚举 (Enums)**：`E`, `e`
- **接口 (Interface)**：`I`, `i`
- **结构体 (Struct)**：`S`, `s`
- **函数 (Function)**：`F`, `f`

### 文件命名
- **包名**：`{项目首字母小写}_{包名}`
- **文件名**：`{大类型}_{小类型(可选)}_{Golang类型缩写(小写)}_{Golang类型缩写(小写)(可选)}.go`

### 代码命名
- **枚举**：`{Golang类型缩写(大写)}{Name}`
- **接口**：`{Golang类型缩写(大写)}{Name}`
- **结构体**：`{Golang类型缩写(大写)}{Name}`
