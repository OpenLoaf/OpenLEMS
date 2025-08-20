package entity

import "common/c_base"

type SSingleDeviceValue struct {
	Meta       *c_base.Meta `json:"meta,omitempty"`
	Value      string       `json:"value,omitempty"`
	HappenTime string       `json:"happenTime,omitempty"`
}
