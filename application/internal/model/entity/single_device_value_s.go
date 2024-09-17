package entity

import "common/c_base"

type SSingleDeviceValue struct {
	DeviceId   string             `json:"deviceId,omitempty"`
	DeviceType c_base.EDeviceType `json:"deviceType,omitempty"`
	Meta       *c_base.Meta       `json:"meta,omitempty"`
	Value      string             `json:"value,omitempty"`
	HappenTime string             `json:"happenTime,omitempty"`
}
