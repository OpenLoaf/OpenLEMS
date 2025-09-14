package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// WritePidFile 将PID写入文件
func WritePidFile(pidFile string, pid int) error {
	// 确保目录存在
	dir := filepath.Dir(pidFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建PID文件目录失败: %v", err)
	}

	// 写入PID文件
	if err := os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("写入PID文件失败: %v", err)
	}

	return nil
}

// RemovePidFile 删除PID文件
func RemovePidFile(pidFile string) error {
	if _, err := os.Stat(pidFile); err == nil {
		if err := os.Remove(pidFile); err != nil {
			return fmt.Errorf("删除PID文件失败: %v", err)
		}
	}
	return nil
}
