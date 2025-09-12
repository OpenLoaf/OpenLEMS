# 网络接口管理模块

本模块提供了完整的网络接口管理功能，支持获取网络接口列表和更新网络接口配置。

## 架构设计

### 核心组件

1. **NetworkManager** - 网络管理器
   - 负责执行具体的网络操作
   - 优先使用 nmcli 命令管理网络配置
   - 提供备用方法确保兼容性

2. **NetworkValidator** - 网络配置验证器
   - 验证网络配置参数的合法性
   - 检查IP地址、子网掩码、网关等参数
   - 提供详细的错误信息

3. **Controller** - 控制器层
   - 处理HTTP请求和响应
   - 协调各个组件完成业务逻辑
   - 提供统一的错误处理和日志记录

### 文件结构

```text
network/
├── README.md                           # 本文档
├── network_manager.go                  # 网络管理器 - 核心网络操作
├── network_validator.go               # 网络验证器 - 参数验证
├── network_types.go                   # 数据类型定义
├── network_v1_get_interface_list.go   # 获取网络接口列表接口
└── network_v1_update_interface.go     # 更新网络接口配置接口
```

## 功能特性

### 网络接口获取

- **优先使用 nmcli**：在支持的系统上使用 NetworkManager 获取详细信息
- **备用方法**：当 nmcli 不可用时，使用系统原生 API
- **跨平台支持**：支持 Linux、macOS、Windows 系统
- **完整信息**：获取接口名称、类型、MAC地址、IP地址、网关、DHCP状态等

### 网络接口配置

- **nmcli 优先**：使用 NetworkManager 进行网络配置管理
- **DHCP 支持**：支持启用/禁用 DHCP 自动获取IP
- **静态配置**：支持配置静态IP地址、子网掩码、网关
- **DNS 管理**：集成DNS服务器配置，支持DHCP和静态模式下的DNS设置
- **配置验证**：自动验证配置结果，确保配置生效
- **错误恢复**：提供备用配置方法，提高成功率

### 参数验证

- **全面验证**：验证所有网络配置参数的合法性
- **格式检查**：检查IP地址、子网掩码格式
- **网段验证**：验证网关与IP地址是否在同一网段
- **接口检查**：验证网络接口是否存在且可用

## 使用方法

### 获取网络接口列表

```go
// 创建网络管理器
networkManager := NewNetworkManager()

// 获取接口列表（包含loopback接口）
interfaces, err := networkManager.GetInterfaceList(ctx, true)
if err != nil {
    // 处理错误
}

// 获取DNS服务器列表
dnsServers := networkManager.GetDNSServers(ctx)
```

### 更新网络接口配置

```go
// 创建验证器和网络管理器
validator := NewNetworkValidator()
networkManager := NewNetworkManager()

// 构建更新请求
updateReq := &UpdateInterfaceRequest{
    Name:        "eth0",
    DHCP:        false,
    IPAddresses: []string{"192.168.1.100"},
    Netmask:     "ffffff00", // 255.255.255.0
    Gateway:     "192.168.1.1",
    DNS:         []string{"8.8.8.8", "8.8.4.4"}, // DNS服务器
}

// 验证请求参数
if err := validator.ValidateUpdateRequest(updateReq); err != nil {
    // 处理验证错误
}

// 更新网络配置
if err := networkManager.UpdateInterface(ctx, updateReq); err != nil {
    // 处理更新错误
}
```

### 参数验证

```go
validator := NewNetworkValidator()

// 验证单个配置
if err := validator.ValidateUpdateRequest(updateReq); err != nil {
    // 处理验证错误
}

// 获取详细验证错误列表
errors := validator.ValidateConfiguration(updateReq)
for _, err := range errors {
    log.Printf("字段 %s 验证失败: %s", err.Field, err.Message)
}
```

## 配置说明

### 子网掩码格式

本模块使用十六进制格式的子网掩码，例如：

- `ffffff00` = 255.255.255.0 = /24
- `ffff0000` = 255.255.0.0 = /16
- `ff000000` = 255.0.0.0 = /8

### nmcli 命令使用

系统需要安装 NetworkManager 并确保 nmcli 命令可用：

```bash
# 检查 nmcli 是否可用
which nmcli

# 检查 NetworkManager 服务状态
systemctl status NetworkManager
```

### 权限要求

网络配置操作需要管理员权限，确保应用程序以适当的权限运行。

## 错误处理

### 错误类型

1. **ValidationError** - 参数验证错误
   - 包含具体的字段名和错误信息
   - 用于客户端参数校正

2. **NetworkError** - 网络操作错误
   - 包含操作类型、命令和输出信息
   - 用于问题诊断和调试

### 错误处理策略

1. **参数验证失败** - 立即返回错误，不执行网络操作
2. **nmcli 不可用** - 自动切换到备用方法
3. **配置验证失败** - 记录警告但不阻断操作
4. **网络操作失败** - 返回详细错误信息供调试

## 日志记录

### 日志级别

- **Info** - 正常操作和状态信息
- **Warning** - 非致命错误和降级操作
- **Error** - 操作失败和严重错误
- **Debug** - 详细的命令执行信息

### 日志内容

- 操作开始和完成时间
- 命令执行详情
- 配置验证结果
- 错误详细信息和堆栈

## 性能优化

1. **并行操作** - 在安全的情况下并行执行网络检查
2. **缓存结果** - 适当缓存网络状态信息
3. **快速失败** - 参数验证失败时立即返回
4. **资源清理** - 及时释放网络资源和进程句柄

## 测试建议

1. **单元测试** - 测试各个组件的独立功能
2. **集成测试** - 测试完整的网络配置流程
3. **错误测试** - 测试各种错误情况的处理
4. **兼容性测试** - 在不同系统和网络环境下测试

## 注意事项

1. **系统兼容性** - 不同Linux发行版的网络管理方式可能不同
2. **权限要求** - 网络配置需要适当的系统权限
3. **网络中断** - 配置过程中可能导致短暂的网络中断
4. **配置持久化** - 确保配置在系统重启后仍然有效

## 扩展功能

### 计划中的功能

1. **IPv6 支持** - 添加IPv6地址配置支持
2. **VLAN 配置** - 支持虚拟LAN配置
3. **网络监控** - 实时监控网络接口状态
4. **配置备份** - 自动备份和恢复网络配置
5. **批量操作** - 支持批量配置多个网络接口

### 扩展点

1. **自定义验证器** - 添加特定业务场景的验证规则
2. **配置模板** - 预定义常用的网络配置模板
3. **事件通知** - 网络状态变化时的事件通知机制
4. **配置同步** - 与外部配置管理系统的同步
