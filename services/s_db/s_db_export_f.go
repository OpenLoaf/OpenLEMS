package s_db

import (
	"s_db/internal"
	"s_db/internal/impl"
	"s_db/s_db_basic"
)

func Init() error {
	return internal.Init()
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

// GetAlarmService 获取告警service对象
func GetAlarmService() s_db_basic.IAlarmService {
	return impl.GetAlarmService()
}

// GetLogService 获取日志service对象
func GetLogService() s_db_basic.ILogService {
	return impl.GetLogService()
}

// GetAutomationService 获取自动化service对象
func GetAutomationService() s_db_basic.IAutomationService {
	return impl.GetAutomationService()
}

// GetEnergyStorageService 获取储能策略service对象
func GetEnergyStorageService() s_db_basic.IEnergyStorageStrategyService {
	return impl.GetEnergyStorageStrategyService()
}

// GetEnergyStorageStrategyService 获取储能策略service对象（已废弃，使用GetEnergyStorageService）
func GetEnergyStorageStrategyService() s_db_basic.IEnergyStorageStrategyService {
	return impl.GetEnergyStorageStrategyService()
}

// GetPriceService 获取电价service对象
func GetPriceService() s_db_basic.IPriceService {
	return impl.GetPriceService()
}
