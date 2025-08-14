package s_db

import (
	"common"
	"s_db/internal"
	"s_db/internal/impl"
	"s_db/s_db_interface"
)

func Init() {
	internal.Init()
}

// GetDbDriverConfigService 创建基础包中获取驱动配置服务的实现
func GetDbDriverConfigService() common.IDriverConfigServ {
	return impl.GetDriverConfigService()
}

// GetDeviceService 获取设备service对象
func GetDeviceService() s_db_interface.IDeviceService {
	return impl.GetDeviceService()
}

func GetConfigService() s_db_interface.IConfigService {
	return impl.GetConfigService()
}

func GetProtocolService() s_db_interface.IProtocolService {
	return impl.GetProtocolService()
}
