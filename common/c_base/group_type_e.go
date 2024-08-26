package c_base

type EGroupType string

const (
	EGroupNan         EGroupType = "" // 空
	EGroupPv          EGroupType = "pv-group"
	EGroupLoad        EGroupType = "load-group"
	EGroupGenerator   EGroupType = "generator-group"
	EGroupEntrance    EGroupType = "entrance-group"
	EGroupCharge      EGroupType = "charge-group"
	EGroupEnergyStore EGroupType = "energy-store-group"
)
