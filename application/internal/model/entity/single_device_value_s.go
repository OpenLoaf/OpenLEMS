package entity

import (
	"common/c_base"
	"time"
)

type SSingleDeviceValue struct {
	Meta       *c_base.Meta `json:"meta,omitempty"`
	Value      string       `json:"value,omitempty"`
	HappenTime *time.Time   `json:"happenTime,omitempty"`
}
