# EMS Plan 架构设计优化方案

## 1. 设备类型系统设计

### 1.1 设备类型和性质分离

```go
// 设备类型（功能类型）
type DeviceType string

const (
    DeviceTypeAmmeter     DeviceType = "ammeter"      // 电表
    DeviceTypeBms         DeviceType = "bms"          // 电池管理系统
    DeviceTypePcs         DeviceType = "pcs"          // 功率转换系统
    DeviceTypeGpio        DeviceType = "gpio"         // GPIO设备
    DeviceTypeEnergyStore DeviceType = "energy_store" // 储能系统
    DeviceTypeStation     DeviceType = "station"      // 电站
    DeviceTypeVirtual     DeviceType = "virtual"      // 虚拟设备
)

// 设备性质（实现方式）
type DeviceNature string

const (
    DeviceNaturePhysical DeviceNature = "physical" // 物理设备：直接通过协议连接
    DeviceNatureVirtual  DeviceNature = "virtual"  // 虚拟设备：由子设备组成
)
```

### 1.2 设备配置结构

```go
type DeviceConfig struct {
    // 基础信息
    Id          string            `json:"id"`
    Name        string            `json:"name"`
    Type        DeviceType        `json:"type"`        // 设备类型
    Nature      DeviceNature      `json:"nature"`      // 设备性质
    ParentId    string            `json:"parent_id"`
    
    // 驱动配置
    Driver      string            `json:"driver"`
    DriverParams map[string]any   `json:"driver_params"`
    
    // 协议配置（仅物理设备需要）
    ProtocolId  string            `json:"protocol_id"`
    ProtocolParams map[string]any `json:"protocol_params"`
    
    // 子设备配置（仅虚拟设备需要）
    Children    []*DeviceConfig   `json:"children"`
    
    // 其他配置
    Enabled     bool              `json:"enabled"`
    LogLevel    string            `json:"log_level"`
    StorageEnable bool            `json:"storage_enable"`
    StorageInterval int32         `json:"storage_interval"`
    
    Extensions  map[string]any    `json:"extensions"`
}
```

### 1.3 设备基类设计

```go
// 设备基础接口
type IDevice interface {
    IDeviceLifecycle
    IDeviceCapability
    IDeviceState
    IDeviceHierarchy
    IDeviceDataAccess
}

// 设备生命周期
type IDeviceLifecycle interface {
    Init(config *DeviceConfig) error
    Start() error
    Stop() error
    Destroy() error
}

// 设备能力接口
type IDeviceCapability interface {
    CanRead() bool
    CanWrite() bool
    CanControl() bool
    GetDeviceType() DeviceType
    GetNature() DeviceNature
    ImplementsInterface(interfaceType string) bool
    HasProtocol() bool
    GetProtocolType() ProtocolType
}

// 设备状态管理
type IDeviceState interface {
    GetStatus() DeviceStatus
    GetHealth() DeviceHealth
    GetLastUpdateTime() time.Time
    GetError() error
    OnStateChange(callback func(old, new DeviceState))
}

// 设备层次结构
type IDeviceHierarchy interface {
    GetParent() IDevice
    SetParent(parent IDevice)
    GetChildren() []IDevice
    AddChild(child IDevice) error
    RemoveChild(childId string) error
    FindChildByType(deviceType DeviceType) []IDevice
    FindChildById(deviceId string) IDevice
    FindChildByInterface(interfaceType string) []IDevice
}

// 设备数据访问
type IDeviceDataAccess interface {
    GetMetaValueList() []*MetaValueWrapper
    GetValue(meta *Meta) (interface{}, error)
    GetCachedValue(meta *Meta) (interface{}, error)
    SetValue(meta *Meta, value interface{}) error
}
```

### 1.4 物理设备基类

