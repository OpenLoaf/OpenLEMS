package t_daemon

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// SProcessMonitorImpl 进程监控实现结构体
type SProcessMonitorImpl struct {
	monitoring   bool
	monitoredPid int
	onExit       func(ctx context.Context, exitCode int)
	mutex        sync.RWMutex
	stopChan     chan struct{}
	wg           sync.WaitGroup
}

// NewProcessMonitor 创建一个新的进程监控实例
func NewProcessMonitor() IProcessMonitor {
	return &SProcessMonitorImpl{
		stopChan: make(chan struct{}),
	}
}

// Start 启动进程监控
func (p *SProcessMonitorImpl) Start(ctx context.Context, mainPid int, onExit func(ctx context.Context, exitCode int)) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.monitoring {
		return errors.New("进程监控已启动")
	}

	if mainPid <= 0 {
		return errors.New("无效的进程ID")
	}

	p.monitoredPid = mainPid
	p.onExit = onExit
	p.monitoring = true

	p.wg.Add(1)
	go p.monitorProcess(ctx)

	g.Log().Infof(ctx, "启动进程监控，监控PID: %d", mainPid)

	return nil
}

// Stop 停止进程监控
func (p *SProcessMonitorImpl) Stop(ctx context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if !p.monitoring {
		return nil
	}

	g.Log().Infof(ctx, "停止进程监控")

	p.monitoring = false
	close(p.stopChan)
	p.wg.Wait()

	return nil
}

// IsProcessRunning 检查进程是否仍在运行
func (p *SProcessMonitorImpl) IsProcessRunning(pid int) bool {
	if pid <= 0 {
		return false
	}

	// 尝试向进程发送信号0（不实际发送任何信号，但会检查进程是否存在）
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// 在Windows上，FindProcess总是成功的，需要尝试其他方法
	if runtime.GOOS == "windows" {
		// 尝试打开进程句柄来验证进程是否存在
		cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid))
		if err := cmd.Run(); err != nil {
			return false
		}
		return true
	}

	// 在Unix系统上，发送信号0来检查进程
	if err := process.Signal(syscall.Signal(0)); err != nil {
		return false
	}

	return true
}

// KillProcess 杀死指定的进程
func (p *SProcessMonitorImpl) KillProcess(ctx context.Context, pid int, force bool) error {
	if pid <= 0 {
		return errors.New("无效的进程ID")
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return errors.Wrapf(err, "找不到进程: %d", pid)
	}

	g.Log().Infof(ctx, "准备杀死进程 %d (强制: %v)", pid, force)

	if force {
		// 强制杀死进程
		if runtime.GOOS == "windows" {
			// Windows 环境
			cmd := exec.Command("taskkill", "/PID", fmt.Sprintf("%d", pid), "/F")
			if err := cmd.Run(); err != nil {
				return errors.Wrapf(err, "强制杀死进程失败: %d", pid)
			}
		} else {
			// Unix系统
			if err := process.Signal(syscall.SIGKILL); err != nil {
				return errors.Wrapf(err, "强制杀死进程失败: %d", pid)
			}
		}
	} else {
		// 优雅关闭进程
		if runtime.GOOS == "windows" {
			// Windows 环境
			cmd := exec.Command("taskkill", "/PID", fmt.Sprintf("%d", pid))
			if err := cmd.Run(); err != nil {
				// 如果优雅关闭失败，尝试强制杀死
				return p.KillProcess(ctx, pid, true)
			}
		} else {
			// Unix系统
			if err := process.Signal(syscall.SIGTERM); err != nil {
				// 如果优雅关闭失败，尝试强制杀死
				return p.KillProcess(ctx, pid, true)
			}

			// 等待进程关闭
			timeout := time.NewTimer(5 * time.Second)
			defer timeout.Stop()

			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-timeout.C:
					// 超时后强制杀死
					g.Log().Warningf(ctx, "进程优雅关闭超时，准备强制杀死")
					return p.KillProcess(ctx, pid, true)
				case <-ticker.C:
					if !p.IsProcessRunning(pid) {
						return nil
					}
				}
			}
		}
	}

	return nil
}

// WaitForProcessExit 等待进程退出
func (p *SProcessMonitorImpl) WaitForProcessExit(ctx context.Context, pid int) (int, error) {
	if pid <= 0 {
		return -1, errors.New("无效的进程ID")
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return -1, errors.Wrapf(err, "找不到进程: %d", pid)
	}

	state, err := process.Wait()
	if err != nil {
		return -1, errors.Wrap(err, "等待进程退出失败")
	}

	exitCode := state.ExitCode()

	return exitCode, nil
}

// 私有方法

// monitorProcess 监控进程的运行状态
func (p *SProcessMonitorImpl) monitorProcess(ctx context.Context) {
	defer p.wg.Done()

	p.mutex.RLock()
	monitoredPid := p.monitoredPid
	onExit := p.onExit
	p.mutex.RUnlock()

	// 使用 Wait() 方法等待进程退出
	exitCode, err := p.WaitForProcessExit(ctx, monitoredPid)
	if err != nil {
		g.Log().Errorf(ctx, "监控进程失败: %v", err)
		if onExit != nil {
			onExit(ctx, -1)
		}
		return
	}

	g.Log().Infof(ctx, "监控进程已退出，PID: %d, 退出码: %d", monitoredPid, exitCode)

	if onExit != nil {
		onExit(ctx, exitCode)
	}
}
