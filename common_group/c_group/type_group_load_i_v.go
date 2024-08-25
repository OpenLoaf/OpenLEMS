package c_group

import "ems-plan/c_device"

type IGroupLoad interface {
	IInfo
	c_device.ILoadBasic

	GetChildren() []c_device.ILoad
}
