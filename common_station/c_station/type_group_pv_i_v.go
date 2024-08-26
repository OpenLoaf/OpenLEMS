package c_station

import "ems-plan/c_device"

type IGroupPv interface {
	IGroup
	c_device.IPvBase

	GetChildren() []c_device.IPv
}
