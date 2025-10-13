//go:build windows

package utils

import (
	"os"
	"syscall"
)

// checkProcessStatusWindows 在Windows系统上检查进程状态
// 通过Windows API检查进程是否存在且正在运行
func checkProcessStatusWindows(pid int) bool {
	// 在Windows上，我们使用os.FindProcess来检查进程是否存在
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// 尝试向进程发送信号0来检查进程是否响应
	// 在Windows上，这通常能正确检测进程是否运行
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}

	// 额外检查：尝试获取进程状态
	// 在Windows上，如果进程不存在或已退出，Signal会失败
	// 如果进程存在但已退出，Signal也可能失败
	return true
}
