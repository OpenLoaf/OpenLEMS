package cb_humiture

import (
	"context"
	"ems-plan/c_device"
)

// CabinetHumiture 实现 c_device.ICabinetHumiture 接口
type CabinetHumiture struct {
	cabinetId uint8
	ctx       context.Context
	humiture  c_device.IHumiture
}

func (c *CabinetHumiture) GetTemperature() (float64, error) {
	return c.humiture.GetTemperature()
}

func (c *CabinetHumiture) GetHumidity() (float64, error) {
	return c.humiture.GetHumidity()
}
