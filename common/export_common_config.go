package common

import (
	"common/c_base"
	"common/internal/internal_config_with_file"
	"common/internal/internal_config_with_sqlite"

	"context"
	"fmt"
	"plugin"

	"github.com/gogf/gf/v2/os/gcmd"
)

// SystemInitConfigInstance 系统启动后需要初始化配置，接收命令行参数
func SystemInitConfigInstance(ctx context.Context) {
	all := gcmd.GetOptAll()
	fmt.Println(all)
	internal_config_with_sqlite.NewConfigInstance(ctx, 1)
}

// GetDriverConfig 获取设备配置，失败将直接系统退出
func GetDriverConfig(ctx context.Context) *c_base.SDriverConfig {
	return internal_config_with_sqlite.ConfigInstance.GetDriverConfig(ctx)
}

// GetStorageConfig 获取存储配置，失败将直接系统退出
func GetStorageConfig(ctx context.Context) *c_base.SStorageConfig {
	return internal_config_with_file.ConfigInstance.GetStorageConfig(ctx)
}

// GetProtocolsConfigList 获取协议配置，失败将直接系统退出
func GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig {
	return internal_config_with_sqlite.ConfigInstance.GetProtocolsConfig(ctx)
}

// GetProtocolsConfigMap 获取协议配置，失败将直接系统退出
func GetProtocolsConfigMap(ctx context.Context) map[string]*c_base.SProtocolConfig {
	list := internal_config_with_file.ConfigInstance.GetProtocolsConfig(ctx)
	m := make(map[string]*c_base.SProtocolConfig)
	for _, protocol := range list {
		m[protocol.Id] = protocol
	}
	return m
}

// GetProtocolById 根据协议ID获取协议配置,可能为空
func GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
	return internal_config_with_file.ConfigInstance.GetProtocolById(ctx, id)
}

// Deprecated: GetLatestDriverPath 获取最新的驱动文件路径，失败将直接系统退出
func GetLatestDriverPath(ctx context.Context, driverName string) string {
	return internal_config_with_file.ConfigInstance.GetLatestDriverPath(ctx, driverName)
}

// OpenPlugin 打开插件
func OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, error) {
	return internal_config_with_file.ConfigInstance.OpenPlugin(ctx, path, symName...)
}
