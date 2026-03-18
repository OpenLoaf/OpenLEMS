<p align="center">
  <img src="docs/images/lems-logo.png" alt="LEMS Logo" width="112" />
</p>

<h1 align="center">LEMS</h1>

<p align="center">
  Local Energy Management System · 本地能源管理系统
</p>

<p align="center">
  面向分布式能源、微电网与储能柜场景的轻量级本地 EMS 运行平台
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25%2B-00ADD8?logo=go&logoColor=white" alt="Go 1.25+" />
  <img src="https://img.shields.io/badge/License-AGPL--3.0--only-0E8A16" alt="License AGPL-3.0-only" />
  <img src="https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20ARM-1F6FEB" alt="Platform Linux macOS ARM" />
</p>

<p align="center">
  <a href="https://lems.hexems.com/">在线体验</a>
  ·
  <a href="docs/README.en.md">English README</a>
  ·
  <a href="#快速开始">快速开始</a>
  ·
  <a href="#界面预览">界面预览</a>
  ·
  <a href="CONTRIBUTING.md">参与贡献</a>
</p>

`LEMS` 是一套面向分布式能源、微电网、储能柜和边缘能源网关场景的本地能源管理系统。它把 PCS、BMS、电表、光伏、负荷、GPIO 等异构设备接入同一套运行时中，通过插件化的驱动、协议、策略和数据服务，完成设备接入、实时监控、告警处理、策略控制、自动化联动、电价管理以及对外数据集成。

它不是一个只看图表的“监控页面”，而是一套可以真实部署在边缘控制器、工业网关、工控机或现场服务器上的 EMS 运行平台。

## 项目预览

<p align="center">
  <img src="docs/images/energy-system-overview.png" alt="LEMS 能源总览" width="88%" />
</p>

## 项目简介

在一个典型的分布式能源项目里，现场常常同时存在：

- 不同厂家的 PCS、BMS、电表、光伏、负荷等设备
- 不同通信方式和工业协议
- 不同控制逻辑、运行约束和策略目标
- 需要上送到上层平台的遥测、告警和运行数据

`LEMS` 解决的正是“设备多、协议杂、数据散、控制难”的问题。它把这些能力收敛到一套统一系统里：

- 向下对接设备：通过驱动和协议插件接入现场设备
- 向中统一模型：把设备点位、状态、告警、策略放到统一运行时
- 向上提供能力：通过 Web、MQTT、Modbus TCP、日志和统计接口对外服务

## 典型架构

### 储能模式架构

下图展示了储能场景下的本地 EMS 组网方式，可用于储能柜、多柜联机和边缘采集网关部署。

<p align="center">
  <img src="docs/images/energy-storage-architecture.png" alt="储能模式架构图" width="88%" />
</p>

### 微电网模式架构

下图展示了微电网场景下风、光、储、荷与上层触控屏、云平台、APP 的协同接入方式。

<p align="center">
  <img src="docs/images/microgrid-architecture.png" alt="微电网模式架构图" width="88%" />
</p>

## 界面预览

### 储能柜详情

柜级详情用于展示储能柜内部拓扑、运行状态与趋势图。

<p align="center">
  <img src="docs/images/energy-system-cabinet-detail.png" alt="储能柜详情" width="80%" />
</p>

### 功能入口与实时数据

系统提供统一功能入口；实时数据页面向运维人员展示设备状态、子设备结构与关键测点。

<p align="center">
  <img src="docs/images/programs-feature-menu.png" alt="功能入口" width="48%" />
  <img src="docs/images/realtime-data.png" alt="实时数据" width="48%" />
</p>

### 告警与日志

告警页适合查看当前故障与历史问题，日志页适合定位通信、推送与运行过程中的细节。

<p align="center">
  <img src="docs/images/alarm-list.png" alt="告警列表" width="48%" />
  <img src="docs/images/log-viewer.png" alt="日志查看" width="48%" />
</p>

### 统计分析与远程管理

统计分析页用于设备对比与趋势分析；远程管理页可管理 MQTT / Modbus 等对外服务。

<p align="center">
  <img src="docs/images/statistics-analysis.png" alt="统计分析" width="48%" />
  <img src="docs/images/remote-management-mqtt.png" alt="远程管理 MQTT" width="48%" />
