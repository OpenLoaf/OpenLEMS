//go:build !windows

package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// IsProcessRunning 检查指定PID的进程是否正在运行
// 改进版本：能够正确识别僵尸进程
func IsProcessRunning(pid int) bool {
	// 方法1：使用 syscall.Kill(pid, 0) 检查进程是否存在
	// 注意：信号 0 不发送实际信号，只是检查进程是否存在
	err := syscall.Kill(pid, 0)
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

	// 默认方法：使用原有的 Signal 检查
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	return err == nil
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
