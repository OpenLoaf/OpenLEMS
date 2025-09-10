package c_device

import (
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
)

type SBasePolicyImpl[T c_base.IDriver] struct {
	Ctx          context.Context
	mode         c_enum.EPolicyMode
	deviceId     string
	manualAction func()
}

func (s *SBasePolicyImpl[T]) SetPolicyMode(mode c_enum.EPolicyMode) {
	// 执行手动触发方法
	if s.mode != c_enum.EPolicyModeManual && mode == c_enum.EPolicyModeManual && s.manualAction != nil {
		s.manualAction()
	}
	s.mode = mode
}

func (s *SBasePolicyImpl[T]) ExecuteDriverFunc(f func(driver T)) {
	device := common.GetDeviceManager().GetDeviceById(s.deviceId)
	if device == nil {
		return
	}
	// 转换类型
	if driver, ok := device.(T); ok {
		f(driver)
	}

}

func (s *SBasePolicyImpl[T]) GetPolicyMode() c_enum.EPolicyMode {
	return s.mode
}

func (s *SBasePolicyImpl[T]) RegisterMonitor(monitor *c_base.SPolicyMonitor) {

}

func (s *SBasePolicyImpl[T]) RegisterActiveManualAction(f func()) {
	s.manualAction = f
}
