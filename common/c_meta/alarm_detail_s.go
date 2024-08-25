package c_meta

import "time"

type SAlarmDetail struct {
	DeviceId   string     `json:"deviceId" dc:"设备ID"`
	Level      AlarmLevel `json:"level" dc:"告警级别"`
	Meta       *Meta      `json:"meta" dc:"告警元数据"`
	HappenTime time.Time  `json:"happenTime" dc:"发生时间"`
	Value      any        `json:"value" dc:"数值"`
}
