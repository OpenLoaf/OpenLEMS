package s_db

import (
	"s_db/internal"
	"s_db/internal/impl"
	"s_db/s_db_basic"
)

func Init() {
	internal.Init()
}

// GetDeviceService 获取设备service对象
func GetDeviceService() s_db_basic.IDeviceService {
	return impl.GetDeviceService()
}

func GetSettingService() s_db_basic.ISettingService {
	return impl.GetSettingService()
}

func GetProtocolService() s_db_basic.IProtocolService {
	return impl.GetProtocolService()
}
