package c_group

import "ems-plan/c_device"

type IGroupPv interface {
	IInfo
	c_device.IPvBase

	GetChildren() []c_device.IPv
}
