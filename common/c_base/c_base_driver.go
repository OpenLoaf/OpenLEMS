package c_base

type IDevice interface {
	IProtocol
	InitDevice(deviceConfig *SDeviceConfig, protocol IProtocol, childDevice []IDevice)
	Shutdown() // 销毁

	GetDriverType() EDeviceType // 获取实现驱动的设备类型

	GetMetaValueList() []*MetaValueWrapper     // 获取元数据列表
	GetDriverDescription() *SDriverDescription // 获取驱动的详情
}

type IDeviceWrapper interface {
	GetDeviceConfig() *SDeviceConfig
	GetDriverInfo() *SDriverInfo
	GetProtocolConfig() *SProtocolConfig

	GetDeviceInstance() IDevice

	GetDeviceState() EServerState

	Shutdown()
}