</p>

<p align="center">
  <img src="docs/images/remote-management-modbus.png" alt="远程管理 Modbus" width="48%" />
  <img src="docs/images/automation-control.png" alt="自动化控制" width="48%" />
</p>

### 电价、协议、驱动与策略

这些页面体现了系统不仅能“看”，还能“配”“控”“管”。

<p align="center">
  <img src="docs/images/price-management.png" alt="电价管理" width="48%" />
  <img src="docs/images/protocol-configuration.png" alt="协议配置" width="48%" />
</p>

<p align="center">
  <img src="docs/images/driver-manager.png" alt="驱动管理" width="48%" />
  <img src="docs/images/strategy-management.png" alt="策略管理" width="48%" />
</p>

### 系统运行状态

系统页可查看 CPU、内存、磁盘、网络流量和运行时长等信息，适合部署后的运维巡检。

<p align="center">
  <img src="docs/images/system-management.png" alt="系统管理" width="80%" />
</p>

## 核心能力

- 插件化驱动体系：便于扩展不同品牌、型号和物模型的现场设备
- 插件化协议体系：当前仓库已包含 `Modbus(TCP/RTU)`、`CANBus`、`GPIO` 等能力
- 实时监控与告警：支持设备状态、功率、电量、SOC、时间戳、告警记录和历史日志查看
- 策略与自动化：支持微电网策略、储能策略、自动化任务和本地控制参数配置
- 电价与成本优化：支持时段电价配置，便于削峰填谷、峰谷套利和成本管理
- 对外集成能力：支持 MQTT 数据上送、Modbus TCP 数据导出和远程服务管理
- 边缘部署友好：适合部署在工业网关、工控机、边缘服务器等本地环境
- 内置 Web 管理界面：构建完成后可直接访问统一管理页面，并针对平板触控场景做了交互优化

## 适用场景

- 储能柜 EMS：单柜运行、柜内设备联动、状态监控与功率控制
- 多柜联机 EMS：边缘网关统一管理多套储能系统
- 微电网 EMS：风、光、储、荷等设备统一接入和协调调度
- 集成网关：将现场异构设备数据整理后上送第三方平台
- 边缘运维平台：提供告警、日志、自动化任务、电价管理和系统状态查看

## 已包含的代表性组件

### 设备驱动

- `Pylon Tech US108` BMS
- `Pylon Checkwatt` 储能相关驱动
- `Elecod MAC` PCS
- `Star Charge 100E` PCS
- `GPIO In/Out` 基础驱动
- `Ammeter / PV / ESS / Load` 演示驱动

### 协议与服务

- `plugins/plug_protocols/p_modbus`
- `plugins/plug_protocols/p_canbus`
- `plugins/plug_protocols/p_gpiod`
- `services/s_export_mqtt`
- `services/s_export_modbus`
- `services/s_automation`
- `services/s_price`
- `services/s_policy`

## 算法与策略能力

当前仓库里的“算法”并不只是几条定时规则，而是由时段命中、策略参数模型、策略执行框架和预测算法接口共同组成。

- 时段与优先级命中：电价服务会先过滤启用配置，再按日期范围、时间范围和优先级选出当前生效的电价；跨天区间也能正确命中
- 本地策略参数模型：已暴露 `SOC` 上下限、充放电效率、变压器容量、安全系数、需量控制、动态扩容等典型 EMS 决策参数
- 储能策略执行框架：支持日期范围、时间范围和按小时目标点位配置，执行时可结合设备列表、设备类型和当前状态进行控制分发
- 自动化联动机制：支持按时间和条件触发任务，将策略判断与现场动作联动起来
- 预测算法扩展接口：`@cpp/hexlib` 提供基于 `C++ / CGO` 的 `MPC + Kalman` 预测器接口，适合用于功率、负荷、发电量等时间序列预测

当前更准确的表述是：`LEMS` 已具备算法基础、策略框架和可扩展的预测能力，并为更复杂的微电网闭环优化和储能调度预留了扩展接口。

## 在线体验

