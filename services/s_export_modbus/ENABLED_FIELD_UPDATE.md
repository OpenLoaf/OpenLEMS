# Modbus 服务 Enabled 字段更新

## 更新概述

为 `s_export_modbus` 服务添加了 `enabled` 字段，允许通过配置控制 Modbus 服务的启动状态。

## 修改内容

### 1. 配置结构体更新

**文件**: `services/s_export_modbus/internal/s_modbus_config_s.go`

```go
// SModbusConfig Modbus配置结构体
type SModbusConfig struct {
	Enabled    bool     `json:"enabled"`     // 是否启用Modbus服务
	ListenPort int      `json:"listenPort"` // Modbus TCP服务器监听端口
	DeviceIds  []string `json:"deviceIds"`  // 设备ID列表
}
```

**新增字段**：
- `Enabled`: 布尔类型，控制是否启用 Modbus 服务

### 2. 服务启动逻辑更新

**文件**: `services/s_export_modbus/internal/s_modbus_manager_s.go`

#### 配置加载逻辑
```go
// 检查是否启用Modbus服务
if !config.Enabled {
    c_log.Info(ctx, "Modbus服务已禁用，跳过启动")
    // 清空现有设备映射
    m.deviceMaps = make(map[string]*SDeviceRegisterMap)
    m.deviceStatus = make(map[string]*SModbusDeviceStatus)
    return nil
}
```

#### 服务启动逻辑
```go
// 检查是否有设备映射（如果enabled为false，设备映射为空）
if len(m.deviceMaps) == 0 {
    c_log.Info(m.ctx, "Modbus服务未启用或无设备配置，跳过服务器启动")
    m.isRunning = true // 标记为运行状态，但实际不启动服务器
    return nil
}
```

#### 状态查询逻辑
```go
// 如果服务器为nil（服务被禁用），返回运行状态但连接数为0
if m.server == nil {
    return true, 0, len(m.deviceMaps)
}
```

### 3. 文档更新

**文件**: `services/s_export_modbus/README.md`

更新了配置示例和字段说明：

```json
{
  "enabled": true,
  "listenPort": 502,
  "deviceIds": ["pylon_bms", "pcs_1"]
}
```

**配置字段说明**：
- `enabled`: 是否启用 Modbus 服务（true/false）
- `listenPort`: Modbus TCP 服务器监听端口（默认 502）
- `deviceIds`: 要对外提供 Modbus 服务的设备ID列表

### 4. 测试验证

**文件**: `services/s_export_modbus/internal/s_modbus_config_test.go`

添加了配置结构体的单元测试，验证：
- JSON 序列化/反序列化
- 启用/禁用状态的正确处理
- 配置字段的完整性

## 功能特性

### 1. 服务控制
- **启用状态**: 当 `enabled: true` 时，服务正常启动并监听端口
- **禁用状态**: 当 `enabled: false` 时，服务不启动但管理器状态为运行中
- **动态配置**: 支持通过配置重载动态启用/禁用服务

### 2. 状态管理
- **运行状态**: 无论是否启用，管理器都标记为运行状态
- **服务器状态**: 禁用时服务器为 nil，但状态查询正常返回
- **设备映射**: 禁用时清空设备映射，避免资源占用

### 3. 日志记录
- **启用日志**: 记录服务启动和配置加载过程
- **禁用日志**: 明确记录服务被禁用的原因
- **状态日志**: 提供清晰的服务状态信息

## 使用示例

### 启用 Modbus 服务
```json
{
  "enabled": true,
  "listenPort": 502,
  "deviceIds": ["pylon_bms", "pcs_1"]
}
```

### 禁用 Modbus 服务
```json
{
  "enabled": false,
  "listenPort": 502,
  "deviceIds": ["pylon_bms", "pcs_1"]
}
```

## 向后兼容性

- **默认行为**: 如果配置中没有 `enabled` 字段，服务会正常启动（向后兼容）
- **配置迁移**: 现有配置无需修改，新字段为可选
- **API 兼容**: 所有现有 API 接口保持不变

## 测试验证

### 单元测试
- ✅ `TestSModbusConfig`: 配置结构体序列化/反序列化测试
- ✅ `TestCalculateRegisterCount`: 寄存器数量计算测试
- ✅ `TestEncodeValueToRegisters`: 值转寄存器测试
- ✅ `TestDecodeRegistersToValue`: 寄存器转值测试
- ✅ `TestCheckAddressOverlap`: 地址重叠检测测试

### 构建测试
- ✅ 项目构建成功
- ✅ 无编译错误
- ✅ 依赖关系正确

## 注意事项

1. **配置优先级**: `enabled` 字段优先级最高，即使其他配置正确，禁用时也不会启动服务
2. **资源管理**: 禁用时不会创建服务器实例，节省系统资源
3. **状态一致性**: 禁用时管理器状态仍为运行中，但服务器状态为未运行
4. **日志记录**: 禁用时会记录相应的日志信息，便于调试和监控

## 总结

成功为 Modbus 服务添加了 `enabled` 字段控制功能，实现了：

- ✅ 配置级别的服务控制
- ✅ 向后兼容性
- ✅ 完整的测试覆盖
- ✅ 详细的文档更新
- ✅ 优雅的错误处理

该功能为系统提供了更灵活的服务管理能力，允许根据需要动态启用或禁用 Modbus 服务。
