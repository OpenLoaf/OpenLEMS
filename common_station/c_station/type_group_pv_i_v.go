package c_station

import "ems-plan/c_device"

type IStationPv interface {
	IStation
	c_device.IPvBase

	GetChildren() []c_device.IPv
}
