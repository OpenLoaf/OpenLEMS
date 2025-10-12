# Modbus TCP 服务器实现总结

## 实现概述

成功实现了 `s_export_modbus` 服务，提供 Modbus TCP 服务器功能，允许外部设备通过 Modbus 协议读取系统中的设备数据。

## 核心特性

### 1. 寄存器映射规则
- **固定点位**: 每个设备前两个点位为系统保留
  - 第1个寄存器：设备在线状态（Bool，1个寄存器）
  - 第2-3个寄存器：通讯时间戳（Uint32，2个寄存器）
- **设备点位**: 从第4个寄存器开始，按顺序映射 `GetExportModbusPoints()` 返回的点位
- **数据类型支持**: 
  - Bool/Int8/Uint8: 1个寄存器
  - Int16/Uint16: 1个寄存器  
  - Int32/Uint32/Float32: 2个寄存器
  - Int64/Uint64/Float64: 4个寄存器
  - String: 跳过（不支持）

### 2. 地址空间管理
- **最小地址间距**: 固定100个寄存器
- **最大寄存器数**: 每个设备不允许超过100个寄存器
- **地址重叠检测**: 相同ModbusId的设备如果地址重叠，两个设备都标记为启动失败
- **失败处理**: 设备启动失败后不退出循环，继续处理其他设备

### 3. 多设备支持
- **一对多关系**: 一个ModbusId可以对应多个设备
- **独立寻址**: 通过寄存器起始地址区分不同设备
- **状态查询**: 可以查询所有设备状态（包括失败设备）及失败原因

## 技术实现

### 文件结构
```
services/s_export_modbus/
├── go.mod                                    # 模块依赖
├── README.md                                 # 服务文档
├── s_modbus_f.go                            # 导出接口
├── internal/
│   ├── s_modbus_config_s.go                 # 配置结构体
│   ├── s_modbus_manager_s.go                # 管理器（单例）
│   ├── s_modbus_server_s.go                 # Modbus TCP 服务器
│   ├── s_modbus_device_handler_s.go         # 设备处理器
│   ├── s_modbus_register_mapper_f.go        # 寄存器映射工具
│   └── s_modbus_register_mapper_f_test.go    # 单元测试
└── IMPLEMENTATION_SUMMARY.md                # 实现总结
```

### 核心组件

#### 1. SModbusManager（管理器）
- 单例模式管理 Modbus 服务
- 从数据库加载配置
- 构建所有设备的寄存器映射
- 启动/停止/重载服务
- 提供状态查询接口

#### 2. SModbusServer（服务器）
- 启动 Modbus TCP 服务器（监听指定端口）
- 处理客户端连接和请求
- 支持优雅停止
- 实现完整的 Modbus TCP 协议栈

#### 3. SModbusDeviceHandler（设备处理器）
- 处理 Modbus 读保持寄存器请求（功能码 0x03）
- 根据从站ID和寄存器地址查找对应设备和点位
- 从设备缓存中获取实时数据并返回
- 支持固定点位和设备点位的处理

#### 4. SModbusRegisterMapper（寄存器映射工具）
- 计算数据类型占用寄存器数
- 构建设备寄存器映射
- 值转寄存器和寄存器转值
- 地址重叠检测

### 数据流程

1. **配置加载**: 从数据库 `SystemSettingModbusConfig` 读取配置
2. **设备映射**: 遍历配置的设备列表，构建设备寄存器映射
3. **地址验证**: 检查地址重叠，标记冲突设备
4. **服务启动**: 启动 Modbus TCP 服务器
5. **请求处理**: 处理客户端请求，返回实时数据

## 集成到主程序

### 1. 服务初始化
在 `application/internal/cmd/ems.go` 中添加：
```go
// 初始化Modbus服务
s_export_modbus.Init()
```

### 2. 服务启动
在 `StartServices` 函数中添加：
```go
// 启动Modbus服务
go func() {
    // 等待设备管理器启动完成
    time.Sleep(4 * time.Second)

    err := s_export_modbus.StartModbus(ctx)
    if err != nil {
        g.Log().Errorf(ctx, "启动Modbus服务失败: %+v", err)
    } else {
        c_log.BizInfof(ctx, "Modbus服务启动成功！")
    }
}()
```

