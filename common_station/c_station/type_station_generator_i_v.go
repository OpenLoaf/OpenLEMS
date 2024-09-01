package c_station

import "ems-plan/c_device"

type IStationGenerator interface {
	IStation

	GetChildren() []c_device.IGenerator
}
