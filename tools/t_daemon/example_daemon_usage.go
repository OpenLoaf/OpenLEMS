package t_daemon

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// ExampleDaemonUsage 演示守护进程的基本使用方式
// 此函数展示了如何创建和配置一个守护进程来管理 EMS 主程序
func ExampleDaemonUsage() {
	// 创建上下文
	ctx := context.Background()

	// 第一步：创建守护进程配置
	config := createDaemonConfig()

	// 第二步：创建守护进程管理器
	manager := NewDaemonManagerInstance()

	// 第三步：启动守护进程
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动守护进程失败: %v\n", err)
		return
	}

	fmt.Println("守护进程已启动")

	// 守护进程现在在后台运行，会自动监控和重启主程序
	// 主程序的 PID
	fmt.Printf("主程序 PID: %d\n", manager.GetMainProgramPid())

	// 可以通过以下方式进行操作：
	// - manager.CheckMainProgramHealth() 检查健康状态
	// - manager.Restart() 重启主程序
	// - manager.UpdateMainBinary() 更新二进制文件
	// - manager.Stop() 停止守护进程

	// 关闭守护进程
	// if err := manager.Stop(ctx); err != nil {
	//     fmt.Printf("停止守护进程失败: %v\n", err)
	// }
}

// createDaemonConfig 创建一个示例的守护进程配置
func createDaemonConfig() IDaemonConfig {
	// 获取当前工作目录
	workDir, _ := os.Getwd()

	config := &SDaemonConfig{
		// 主程序配置
		MainBinaryPath:     filepath.Join(workDir, "application", "ems"),
		WorkDirectory:      filepath.Join(workDir, "application"),
		LogDirectory:       filepath.Join(workDir, "out", "logs"),
		DaemonPidFile:      filepath.Join(workDir, "out", "daemon.pid"),
		MainProgramPidFile: filepath.Join(workDir, "out", "ems.pid"),

		// 重启策略
		MaxRestarts:  3,          // 时间窗口内最多重启3次
		RestartWindow: 300,        // 5分钟时间窗口
		RestartDelay: 5,           // 重启延迟5秒

		// 主程序命令行参数
		MainProgramArgs: []string{
			"--web",           // 启动Web服务
			"--profile", "prod", // 生产环境配置
		},

		// 环境变量配置
		// 这些变量会在启动主程序前设置
		EnvironmentVariables: map[string]string{
			// macOS 特定的库路径
			"DYLD_LIBRARY_PATH": filepath.Join(workDir, "@cpp", "build") + ":" +
				"/usr/local/lib:/opt/local/lib",
			// 语言设置
			"LANG": "zh_CN.UTF-8",
		},

		// 版本更新配置
		VersionCheckUrl:      "http://example.com/api/version",
		VersionCheckInterval: 3600, // 1小时检查一次
		AutoUpdate:           false, // 禁用自动更新（可根据需要启用）
		BackupBinaryPath:     filepath.Join(workDir, "out", "ems.backup"),
	}

	return config
}

// ExampleWithEnvironmentVariables 演示如何配置特殊的环境变量
func ExampleWithEnvironmentVariables() {
	ctx := context.Background()

	config := &SDaemonConfig{
		MainBinaryPath: "./application/ems",
		WorkDirectory:  "./application",

		// 配置特定的环境变量
		EnvironmentVariables: map[string]string{
			// 库路径 - 这是最常见的需求
			"DYLD_LIBRARY_PATH": "/path/to/cpp/build",
			"LD_LIBRARY_PATH":   "/path/to/cpp/build",

			// 其他常见变量
			"LANG":           "zh_CN.UTF-8",
			"TZ":             "Asia/Shanghai",
			"GO_DEBUG_LEVEL": "1",
		},

		MaxRestarts:   3,
		RestartWindow: 300,
		RestartDelay:  5,
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
	}
}

