# t_daemon 守护进程使用指南

## 概述

`t_daemon` 是一个完整的守护进程解决方案，用于管理和保护 EMS 应用程序的运行。它提供了以下功能：

- **自动重启**：当主程序崩溃时自动重启
- **环境变量管理**：在启动主程序前自动设置必要的环境变量（如 DYLD_LIBRARY_PATH）
- **版本管理**：支持程序更新和回滚
- **进程监控**：实时监控主程序状态
- **跨平台支持**：支持 macOS、Linux 和 Windows

## 快速开始

### 1. 将 t_daemon 加入你的项目

在你的 `go.work` 或 `go.mod` 中添加 t_daemon 模块：

```bash
# 在项目根目录下修改 go.work
replace t_daemon => ./tools/t_daemon
```

### 2. 创建守护进程启动程序

创建一个新的可执行程序来启动守护进程，例如 `cmd/daemon/main.go`：

```go
package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/frame/g"
	"t_daemon"
)

func main() {
	ctx := context.Background()

	// 获取项目根目录
	workDir, _ := os.Getwd()

	// 配置守护进程
	config := &t_daemon.SDaemonConfig{
		MainBinaryPath: filepath.Join(workDir, "application", "ems"),
		WorkDirectory:  filepath.Join(workDir, "application"),
		LogDirectory:   filepath.Join(workDir, "out", "logs"),

		// 重启策略
		MaxRestarts:  3,      // 5分钟内最多重启3次
		RestartWindow: 300,    // 5分钟
		RestartDelay: 5,       // 每次重启延迟5秒

		// 主程序参数
		MainProgramArgs: []string{
			"--web",           // 启动Web端
			"--profile", "prod", // 生产环境
		},

		// 环境变量 - 这里设置 DYLD_LIBRARY_PATH
		EnvironmentVariables: map[string]string{
			"DYLD_LIBRARY_PATH": filepath.Join(workDir, "@cpp", "build"),
			"LANG": "zh_CN.UTF-8",
		},
	}

	// 创建和启动守护进程
	manager := t_daemon.NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		g.Log().Fatalf(ctx, "启动守护进程失败: %v", err)
	}

	g.Log().Infof(ctx, "守护进程已启动，主程序 PID: %d", manager.GetMainProgramPid())

	// 守护进程现在在后台运行
	// 按 Ctrl+C 可以停止
	select {}
}
```

### 3. 编译并运行

```bash
# 编译守护进程
go build -o daemon ./cmd/daemon

# 运行守护进程
./daemon
```

## 配置详解

### 基础配置

```go
config := &t_daemon.SDaemonConfig{
	// 必须配置
	MainBinaryPath: "./application/ems",  // 主程序路径
	WorkDirectory:  "./application",       // 主程序工作目录

	// 可选配置
	LogDirectory:   "./out/logs",          // 日志目录
	DaemonPidFile: "./out/daemon.pid",    // 守护进程PID文件
}
```

### 重启策略

```go
config := &t_daemon.SDaemonConfig{
	MaxRestarts:   3,      // 最多重启次数
	RestartWindow: 300,    // 时间窗口（秒）
	RestartDelay:  5,      // 重启延迟（秒）
}
```

**示例说明**：
- 设置 `MaxRestarts=3` 和 `RestartWindow=300` 意味着在 5 分钟内最多重启 3 次
- 如果在这个时间窗口内超过 3 次重启，守护进程会停止并记录错误
- `RestartDelay` 设置每次重启的延迟时间，防止频繁重启

### 环境变量配置

```go
config := &t_daemon.SDaemonConfig{
	EnvironmentVariables: map[string]string{
		// macOS 特定
		"DYLD_LIBRARY_PATH": "/path/to/cpp/build",

		// Linux 特定
		"LD_LIBRARY_PATH": "/path/to/cpp/build",

		// 其他常见变量
		"LANG": "zh_CN.UTF-8",
		"TZ":   "Asia/Shanghai",
	},
}
```

**重要提示**：这些环境变量会在启动主程序前被设置，但不会影响守护进程本身的环境。

### 主程序参数

```go
config := &t_daemon.SDaemonConfig{
	MainProgramArgs: []string{
		"--web",           // 启动Web服务
		"--profile", "prod", // 使用生产环境配置
		"-db-path", "./data", // 数据库路径
	},
}
```

## 常见使用场景

### 场景 1：带环境变量的启动

这是最常见的场景，需要在启动主程序前设置特定的环境变量：

```go
config := &t_daemon.SDaemonConfig{
	MainBinaryPath: "./application/ems",
	WorkDirectory:  "./application",

	EnvironmentVariables: map[string]string{
		"DYLD_LIBRARY_PATH": "/Users/zhao/Documents/01.Code/Hex/ems-plan/@cpp/build",
		"LANG": "zh_CN.UTF-8",
	},

	MaxRestarts:   3,
	RestartWindow: 300,
	RestartDelay:  5,
}

manager := t_daemon.NewDaemonManagerInstance()
manager.Start(ctx, config)
```

### 场景 2：启用自动更新

```go
config := &t_daemon.SDaemonConfig{
	MainBinaryPath: "./application/ems",

	// 启用自动更新
	AutoUpdate:           true,
	VersionCheckUrl:      "http://api.example.com/version",
	VersionCheckInterval: 3600, // 每小时检查一次
	BackupBinaryPath:     "./out/ems.backup",

	MaxRestarts:   3,
	RestartWindow: 300,
	RestartDelay:  5,
}

manager := t_daemon.NewDaemonManagerInstance()
manager.Start(ctx, config)
```

### 场景 3：Web 端监控和管理

