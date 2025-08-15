package common

import (
	"common/c_base"
	"common/internal/internal_config_with_file"
	"context"
	"plugin"
)

// IDriverConfigServ 获取驱动配置服务类(依赖注入)
type IDriverConfigServ interface {

	// GetDriverConfig 获取设备配置，失败将直接系统退出
	GetDriverConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig

	// GetStorageConfig 获取存储配置，失败将直接系统退出
	//GetStorageConfig(ctx context.Context) *c_base.SStorageConfig

	// GetProtocolsConfigList 获取协议配置，失败将直接系统退出
	GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig

	// GetProtocolsConfigMap 获取协议配置，失败将直接系统退出
	GetProtocolsConfigMap(ctx context.Context) map[string]*c_base.SProtocolConfig

	// GetProtocolById 根据协议ID获取协议配置,可能为空
	GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig
}

var driverConfigInstance IDriverConfigServ

// RegisterDriverConfig 注册服务
func RegisterDriverConfig(d IDriverConfigServ) {
	driverConfigInstance = d
}

// GetDriverConfigServ 获取驱动配置服务
func GetDriverConfigServ() IDriverConfigServ {
	if driverConfigInstance == nil {
		panic("StringProcessor not initialized. Call mylib.Init() first.")
	}
	return driverConfigInstance
}

//
//// Deprecated: GetLatestDriverPath 获取最新的驱动文件路径，失败将直接系统退出
//func GetLatestDriverPath(ctx context.Context, driverName string) string {
//	return internal_config_with_file.ConfigInstance.GetLatestDriverPath(ctx, driverName)
//}

// OpenPlugin 打开插件
func OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, string, error) {
	return internal_config_with_file.ConfigInstance.OpenPlugin(ctx, path, symName...)
}
