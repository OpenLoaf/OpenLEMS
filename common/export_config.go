package common

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/internal/internal_config"
	"plugin"
)

// SystemInitConfigInstance 系统启动后需要初始化配置，接收命令行参数
func SystemInitConfigInstance(deviceCfgName, driverPath string) {
	internal_config.NewConfigInstance(deviceCfgName, driverPath)
}

// GetDriverConfig 获取设备配置，失败将直接系统退出
func GetDriverConfig(ctx context.Context) *c_base.SDriverConfig {
	return internal_config.ConfigInstance.GetDriverConfig(ctx)
}

// GetProtocolsConfig 获取协议配置，失败将直接系统退出
func GetProtocolsConfig(ctx context.Context) []*c_base.SProtocolConfig {
	return internal_config.ConfigInstance.GetProtocolsConfig(ctx)
}

// GetProtocolById 根据协议ID获取协议配置,可能为空
func GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
	return internal_config.ConfigInstance.GetProtocolById(ctx, id)
}

// Deprecated: GetLatestDriverPath 获取最新的驱动文件路径，失败将直接系统退出
func GetLatestDriverPath(ctx context.Context, driverName string) string {
	return internal_config.ConfigInstance.GetLatestDriverPath(ctx, driverName)
}

// OpenPlugin 打开插件
func OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, error) {
	return internal_config.ConfigInstance.OpenPlugin(ctx, path, symName...)
}
