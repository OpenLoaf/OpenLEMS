package utils

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// WritePidFile 将PID写入文件
func WritePidFile(pidFile string, pid int) error {
	// 确保目录存在
	dir := filepath.Dir(pidFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Errorf("创建PID文件目录失败: %v", err)
	}

	// 写入PID文件
	if err := os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return errors.Errorf("写入PID文件失败: %v", err)
	}

	return nil
}

// RemovePidFile 删除PID文件
func RemovePidFile(pidFile string) error {
	if _, err := os.Stat(pidFile); err == nil {
		if err := os.Remove(pidFile); err != nil {
			return errors.Errorf("删除PID文件失败: %v", err)
		}
	}
	return nil
}

// IsProcessRunning 检查指定PID的进程是否正在运行
// 改进版本：能够正确识别僵尸进程，支持跨平台
func IsProcessRunning(pid int) bool {
	// 方法1：使用跨平台的 os.FindProcess + process.Signal 检查进程是否存在
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// 使用 Signal(0) 检查进程是否存在（跨平台兼容）
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}

	// 方法2：在支持 /proc 的系统上检查进程状态
	if runtime.GOOS == "linux" {
		return checkProcessStatusLinux(pid)
	}

	// 方法3：在 macOS 上使用进程信息检查
	if runtime.GOOS == "darwin" {
		return checkProcessStatusDarwin(pid)
	}

	// 方法4：在 Windows 上使用特定检查
	if runtime.GOOS == "windows" {
		return checkProcessStatusWindows(pid)
	}

	// 默认情况：如果前面的检查都通过，认为进程在运行
	return true
}

// CheckPidFile 检查PID文件是否存在且对应的进程正在运行
// 返回: (进程是否运行, PID值, 错误信息)
func CheckPidFile(pidFile string) (bool, int, error) {
	// 检查PID文件是否存在
	if _, err := os.Stat(pidFile); os.IsNotExist(err) {
		return false, 0, nil
	}

	// 读取PID文件内容
	content, err := os.ReadFile(pidFile)
	if err != nil {
		return false, 0, errors.Errorf("读取PID文件失败: %v", err)
	}

	// 解析PID
	pidStr := string(content)
	// 去除可能的换行符
	pidStr = strings.TrimSpace(pidStr)
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return false, 0, errors.Errorf("PID文件格式错误: %v", err)
	}

	// 检查进程是否运行
	isRunning := IsProcessRunning(pid)
	return isRunning, pid, nil
}

// WritePidFileWithCheck 检查PID文件后写入新的PID文件
// 如果已有进程运行且force为false，则返回错误
func WritePidFileWithCheck(pidFile string, pid int, force bool) error {
	// 检查是否已有进程运行
	isRunning, existingPid, err := CheckPidFile(pidFile)
	if err != nil {
		return errors.Errorf("检查PID文件失败: %v", err)
	}

	if isRunning {
		if !force {
			return errors.Errorf("进程已在运行中 (PID: %d)，请先停止现有进程或使用 --force 参数强制启动", existingPid)
		}
		// 强制启动时，删除旧的PID文件
		if err := RemovePidFile(pidFile); err != nil {
			return errors.Errorf("删除旧PID文件失败: %v", err)
		}
	}

	// 写入新的PID文件
	return WritePidFile(pidFile, pid)
}

// GetPidFilePath 从配置中获取PID文件路径，如果未配置则使用默认值
func GetPidFilePath(ctx context.Context) string {
	// 从配置中读取PID文件路径
	pidFile := g.Cfg().MustGet(ctx, "pid.file", "out/ems.pid").String()
	return pidFile
}

// checkProcessStatusLinux 在Linux系统上检查进程状态
// 通过读取 /proc/[pid]/stat 文件来检测僵尸进程
func checkProcessStatusLinux(pid int) bool {
	statFile := filepath.Join("/proc", strconv.Itoa(pid), "stat")
	content, err := os.ReadFile(statFile)
	if err != nil {
		return false
	}

	// 解析 stat 文件内容
	// 格式：pid comm state ppid pgrp session tty_nr tpgid flags minflt cminflt majflt cmajflt utime stime cutime cstime priority nice num_threads itrealvalue starttime vsize rss rsslim startcode endcode startstack kstkesp kstkeip signal blocked sigignore sigcatch wchan nswap cnswap exit_signal processor rt_priority policy delayacct_blkio_ticks guest_time cguest_time start_data end_data start_brk arg_start arg_end env_start env_end exit_code
	fields := strings.Fields(string(content))
	if len(fields) < 3 {
		return false
	}

	// 第3个字段是进程状态
	// Z = 僵尸进程，D = 不可中断睡眠，R = 运行，S = 可中断睡眠，T = 停止，W = 分页，X = 死亡
	state := fields[2]

	// 如果是僵尸进程(Z)或死亡进程(X)，认为进程不在运行
	if state == "Z" || state == "X" {
		return false
	}

	return true
}

// checkProcessStatusDarwin 在macOS系统上检查进程状态
// 通过执行 ps 命令来检测僵尸进程
func checkProcessStatusDarwin(pid int) bool {
	// 使用 ps 命令检查进程状态
	// 僵尸进程的状态为 Z，我们通过检查状态来过滤
	cmd := exec.Command("ps", "-o", "pid,stat", "-p", strconv.Itoa(pid))
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return false
	}

	// 解析输出，检查状态字段
	// 输出格式：PID STAT
	// 例如：68558 Z
	fields := strings.Fields(lines[1])
	if len(fields) < 2 {
		return false
	}

	state := fields[1]
	// 如果状态包含 Z（僵尸），认为进程不在运行
	if strings.Contains(state, "Z") {
		return false
	}

	return true
}

// checkProcessStatusWindows 在Windows系统上检查进程状态
// 通过执行 tasklist 命令来检测进程是否存在
func checkProcessStatusWindows(pid int) bool {
	// 使用 tasklist 命令检查进程是否存在
	// /FI "PID eq [pid]" 过滤指定PID的进程
	cmd := exec.Command("tasklist", "/FI", "PID eq "+strconv.Itoa(pid), "/FO", "CSV")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	// 解析输出，检查是否包含目标PID
	// 输出格式：CSV格式，包含进程信息
	lines := strings.Split(string(output), "\n")

	// 如果输出行数大于1（除了标题行），说明找到了进程
	// 检查输出中是否包含目标PID
	pidStr := strconv.Itoa(pid)
	for _, line := range lines {
		if strings.Contains(line, pidStr) && !strings.Contains(line, "PID") {
			return true
		}
	}

	return false
}
