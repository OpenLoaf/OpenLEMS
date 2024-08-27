package internal_humiture

import (
	"context"
	"ems-plan/c_device"
)

// sCabinetHumiture 实现 c_base.ICabinetHumiture 接口
type sCabinetHumiture struct {
	cabinetId uint8
	ctx       context.Context
	humiture  c_device.IHumiture
}

func (c *sCabinetHumiture) GetTemperature() (float64, error) {
	return c.humiture.GetTemperature()
}

func (c *sCabinetHumiture) GetHumidity() (float64, error) {
	return c.humiture.GetHumidity()
}
