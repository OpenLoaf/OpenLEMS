# GPIO协议插件 (p_gpiod)

基于Linux GPIO字符设备直接实现的GPIO协议插件，支持Linux GPIO字符设备的输入和输出操作。使用syscall直接操作GPIO设备，避免第三方库依赖。

## 功能特性

- ✅ GPIO芯片初始化和管理
- ✅ 输入模式：支持状态读取和边缘检测
- ✅ 输出模式：支持高/低电平设置
- ✅ 低电平有效/高电平有效配置
- ✅ 状态变化事件监听
- ✅ 线程安全的状态管理
- ✅ 资源自动清理
- ✅ 错误处理和告警支持

## 配置结构

```go
type SGpiodProtocolConfig struct {
    Direction c_enum.EGpioDirection `json:"direction" required:"true" name:"方向" ct:"singleSelect" vt:"string" selectOptions:"in:输入,out:输出" required:"true" dc:"GPIO引脚方向：输入用于读取状态，输出用于控制"`
    ChipIndex uint8                 `json:"chipIndex" dc:"GPIO芯片名称，如gpiochip0"`
    Pin       uint8                 `json:"pin" dc:"GPIO引脚编号，范围0-99"`
    LowActive bool                  `json:"lowActive" dc:"低电平有效"`
}
```

### 配置参数说明

- `direction`: GPIO引脚方向
  - `in`: 输入模式，用于读取外部信号状态
  - `out`: 输出模式，用于控制外部设备
- `chipIndex`: GPIO芯片索引，对应 `/dev/gpiochip{N}` 设备
- `pin`: GPIO引脚编号，具体范围取决于硬件
- `lowActive`: 电平有效模式
  - `false`: 高电平有效（默认）
  - `true`: 低电平有效

## 使用方法

### 1. 创建GPIO提供者

```go
import (
    "common/c_enum"
    "common/c_proto"
    "p_gpiod/internal"
)

// 创建配置
config := &c_proto.SGpiodProtocolConfig{
    Direction: c_enum.EGpioDirectionOut, // 输出模式
    ChipIndex: 0,                        // gpiochip0
    Pin:       18,                       // GPIO18引脚
    LowActive: false,                    // 高电平有效
}

// 创建提供者
provider := internal.NewGpiodProvider(alarm, config)
```

### 2. 初始化GPIO

```go
// 初始化GPIO芯片和引脚
provider.ProtocolListen()

// 检查连接状态
status := provider.GetProtocolStatus()
if status == c_enum.EProtocolConnected {
    fmt.Println("GPIO初始化成功")
}
```

### 3. 输出模式操作

```go
// 设置高电平
err := provider.SetHigh()
if err != nil {
    log.Printf("设置高电平失败: %v", err)
}

// 设置低电平
err = provider.SetLow()
if err != nil {
    log.Printf("设置低电平失败: %v", err)
}

// 读取当前状态
status := provider.GetGpioStatus()
if status != nil {
    fmt.Printf("当前GPIO状态: %v\n", *status)
}
```

### 4. 输入模式操作

```go
// 注册状态变化处理函数
provider.RegisterHandler(func(status bool) {
    fmt.Printf("GPIO状态变化: %v\n", status)
})

// 读取当前状态
status := provider.GetGpioStatus()
if status != nil {
    fmt.Printf("当前GPIO状态: %v\n", *status)
}
```

### 5. 资源清理

```go
// 关闭GPIO提供者，释放资源
err := provider.Close()
if err != nil {
    log.Printf("关闭失败: %v", err)
}
```

## 接口方法

### IGpiodProtocol 接口

```go
type IGpiodProtocol interface {
    c_base.IProtocol
    
    RegisterHandler(handler func(status bool)) // 注册状态变化处理函数
    GetGpioStatus() *bool                      // 获取当前GPIO状态
    SetHigh() error                            // 设置为高电平
    SetLow() error                             // 设置为低电平
}
```

### 继承的 IProtocol 接口

```go
type IProtocol interface {
    GetProtocolStatus() c_enum.EProtocolStatus // 获取协议连接状态
    GetLastUpdateTime() *time.Time             // 获取最后更新时间
    GetPointValueList() []*SPointValue         // 获取点位值列表
    GetValue(point IPoint) (any, error)        // 获取指定点位的值
    RegisterTask(task IPointTask, tasks ...IPointTask) // 注册任务
    ProtocolListen()                           // 协议监听（初始化）
}
```

## 硬件要求

- Linux系统
- 支持GPIO字符设备的硬件
- 适当的GPIO权限（通常需要root权限或gpio组权限）

## 权限配置

确保运行程序有访问GPIO设备的权限：

```bash
# 方法1: 使用root权限运行
sudo ./your_program

# 方法2: 将用户添加到gpio组
sudo usermod -a -G gpio $USER

# 方法3: 设置设备权限
sudo chmod 666 /dev/gpiochip*
```

## 示例代码

参考 `example_usage.go` 文件中的完整示例，包括：

- 输出模式示例：控制LED或继电器
- 输入模式示例：读取按钮或传感器状态
- 状态变化监听示例

## 错误处理

- 初始化失败时会记录错误信息
- GPIO操作失败会返回详细的错误信息
- 支持告警系统集成

## 注意事项

1. **线程安全**: 所有方法都是线程安全的
2. **资源管理**: 使用完毕后务必调用 `Close()` 方法
3. **权限要求**: 需要GPIO设备访问权限
4. **硬件兼容**: 确保硬件支持GPIO字符设备
5. **引脚冲突**: 避免同时使用已被其他程序占用的GPIO引脚

## 依赖库

- `golang.org/x/sys v0.18.0` (仅用于系统调用)

## 版本历史

- v1.0.0: 初始版本，支持基本的GPIO输入输出功能
