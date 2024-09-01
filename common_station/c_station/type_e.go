package c_station

import (
	"ems-plan/c_base"
)

type EType = c_base.EStationType

const (
	EGroupNan         = c_base.EStationNan
	EGroupPv          = c_base.EStationPv
	EGroupLoad        = c_base.EStationLoad
	EGroupGenerator   = c_base.EStationGenerator
	EGroupEntrance    = c_base.EStationEntrance
	EGroupCharge      = c_base.EStationCharge
	EGroupEnergyStore = c_base.EStationEnergyStore
)
