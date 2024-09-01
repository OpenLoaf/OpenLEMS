package internal

import (
	"common_station/c_station"
)

var (
	Instance c_station.IGroupInstance
)

type sStationInstance struct {
	generator   c_station.IStationGenerator
	entrance    c_station.IStationEntrance
	load        c_station.IStationLoad
	pv          c_station.IStationPv
	energyStore c_station.IStationEnergyStore
}

func init() {
	Instance = &sStationInstance{}
}

func (s *sStationInstance) RegisterInstance(info c_station.IStation) {
	switch info.GetGroupType() {
	case c_station.EGroupPv:
		if s.pv != nil {
			panic("pv instance already registered")
		}
		s.pv = info.(c_station.IStationPv)

	case c_station.EGroupEntrance:
		if s.entrance != nil {
			panic("entrance instance already registered")
		}
		s.entrance = info.(c_station.IStationEntrance)
	case c_station.EGroupGenerator:
		if s.generator != nil {
			panic("generator instance already registered")
		}
		s.generator = info.(c_station.IStationGenerator)
	case c_station.EGroupLoad:
		if s.load != nil {
			panic("load instance already registered")
		}
		s.load = info.(c_station.IStationLoad)
	case c_station.EGroupEnergyStore:
		if s.energyStore != nil {
			panic("energyStore instance already registered")
		}

	default:
		panic("unknown group type: " + info.GetGroupType())
	}
}

func (s *sStationInstance) FindAll() []c_station.IStation {
	list := make([]c_station.IStation, 0, 5)

	if s.entrance != nil {
		list = append(list, s.entrance)
	}
	if s.energyStore != nil {
		list = append(list, s.energyStore)
	}
	if s.pv != nil {
		list = append(list, s.pv)
	}
	if s.generator != nil {
		list = append(list, s.generator)
	}
	if s.load != nil {
		list = append(list, s.load)
	}

	return list
}

func (s *sStationInstance) GetStationLoad() c_station.IStationLoad {
	return s.load
}

func (s *sStationInstance) GetEntrance() c_station.IStationEntrance {
	return s.entrance
}

func (s *sStationInstance) GetEnergyStore() c_station.IStationEnergyStore {
	return s.energyStore
}

func (s *sStationInstance) GetLoad() c_station.IStationLoad {
	return s.load
}

func (s *sStationInstance) GetPv() c_station.IStationPv {
	return s.pv
}

func (s *sStationInstance) GetGenerator() c_station.IStationGenerator {
	return s.generator
}
