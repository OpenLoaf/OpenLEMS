package ess_demo_v1

import (
	"common/c_log"
	"errors"
	"math/rand"
)

func (s *sEssDemo) CustomPowerZeroService() error {
	return s.SetPower(0)
}
func (s *sEssDemo) CustomPowerRandomService() error {
	mx, err := s.GetMaxInputPower()
	if err != nil {
		return err
	}
	if mx == nil {
		return errors.New("无法获取最大输入功率")
	}

	mi, err := s.GetMaxOutputPower()
	if err != nil {
		return err
	}
	if mi == nil {
		return errors.New("无法获取最大输出功率")
	}

	// 生成在最大输入功率到最大输出功率之间的随机数
	random := rand.Intn(int(*mx-*mi)+1) + int(*mi)

	c_log.BizDebugf(s.DeviceCtx, "设置随机功率为：%v", random)
	return s.SetPower(int32(random))
}
