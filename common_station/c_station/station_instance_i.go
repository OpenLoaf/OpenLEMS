package c_station

type IGroupInstance interface {
	RegisterInstance(info IStation)
	FindAll() []IStation

	GetStationLoad() IStationLoad
	GetEntrance() IStationEntrance
	GetEnergyStore() IStationEnergyStore
	GetLoad() IStationLoad
	GetPv() IStationPv
	GetGenerator() IStationGenerator
}
