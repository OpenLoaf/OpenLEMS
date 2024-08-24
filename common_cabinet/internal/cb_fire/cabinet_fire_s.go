package cb_fire

import (
	"context"
	"ems-plan/c_device"
)

// CabinetFire 实现 c_device.IFireBasic 接口
type CabinetFire struct {
	ctx  context.Context
	fire c_device.IFire // 消防设备
}

func (c *CabinetFire) GetFireEnvTemperature() (float64, error) {
	return c.fire.GetFireEnvTemperature()
}

func (c *CabinetFire) GetCarbonMonoxideConcentration() (float64, error) {
	return c.fire.GetCarbonMonoxideConcentration()
}

func (c *CabinetFire) HasSmoke() (bool, error) {
	return c.fire.HasSmoke()
}
