package t_daemon

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// SDaemonManagerImpl 守护进程管理器实现结构体
type SDaemonManagerImpl struct {
	config              IDaemonConfig
	mainProgramPid      int
	running             bool
	restartTimes        []int64 // 记录重启时间戳（毫秒）
	lastRestartTime     int64
	mutex               sync.RWMutex
	stopChan            chan struct{}
	processMonitor      IProcessMonitor
	environmentManager  IEnvironmentManager
	versionManager      IVersionManager
}

// NewDaemonManager 创建一个新的守护进程管理器实例
func NewDaemonManager() IDaemonManager {
	return &SDaemonManagerImpl{
		restartTimes: make([]int64, 0),
		stopChan:     make(chan struct{}),
	}
}

// Start 启动守护进程
func (d *SDaemonManagerImpl) Start(ctx context.Context, config IDaemonConfig) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.running {
		return errors.New("守护进程已经运行中")
	}

	d.config = config

	// 验证配置
	if err := d.validateConfig(ctx); err != nil {
		return errors.Wrap(err, "配置验证失败")
	}

	g.Log().Infof(ctx, "启动守护进程，主程序: %s", config.GetMainBinaryPath())

	// 初始化环境管理器
	d.environmentManager = NewEnvironmentManager()
	if err := d.environmentManager.SetupEnvironment(ctx, config.GetEnvironmentVariables()); err != nil {
		return errors.Wrap(err, "环境变量设置失败")
	}

	// 初始化版本管理器
	d.versionManager = NewVersionManager()

	// 初始化进程监控器
	d.processMonitor = NewProcessMonitor()

	// 启动主程序
	if err := d.startMainProgram(ctx); err != nil {
		return errors.Wrap(err, "启动主程序失败")
	}

	d.running = true

	// 启动监控 goroutine
	go d.monitorMainProgram(ctx)

	// 启动版本检查 goroutine（如果启用自动更新）
	if config.GetAutoUpdate() {
		go d.checkForUpdates(ctx)
	}

	// 设置信号处理
	d.setupSignalHandlers(ctx)

	g.Log().Infof(ctx, "守护进程启动成功，主程序PID: %d", d.mainProgramPid)

	return nil
}

// Stop 停止守护进程
func (d *SDaemonManagerImpl) Stop(ctx context.Context) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.running {
		return errors.New("守护进程未运行")
	}

	g.Log().Infof(ctx, "停止守护进程")

	// 停止进程监控
	if d.processMonitor != nil {
		if err := d.processMonitor.Stop(ctx); err != nil {
			g.Log().Warningf(ctx, "停止进程监控失败: %v", err)
		}
	}

	// 停止主程序
	if err := d.processMonitor.KillProcess(ctx, d.mainProgramPid, false); err != nil {
		g.Log().Warningf(ctx, "停止主程序失败: %v", err)
	}

	// 清理环境变量
	if d.environmentManager != nil {
		if err := d.environmentManager.CleanupEnvironment(ctx); err != nil {
			g.Log().Warningf(ctx, "清理环境变量失败: %v", err)
		}
	}

	d.running = false
	close(d.stopChan)

	return nil
}

// Shutdown 优雅关闭守护进程
func (d *SDaemonManagerImpl) Shutdown(ctx context.Context) error {
	return d.Stop(ctx)
}

// IsRunning 检查守护进程是否正在运行
func (d *SDaemonManagerImpl) IsRunning() bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.running
}

// Restart 重启主程序
func (d *SDaemonManagerImpl) Restart(ctx context.Context) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.running {
		return errors.New("守护进程未运行")
	}

	// 检查重启频率
	if !d.canRestart() {
		return errors.New("重启次数过于频繁，请稍后再试")
	}

	g.Log().Infof(ctx, "准备重启主程序，当前PID: %d", d.mainProgramPid)

	// 停止当前的主程序
	if err := d.processMonitor.KillProcess(ctx, d.mainProgramPid, false); err != nil {
		g.Log().Warningf(ctx, "停止旧程序失败: %v", err)
	}

	// 等待一段时间后重启
	restartDelay := time.Duration(d.config.GetRestartDelay()) * time.Second
	time.Sleep(restartDelay)

	// 启动新的主程序
	if err := d.startMainProgram(ctx); err != nil {
		return errors.Wrap(err, "重启主程序失败")
	}

	g.Log().Infof(ctx, "主程序重启成功，新PID: %d", d.mainProgramPid)

	return nil
}

