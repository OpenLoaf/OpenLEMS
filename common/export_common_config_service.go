package common

import (
	"common/c_base"
	"common/internal/internal_config_with_file"
	"context"
	"plugin"
)

type IDriverConfigService interface {
	// SystemInitConfigInstance 系统启动后需要初始化配置，接收命令行参数
	SystemInitConfigInstance(ctx context.Context)

	// GetDriverConfig 获取设备配置，失败将直接系统退出
	GetDriverConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig

	// GetStorageConfig 获取存储配置，失败将直接系统退出
	GetStorageConfig(ctx context.Context) *c_base.SStorageConfig

	// GetProtocolsConfigList 获取协议配置，失败将直接系统退出
	GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig

	// GetProtocolsConfigMap 获取协议配置，失败将直接系统退出
	GetProtocolsConfigMap(ctx context.Context) map[string]*c_base.SProtocolConfig

	// GetProtocolById 根据协议ID获取协议配置,可能为空
	GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig
}

var driverConfigInstance IDriverConfigService

// RegisterDriverConfig 注册服务
func RegisterDriverConfig(d IDriverConfigService) {
	driverConfigInstance = d
}

// GetDriverConfigService 获取驱动配置服务
func GetDriverConfigService() IDriverConfigService {
	if driverConfigInstance == nil {
		panic("StringProcessor not initialized. Call mylib.Init() first.")
	}
	return driverConfigInstance
}

//// SystemInitConfigInstance 系统启动后需要初始化配置，接收命令行参数
//func SystemInitConfigInstance(ctx context.Context) {
//	all := gcmd.GetOptAll()
//	fmt.Println(all)
//	internal_config_with_sqlite.NewConfigInstance(ctx)
//}
//
//// GetDriverConfig 获取设备配置，失败将直接系统退出
//func GetDriverConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig {
//	return internal_config_with_sqlite.ConfigInstance.GetDriverConfig(ctx, activeDeviceRootId)
//}
//
//// GetStorageConfig 获取存储配置，失败将直接系统退出
//func GetStorageConfig(ctx context.Context) *c_base.SStorageConfig {
//	return internal_config_with_file.ConfigInstance.GetStorageConfig(ctx)
//}

//// GetProtocolsConfigList 获取协议配置，失败将直接系统退出
//func GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig {
//	return internal_config_with_sqlite.ConfigInstance.GetProtocolsConfig(ctx)
//}
//
//// GetProtocolsConfigMap 获取协议配置，失败将直接系统退出
//func GetProtocolsConfigMap(ctx context.Context) map[string]*c_base.SProtocolConfig {
//	list := internal_config_with_file.ConfigInstance.GetProtocolsConfig(ctx)
//	m := make(map[string]*c_base.SProtocolConfig)
//	for _, protocol := range list {
//		m[protocol.Id] = protocol
//	}
//	return m
//}
//
//// GetProtocolById 根据协议ID获取协议配置,可能为空
//func GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
//	return internal_config_with_file.ConfigInstance.GetProtocolById(ctx, id)
//}
//
//// Deprecated: GetLatestDriverPath 获取最新的驱动文件路径，失败将直接系统退出
//func GetLatestDriverPath(ctx context.Context, driverName string) string {
//	return internal_config_with_file.ConfigInstance.GetLatestDriverPath(ctx, driverName)
//}

// OpenPlugin 打开插件
func OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, error) {
	return internal_config_with_file.ConfigInstance.OpenPlugin(ctx, path, symName...)
}
