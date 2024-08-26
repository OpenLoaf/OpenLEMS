package c_station

import "ems-plan/c_device"

type IGroupLoad interface {
	IGroup
	c_device.ILoadBasic

	GetChildren() []c_device.ILoad
}
