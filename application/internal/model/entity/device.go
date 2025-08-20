package entity

import (
	"common/c_base"
)

type SDevice struct {
	DeviceId       string             `json:"deviceId" dc:"设备Id"`
	DeviceType     string             `json:"deviceType" dc:"设备类型"`
	DeviceName     string             `json:"deviceName" dc:"设备名称"`
	LastUpdateTime string             `json:"lastUpdateTime" dc:"最后更新时间"`
	IsVirtual      bool               `json:"isVirtual"  dc:"是否虚拟设备"`
	AlarmLevel     c_base.EAlarmLevel `json:"alarmLevel" dc:"告警级别"`
}

type SDeviceTree struct {
	//DeviceId       string         `json:"deviceId" dc:"设备Id"`
	//DevicePid      string         `json:"devicePid" dc:"父设备Id"`
	//ProtocolId     string         `json:"protocolId" dc:"协议Id"`
	//DeviceName     string         `json:"deviceName" dc:"设备名称"`
	//DeviceDriver   string         `json:"deviceDriver" dc:"设备驱动"`
	//LogLevel       string         `json:"logLevel" dc:"日志级别"`
	//Enable         bool           `json:"enable" dc:"是否启用"`
	//IsRunning      bool           `json:"isRunning" dc:"是否运行中"`
	//LastUpdateTime string         `json:"lastUpdateTime" dc:"最后通讯时间"`
	//Sort           int            `json:"sort" dc:"排序"`
	*c_base.SDeviceDetail
	Children []*SDeviceTree `json:"deviceChildren" dc:"子设备列表"`
}
