package internal

import (
	"common/c_base"
	"context"
	"s_driver/s_driver_interface"
	"sync"
)

// DriverManager 通用驱动管理器实现
type DriverManager struct {
	deviceCmd *SDeviceCmd
}

var (
	// 编译期断言内部实现满足对外接口
	_ s_driver_interface.IDriverManager = (*DriverManager)(nil)

	driverManagerInstance *DriverManager
	driverManagerInitOnce sync.Once
)

// GetDriverManager 创建驱动管理器
func GetDriverManager() *DriverManager {
	driverManagerInitOnce.Do(func() {
		driverManagerInstance = &DriverManager{
			deviceCmd: &SDeviceCmd{},
		}
	})
	return driverManagerInstance
}

// GetAllDriverNames 获取所有驱动名称
func (dm *DriverManager) GetAllDriverNames() []string {
	return GetAllDriverNames()
}

// GetAllDriversInfo 获取所有驱动的详细信息
func (dm *DriverManager) GetAllDriversInfo(ctx context.Context) []c_base.DriverInfo {
	return GetAllDriversInfo(ctx)
}

// GetDriverInfo 获取指定驱动的详细信息
func (dm *DriverManager) GetDriverInfo(ctx context.Context, driverName string) (*c_base.DriverInfo, error) {
	return GetDriverInfo(ctx, driverName)
}

// GetDriversByType 根据设备类型获取驱动信息
func (dm *DriverManager) GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.DriverInfo {
	return GetDriversByType(ctx, deviceType)
}

// CreateDriver 创建驱动实例
func (dm *DriverManager) CreateDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) (c_base.IDriver, error) {
	defer func() {
		if r := recover(); r != nil {
			// 将panic转换为error
			if err, ok := r.(error); ok {
				panic(err)
			} else {
				panic(r)
			}
		}
	}()

	driver := dm.deviceCmd.getDriver(ctx, deviceConfig)
	return driver, nil
}

// IsDriverAvailable 检查驱动是否可用
func (dm *DriverManager) IsDriverAvailable(ctx context.Context, driverName string) bool {
	driverInfo, err := dm.GetDriverInfo(ctx, driverName)
	if err != nil {
		return false
	}
	return driverInfo.Available
}

// GetDriverDescription 获取驱动描述信息
func (dm *DriverManager) GetDriverDescription(ctx context.Context, driverName string) (*c_base.SDescription, error) {
	driverInfo, err := dm.GetDriverInfo(ctx, driverName)
	if err != nil {
		return nil, err
	}
	return driverInfo.Description, nil
}

// GetSupportedDeviceTypes 获取支持的设备类型列表
func (dm *DriverManager) GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType {
	driversInfo := dm.GetAllDriversInfo(ctx)
	typeMap := make(map[c_base.EDeviceType]bool)

	for _, driver := range driversInfo {
		if driver.Available && driver.Type != "" {
			typeMap[driver.Type] = true
		}
	}

	var deviceTypes []c_base.EDeviceType
	for deviceType := range typeMap {
		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes
}
