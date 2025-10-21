package t_daemon

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// SEnvironmentManagerImpl 环境变量管理器实现结构体
type SEnvironmentManagerImpl struct {
	originalEnv map[string]string // 保存原始环境变量，用于清理
}

// NewEnvironmentManager 创建一个新的环境变量管理器实例
func NewEnvironmentManager() IEnvironmentManager {
	return &SEnvironmentManagerImpl{
		originalEnv: make(map[string]string),
	}
}

// SetupEnvironment 设置指定的环境变量
func (e *SEnvironmentManagerImpl) SetupEnvironment(ctx context.Context, variables map[string]string) error {
	if err := e.ValidateEnvironment(ctx, variables); err != nil {
		return errors.Wrap(err, "环境变量验证失败")
	}

	// 保存原始环境变量值
	for key := range variables {
		if originalValue, exists := os.LookupEnv(key); exists {
			e.originalEnv[key] = originalValue
		}
	}

	// 设置新的环境变量
	for key, value := range variables {
		if err := os.Setenv(key, value); err != nil {
			return errors.Wrapf(err, "设置环境变量失败: %s=%s", key, value)
		}
		c_log.Infof(ctx, "设置环境变量: %s=%s", key, value)
	}

	return nil
}

// CleanupEnvironment 清理环境变量
func (e *SEnvironmentManagerImpl) CleanupEnvironment(ctx context.Context) error {
	// 恢复原始环境变量
	for key, originalValue := range e.originalEnv {
		if err := os.Setenv(key, originalValue); err != nil {
			return errors.Wrapf(err, "恢复环境变量失败: %s", key)
		}
		c_log.Infof(ctx, "恢复环境变量: %s=%s", key, originalValue)
	}

	return nil
}

// BuildProcessEnv 根据配置构建进程的环境变量列表
func (e *SEnvironmentManagerImpl) BuildProcessEnv(ctx context.Context) []string {
	// 获取当前进程的环境变量
	currentEnv := os.Environ()

	// 将环境变量转换为 map，便于合并
	envMap := make(map[string]string)
	for _, env := range currentEnv {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	// 构建最终的环境变量列表
	var result []string
	for key, value := range envMap {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}

	return result
}

// ValidateEnvironment 验证环境变量的有效性
func (e *SEnvironmentManagerImpl) ValidateEnvironment(ctx context.Context, variables map[string]string) error {
	for key, value := range variables {
		if key == "" {
			return errors.New("环境变量名称不能为空")
		}

		// 特殊处理路径相关的环境变量
		if strings.Contains(key, "PATH") || strings.Contains(key, "path") {
			if err := e.validatePathVariable(ctx, key, value); err != nil {
				return errors.Wrapf(err, "验证路径环境变量失败: %s", key)
			}
		}
	}

	return nil
}

// validatePathVariable 验证路径相关的环境变量
func (e *SEnvironmentManagerImpl) validatePathVariable(ctx context.Context, key string, value string) error {
	// 获取路径分隔符
	separator := string(filepath.ListSeparator)

	// 分割路径
	paths := strings.Split(value, separator)

	for _, path := range paths {
		if path == "" {
			continue
		}

		// 检查路径是否存在（仅对绝对路径进行检查）
		if filepath.IsAbs(path) {
			if _, err := os.Stat(path); err != nil {
				c_log.Warningf(ctx, "路径不存在或无法访问: %s=%s, 错误: %v", key, path, err)
				// 不返回错误，仅记录警告
			}
		}
	}

	return nil
}

// AppendToPathVariable 在现有路径变量后追加新路径
func (e *SEnvironmentManagerImpl) AppendToPathVariable(ctx context.Context, key string, pathToAppend string) error {
	currentValue := os.Getenv(key)
	separator := string(filepath.ListSeparator)

	var newValue string
	if currentValue == "" {
		newValue = pathToAppend
	} else {
		newValue = currentValue + separator + pathToAppend
	}

	return os.Setenv(key, newValue)
}

// PrependToPathVariable 在现有路径变量前添加新路径
func (e *SEnvironmentManagerImpl) PrependToPathVariable(ctx context.Context, key string, pathToPrepend string) error {
	currentValue := os.Getenv(key)
	separator := string(filepath.ListSeparator)

	var newValue string
	if currentValue == "" {
		newValue = pathToPrepend
	} else {
		newValue = pathToPrepend + separator + currentValue
	}

	return os.Setenv(key, newValue)
}

// GetPlatformSpecificLibraryPath 获取平台特定的库路径环境变量名
func (e *SEnvironmentManagerImpl) GetPlatformSpecificLibraryPath() string {
	switch runtime.GOOS {
	case "darwin":
		return "DYLD_LIBRARY_PATH"
	case "linux":
		return "LD_LIBRARY_PATH"
	case "windows":
		return "PATH"
	default:
		return "LD_LIBRARY_PATH"
	}
}
