package t_daemon

// SDaemonConfig 守护进程配置结构体，包含守护进程的所有配置参数
type SDaemonConfig struct {
	// MainBinaryPath 主程序的二进制文件路径
	MainBinaryPath string `json:"mainBinaryPath"` // 主程序二进制文件路径

	// WorkDirectory 工作目录
	WorkDirectory string `json:"workDirectory"` // 工作目录

	// LogDirectory 日志目录
	LogDirectory string `json:"logDirectory"` // 日志目录

	// MaxRestarts 最大重启次数（在时间窗口内）
	MaxRestarts int `json:"maxRestarts"` // 最大重启次数

	// RestartWindow 重启时间窗口（秒），超过此时间窗口，重启计数器将重置
	RestartWindow int `json:"restartWindow"` // 重启时间窗口（秒）

	// RestartDelay 重启延迟时间（秒），防止频繁重启
	RestartDelay int `json:"restartDelay"` // 重启延迟时间（秒）

	// EnvironmentVariables 环境变量映射，例如 DYLD_LIBRARY_PATH
	EnvironmentVariables map[string]string `json:"environmentVariables"` // 环境变量

	// MainProgramArgs 主程序的命令行参数
	MainProgramArgs []string `json:"mainProgramArgs"` // 主程序命令行参数

	// DaemonPidFile 守护进程的PID文件路径
	DaemonPidFile string `json:"daemonPidFile"` // 守护进程PID文件

	// MainProgramPidFile 主程序的PID文件路径
	MainProgramPidFile string `json:"mainProgramPidFile"` // 主程序PID文件

	// VersionCheckUrl 版本检查的URL，用于检查是否有新版本可用
	VersionCheckUrl string `json:"versionCheckUrl"` // 版本检查URL

	// VersionCheckInterval 版本检查间隔（秒）
	VersionCheckInterval int `json:"versionCheckInterval"` // 版本检查间隔（秒）

	// AutoUpdate 是否启用自动更新
	AutoUpdate bool `json:"autoUpdate"` // 是否启用自动更新

	// BackupBinaryPath 备份二进制文件的路径
	BackupBinaryPath string `json:"backupBinaryPath"` // 备份二进制文件路径
}

// GetMainBinaryPath 获取主程序的二进制文件路径
func (s *SDaemonConfig) GetMainBinaryPath() string {
	return s.MainBinaryPath
}

// GetWorkDirectory 获取工作目录
func (s *SDaemonConfig) GetWorkDirectory() string {
	return s.WorkDirectory
}

// GetLogDirectory 获取日志目录
func (s *SDaemonConfig) GetLogDirectory() string {
	return s.LogDirectory
}

// GetMaxRestarts 获取最大重启次数
func (s *SDaemonConfig) GetMaxRestarts() int {
	return s.MaxRestarts
}

// GetRestartWindow 获取重启时间窗口
func (s *SDaemonConfig) GetRestartWindow() int {
	return s.RestartWindow
}

// GetRestartDelay 获取重启延迟时间
func (s *SDaemonConfig) GetRestartDelay() int {
	return s.RestartDelay
}

// GetEnvironmentVariables 获取环境变量映射
func (s *SDaemonConfig) GetEnvironmentVariables() map[string]string {
	return s.EnvironmentVariables
}

// GetMainProgramArgs 获取主程序的命令行参数
func (s *SDaemonConfig) GetMainProgramArgs() []string {
	return s.MainProgramArgs
}

// GetDaemonPidFile 获取守护进程的PID文件路径
func (s *SDaemonConfig) GetDaemonPidFile() string {
	return s.DaemonPidFile
}

// GetMainProgramPidFile 获取主程序的PID文件路径
func (s *SDaemonConfig) GetMainProgramPidFile() string {
	return s.MainProgramPidFile
}

// GetVersionCheckUrl 获取版本检查的URL
func (s *SDaemonConfig) GetVersionCheckUrl() string {
	return s.VersionCheckUrl
}

// GetVersionCheckInterval 获取版本检查间隔
func (s *SDaemonConfig) GetVersionCheckInterval() int {
	return s.VersionCheckInterval
}

// GetAutoUpdate 获取是否启用自动更新
func (s *SDaemonConfig) GetAutoUpdate() bool {
	return s.AutoUpdate
}

// GetBackupBinaryPath 获取备份二进制文件的路径
func (s *SDaemonConfig) GetBackupBinaryPath() string {
	return s.BackupBinaryPath
}