// CheckMainProgramHealth 检查主程序的健康状态
func (d *SDaemonManagerImpl) CheckMainProgramHealth(ctx context.Context) bool {
	d.mutex.RLock()
	mainPid := d.mainProgramPid
	d.mutex.RUnlock()

	if mainPid == 0 {
		return false
	}

	return d.processMonitor.IsProcessRunning(mainPid)
}

// GetMainProgramPid 获取主程序的进程ID
func (d *SDaemonManagerImpl) GetMainProgramPid() int {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.mainProgramPid
}

// GetRestartCount 获取在当前时间窗口内的重启次数
func (d *SDaemonManagerImpl) GetRestartCount() int {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	now := time.Now().UnixMilli()
	windowStart := now - int64(d.config.GetRestartWindow())*1000

	count := 0
	for _, t := range d.restartTimes {
		if t >= windowStart {
			count++
		}
	}

	return count
}

// GetLastRestartTime 获取上一次重启的时间戳（毫秒）
func (d *SDaemonManagerImpl) GetLastRestartTime() int64 {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.lastRestartTime
}

// UpdateMainBinary 更新主程序二进制文件
func (d *SDaemonManagerImpl) UpdateMainBinary(ctx context.Context, newBinaryPath string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.running {
		return errors.New("守护进程未运行")
	}

	g.Log().Infof(ctx, "准备更新主程序二进制文件: %s", newBinaryPath)

	// 备份当前的二进制文件
	if err := d.backupCurrentBinary(ctx); err != nil {
		return errors.Wrap(err, "备份当前二进制文件失败")
	}

	// 应用更新
	if err := d.versionManager.ApplyUpdate(ctx, newBinaryPath, d.config.GetMainBinaryPath()); err != nil {
		return errors.Wrap(err, "应用更新失败")
	}

	// 重启主程序以应用更新
	d.mainProgramPid = 0 // 重置PID，以便重启
	if err := d.startMainProgram(ctx); err != nil {
		// 如果启动失败，回滚更新
		if rollbackErr := d.versionManager.ApplyUpdate(ctx, d.config.GetBackupBinaryPath(), d.config.GetMainBinaryPath()); rollbackErr != nil {
			g.Log().Errorf(ctx, "回滚二进制文件失败: %v", rollbackErr)
		}
		return errors.Wrap(err, "启动新程序失败")
	}

	g.Log().Infof(ctx, "主程序更新成功，新PID: %d", d.mainProgramPid)

	return nil
}

// RollbackMainBinary 回滚主程序二进制文件到备份版本
func (d *SDaemonManagerImpl) RollbackMainBinary(ctx context.Context) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.running {
		return errors.New("守护进程未运行")
	}

	g.Log().Infof(ctx, "准备回滚主程序二进制文件")

	// 停止当前程序
	if err := d.processMonitor.KillProcess(ctx, d.mainProgramPid, false); err != nil {
		g.Log().Warningf(ctx, "停止旧程序失败: %v", err)
	}

	// 应用备份版本
	if err := d.versionManager.ApplyUpdate(ctx, d.config.GetBackupBinaryPath(), d.config.GetMainBinaryPath()); err != nil {
		return errors.Wrap(err, "应用备份版本失败")
	}

	// 重启主程序
	if err := d.startMainProgram(ctx); err != nil {
		return errors.Wrap(err, "启动程序失败")
	}

	g.Log().Infof(ctx, "主程序回滚成功，新PID: %d", d.mainProgramPid)

	return nil
}

// 私有方法

// validateConfig 验证配置的有效性
func (d *SDaemonManagerImpl) validateConfig(ctx context.Context) error {
	if d.config.GetMainBinaryPath() == "" {
		return errors.New("主程序路径不能为空")
	}

	if _, err := os.Stat(d.config.GetMainBinaryPath()); err != nil {
		return errors.Wrapf(err, "主程序文件不存在: %s", d.config.GetMainBinaryPath())
	}

	if d.config.GetMaxRestarts() < 1 {
		return errors.New("最大重启次数必须至少为1")
	}

	if d.config.GetRestartWindow() < 1 {
		return errors.New("重启时间窗口必须至少为1秒")
	}

	if d.config.GetRestartDelay() < 0 {
		return errors.New("重启延迟时间不能为负数")
	}

	return nil
}