// ExampleWithAutoUpdate 演示如何启用自动更新功能
func ExampleWithAutoUpdate() {
	ctx := context.Background()

	config := &SDaemonConfig{
		MainBinaryPath: "./application/ems",
		WorkDirectory:  "./application",

		// 启用自动更新
		AutoUpdate:           true,
		VersionCheckUrl:      "http://api.example.com/version",
		VersionCheckInterval: 3600, // 每小时检查一次

		MaxRestarts:      3,
		RestartWindow:    300,
		RestartDelay:     5,
		BackupBinaryPath: "./out/ems.backup",
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
		return
	}

	// 系统现在会定期检查新版本
	// 如果有新版本可用，可以调用 UpdateMainBinary 进行更新
}

// ExampleManualRestart 演示如何手动重启主程序
func ExampleManualRestart() {
	ctx := context.Background()

	config := &SDaemonConfig{
		MainBinaryPath: "./application/ems",
		MaxRestarts:    3,
		RestartWindow:  300,
		RestartDelay:   5,
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
		return
	}

	// 检查主程序是否健康
	if !manager.CheckMainProgramHealth(ctx) {
		fmt.Println("主程序不健康，准备重启")

		if err := manager.Restart(ctx); err != nil {
			fmt.Printf("重启失败: %v\n", err)
		} else {
			fmt.Printf("重启成功，新 PID: %d\n", manager.GetMainProgramPid())
		}
	}
}

// ExampleBinaryUpdate 演示如何更新主程序二进制文件
func ExampleBinaryUpdate() {
	ctx := context.Background()

	config := &SDaemonConfig{
		MainBinaryPath:   "./application/ems",
		BackupBinaryPath: "./out/ems.backup",
		MaxRestarts:      3,
		RestartWindow:    300,
		RestartDelay:     5,
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
		return
	}

	// 更新主程序
	newBinaryPath := "./out/ems_new"
	if err := manager.UpdateMainBinary(ctx, newBinaryPath); err != nil {
		fmt.Printf("更新失败: %v\n", err)

		// 尝试回滚
		if err := manager.RollbackMainBinary(ctx); err != nil {
			fmt.Printf("回滚也失败了: %v\n", err)
		}
	} else {
		fmt.Printf("更新成功，新 PID: %d\n", manager.GetMainProgramPid())
	}
}

// ExampleMonitoringRestarts 演示如何监控重启统计信息
func ExampleMonitoringRestarts() {
	ctx := context.Background()

	config := &SDaemonConfig{
		MainBinaryPath: "./application/ems",
		MaxRestarts:    3,
		RestartWindow:  300,
		RestartDelay:   5,
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
		return
	}

	// 获取重启统计信息
	restartCount := manager.GetRestartCount()
	lastRestartTime := manager.GetLastRestartTime()

	fmt.Printf("当前时间窗口内的重启次数: %d\n", restartCount)
	fmt.Printf("上次重启时间戳: %d\n", lastRestartTime)
	fmt.Printf("主程序运行状态: %v\n", manager.CheckMainProgramHealth(ctx))
}

// ExamplePlatformSpecific 演示平台特定的配置
func ExamplePlatformSpecific() {
	ctx := context.Background()

	// 根据操作系统设置不同的环境变量
	envVars := make(map[string]string)

	// 根据运行时的操作系统进行条件配置
	// runtime.GOOS 会返回 "darwin"（macOS）、"linux" 或 "windows"

	// 这里只是示例，实际应该根据 runtime.GOOS 进行判断
	envVars["DYLD_LIBRARY_PATH"] = "/path/to/macos/libs"

	config := &SDaemonConfig{
		MainBinaryPath:       "./application/ems",
		EnvironmentVariables: envVars,
		MaxRestarts:          3,
		RestartWindow:        300,
		RestartDelay:         5,
	}

	manager := NewDaemonManagerInstance()
	if err := manager.Start(ctx, config); err != nil {
		fmt.Printf("启动失败: %v\n", err)
	}
}
