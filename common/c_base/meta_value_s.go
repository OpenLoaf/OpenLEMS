package c_base

import "time"

type MetaValue struct {
	Value      any        `dc:"数值"`
	HappenTime *time.Time `dc:"发生时间"`
}
