package t_daemon

import (
	"context"
)

// IDaemonConfig 守护进程配置接口，定义守护进程的配置参数
type IDaemonConfig interface {
	// GetMainBinaryPath 获取主程序的二进制文件路径
	GetMainBinaryPath() string

	// GetWorkDirectory 获取工作目录
	GetWorkDirectory() string

	// GetLogDirectory 获取日志目录
	GetLogDirectory() string

	// GetMaxRestarts 获取最大重启次数（在时间窗口内）
	GetMaxRestarts() int

	// GetRestartWindow 获取重启时间窗口（秒）
	GetRestartWindow() int

	// GetRestartDelay 获取重启延迟时间（秒）
	GetRestartDelay() int

	// GetEnvironmentVariables 获取环境变量映射
	GetEnvironmentVariables() map[string]string

	// GetMainProgramArgs 获取主程序的命令行参数
	GetMainProgramArgs() []string

	// GetDaemonPidFile 获取守护进程的PID文件路径
	GetDaemonPidFile() string

	// GetMainProgramPidFile 获取主程序的PID文件路径
	GetMainProgramPidFile() string

	// GetVersionCheckUrl 获取版本检查的URL
	GetVersionCheckUrl() string

	// GetVersionCheckInterval 获取版本检查间隔（秒）
	GetVersionCheckInterval() int

	// GetAutoUpdate 获取是否启用自动更新
	GetAutoUpdate() bool

	// GetBackupBinaryPath 获取备份二进制文件的路径
	GetBackupBinaryPath() string
}

// IDaemonManager 守护进程管理器接口，负责主程序的生命周期管理
type IDaemonManager interface {
	// Start 启动守护进程
	Start(ctx context.Context, config IDaemonConfig) error

	// Stop 停止守护进程
	Stop(ctx context.Context) error

	// Shutdown 优雅关闭守护进程
	Shutdown(ctx context.Context) error

	// IsRunning 检查守护进程是否正在运行
	IsRunning() bool

	// Restart 重启主程序
	Restart(ctx context.Context) error

	// CheckMainProgramHealth 检查主程序的健康状态
	CheckMainProgramHealth(ctx context.Context) bool

	// GetMainProgramPid 获取主程序的进程ID
	GetMainProgramPid() int

	// GetRestartCount 获取在当前时间窗口内的重启次数
	GetRestartCount() int

	// GetLastRestartTime 获取上一次重启的时间戳（毫秒）
	GetLastRestartTime() int64

	// UpdateMainBinary 更新主程序二进制文件
	UpdateMainBinary(ctx context.Context, newBinaryPath string) error

	// RollbackMainBinary 回滚主程序二进制文件到备份版本
	RollbackMainBinary(ctx context.Context) error
}

// IProcessMonitor 进程监控接口，负责监控主程序的运行状态
type IProcessMonitor interface {
	// Start 启动进程监控
	Start(ctx context.Context, mainPid int, onExit func(ctx context.Context, exitCode int)) error

	// Stop 停止进程监控
	Stop(ctx context.Context) error

	// IsProcessRunning 检查进程是否仍在运行
	IsProcessRunning(pid int) bool

	// KillProcess 杀死指定的进程
	KillProcess(ctx context.Context, pid int, force bool) error

	// WaitForProcessExit 等待进程退出
	WaitForProcessExit(ctx context.Context, pid int) (int, error)
}

// IVersionManager 版本管理接口，负责主程序的版本检查和更新
type IVersionManager interface {
	// CheckForUpdates 检查是否有新版本可用
	CheckForUpdates(ctx context.Context) (bool, string, error)

	// DownloadUpdate 下载更新的二进制文件
	DownloadUpdate(ctx context.Context, downloadUrl string, targetPath string) error

	// VerifyBinary 验证二进制文件的完整性（使用哈希校验）
	VerifyBinary(ctx context.Context, binaryPath string, expectedHash string) (bool, error)

	// BackupCurrentBinary 备份当前的二进制文件
	BackupCurrentBinary(ctx context.Context, binaryPath string) (string, error)

	// ApplyUpdate 应用更新，替换当前的二进制文件
	ApplyUpdate(ctx context.Context, newBinaryPath string, currentBinaryPath string) error

	// GetCurrentVersion 获取当前二进制文件的版本信息
	GetCurrentVersion(ctx context.Context) (string, error)
}

// IEnvironmentManager 环境变量管理接口，负责设置和管理环境变量
type IEnvironmentManager interface {
	// SetupEnvironment 设置指定的环境变量
	SetupEnvironment(ctx context.Context, variables map[string]string) error

	// CleanupEnvironment 清理环境变量
	CleanupEnvironment(ctx context.Context) error

	// BuildProcessEnv 根据配置构建进程的环境变量列表
	BuildProcessEnv(ctx context.Context) []string

	// ValidateEnvironment 验证环境变量的有效性
	ValidateEnvironment(ctx context.Context, variables map[string]string) error
}
