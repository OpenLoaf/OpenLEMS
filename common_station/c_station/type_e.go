package c_station

import (
	"ems-plan/c_base"
)

type EType = c_base.EGroupType

const (
	EGroupNan         = c_base.EGroupNan
	EGroupPv          = c_base.EGroupPv
	EGroupLoad        = c_base.EGroupLoad
	EGroupGenerator   = c_base.EGroupGenerator
	EGroupEntrance    = c_base.EGroupEntrance
	EGroupCharge      = c_base.EGroupCharge
	EGroupEnergyStore = c_base.EGroupEnergyStore
)
