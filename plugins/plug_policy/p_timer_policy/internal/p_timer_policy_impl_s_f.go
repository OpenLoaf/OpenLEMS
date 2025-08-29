package internal

import (
	"common/c_base"
	"common/c_type"
)

type sTimerPolicyImpl struct {
	c_base.IPolicy[c_type.IEnergyStore]
}

func (s *sTimerPolicyImpl) Init() {

	s.RegisterMonitor(&c_base.SPolicyMonitor{
		Name:     "充电",
		Duration: nil,
		Modes:    nil,
		TriggerFunc: func() bool {
			return true
		},
		HandleFunc: func() {

		},
	})

}

//func (s *sTimerPolicyImpl) GetDevice
