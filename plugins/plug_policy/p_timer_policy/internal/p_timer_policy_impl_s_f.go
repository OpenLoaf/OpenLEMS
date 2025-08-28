package internal

import (
	"common/c_base"
	"common/c_type"
)

type sTimerPolicyImpl struct {
	c_base.IPolicy[c_type.IEnergyStore]
}

func (s *sTimerPolicyImpl) Init() {

	//s.RegisterMonitor()

}

//func (s *sTimerPolicyImpl) GetDevice
