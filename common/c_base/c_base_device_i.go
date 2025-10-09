package c_base

import (
	"common/c_enum"
	"time"
)

type IDevice interface {
	IAlarm // 告警

	GetConfig() *SDeviceConfig
	GetProtocolStatus() c_enum.EProtocolStatus // 获取协议连接状态

	GetPointValueList() []*SPointValue // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time     // 获取最后更新时间

	IsVirtualDevice() bool // 是否是虚拟设备

	// 统一点位管理方法
	GetTelemetryPoints() []IPoint // 获取所有点位列表（遥测点位+协议点位）
}
