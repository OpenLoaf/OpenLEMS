package c_station

import "ems-plan/c_device"

type IStationLoad interface {
	IStation
	c_device.ILoadBasic

	GetChildren() []c_device.ILoad
}
