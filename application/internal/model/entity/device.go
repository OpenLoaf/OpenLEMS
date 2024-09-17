package entity

import "common/c_base"

type SDevice struct {
	DeviceId       string             `json:"deviceId" dc:"设备Id"`
	DeviceType     string             `json:"deviceType" dc:"设备类型"`
	DeviceName     string             `json:"deviceName" dc:"设备名称"`
	LastUpdateTime string             `json:"lastUpdateTime" dc:"最后更新时间"`
	IsVirtual      bool               `json:"isVirtual"  dc:"是否虚拟设备"`
	AlarmLevel     c_base.EAlarmLevel `json:"alarmLevel" dc:"告警级别"`
}
