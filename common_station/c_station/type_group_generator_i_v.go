package c_station

import "ems-plan/c_device"

type IGroupGenerator interface {
	IGroup

	GetChildren() []c_device.IGenerator
}
