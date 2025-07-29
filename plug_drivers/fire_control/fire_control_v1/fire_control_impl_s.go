package fire_control_v1

import (
	"canbus/p_canbus"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sFireControlBasic struct {
	p_canbus.ICanbusProtocol
	ctx          context.Context
	log          *glog.Logger
	deviceConfig *c_base.SDriverConfig
	*c_base.SDescription
}

func (s *sFireControlBasic) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.ICanbusProtocol = protocol.(p_canbus.ICanbusProtocol)

	s.RegisterRead(&Detail)
}

func (s *sFireControlBasic) Destroy() {
	g.Log().Info(s.ctx, "Destroy")
}

func (s *sFireControlBasic) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceFire
}

func (s *sFireControlBasic) GetFireEnvTemperature() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sFireControlBasic) GetCarbonMonoxideConcentration() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sFireControlBasic) HasSmoke() (bool, error) {
	//TODO implement me
	panic("implement me")
}
