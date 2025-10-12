# Modbus TCP 数据导出服务 (s_export_modbus)

## 概述

Modbus TCP 数据导出服务负责将系统中的设备数据通过 Modbus TCP 协议对外提供，允许外部设备通过 Modbus 协议读取系统中的设备数据。

## 功能特性

- **多设备支持**: 支持多个设备同时对外提供 Modbus 服务
- **独立寻址**: 每个设备使用独立的 Modbus 从站ID和寄存器地址空间
- **固定点位**: 每个设备前两个点位为系统保留（在线状态、时间戳）
- **动态映射**: 根据设备配置动态构建寄存器映射关系
- **地址验证**: 自动检测地址重叠，防止冲突
- **实时数据**: 从设备缓存中获取实时数据
- **错误处理**: 设备启动失败不影响其他设备运行

## 配置格式

Modbus 配置存储在数据库的 `modbus_config` 设置中，JSON 格式如下：

```json
{
  "enabled": true,
  "listenPort": 502,
  "deviceIds": ["pylon_bms", "pcs_1"]
}
```

### 配置字段说明

- `enabled`: 是否启用 Modbus 服务（true/false）
- `listenPort`: Modbus TCP 服务器监听端口（默认 502）
- `deviceIds`: 要对外提供 Modbus 服务的设备ID列表

### 设备配置

每个设备需要在 `ExternalParam` 中配置 Modbus 相关参数：

```json
{
  "modbusId": 1,
  "modbusRegisterAddr": 40000,
  "modbusAllowControl": false
}
```

- `modbusId`: Modbus 从站ID（1-255）
- `modbusRegisterAddr`: 寄存器起始地址
- `modbusAllowControl`: 是否允许控制（暂未实现）

## 寄存器映射规则

### 固定点位

每个设备的前两个点位是系统保留的：

| 寄存器地址 | 功能 | 数据类型 | 说明 |
|-----------|------|---------|------|
| 起始地址 + 0 | 设备在线状态 | Bool | 0=离线，1=在线 |
| 起始地址 + 1-2 | 通讯时间戳 | Uint32 | Unix时间戳（秒） |

### 设备点位

从第三个寄存器开始，按顺序映射 `GetExportModbusPoints()` 返回的点位：

| 数据类型 | 寄存器数量 | 说明 |
|---------|----------|------|
| EBool, EInt8, EUint8 | 1 | 8位数据 |
| EInt16, EUint16 | 1 | 16位数据 |
| EInt32, EUint32, EFloat32 | 2 | 32位数据 |
| EInt64, EUint64, EFloat64 | 4 | 64位数据 |
| EString | **跳过** | 不支持字符串类型 |

### 地址空间管理

- **最小地址间距**: 固定100个寄存器
- **最大寄存器数**: 每个设备不允许超过100个寄存器
- **地址重叠检测**: 相同ModbusId的设备如果地址重叠，两个设备都标记为启动失败
- **失败处理**: 设备启动失败后不退出循环，继续处理其他设备

## 使用方法

### 1. 初始化服务

```go
import "s_export_modbus"

// 初始化Modbus服务
s_export_modbus.Init()
```

### 2. 启动服务

```go
// 启动Modbus服务
err := s_export_modbus.StartModbus(ctx)
if err != nil {
    log.Fatalf("启动Modbus服务失败: %v", err)
}
```

### 3. 停止服务

```go
// 停止Modbus服务
err := s_export_modbus.StopModbus(ctx)
if err != nil {
    log.Errorf("停止Modbus服务失败: %v", err)
}
```

### 4. 重新加载配置

```go
// 重新加载Modbus配置
err := s_export_modbus.ReloadModbus(ctx)
if err != nil {
    log.Errorf("重新加载Modbus配置失败: %v", err)
}
```

### 5. 获取服务状态

```go
// 获取Modbus服务状态
isRunning, port, deviceCount := s_export_modbus.GetModbusStatus()
fmt.Printf("服务运行状态: %v, 监听端口: %d, 设备数量: %d\n", isRunning, port, deviceCount)
```

## 架构设计

### 核心组件

1. **SModbusConfig**: Modbus 配置结构体
2. **SModbusManager**: Modbus 管理器（单例）
3. **SModbusServer**: Modbus TCP 服务器
4. **SModbusDeviceHandler**: 设备处理器
5. **SModbusRegisterMapper**: 寄存器映射工具

### 数据流程

1. 从数据库加载配置
2. 遍历配置的设备列表
3. 获取设备的 `ExternalParam` 配置
4. 构建寄存器映射关系
5. 验证地址空间不重叠
6. 启动 Modbus TCP 服务器
7. 处理客户端请求，返回实时数据

## 错误处理

- 设备不存在：返回 Modbus 异常码 0x0B (Gateway Target Device Failed to Respond)
- 地址越界：返回 Modbus 异常码 0x02 (Illegal Data Address)
- 数据转换失败：返回异常码 0x04 (Slave Device Failure)
- 地址重叠：设备启动失败，记录错误日志

## 日志记录

服务会记录以下日志：

- Modbus 服务启动/停止
- 设备映射构建成功/失败
- 地址重叠检测结果
- 客户端连接和请求处理
- 错误和异常信息

## 依赖

- `github.com/grid-x/modbus`: Modbus 协议库
- `common`: 通用库（设备管理、日志等）
- `s_db`: 数据库服务（配置读取）

## 注意事项

1. **地址空间管理**: 确保相同ModbusId的设备地址不重叠
2. **数据类型支持**: 不支持字符串类型，会自动跳过
3. **实时数据**: 数据从设备缓存中获取，确保设备正常运行
4. **错误恢复**: 单个设备失败不影响其他设备运行
5. **线程安全**: 支持并发访问，使用读写锁保护共享数据
