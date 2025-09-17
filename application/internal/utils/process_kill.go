//go:build !windows

package utils

import (
	"os"
	"syscall"
)

// KillProcess 发送终止信号给当前进程
// 在非Windows系统上使用syscall.Kill
func KillProcess() error {
	return syscall.Kill(os.Getpid(), syscall.SIGTERM)
}
