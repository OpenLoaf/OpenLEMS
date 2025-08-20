package c_base

import "time"

type IDeviceRuntimeInfo interface {
	IsPhysical() bool // 是否是物理设备

	GetDeviceConfig() *SDeviceConfig       // 获取设备配置
	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time         // 获取最后更新时间
}

type IDevice interface {
	IAlarm             // 告警
	IDeviceRuntimeInfo // 设备运行信息

	InitDevice(deviceConfig *SDeviceConfig, protocol IProtocol, childDevice []IDevice)
	Shutdown() // 销毁

	GetDriverType() EDeviceType                // 获取实现驱动的设备类型
	GetDriverDescription() *SDriverDescription // 获取驱动的详情
}

type IDeviceWrapper interface {
	GetDeviceDetail() *SDeviceDetail // 获取设备详情
	GetDeviceInstance() IDevice      // 获取设备实例
	GetDeviceState() EServerState    // 获取设备运行状态（系统中的）
	Shutdown()                       // 停机
}