```go
type PhysicalDevice struct {
    BaseDevice
    protocol IProtocol
    dataCache IDataCache
}

func (p *PhysicalDevice) Init(config *DeviceConfig) error {
    if config.Nature != DeviceNaturePhysical {
        return errors.New("配置的设备性质与物理设备不匹配")
    }
    
    p.config = config
    p.nature = DeviceNaturePhysical
    p.deviceType = config.Type
    
    if config.ProtocolId == "" {
        return errors.New("物理设备必须配置协议")
    }
    
    return nil
}

func (p *PhysicalDevice) GetValue(meta *Meta) (interface{}, error) {
    // 先从缓存获取
    if value, err := p.dataCache.Get(meta.Name); err == nil && value != nil {
        return value, nil
    }
    
    // 缓存没有则从协议读取
    if p.protocol != nil {
        return p.protocol.ReadValue(meta)
    }
    
    return nil, errors.New("协议未初始化")
}

func (p *PhysicalDevice) GetMetaValueList() []*MetaValueWrapper {
    return p.dataCache.GetAllValues()
}
```

### 1.5 虚拟设备基类

```go
type VirtualDevice struct {
    BaseDevice
    children map[string]IDevice
    childTypes map[DeviceType][]IDevice
    dataAggregator IDataAggregator
}

func (v *VirtualDevice) Init(config *DeviceConfig) error {
    if config.Nature != DeviceNatureVirtual {
        return errors.New("配置的设备性质与虚拟设备不匹配")
    }
    
    v.config = config
    v.nature = DeviceNatureVirtual
    v.deviceType = config.Type
    v.children = make(map[string]IDevice)
    v.childTypes = make(map[DeviceType][]IDevice)
    
    return nil
}

func (v *VirtualDevice) GetValue(meta *Meta) (interface{}, error) {
    // 虚拟设备通过聚合子设备数据获取值
    return v.dataAggregator.AggregateValue(meta, v.children)
}

func (v *VirtualDevice) GetMetaValueList() []*MetaValueWrapper {
    // 聚合所有子设备的数据
    var allValues []*MetaValueWrapper
    for _, child := range v.children {
        allValues = append(allValues, child.GetMetaValueList()...)
    }
    return allValues
}
```

## 2. 协议系统设计

### 2.1 协议抽象层

```go
// 协议类型
type ProtocolType string

const (
    ProtocolTypeModbus    ProtocolType = "modbus"
    ProtocolTypeCanbus    ProtocolType = "canbus"
    ProtocolTypeGpioSysfs ProtocolType = "gpio_sysfs"
    ProtocolTypeMqtt      ProtocolType = "mqtt"
    ProtocolTypeHttp      ProtocolType = "http"
)

// 协议接口
type IProtocol interface {
    IProtocolConnection
    IProtocolDataAccess
    IProtocolLifecycle
}

// 协议连接管理
type IProtocolConnection interface {
    Connect() error
    Disconnect() error
    IsConnected() bool
    GetConnectionInfo() ConnectionInfo
    GetPool() ConnectionPool
}

// 协议数据访问
type IProtocolDataAccess interface {
    ReadValue(meta *Meta) (interface{}, error)
    WriteValue(meta *Meta, value interface{}) error
    ReadBatch(metas []*Meta) ([]interface{}, error)
    WriteBatch(metas []*Meta, values []interface{}) error
    Subscribe(meta *Meta, callback func(interface{})) error
    Unsubscribe(meta *Meta) error
}

// 协议生命周期
type IProtocolLifecycle interface {
    Init(config *ProtocolConfig) error
    Start() error
    Stop() error
    Destroy() error
}
```

### 2.2 协议连接池

```go
type IProtocolConnectionPool interface {
    GetConnection(protocolId string) (IProtocolConnection, error)
    ReturnConnection(protocolId string, conn IProtocolConnection)
    CloseConnection(protocolId string) error
    CloseAll() error
    GetPoolStats() ConnectionPoolStats
    HealthCheck() error
}

type ConnectionPool struct {
    connections map[string]*ConnectionWrapper
    mutex       sync.RWMutex
    maxConnections int
    timeout     time.Duration
}

type ConnectionWrapper struct {
    connection IProtocolConnection
    lastUsed   time.Time
    usageCount int64
    isHealthy  bool
}
```

### 2.3 协议适配器

