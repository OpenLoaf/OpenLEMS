# t_daemon - 守护进程工具模块

t_daemon 是一个完整的守护进程工具模块，用于确保 EMS 主程序的可靠运行。主要功能包括：

## 核心功能

### 1. 自动重启机制
- 监控主程序运行状态
- 主程序异常退出时自动重启
- 支持重启频率限制（防止频繁重启）
- 可配置重启延迟时间

### 2. 版本管理与更新
- 检查新版本可用性
- 下载和验证新版本
- 原子性更新（备份+替换）
- 更新失败时自动回滚

### 3. 环境变量管理
- 启动前配置环境变量（如 DYLD_LIBRARY_PATH）
- 支持平台特定的库路径管理
- 自动备份和恢复原始环境变量

### 4. 进程管理
- 启动和停止进程
- 跨平台进程监控（Windows/Linux/macOS）
- 进程健康检查
- 优雅关闭和强制终止

## 模块结构

```
t_daemon/
├── go.mod                              # Go 模块定义
├── t_daemon_i.go                       # 核心接口定义
├── t_daemon_config_s.go                # 配置结构体
├── t_daemon_manager_impl_s.go          # 守护进程管理器实现
├── t_process_monitor_impl_s.go         # 进程监控器实现
├── t_environment_manager_impl_s.go     # 环境变量管理器实现
├── t_version_manager_impl_s.go         # 版本管理器实现
├── t_daemon_export_f.go                # 导出函数
└── README.md                           # 本文件
```

## 接口说明

### IDaemonManager - 守护进程管理器
主要职责是管理整个守护进程的生命周期。

#### 主要方法

| 方法 | 说明 |
|------|------|
| `Start(ctx, config)` | 启动守护进程 |
| `Stop(ctx)` | 停止守护进程 |
| `Shutdown(ctx)` | 优雅关闭守护进程 |
| `Restart(ctx)` | 重启主程序 |
| `UpdateMainBinary(ctx, newPath)` | 更新主程序二进制文件 |
| `RollbackMainBinary(ctx)` | 回滚到备份版本 |
| `CheckMainProgramHealth(ctx)` | 检查主程序健康状态 |

### IProcessMonitor - 进程监控器
负责监控主程序的运行状态。

#### 主要方法

| 方法 | 说明 |
|------|------|
| `Start(ctx, pid, onExit)` | 启动进程监控 |
| `Stop(ctx)` | 停止进程监控 |
| `IsProcessRunning(pid)` | 检查进程是否运行 |
| `KillProcess(ctx, pid, force)` | 杀死进程 |

### IVersionManager - 版本管理器
负责版本检查和更新。

#### 主要方法

| 方法 | 说明 |
|------|------|
| `CheckForUpdates(ctx)` | 检查新版本 |
| `DownloadUpdate(ctx, url, targetPath)` | 下载更新 |
| `VerifyBinary(ctx, path, hash)` | 验证二进制文件 |
| `BackupCurrentBinary(ctx, path)` | 备份当前版本 |
| `ApplyUpdate(ctx, newPath, currentPath)` | 应用更新 |

### IEnvironmentManager - 环境变量管理器
负责管理环境变量。

#### 主要方法

| 方法 | 说明 |
|------|------|
| `SetupEnvironment(ctx, variables)` | 设置环境变量 |
| `CleanupEnvironment(ctx)` | 清理环境变量 |
| `BuildProcessEnv(ctx)` | 构建进程环境变量列表 |
| `ValidateEnvironment(ctx, variables)` | 验证环境变量 |

## 使用示例

### 1. 基本使用

```go
package main

import (
    "context"
    "t_daemon"
)

func main() {
    ctx := context.Background()

    // 创建配置
    config := &t_daemon.SDaemonConfig{
        MainBinaryPath: "/path/to/ems",
        WorkDirectory: "/path/to/work",
        LogDirectory: "/path/to/logs",
        MaxRestarts: 3,
        RestartWindow: 300,
        RestartDelay: 5,
        EnvironmentVariables: map[string]string{
            "DYLD_LIBRARY_PATH": "/path/to/libs",
        },
        MainProgramArgs: []string{"--web", "--profile", "prod"},
    }

    // 创建守护进程管理器
    manager := t_daemon.NewDaemonManagerInstance()

    // 启动守护进程
    if err := manager.Start(ctx, config); err != nil {
        panic(err)
    }

    // 守护进程现在在后台运行
    // 主程序会自动监控和重启
}
```

### 2. 与环境变量配置

