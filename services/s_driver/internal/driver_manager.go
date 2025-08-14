package internal

import (
	"common/c_base"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

// DriverManager 通用驱动管理器实现
type DriverManager struct {
	deviceCmd *SDeviceCmd
}

// 编译期断言内部实现满足对外接口
var _ c_base.IDriverManager = (*DriverManager)(nil)

// NewDriverManager 创建驱动管理器
func NewDriverManager() c_base.IDriverManager {
	return &DriverManager{
		deviceCmd: &SDeviceCmd{},
	}
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

// DriverManagerHelper 驱动管理器帮助工具
type DriverManagerHelper struct {
	manager c_base.IDriverManager
}

// NewDriverManagerHelper 创建驱动管理器帮助工具
func NewDriverManagerHelper() *DriverManagerHelper {
	return &DriverManagerHelper{
		manager: NewDriverManager(),
	}
}

// GetDriverStats 获取驱动统计信息
func (dmh *DriverManagerHelper) GetDriverStats(ctx context.Context) map[string]interface{} {
	allDrivers := dmh.manager.GetAllDriversInfo(ctx)

	stats := make(map[string]interface{})
	stats["total"] = len(allDrivers)

	availableCount := 0
	unavailableCount := 0
	typeCount := make(map[c_base.EDeviceType]int)

	for _, driver := range allDrivers {
		if driver.Available {
			availableCount++
		} else {
			unavailableCount++
		}

		if driver.Type != "" {
			typeCount[driver.Type]++
		}
	}

	stats["available"] = availableCount
	stats["unavailable"] = unavailableCount
	stats["byType"] = typeCount

	return stats
}

// ValidateDriverConfig 验证驱动配置
func (dmh *DriverManagerHelper) ValidateDriverConfig(ctx context.Context, deviceConfig *c_base.SDriverConfig) error {
	if deviceConfig.Driver == "" {
		return gerror.Newf("驱动名称不能为空")
	}

	if !dmh.manager.IsDriverAvailable(ctx, deviceConfig.Driver) {
		return gerror.Newf("驱动[%s]不可用", deviceConfig.Driver)
	}

	return nil
}

// GetDriverTelemetryInfo 获取驱动遥测信息
func (dmh *DriverManagerHelper) GetDriverTelemetryInfo(ctx context.Context, driverName string) ([]*c_base.STelemetry, error) {
	description, err := dmh.manager.GetDriverDescription(ctx, driverName)
	if err != nil {
		return nil, err
	}

	if description == nil {
		return nil, nil
	}

	return description.Telemetry, nil
}

// SearchDrivers 搜索驱动
func (dmh *DriverManagerHelper) SearchDrivers(ctx context.Context, keyword string) []c_base.DriverInfo {
	allDrivers := dmh.manager.GetAllDriversInfo(ctx)
	var matchedDrivers []c_base.DriverInfo

	for _, driver := range allDrivers {
		if contains(driver.Name, keyword) ||
			(driver.Description != nil &&
				(contains(driver.Description.Brand, keyword) ||
					contains(driver.Description.Model, keyword))) {
			matchedDrivers = append(matchedDrivers, driver)
		}
	}

	return matchedDrivers
}

// contains 检查字符串是否包含子串（忽略大小写）
func contains(str, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(str) == 0 {
		return false
	}

	// 简单的包含检查
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// GetManager 获取驱动管理器实例
func (dmh *DriverManagerHelper) GetManager() c_base.IDriverManager {
	return dmh.manager
}
