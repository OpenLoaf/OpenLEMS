package c_group

import "ems-plan/c_device"

type EType = c_device.EGroupType

const (
	EGroupNan         = c_device.EGroupNan
	EGroupPv          = c_device.EGroupPv
	EGroupLoad        = c_device.EGroupLoad
	EGroupGenerator   = c_device.EGroupGenerator
	EGroupEntrance    = c_device.EGroupEntrance
	GroupCharge       = c_device.EGroupCharge
	EGroupEnergyStore = c_device.EGroupEnergyStore
)
