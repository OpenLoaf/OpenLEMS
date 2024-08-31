package c_base

import (
	"github.com/gogf/gf/v2/container/gvar"
	"time"
)

type MetaValue struct {
	Value      *gvar.Var  `dc:"数值"`
	HappenTime *time.Time `dc:"发生时间"`
}

type MetaValueWrapper struct {
	DeviceId   string
	DeviceType EDeviceType
	Meta       *Meta
	Value      *gvar.Var  `dc:"数值"`
	HappenTime *time.Time `dc:"发生时间"`
}