### 3. 服务停止
在 `SetupShutdownHandler` 函数中添加：
```go
// 停止Modbus服务
err = s_export_modbus.StopModbus(ctx)
if err != nil {
    g.Log().Errorf(ctx, "停止Modbus服务失败: %+v", err)
} else {
    g.Log().Infof(ctx, "Modbus服务已停止")
    c_log.BizInfof(ctx, "Modbus服务已停止")
}
```

## 配置说明

### 数据库配置
```json
{
  "enabled": true,
  "listenPort": 502,
  "deviceIds": ["pylon_bms", "pcs_1"]
}
```

**配置字段说明**：
- `enabled`: 是否启用 Modbus 服务（true/false），当设置为 false 时，服务不会启动
- `listenPort`: Modbus TCP 服务器监听端口
- `deviceIds`: 要暴露的设备ID列表

### 设备配置
每个设备需要在 `ExternalParam` 中配置：
```json
{
  "modbusId": 1,
  "modbusRegisterAddr": 40000,
  "modbusAllowControl": false
}
```

## 测试验证

### 单元测试
- ✅ `TestCalculateRegisterCount`: 测试寄存器数量计算
- ✅ `TestEncodeValueToRegisters`: 测试值转寄存器
- ✅ `TestDecodeRegistersToValue`: 测试寄存器转值
- ✅ `TestCheckAddressOverlap`: 测试地址重叠检测

### 构建测试
- ✅ 项目构建成功
- ✅ 依赖关系正确
- ✅ 无编译错误

## 使用示例

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

### 3. 获取服务状态
```go
// 获取Modbus服务状态
isRunning, port, deviceCount := s_export_modbus.GetModbusStatus()
fmt.Printf("服务运行状态: %v, 监听端口: %d, 设备数量: %d\n", isRunning, port, deviceCount)
```

### 4. 获取设备状态
```go
// 获取所有设备状态
deviceStatusList := s_export_modbus.GetModbusDeviceStatus()
for _, status := range deviceStatusList {
    fmt.Printf("设备: %s, ModbusId: %d, 状态: %v, 错误: %s\n", 
        status.DeviceId, status.ModbusId, status.IsOnline, status.Error)
}
```

## 技术要点

### 1. 线程安全
- 使用 `sync.RWMutex` 保护共享数据
- 设备映射表支持并发读取
- 配置重载时加写锁

### 2. 数据获取
- 从设备缓存获取点位值
- 支持实时数据查询
- 处理设备离线情况

### 3. 字节序处理
- Modbus 使用大端字节序（Big Endian）
- 多寄存器数据需要正确拼接
- 使用 `encoding/binary` 标准库处理

### 4. 错误处理
- 设备不存在：返回 Modbus 异常码 0x0B
- 地址越界：返回 Modbus 异常码 0x02
- 数据转换失败：返回异常码 0x04

## 注意事项

1. **地址空间管理**: 确保相同ModbusId的设备地址不重叠
2. **数据类型支持**: 不支持字符串类型，会自动跳过
3. **实时数据**: 数据从设备缓存中获取，确保设备正常运行
4. **错误恢复**: 单个设备失败不影响其他设备运行
5. **线程安全**: 支持并发访问，使用读写锁保护共享数据

## 后续优化建议

1. **性能优化**: 考虑使用连接池管理客户端连接
2. **监控增强**: 添加更详细的性能监控和统计
3. **配置热更新**: 支持运行时配置更新
4. **日志优化**: 添加更详细的调试日志
5. **错误恢复**: 增强设备故障恢复机制

## 总结

成功实现了完整的 Modbus TCP 服务器功能，满足所有需求：

- ✅ 支持多设备独立寻址
- ✅ 实现固定点位和设备点位映射
- ✅ 支持地址重叠检测和冲突处理
- ✅ 提供完整的服务生命周期管理
- ✅ 集成到主程序启动流程
- ✅ 通过单元测试验证
- ✅ 项目构建成功

该实现为系统提供了强大的 Modbus 数据导出能力，允许外部设备通过标准 Modbus 协议访问系统数据。
