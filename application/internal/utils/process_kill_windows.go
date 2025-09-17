//go:build windows

package utils

import (
	"os"
	"syscall"
)

// KillProcess 发送终止信号给当前进程
// 在Windows系统上使用syscall.Kill，但使用Windows特定的信号
func KillProcess() error {
	// Windows上使用SIGTERM，但需要导入syscall包
	// 在Windows上，SIGTERM实际上会触发进程退出
	// 注意：在Windows上，syscall.Kill可能不可用，我们使用os.Exit作为备选方案
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		// 如果无法找到进程，直接退出
		os.Exit(0)
		return nil
	}

	// 尝试发送SIGTERM信号
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		// 如果发送信号失败，直接退出
		os.Exit(0)
		return nil
	}

	return nil
}
