package c_device

import (
	"common/c_base"
)

type SPolicyManager struct {
	mode         c_base.EPolicyMode
	manualAction func()
}

func (s *SPolicyManager) SetPolicyMode(mode c_base.EPolicyMode) {
	// 执行手动触发方法
	if s.mode != c_base.EPolicyModeManual && mode == c_base.EPolicyModeManual && s.manualAction != nil {
		s.manualAction()
	}
	s.mode = mode
}

func (s *SPolicyManager) GetPolicyMode() c_base.EPolicyMode {
	return s.mode
}

func (s *SPolicyManager) RegisterMonitor(monitor *c_base.SPolicyMonitor) {

}

func (s *SPolicyManager) RegisterActiveManualAction(f func()) {
	s.manualAction = f
}

func (s *SPolicyManager) Init() {

	//c := &c_base.SPolicyMonitor{
	//	Name:        "",
	//	Duration:    nil,
	//	Modes:       []c_base.EPolicyMode{c_base.EPolicyModeAuto},
	//	TriggerFunc: nil,
	//	HandleFunc:  nil,
	//}

}