```go
config := &t_daemon.SDaemonConfig{
    MainBinaryPath: "./ems",
    EnvironmentVariables: map[string]string{
        "DYLD_LIBRARY_PATH": "/Users/zhao/Documents/01.Code/Hex/ems-plan/@cpp/build",
        "LANG": "zh_CN.UTF-8",
    },
}

manager := t_daemon.NewDaemonManagerInstance()
manager.Start(ctx, config)
```

### 3. 主程序更新

```go
// 更新主程序
newBinaryPath := "/path/to/new/ems"
if err := manager.UpdateMainBinary(ctx, newBinaryPath); err != nil {
    // 处理错误
}

// 或回滚到上一个版本
if err := manager.RollbackMainBinary(ctx); err != nil {
    // 处理错误
}
```

### 4. 手动重启

```go
// 检查主程序是否健康
if !manager.CheckMainProgramHealth(ctx) {
    // 重启主程序
    if err := manager.Restart(ctx); err != nil {
        // 处理错误
    }
}
```

## 配置参数说明

### SDaemonConfig 字段说明

| 字段 | 类型 | 说明 | 默认值 |
|------|------|------|--------|
| `MainBinaryPath` | string | 主程序的二进制文件路径 | - |
| `WorkDirectory` | string | 主程序的工作目录 | - |
| `LogDirectory` | string | 日志目录 | - |
| `MaxRestarts` | int | 时间窗口内最大重启次数 | 3 |
| `RestartWindow` | int | 重启时间窗口（秒） | 300 |
| `RestartDelay` | int | 重启延迟时间（秒） | 5 |
| `EnvironmentVariables` | map[string]string | 环境变量映射 | 空 |
| `MainProgramArgs` | []string | 主程序命令行参数 | 空 |
| `DaemonPidFile` | string | 守护进程PID文件路径 | - |
| `MainProgramPidFile` | string | 主程序PID文件路径 | - |
| `VersionCheckUrl` | string | 版本检查URL | - |
| `VersionCheckInterval` | int | 版本检查间隔（秒） | 3600 |
| `AutoUpdate` | bool | 是否启用自动更新 | false |
| `BackupBinaryPath` | string | 备份二进制文件路径 | - |

## 重启机制说明

### 防止频繁重启

守护进程实现了防止频繁重启的机制：

- 使用时间窗口（`RestartWindow`）来限制重启次数
- 在指定时间窗口内，最多重启 `MaxRestarts` 次
- 如果超过限制，守护进程将停止并记录错误

示例配置：
```
MaxRestarts: 3
RestartWindow: 300  // 5分钟
```

这意味着：在5分钟内最多重启3次，如果第4次重启尝试在5分钟内发生，守护进程将停止。

### 重启延迟

主程序退出后，守护进程会等待 `RestartDelay` 秒再重启，这样可以：
- 避免立即重启导致资源未释放
- 给系统一些恢复时间
- 便于问题诊断

## 环境变量配置示例

### macOS

```go
config.EnvironmentVariables = map[string]string{
    "DYLD_LIBRARY_PATH": "/usr/local/lib:/opt/local/lib",
}
```

### Linux

```go
config.EnvironmentVariables = map[string]string{
    "LD_LIBRARY_PATH": "/usr/local/lib:/opt/local/lib",
}
```

### Windows

```go
config.EnvironmentVariables = map[string]string{
    "PATH": "C:\\Program Files\\App\\lib;" + os.Getenv("PATH"),
}
```

## 跨平台支持

t_daemon 支持多个操作系统：

- **macOS**: 使用 DYLD_LIBRARY_PATH，支持 SIGTERM/SIGKILL
- **Linux**: 使用 LD_LIBRARY_PATH，支持 SIGTERM/SIGKILL
- **Windows**: 使用 PATH，使用 taskkill 命令

## 日志和监控

所有重要操作都会被记录到日志中：

- 守护进程启动/停止
- 主程序启动/重启/退出
- 环境变量设置
- 版本检查和更新
- 错误和警告

## 错误处理

模块使用统一的错误处理机制：

```go
import "github.com/pkg/errors"

// 所有错误都使用 errors.Wrap 添加上下文
if err != nil {
    return errors.Wrap(err, "操作描述")
}
```

## 性能考虑

1. **进程监控**: 使用高效的 os.FindProcess 和信号检查
2. **版本检查**: 可配置检查间隔，默认1小时
3. **内存使用**: 最小化内存占用，只保存必要的状态

## 未来改进

- [ ] 支持集群模式（多个守护进程协调）
- [ ] 更智能的重启策略（退避算法）
- [ ] Web API 用于远程管理
- [ ] 健康检查端点集成
- [ ] 更详细的监控指标

## 许可证

本模块遵循项目整体许可证。

## 相关文档

- [EMS 项目规范](../../.cursor/rules/code-common-library_v1.0.mdc)
- [应用主程序](../../application/README.md)
