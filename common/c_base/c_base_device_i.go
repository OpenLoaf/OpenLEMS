package c_base

import "time"

type IDevice interface {
	IAlarm // 告警
	//IPolicy // 策略

	GetConfig() *SDeviceConfig
	GetProtocolStatus() EProtocolStatus // 获取协议连接状态

	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time         // 获取最后更新时间

	IsVirtualDevice() bool // 是否是虚拟设备
}
