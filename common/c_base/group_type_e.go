package c_base

type EStationType string

const (
	EStationNan         EStationType = "" // 空
	EStationPv          EStationType = "pv"
	EStationLoad        EStationType = "load"
	EStationGenerator   EStationType = "generator"
	EStationEntrance    EStationType = "entrance"
	EStationCharge      EStationType = "charge"
	EStationEnergyStore EStationType = "energy-store"
)
