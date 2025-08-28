package internal

import (
	"common/c_base"
)

type sTimerPolicyImpl struct{}

func (s *sTimerPolicyImpl) SetPolicyMode(mode c_base.EPolicyMode) {
	//TODO implement me
	panic("implement me")
}

func (s *sTimerPolicyImpl) GetPolicyMode() c_base.EPolicyMode {
	//TODO implement me
	panic("implement me")
}

func (s *sTimerPolicyImpl) RegisterMonitor(monitor *c_base.SPolicyMonitor) {
	//TODO implement me
	panic("implement me")
}

func (s *sTimerPolicyImpl) RegisterActiveManualAction(f func()) {
	//TODO implement me
	panic("implement me")
}