```go
// 启动守护进程后，可以通过 API 进行管理
manager := t_daemon.NewDaemonManagerInstance()
manager.Start(ctx, config)

// 检查主程序状态
if manager.CheckMainProgramHealth(ctx) {
	fmt.Println("主程序运行正常")
} else {
	fmt.Println("主程序异常，尝试重启")
	manager.Restart(ctx)
}

// 获取统计信息
fmt.Printf("重启次数: %d\n", manager.GetRestartCount())
fmt.Printf("主程序 PID: %d\n", manager.GetMainProgramPid())
```

## API 参考

### IDaemonManager 接口

主要操作方法：

| 方法 | 说明 | 示例 |
|------|------|------|
| `Start(ctx, config)` | 启动守护进程 | `manager.Start(ctx, config)` |
| `Stop(ctx)` | 停止守护进程 | `manager.Stop(ctx)` |
| `IsRunning()` | 检查守护进程是否运行 | `if manager.IsRunning() {...}` |
| `Restart(ctx)` | 重启主程序 | `manager.Restart(ctx)` |
| `CheckMainProgramHealth(ctx)` | 检查主程序健康状态 | `if manager.CheckMainProgramHealth(ctx) {...}` |
| `GetMainProgramPid()` | 获取主程序 PID | `pid := manager.GetMainProgramPid()` |
| `GetRestartCount()` | 获取重启次数 | `count := manager.GetRestartCount()` |
| `UpdateMainBinary(ctx, path)` | 更新主程序 | `manager.UpdateMainBinary(ctx, newPath)` |
| `RollbackMainBinary(ctx)` | 回滚主程序 | `manager.RollbackMainBinary(ctx)` |

## 平台特定的注意事项

### macOS

```go
EnvironmentVariables: map[string]string{
	"DYLD_LIBRARY_PATH": "/path/to/libs",
}
```

### Linux

```go
EnvironmentVariables: map[string]string{
	"LD_LIBRARY_PATH": "/path/to/libs",
}
```

### Windows

```go
EnvironmentVariables: map[string]string{
	"PATH": "C:\\path\\to\\libs;" + os.Getenv("PATH"),
}
```

## 日志输出

守护进程会输出详细的日志信息，包括：

- 启动/停止事件
- 主程序重启事件
- 环境变量设置
- 错误和警告

日志示例：

```
2024-10-20 15:30:45 启动守护进程，主程序: ./application/ems
2024-10-20 15:30:45 设置环境变量: DYLD_LIBRARY_PATH=/path/to/libs
2024-10-20 15:30:45 守护进程启动成功，主程序PID: 12345
2024-10-20 15:30:50 主程序已退出，退出码: 1
2024-10-20 15:30:55 准备重启主程序
2024-10-20 15:30:56 主程序重启成功，新PID: 12346
```

## 故障排查

### 主程序不能启动

**症状**：守护进程启动但主程序没有运行

**解决方案**：
1. 检查 `MainBinaryPath` 是否正确
2. 检查主程序是否有执行权限：`chmod +x ./application/ems`
3. 检查 `WorkDirectory` 是否存在
4. 查看日志输出了解具体错误

### 环境变量没有生效

**症状**：主程序无法找到库文件

**解决方案**：
1. 确认 `EnvironmentVariables` 配置正确
2. 检查路径是否存在：`ls /path/to/libs`
3. 检查是否使用了正确的环境变量名（macOS 用 DYLD_LIBRARY_PATH，Linux 用 LD_LIBRARY_PATH）
4. 确认没有拼写错误

### 主程序频繁重启

**症状**：主程序不断重启

**解决方案**：
1. 检查主程序是否存在 bug，导致立即退出
2. 增加 `RestartDelay` 给主程序更多启动时间
3. 检查是否正确设置了 `MaxRestarts` 和 `RestartWindow`
4. 查看主程序日志了解具体错误

### 守护进程停止了

**症状**：守护进程自己停止了

**解决方案**：
1. 检查重启次数是否超过限制（在 5 分钟内超过 3 次）
2. 增加 `MaxRestarts` 的值
3. 增加 `RestartWindow` 的时间窗口
4. 检查是否有内存或其他资源耗尽的问题

## 最佳实践

1. **设置合理的重启策略**
   - 不要设置太短的时间窗口
   - 不要设置太少的最大重启次数
   - 示例：`MaxRestarts=5, RestartWindow=600`（10分钟内最多5次）

2. **使用适当的延迟**
   - `RestartDelay` 应该至少为 5 秒
   - 如果主程序启动较慢，增加这个值
   - 这样可以避免资源耗尽

3. **记录日志**
   - 始终将日志输出到文件
   - 定期检查日志以发现问题
   - 使用适当的日志级别

4. **监控和告警**
   - 定期检查守护进程状态
   - 设置告警规则（如连续重启 N 次）
   - 集成到监控系统

5. **测试和验证**
   - 在生产部署前充分测试
   - 测试各种故障场景
   - 验证环境变量设置正确

## 相关文件

- `t_daemon_i.go` - 接口定义
- `t_daemon_config_s.go` - 配置结构体
- `t_daemon_manager_impl_s.go` - 管理器实现
- `t_process_monitor_impl_s.go` - 进程监控器
- `t_environment_manager_impl_s.go` - 环境变量管理器
- `t_version_manager_impl_s.go` - 版本管理器
- `example_daemon_usage.go` - 使用示例
- `README.md` - 详细文档

## 支持

如有问题，请查看：
1. README.md - 详细的 API 文档
2. example_daemon_usage.go - 具体的代码示例
3. 项目日志输出 - 错误信息和诊断信息
