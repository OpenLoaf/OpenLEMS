package ess_demo_v1

import (
	"common/c_log"
	"math/rand"
)

func (s *sEssDemo) CustomPowerZeroService() error {
	return s.SetPower(0)
}
func (s *sEssDemo) CustomPowerRandomService() error {
	mx, err := s.GetMaxInputPower()
	if err != nil {
		mx = 0
	}
	mi, err := s.GetMaxInputPower()
	if err != nil {
		mi = 0
	}

	random := rand.Intn(int(mx-mi)+1) + int(mi)

	c_log.BizDebugf(s.DeviceCtx, "设置随机功率为：%v", random)
	return s.SetPower(int32(random))
}
