package c_base

import (
	"time"
)

type MetaValue struct {
	Value      any        `json:"value,omitempty" dc:"数值"`
	HappenTime *time.Time `json:"happenTime,omitempty" dc:"发生时间"`
}

type MetaValueWrapper struct {
	DeviceId   string      `json:"deviceId,omitempty" dc:"设备ID"`
	DeviceType EDeviceType `json:"deviceType,omitempty" dc:"设备类型"`
	Meta       *Meta       `json:"meta,omitempty" dc:"点位信息"`
	Value      any         `json:"value,omitempty" dc:"数值"`
	HappenTime *time.Time  `json:"happenTime,omitempty" dc:"发生时间"`
}
