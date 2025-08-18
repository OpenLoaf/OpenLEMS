package s_db

import (
	"s_db/basic"
	"s_db/internal"
	"s_db/internal/impl"
)

func Init() {
	internal.Init()
}

// GetDeviceService 获取设备service对象
func GetDeviceService() basic.IDeviceService {
	return impl.GetDeviceService()
}

func GetSettingService() basic.ISettingService {
	return impl.GetSettingService()
}

func GetProtocolService() basic.IProtocolService {
	return impl.GetProtocolService()
}
