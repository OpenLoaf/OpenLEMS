package c_base

import "context"

// IDriverManager 驱动管理器接口
type IDriverManager interface {
	// GetAllDriverNames 获取所有驱动名称
	GetAllDriverNames() []string

	// GetAllDriversInfo 获取所有驱动的详细信息
	GetAllDriversInfo(ctx context.Context) []DriverInfo

	// GetDriverInfo 获取指定驱动的详细信息
	GetDriverInfo(ctx context.Context, driverName string) (*DriverInfo, error)

	// GetDriversByType 根据设备类型获取驱动信息
	GetDriversByType(ctx context.Context, deviceType EDeviceType) []DriverInfo

	// CreateDriver 创建驱动实例
	CreateDriver(ctx context.Context, deviceConfig *SDriverConfig) (IDriver, error)

	// IsDriverAvailable 检查驱动是否可用
	IsDriverAvailable(ctx context.Context, driverName string) bool

	// GetDriverDescription 获取驱动描述信息
	GetDriverDescription(ctx context.Context, driverName string) (*SDescription, error)

	// GetSupportedDeviceTypes 获取支持的设备类型列表
	GetSupportedDeviceTypes(ctx context.Context) []EDeviceType
}