```go
type IProtocolAdapter interface {
    AdaptProtocol(deviceType DeviceType, protocol IProtocol) (IDeviceProtocol, error)
    GetSupportedProtocols(deviceType DeviceType) []ProtocolType
    ValidateProtocolCompatibility(deviceType DeviceType, protocolType ProtocolType) bool
}

type ProtocolAdapter struct {
    adapters map[ProtocolType]IProtocolImplementation
}

type IProtocolImplementation interface {
    CreateClient(config *ProtocolConfig) (IProtocolClient, error)
    CreateProvider(deviceType DeviceType, config *ProtocolConfig, client IProtocolClient) (IProtocol, error)
    GetSupportedDeviceTypes() []DeviceType
}
```

### 2.4 数据缓存系统

```go
type IDataCache interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}, ttl time.Duration) error
    Delete(key string) error
    Clear() error
    GetAllValues() []*MetaValueWrapper
    GetExpiredKeys() []string
    Cleanup() error
}

type DataCache struct {
    cache map[string]*CacheEntry
    mutex sync.RWMutex
    ttl   time.Duration
}

type CacheEntry struct {
    Value      interface{}
    ExpireTime time.Time
    Meta       *Meta
    DeviceId   string
}

func (c *DataCache) Get(key string) (interface{}, error) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    entry, exists := c.cache[key]
    if !exists {
        return nil, errors.New("key not found")
    }
    
    if time.Now().After(entry.ExpireTime) {
        return nil, errors.New("value expired")
    }
    
    return entry.Value, nil
}

func (c *DataCache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    expireTime := time.Now().Add(ttl)
    c.cache[key] = &CacheEntry{
        Value:      value,
        ExpireTime: expireTime,
    }
    
    return nil
}
```

## 3. 设备管理器设计

### 3.1 设备管理器接口

```go
type IDeviceManager interface {
    // 设备注册管理
    RegisterDevice(device IDevice) error
    UnregisterDevice(deviceId string) error
    GetDevice(deviceId string) IDevice
    
    // 设备树管理
    BuildDeviceTree(configs []*DeviceConfig) error
    GetDeviceTree() *DeviceTree
    GetRootDevices() []IDevice
    
    // 设备操作
    StartDevice(deviceId string) error
    StopDevice(deviceId string) error
    RestartDevice(deviceId string) error
    
    // 批量操作
    StartAllDevices() error
    StopAllDevices() error
    GetDeviceStatus() map[string]DeviceStatus
    
    // 设备发现
    DiscoverDevices() []IDevice
    GetDevicesByType(deviceType DeviceType) []IDevice
    GetDevicesByProtocol(protocolType ProtocolType) []IDevice
    
    // 健康检查
    HealthCheck() DeviceHealthReport
    AutoRecovery() error
}

type DeviceManager struct {
    typeRegistry IDeviceTypeRegistry
    deviceCreator *DeviceCreator
    deviceTree *DeviceTree
    protocolPool IProtocolConnectionPool
    configManager IConfigManager
}
```

### 3.2 设备启动流程

```go
type DeviceStartupFlow struct {
    configManager    IConfigManager
    deviceManager    IDeviceManager
    protocolPool     IProtocolConnectionPool
    typeRegistry     IDeviceTypeRegistry
}

func (f *DeviceStartupFlow) Start() error {
    // 1. 加载配置
    configs, err := f.configManager.LoadConfigs()
    if err != nil {
        return err
    }
    
    // 2. 验证配置
    if err := f.validateConfigs(configs); err != nil {
        return err
    }
    
    // 3. 构建设备树
    if err := f.deviceManager.BuildDeviceTree(configs); err != nil {
        return err
    }
    
    // 4. 初始化设备（从叶子节点开始）
    if err := f.initializeDevices(); err != nil {
        return err
    }
    
    // 5. 启动设备（从根节点开始）
    if err := f.startDevices(); err != nil {
        return err
    }
    
    return nil
}

func (f *DeviceStartupFlow) initializeDevices() error {
    devices := f.deviceManager.GetDeviceTree().GetTopologicalOrder()
    
    for _, device := range devices {
        if err := f.initializeDevice(device); err != nil {
            return err
        }
    }
    
    return nil
}

func (f *DeviceStartupFlow) initializeDevice(device IDevice) error {
    config := device.GetConfig()
    
    // 1. 创建协议连接（仅物理设备）
    if device.IsPhysical() && config.ProtocolId != "" {
        protocol, err := f.protocolPool.GetConnection(config.ProtocolId)
        if err != nil {
            return err
        }
        device.SetProtocol(protocol)
    }
    
    // 2. 初始化设备
    if err := device.Init(config); err != nil {
        return err
    }
    
    // 3. 设置子设备引用（仅虚拟设备）
    if device.IsVirtual() {
        if virtualDevice, ok := device.(IVirtualDevice); ok {
            virtualDevice.SetupChildren()
        }
    }
    
    return nil
}
```