- 在线体验地址：[https://lems.hexems.com/](https://lems.hexems.com/)
- 管理员密码：`888888`

建议使用平板横屏模式体验，当前项目的界面布局、触控操作和运维主流程已针对平板场景做了优化。

## 快速开始

公开仓库克隆后，可以直接在 `application/` 目录执行 `go build`。

### 环境要求

- Go `1.25+`
- 建议使用 Linux / macOS 开发环境
- 若启用 CGO 相关能力，建议安装 `GCC` 或 `Clang`

### 1. 克隆仓库

```bash
git clone git@github.com:OpenLoaf/OpenLEMS.git
cd OpenLEMS
```

### 2. 同步工作区依赖

```bash
go work sync
```

### 3. 构建主程序

```bash
cd application
go build -o ./bin/lems .
```

如果你想一次性校验整个 Go Workspace，可以在仓库根目录执行：

```bash
./script/build-workspace.sh
```

### 4. 启动 Web 管理界面

```bash
cd application
go run . --web=true --driver-path=resources/driver
```

启动后默认可访问：

- `http://localhost:8000`

### 5. 演示模式启动

```bash
cd application
go run . --web=true --demo=true --driver-path=resources/driver
```

## 常用启动参数

| 参数 | 简写 | 说明 |
| --- | --- | --- |
| `--web` | `-w` | 启用 Web 服务 |
| `--demo` | - | 标记为演示模式 |
| `--driver-path` | `-dp` | 驱动文件目录 |
| `--device-name` | `-d` | 设备配置名称 |
| `--runtime-path` | `-rp` | 运行时数据库路径 |
| `--db-path` | `-cp` | 配置数据库路径 |
| `--language` | `-l` | 全局语言 |
| `--profile` | - | 配置环境，如 `dev` / `prod` |
| `--force` | - | 忽略 PID 检查强制启动 |

## 仓库结构

```text
OpenLEMS/
├── application/             # 主程序入口、API、控制器、Web 服务
├── common/                  # 通用接口、设备/协议基础类型、工具能力
├── plugins/
│   ├── plug_drivers/        # 设备驱动插件
│   ├── plug_protocols/      # 通信协议插件
│   ├── plug_policy/         # 策略插件
│   ├── plug_push/           # 推送插件
│   └── plug_storages/       # 存储插件
├── services/                # 服务编排层，如驱动、自动化、电价、远程导出
├── docs/                    # 项目说明与截图资源
└── script/                  # 工作区构建与校验脚本
```

## 开发与扩展

### 构建整个工作区

```bash
./script/build-workspace.sh
```

### 驱动开发

驱动采用插件模式组织。可以优先参考以下目录中的已有实现：

- `plugins/plug_drivers/bms/`
- `plugins/plug_drivers/pcs/`
- `plugins/plug_drivers/ess/`
- `plugins/plug_drivers/gpio/`
- `plugins/plug_drivers/demo/`

### 协议与服务开发

如果你想扩展协议或数据对接能力，建议优先阅读：

- `plugins/plug_protocols/p_modbus`
- `plugins/plug_protocols/p_canbus`
- `services/s_export_mqtt`
- `services/s_export_modbus`
- `services/s_automation`
- `services/s_price`

## 相关文档

- [软件定位与能力概述](docs/ems-plan-software-summary.md)
- [Remote API 文档](application/api/remote/README.md)
- [自动化服务说明](services/s_automation/README.md)
- [MQTT 数据推送说明](services/s_export_mqtt/README.md)
- [Modbus TCP 数据导出说明](services/s_export_modbus/README.md)

## 参与贡献

欢迎提交 Issue 和 Pull Request。

提交前建议先阅读：

- [贡献指南](CONTRIBUTING.md)

本地校验命令：

```bash
./script/build-workspace.sh
./script/test-workspace.sh
```

## 商务合作

- 合作邮箱：`zz@hexems.com`

如需项目合作、方案咨询、私有化部署或技术交流，可通过邮箱联系，或扫码添加微信：

<p align="center">
  <img src="docs/images/business-wechat-qr.png" alt="商务合作微信二维码" width="240" />
</p>

## License

本项目采用 [AGPL-3.0-only](LICENSE) 许可证。
