# 网络接口管理模块 - 更新日志

## v2.0.0 - DNS集成和代码重构

### ✅ 主要更新

#### 🔧 DNS配置集成
- **删除独立DNS接口**：删除了 `network_v1_update_dns.go` 文件
- **DNS集成到网络接口**：将DNS配置集成到网络接口管理中
- **统一管理**：DNS配置现在通过网络接口更新接口统一管理

#### 🚀 nmcli命令优化
- **优先使用nmcli**：所有网络操作优先使用NetworkManager的nmcli命令
- **DNS配置支持**：nmcli命令现在支持DNS服务器配置
- **更好的兼容性**：提供备用方法确保在nmcli不可用时的兼容性

#### 📝 数据结构更新
- **UpdateInterfaceRequest**：添加DNS字段支持
- **ConnectionInfo**：添加DNS字段用于连接信息存储
- **API接口更新**：UpdateNetworkInterfaceReq添加DNS字段

#### ✅ 验证器增强
- **DNS验证**：添加DNS服务器地址格式验证
- **DHCP模式**：支持DHCP模式下的DNS配置验证
- **静态模式**：支持静态模式下的DNS配置验证

#### 🧹 代码精简
- **删除冗余方法**：删除了MacOS和Windows的备用方法
- **专注Linux**：代码现在专注于Linux系统支持
- **简化逻辑**：移除了复杂的跨平台兼容代码

### 📋 具体变更

#### 新增功能
1. **DNS配置集成**
   - DNS服务器地址可以在网络接口配置中设置
   - 支持DHCP和静态模式下的DNS配置
   - 使用nmcli命令进行DNS配置管理

2. **增强的nmcli支持**
   - 获取DNS配置信息
   - 设置DNS服务器地址
   - 自动检测DNS配置状态

#### 删除功能
1. **独立DNS接口**
   - 删除 `/network/dns/update` API端点
   - 删除 `UpdateDNS` 控制器方法
   - 删除相关的请求/响应结构体

2. **跨平台支持**
   - 删除MacOS相关方法
   - 删除Windows相关方法
   - 专注Linux系统支持

#### 修改功能
1. **网络接口更新**
   - API路径：`/network/interface/update`
   - 新增DNS字段：`dns: []string`
   - 支持DNS服务器配置

2. **数据验证**
   - 验证DNS服务器地址格式
   - 支持DHCP和静态模式的DNS验证

### 🔄 API变更

#### 更新的接口

**POST /network/interface/update**
```json
{
  "name": "eth0",
  "dhcp": false,
  "ipAddresses": ["192.168.1.100"],
  "netmask": "ffffff00",
  "gateway": "192.168.1.1",
  "dns": ["8.8.8.8", "8.8.4.4"]  // 新增字段
}
```

#### 删除的接口

**POST /network/dns/update** - 已删除
```json
// 此接口已被删除，DNS配置现在通过网络接口更新接口管理
```

### 💻 使用示例

#### 更新网络接口配置（包含DNS）

```go
// 静态IP配置
updateReq := &UpdateInterfaceRequest{
    Name:        "eth0",
    DHCP:        false,
    IPAddresses: []string{"192.168.1.100"},
    Netmask:     "ffffff00",
    Gateway:     "192.168.1.1",
    DNS:         []string{"8.8.8.8", "8.8.4.4"},
}

// DHCP配置
updateReq := &UpdateInterfaceRequest{
    Name: "eth0",
    DHCP: true,
    DNS:  []string{"1.1.1.1", "1.0.0.1"}, // 可选DNS设置
}
```

#### nmcli命令使用

```bash
# 静态配置（包含DNS）
nmcli connection modify eth0 \
  ipv4.method manual \
  ipv4.addresses 192.168.1.100/24 \
  ipv4.gateway 192.168.1.1 \
  ipv4.dns "8.8.8.8,8.8.4.4"

# DHCP配置（包含DNS）
nmcli connection modify eth0 \
  ipv4.method auto \
  ipv4.dns "1.1.1.1,1.0.0.1"
```

### 🔧 技术改进

1. **更好的错误处理**
   - 统一的错误类型
   - 详细的错误信息
   - 完善的日志记录

2. **性能优化**
   - 减少系统调用
   - 优化命令执行
   - 简化代码逻辑

3. **代码质量**
   - 删除冗余代码
   - 统一代码风格
   - 完善文档注释

### ⚠️ 破坏性变更

1. **API接口变更**
   - 删除了 `POST /network/dns/update` 接口
   - DNS配置现在通过 `POST /network/interface/update` 管理

2. **数据结构变更**
   - `UpdateNetworkInterfaceReq` 添加了 `dns` 字段
   - 删除了 `UpdateDNSReq` 和 `UpdateDNSRes` 结构体

3. **控制器方法变更**
   - 删除了 `UpdateDNS` 方法
   - 修改了 `UpdateNetworkInterface` 方法签名

### 🔄 迁移指南

#### 从独立DNS接口迁移

**之前的代码：**
```go
// 更新网络接口
networkReq := &UpdateNetworkInterfaceReq{
    Name: "eth0",
    DHCP: false,
    IPAddresses: []string{"192.168.1.100"},
    Netmask: "ffffff00",
    Gateway: "192.168.1.1",
}

// 单独更新DNS
dnsReq := &UpdateDNSReq{
    DNS: []string{"8.8.8.8", "8.8.4.4"},
}
```

**现在的代码：**
```go
// 统一更新网络接口和DNS
networkReq := &UpdateNetworkInterfaceReq{
    Name: "eth0",
    DHCP: false,
    IPAddresses: []string{"192.168.1.100"},
    Netmask: "ffffff00",
    Gateway: "192.168.1.1",
    DNS: []string{"8.8.8.8", "8.8.4.4"}, // DNS集成到网络接口配置中
}
```

### 📚 相关文档

- [README.md](./README.md) - 完整的使用文档
- [network_types.go](./network_types.go) - 数据结构定义
- [network_manager.go](./network_manager.go) - 核心网络管理逻辑
- [network_validator.go](./network_validator.go) - 参数验证逻辑

### 🎯 下一步计划

1. **IPv6支持** - 添加IPv6地址配置支持
2. **VLAN配置** - 支持虚拟LAN配置
3. **网络监控** - 实时监控网络接口状态
4. **配置模板** - 预定义常用的网络配置模板
5. **批量操作** - 支持批量配置多个网络接口