## 4. 数据聚合系统

### 4.1 数据聚合器

```go
type IDataAggregator interface {
    AggregateValue(meta *Meta, children map[string]IDevice) (interface{}, error)
    AggregateValues(metas []*Meta, children map[string]IDevice) ([]interface{}, error)
    GetAggregationRules() map[string]AggregationRule
    AddAggregationRule(metaName string, rule AggregationRule) error
}

type DataAggregator struct {
    rules map[string]AggregationRule
}

type AggregationRule struct {
    Type        AggregationType
    SourceDevices []string
    Formula     string
    Transform   func([]interface{}) interface{}
}

type AggregationType string

const (
    AggregationTypeSum     AggregationType = "sum"
    AggregationTypeAverage AggregationType = "average"
    AggregationTypeMax     AggregationType = "max"
    AggregationTypeMin     AggregationType = "min"
    AggregationTypeCustom  AggregationType = "custom"
)

func (a *DataAggregator) AggregateValue(meta *Meta, children map[string]IDevice) (interface{}, error) {
    rule, exists := a.rules[meta.Name]
    if !exists {
        return nil, errors.New("no aggregation rule found")
    }
    
    var values []interface{}
    for _, deviceId := range rule.SourceDevices {
        if child, exists := children[deviceId]; exists {
            if value, err := child.GetValue(meta); err == nil {
                values = append(values, value)
            }
        }
    }
    
    if len(values) == 0 {
        return nil, errors.New("no valid values found")
    }
    
    return rule.Transform(values), nil
}
```

## 5. 配置管理优化

### 5.1 配置验证器

```go
type IConfigValidator interface {
    ValidateDeviceConfig(config *DeviceConfig) error
    ValidateProtocolConfig(config *ProtocolConfig) error
    ValidateHierarchy(configs []*DeviceConfig) error
    ValidateDependencies(config *DeviceConfig, allConfigs []*DeviceConfig) error
    CheckDeviceProtocolCompatibility(deviceType DeviceType, protocolType ProtocolType) bool
}

type ConfigValidator struct {
    typeRegistry IDeviceTypeRegistry
}

func (v *ConfigValidator) ValidateDeviceConfig(config *DeviceConfig) error {
    // 检查设备类型是否支持
    if !v.typeRegistry.IsValidDeviceType(config.Type) {
        return errors.New("不支持的设备类型")
    }
    
    // 检查设备性质是否支持
    if !v.typeRegistry.SupportsNature(config.Type, config.Nature) {
        return errors.New("设备类型不支持该性质")
    }
    
    // 根据设备性质验证配置
    if config.Nature == DeviceNaturePhysical {
        if config.ProtocolId == "" {
            return errors.New("物理设备必须配置协议")
        }
    } else if config.Nature == DeviceNatureVirtual {
        if len(config.Children) == 0 {
            return errors.New("虚拟设备必须配置子设备")
        }
    }
    
    return nil
}
```

### 5.2 配置热更新

```go
type IConfigManager interface {
    LoadConfigs() ([]*DeviceConfig, error)
    SaveConfig(config *DeviceConfig) error
    DeleteConfig(deviceId string) error
    WatchConfigs(callback func(ConfigChange))
    ReloadConfigs() error
    GetConfigVersion() string
    RollbackConfig(version string) error
}

type ConfigManager struct {
    configPath string
    configs    map[string]*DeviceConfig
    watchers   []func(ConfigChange)
    validator  IConfigValidator
}
```

## 6. 错误处理和恢复机制

### 6.1 错误分类

