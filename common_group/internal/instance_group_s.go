package internal

import (
	"common_group/c_group"
	"ems-plan/c_device"
)

var (
	Instance c_group.IInstance
)

type sStationInstance struct {
	generator   c_group.IGroupGenerator
	entrance    c_group.IGroupEntrance
	load        c_group.IGroupLoad
	pv          c_group.IGroupPv
	energyStore c_group.IGroupEnergyStore

	inputMap  map[c_device.EInputType]c_device.IInput   // 输入信号
	outputMap map[c_device.EOutputType]c_device.IOutput // 输出信号
}

func init() {
	Instance = &sStationInstance{}
}

func (s *sStationInstance) RegisterInstance(info c_group.IInfo) {
	switch info.GetGroupType() {
	case c_group.EGroupPv:
		if s.pv != nil {
			panic("pv instance already registered")
		}
		s.pv = info.(c_group.IGroupPv)

	case c_group.EGroupEntrance:
		if s.entrance != nil {
			panic("entrance instance already registered")
		}
		s.entrance = info.(c_group.IGroupEntrance)
	case c_group.EGroupGenerator:
		if s.generator != nil {
			panic("generator instance already registered")
		}
		s.generator = info.(c_group.IGroupGenerator)
	case c_group.EGroupLoad:
		if s.load != nil {
			panic("load instance already registered")
		}
		s.load = info.(c_group.IGroupLoad)
	case c_group.EGroupEnergyStore:
		if s.energyStore != nil {
			panic("energyStore instance already registered")
		}

	default:
		panic("unknown group type: " + info.GetGroupType())
	}
}

func (s *sStationInstance) FindAll() []c_group.IInfo {
	list := make([]c_group.IInfo, 0, 5)

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

func (s *sStationInstance) GetStationLoad() c_group.IGroupLoad {
	return s.load
}

func (s *sStationInstance) GetEntrance() c_group.IGroupEntrance {
	return s.entrance
}

func (s *sStationInstance) GetEnergyStore() c_group.IGroupEnergyStore {
	return s.energyStore
}

func (s *sStationInstance) GetLoad() c_group.IGroupLoad {
	return s.load
}

func (s *sStationInstance) GetPv() c_group.IGroupPv {
	return s.pv
}

func (s *sStationInstance) GetGenerator() c_group.IGroupGenerator {
	return s.generator
}

func (s *sStationInstance) GetInput(inputType c_device.EInputType) c_device.IInput {
	return s.inputMap[inputType]
}

func (s *sStationInstance) GetOutput(outputType c_device.EOutputType) c_device.IOutput {
	return s.outputMap[outputType]
}
