package internal

import "common/c_base"

var gpioPoint = &c_base.SPoint{
	Key:     "pin",
	Name:    "状态",
	Group:   c_base.GroupTotal,
	Precise: 0,
}

var high = true
var low = false
