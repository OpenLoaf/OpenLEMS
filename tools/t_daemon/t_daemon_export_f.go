package t_daemon

// NewDaemonManagerInstance 创建一个新的守护进程管理器实例
// 这是一个工厂函数，用于外部模块创建守护进程管理器
func NewDaemonManagerInstance() IDaemonManager {
	return NewDaemonManager()
}

// NewDaemonConfigInstance 创建一个新的守护进程配置实例
// 这是一个工厂函数，用于外部模块创建守护进程配置
func NewDaemonConfigInstance() IDaemonConfig {
	return &SDaemonConfig{
		MaxRestarts:          3,
		RestartWindow:        300,  // 5分钟
		RestartDelay:         5,    // 5秒
		VersionCheckInterval: 3600, // 1小时
		AutoUpdate:           false,
		EnvironmentVariables: make(map[string]string),
		MainProgramArgs:      make([]string, 0),
	}
}

// NewProcessMonitorInstance 创建一个新的进程监控实例
func NewProcessMonitorInstance() IProcessMonitor {
	return NewProcessMonitor()
}

// NewVersionManagerInstance 创建一个新的版本管理器实例
func NewVersionManagerInstance() IVersionManager {
	return NewVersionManager()
}

// NewEnvironmentManagerInstance 创建一个新的环境变量管理器实例
func NewEnvironmentManagerInstance() IEnvironmentManager {
	return NewEnvironmentManager()
}
