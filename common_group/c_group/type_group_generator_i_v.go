package c_group

import "ems-plan/c_device"

type IGroupGenerator interface {
	IInfo

	GetChildren() []c_device.IGenerator
}
