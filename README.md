

# EMS

一个面向分布式能源场景的轻量级能量管理系统（Energy Management System, EMS）。本项目提供设备驱动、通信协议、数据存储与推送的插件化能力，可快速对接 PCS（变流器/逆变器）、BMS（电池管理系统）、电表、消防等各类设备，完成数据采集、状态监控、策略控制与数据持久化。

## 项目简介
- 插件化架构：驱动（Drivers）、协议（Protocols）、存储（Storages）、推送（Push）均为可插拔模块，便于扩展与替换。
- 多协议支持：内置 Modbus(TCP/RTU)、CANBus（含 UDP 透传）、GPIO(Sysfs) 等。
- 设备生态：已包含多类设备驱动（PCS/BMS/电表/消防/GPIO 等），可通过清单/数据库配置启停与参数化。
- 数据存储：支持 InfluxDB 1.x/2.x；系统也内置 PebbleDB/SQLite 用于轻量存储与配置管理。
- 数据推送：支持 MQTT 等方式将运行数据/事件上报到上层平台。
- Web 开关：可选择是否启用内置 Web 服务以便集成或调试。

## 架构与运行时
应用入口位于 application/main.go，使用 gogf/gf 命令与服务管理：
- application/internal/cmd/cmd.go 定义了主命令与参数，启动时会：
  - 初始化 SQLite 配置实例与 PebbleDB 存储。
  - 加载协议配置与设备配置，按配置树递归初始化设备及其子设备。
  - 按需初始化协议提供者（Modbus/CANBus/GPIO/Virtual）。
  - 注册设备并按配置启用数据存储、推送等功能。

## 快速开始
1. 环境准备：
   - Go 1.20+（建议）
   - 可选：InfluxDB 1.x/2.x（用于时间序列存储）、MQTT Broker
2. 运行（示例）：
   - 直接运行：
     go run ./application --web=true --driver-path=resources/driver
   - 常用参数（均为可选）：
     - --device-name         设备配置名（默认：device）
     - --driver-path         驱动文件路径（默认：./driver）
     - --web                 是否启用内置 Web（默认：false）
     - --language            全局语言（默认：zh-CN）
     - --time-zone           全局时区
     - --pebbledb-path       PebbleDB 路径（默认：./out/pebbledb）
     - --db-path         SQLite 路径（默认：./out/db.sqlite3）

3. Docker 构建与运行：
   - 构建镜像：
     docker build -t ems-go .
   - 运行容器（挂载当前目录以便调试）：
     docker run -v .:/root/work -ti ems-go /bin/bash

## 开发调试运行
```bash
cd appllication
gf run main.go -p ./bin -a "--web=true" -w "api/*.go" -w "internal/*.go"
```

## 目录说明（节选）
- application/             应用入口与命令初始化
- common/                  通用基础能力（配置、枚举、接口、工具）
- plug_drivers/            设备驱动插件（PCS/BMS/电表/消防/GPIO 等）
- plug_protocols/          协议插件（modbus/canbus/gpio_sysfs 等）
- plug_storages/           存储插件（influxdb_1/influxdb_2/pebbledb/sqlite）
- plug_push/               推送插件（mqtt）
- services/                服务编排层（设备启动/停止、协议装配等）
- application/resources/   资源文件（如默认驱动/清单）

## 配置与数据
- 配置来源：
  - 系统使用 internal_config_with_sqlite 初始化配置（默认 ./out/db.sqlite3），也保留了基于文件的读取实现以便迁移/调试。
- 数据存储：
  - 可启用 InfluxDB 1.x/2.x 进行时间序列存储；也支持 PebbleDB/SQLite 用于轻量级场景。
- 推送：
  - 通过 MQTT 将关键数据/事件上报到上层平台。

## 命名规范
### Golang类型缩写
Enums: 
>E,e

Interface: 
>I,i

Struct: 
>S,s

Function: 
>F,f


### 文件命名
包名：{项目第一个字母小写}_包名 
>PS: base/device -> device需要改名为 b_device

文件名：{大类型}_{小类型(可选)...}_{Golang类型缩写(小写)}_{Golang类型缩写(小写)(可选)...}.go  
>PS: type_e.go type_pcs_i.go type_pcs_exec_f.go


### 代码命名
枚举:
>{Golang类型缩写(大写)}{Name}

接口:
>{Golang类型缩写(大写)}{Name}

结构体:
> {Golang类型缩写(大写)}{Name}


## 运行
### 参数

--driver-path=resources/driver --web=true

## 编译
### 创建镜像
docker build -t ems-go .

### 启动容器
docker run -v .:/root/work -ti ems-go /bin/bash


## 环境部署
### influxdb1
https://docs.influxdata.com/influxdb/v1/introduction/install/

```bash
sudo apt install influxdb-client
```

### influxdb2
https://docs.influxdata.com/influxdb/v2/install/#install-influxdb-as-a-service-with-systemd