// startMainProgram 启动主程序
func (d *SDaemonManagerImpl) startMainProgram(ctx context.Context) error {
	// 构建进程环境变量
	env := d.environmentManager.BuildProcessEnv(ctx)

	// 创建命令
	cmd := exec.CommandContext(ctx, d.config.GetMainBinaryPath(), d.config.GetMainProgramArgs()...)
	cmd.Env = env
	cmd.Dir = d.config.GetWorkDirectory()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 启动进程
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "启动进程失败")
	}

	d.mainProgramPid = cmd.Process.Pid

	// 记录重启时间
	d.recordRestart()

	// 启动进程监控
	if err := d.processMonitor.Start(ctx, d.mainProgramPid, d.onMainProgramExit); err != nil {
		return errors.Wrap(err, "启动进程监控失败")
	}

	return nil
}

// canRestart 检查是否可以重启（不超过最大重启次数）
func (d *SDaemonManagerImpl) canRestart() bool {
	now := time.Now().UnixMilli()
	windowStart := now - int64(d.config.GetRestartWindow())*1000

	count := 0
	for _, t := range d.restartTimes {
		if t >= windowStart {
			count++
		}
	}

	return count < d.config.GetMaxRestarts()
}

// recordRestart 记录一次重启事件
func (d *SDaemonManagerImpl) recordRestart() {
	now := time.Now().UnixMilli()
	d.restartTimes = append(d.restartTimes, now)
	d.lastRestartTime = now

	// 清理过期的重启记录
	windowStart := now - int64(d.config.GetRestartWindow())*1000
	validTimes := make([]int64, 0)
	for _, t := range d.restartTimes {
		if t >= windowStart {
			validTimes = append(validTimes, t)
		}
	}
	d.restartTimes = validTimes
}

// backupCurrentBinary 备份当前的二进制文件
func (d *SDaemonManagerImpl) backupCurrentBinary(ctx context.Context) error {
	currentBinaryPath := d.config.GetMainBinaryPath()
	_, err := d.versionManager.BackupCurrentBinary(ctx, currentBinaryPath)
	return err
}

// onMainProgramExit 主程序退出时的回调函数
func (d *SDaemonManagerImpl) onMainProgramExit(ctx context.Context, exitCode int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	g.Log().Warningf(ctx, "主程序已退出，退出码: %d", exitCode)

	if !d.running {
		return
	}

	// 检查是否可以重启
	if !d.canRestart() {
		g.Log().Errorf(ctx, "重启次数过于频繁，停止守护进程")
		d.running = false
		return
	}

	// 延迟后自动重启
	delay := time.Duration(d.config.GetRestartDelay()) * time.Second
	g.Log().Infof(ctx, "将在 %d 秒后自动重启主程序", d.config.GetRestartDelay())

	go func() {
		time.Sleep(delay)
		d.mutex.Lock()
		defer d.mutex.Unlock()

		if d.running && d.canRestart() {
			if err := d.startMainProgram(ctx); err != nil {
				g.Log().Errorf(ctx, "自动重启失败: %v", err)
			}
		}
	}()
}

// monitorMainProgram 监控主程序的运行状态
func (d *SDaemonManagerImpl) monitorMainProgram(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-d.stopChan:
			return
		case <-ticker.C:
			d.mutex.RLock()
			if d.running && d.mainProgramPid > 0 {
				isRunning := d.processMonitor.IsProcessRunning(d.mainProgramPid)
				if !isRunning {
					d.mutex.RUnlock()
					// 进程已退出，将由 onMainProgramExit 处理
					continue
				}
			}
			d.mutex.RUnlock()
		}
	}
}

// checkForUpdates 定期检查是否有新版本可用
func (d *SDaemonManagerImpl) checkForUpdates(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(d.config.GetVersionCheckInterval()) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-d.stopChan:
			return
		case <-ticker.C:
			if hasUpdate, _, err := d.versionManager.CheckForUpdates(ctx); err != nil {
				g.Log().Warningf(ctx, "检查版本更新失败: %v", err)
			} else if hasUpdate {
				g.Log().Infof(ctx, "检测到新版本可用")
				// 这里可以选择自动更新或通知管理员
			}
		}
	}
}

// setupSignalHandlers 设置信号处理器
func (d *SDaemonManagerImpl) setupSignalHandlers(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		g.Log().Infof(ctx, "接收到信号: %v", sig)
		d.Stop(ctx)
	}()
}

// WriteDaemonPid 写入守护进程的PID到文件
func (d *SDaemonManagerImpl) WriteDaemonPid(ctx context.Context) error {
	daemonPidFile := d.config.GetDaemonPidFile()
	if daemonPidFile == "" {
		return nil
	}

	pid := os.Getpid()
	content := strconv.Itoa(pid)

	return os.WriteFile(daemonPidFile, []byte(content), 0644)
}
