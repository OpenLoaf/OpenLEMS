package internal

import (
	"ems-plan/c_device"
)

var (
	Instance c_station.IGroupInstance
)

type sStationInstance struct {
	generator   c_station.IGroupGenerator
	entrance    c_station.IGroupEntrance
	load        c_station.IGroupLoad
	pv          c_station.IGroupPv
	energyStore c_station.IGroupEnergyStore

	inputMap  map[c_base.EInputType]c_base.IInput   // 输入信号
	outputMap map[c_base.EOutputType]c_base.IOutput // 输出信号
}

func init() {
	Instance = &sStationInstance{}
}

func (s *sStationInstance) RegisterInstance(info c_station.IGroup) {
	switch info.GetGroupType() {
	case c_station.EGroupPv:
		if s.pv != nil {
			panic("pv instance already registered")
		}
		s.pv = info.(c_station.IGroupPv)

	case c_station.EGroupEntrance:
		if s.entrance != nil {
			panic("entrance instance already registered")
		}
		s.entrance = info.(c_station.IGroupEntrance)
	case c_station.EGroupGenerator:
		if s.generator != nil {
			panic("generator instance already registered")
		}
		s.generator = info.(c_station.IGroupGenerator)
	case c_station.EGroupLoad:
		if s.load != nil {
			panic("load instance already registered")
		}
		s.load = info.(c_station.IGroupLoad)
	case c_station.EGroupEnergyStore:
		if s.energyStore != nil {
			panic("energyStore instance already registered")
		}

	default:
		panic("unknown group type: " + info.GetGroupType())
	}
}

func (s *sStationInstance) FindAll() []c_station.IGroup {
	list := make([]c_station.IGroup, 0, 5)

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

func (s *sStationInstance) GetStationLoad() c_station.IGroupLoad {
	return s.load
}

func (s *sStationInstance) GetEntrance() c_station.IGroupEntrance {
	return s.entrance
}

func (s *sStationInstance) GetEnergyStore() c_station.IGroupEnergyStore {
	return s.energyStore
}

func (s *sStationInstance) GetLoad() c_station.IGroupLoad {
	return s.load
}

func (s *sStationInstance) GetPv() c_station.IGroupPv {
	return s.pv
}

func (s *sStationInstance) GetGenerator() c_station.IGroupGenerator {
	return s.generator
}

func (s *sStationInstance) GetInput(inputType c_base.EInputType) c_base.IInput {
	return s.inputMap[inputType]
}

func (s *sStationInstance) GetOutput(outputType c_base.EOutputType) c_base.IOutput {
	return s.outputMap[outputType]
}
