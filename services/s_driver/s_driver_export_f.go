package s_driver

import (
	"common"
	"context"
	"s_driver/internal"
)

//type SDeviceManager interface {
//	Start(activeDeviceRootId string) // EMS 服务
//	Shutdown()                           // 停止EMS服务
//
//	initDevice(clientCache map[string]any, config *c_base.SDriverConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDriver // 初始化驱动
//
//	GetAllDriverNames() []string                                                               // 获取所有驱动名称
//	GetAllDriversInfo(ctx context.Context) []c_base.DriverInfo                                 // 获取所有驱动的详细信息
//	GetDriverInfo(ctx context.Context, driverName string) (*c_base.DriverInfo, error)          //获取指定驱动的详细信息
//	GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.DriverInfo   //根据设备类型获取驱动信息
//	GetDriverDescription(ctx context.Context, driverName string) (*c_base.SDriverDescription, error) // 获取驱动描述信息
//	GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType                          // 获取支持的设备类型列表
//
//	CreateDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) (c_base.IDriver, error) // 创建驱动实例
//
//	IsProtocolActive(protocolId string) bool                       // 协议是否激活
//	IsDriverAvailable(ctx context.Context, driverName string) bool // 检查驱动是否可用
//}

func GetDriverManager(ctx context.Context) common.IDeviceManager {
	return internal.GetDriverManager()
}
