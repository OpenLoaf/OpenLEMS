package c_group

type IInstance interface {
	RegisterInstance(info IInfo)
	FindAll() []IInfo

	GetStationLoad() IGroupLoad
	GetEntrance() IGroupEntrance
	GetEnergyStore() IGroupEnergyStore
	GetLoad() IGroupLoad
	GetPv() IGroupPv
	GetGenerator() IGroupGenerator
}
