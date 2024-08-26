package c_station

type IGroupInstance interface {
	RegisterInstance(info IGroup)
	FindAll() []IGroup

	GetStationLoad() IGroupLoad
	GetEntrance() IGroupEntrance
	GetEnergyStore() IGroupEnergyStore
	GetLoad() IGroupLoad
	GetPv() IGroupPv
	GetGenerator() IGroupGenerator
}
