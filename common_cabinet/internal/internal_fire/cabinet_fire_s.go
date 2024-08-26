package internal_fire

import (
	"context"
	"ems-plan/c_device"
)

// sCabinetFire 实现 c_base.IFireBasic 接口
type sCabinetFire struct {
	ctx  context.Context
	fire c_base.IFire // 消防设备
}

func (c *sCabinetFire) GetFireEnvTemperature() (float64, error) {
	return c.fire.GetFireEnvTemperature()
}

func (c *sCabinetFire) GetCarbonMonoxideConcentration() (float64, error) {
	return c.fire.GetCarbonMonoxideConcentration()
}

func (c *sCabinetFire) HasSmoke() (bool, error) {
	return c.fire.HasSmoke()
}
