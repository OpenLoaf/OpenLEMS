package s_driver_interface

import (
	"common/c_base"
	"context"
)

// IDriverManager 驱动管理器接口
type IDriverManager interface {
	// GetAllDriverNames 获取所有驱动名称
	GetAllDriverNames() []string

	// GetAllDriversInfo 获取所有驱动的详细信息
	GetAllDriversInfo(ctx context.Context) []c_base.DriverInfo

	// GetDriverInfo 获取指定驱动的详细信息
	GetDriverInfo(ctx context.Context, driverName string) (*c_base.DriverInfo, error)

	// GetDriversByType 根据设备类型获取驱动信息
	GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.DriverInfo

	// CreateDriver 创建驱动实例
	CreateDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) (c_base.IDriver, error)

	// IsDriverAvailable 检查驱动是否可用
	IsDriverAvailable(ctx context.Context, driverName string) bool

	// GetDriverDescription 获取驱动描述信息
	GetDriverDescription(ctx context.Context, driverName string) (*c_base.SDescription, error)

	// GetSupportedDeviceTypes 获取支持的设备类型列表
	GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType
}
