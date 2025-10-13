package utils

import (
	"context"
	"os"
	"path/filepath"
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
	// 注意：IsProcessRunning 函数在平台特定的文件中定义
	// 这里我们使用一个通用的检查方法
	process, err := os.FindProcess(pid)
	if err != nil {
		return false, pid, nil
	}

	// 尝试向进程发送信号0来检查进程是否响应
	err = process.Signal(os.Signal(syscall.Signal(0)))
	if err != nil {
		return false, pid, nil
	}

	return true, pid, nil
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