```go
type DeviceError interface {
    GetErrorType() ErrorType
    GetErrorCode() ErrorCode
    GetErrorMessage() string
    GetErrorContext() map[string]interface{}
    IsRecoverable() bool
}

type ErrorType string

const (
    ErrorTypeConnection   ErrorType = "connection"
    ErrorTypeProtocol     ErrorType = "protocol"
    ErrorTypeData         ErrorType = "data"
    ErrorTypeConfig       ErrorType = "config"
    ErrorTypeSystem       ErrorType = "system"
)

type ErrorCode int

const (
    ErrorCodeConnectionFailed ErrorCode = 1001
    ErrorCodeProtocolTimeout  ErrorCode = 1002
    ErrorCodeDataInvalid      ErrorCode = 1003
    ErrorCodeConfigInvalid    ErrorCode = 1004
    ErrorCodeSystemError      ErrorCode = 1005
)
```

### 6.2 自动恢复机制

```go
type IAutoRecovery interface {
    EnableAutoRecovery(device IDevice) error
    DisableAutoRecovery(device IDevice) error
    SetRecoveryStrategy(device IDevice, strategy RecoveryStrategy) error
    GetRecoveryHistory(device IDevice) []RecoveryRecord
}

type RecoveryStrategy struct {
    MaxRetries    int
    RetryInterval time.Duration
    BackoffFactor float64
    RecoverySteps []RecoveryStep
}

type RecoveryStep struct {
    Name        string
    Action      func(IDevice) error
    Rollback    func(IDevice) error
    Timeout     time.Duration
}
```

## 7. 性能监控和指标

### 7.1 性能监控器

```go
type IPerformanceMonitor interface {
    StartMonitoring(device IDevice) error
    StopMonitoring(device IDevice) error
    GetDeviceMetrics(device IDevice) DeviceMetrics
    GetProtocolMetrics(protocolId string) ProtocolMetrics
    GetSystemMetrics() SystemMetrics
}

type DeviceMetrics struct {
    DeviceId        string
    ResponseTime    time.Duration
    Throughput      int64
    ErrorRate       float64
    LastUpdateTime  time.Time
    CacheHitRate    float64
}

type ProtocolMetrics struct {
    ProtocolId      string
    ConnectionCount int
    ActiveConnections int
    RequestCount    int64
    ErrorCount      int64
    AverageLatency  time.Duration
}
```

## 8. 插件架构优化

### 8.1 插件管理器

```go
type IPluginManager interface {
    LoadPlugin(pluginPath string) error
    UnloadPlugin(pluginId string) error
    GetPlugin(pluginId string) IPlugin
    DiscoverPlugins() []PluginInfo
    ValidatePlugin(plugin IPlugin) error
    ResolveDependencies(plugin IPlugin) error
    CheckCompatibility(plugin1, plugin2 IPlugin) bool
}

type IPlugin interface {
    GetPluginInfo() PluginInfo
    Init(config PluginConfig) error
    Start() error
    Stop() error
    Destroy() error
    GetCapabilities() []PluginCapability
    SupportsFeature(feature string) bool
}

type PluginInfo struct {
    Id          string
    Name        string
    Version     string
    Type        PluginType
    Dependencies []string
    Capabilities []PluginCapability
}
```

## 9. 接口设计原则总结

### 9.1 核心设计原则

1. **统一接口**：物理设备和虚拟设备使用相同的接口
2. **组合优于继承**：通过组合实现功能扩展
3. **依赖注入**：通过配置注入依赖关系
4. **类型安全**：通过接口确保类型安全
5. **可扩展性**：支持新的设备类型和接口
6. **数据缓存**：统一的数据缓存机制
7. **协议抽象**：统一的协议访问接口

### 9.2 实现优势

1. **简化管理**：统一的设备管理接口
2. **提高可维护性**：清晰的层次结构
3. **增强扩展性**：易于添加新设备类型和协议
4. **改善性能**：优化的启动流程和数据缓存
5. **提升可靠性**：完善的错误处理和自动恢复
6. **支持热更新**：配置和设备的热更新能力
7. **性能监控**：全面的性能指标收集

这个架构设计方案解决了当前架构中的主要问题，提供了更好的可扩展性、可维护性和性能，同时支持物理设备和虚拟设备的统一管理。
