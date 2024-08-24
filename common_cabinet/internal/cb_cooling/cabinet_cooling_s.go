package cb_cooling

import "ems-plan/c_device"

// CabinetCooling 实现 c_device.ICoolingBasic 接口
type CabinetCooling struct {
	cabinetId uint8
	cooling   c_device.ICoolingBasic
}

func (c *CabinetCooling) SetTemperature(temperature float32) error {
	return c.cooling.SetTemperature(temperature)
}

func (c *CabinetCooling) GetTemperature() (float32, error) {
	return c.cooling.GetTemperature()
}

func (c *CabinetCooling) GetTargetTemperature() (float32, error) {
	return c.cooling.GetTargetTemperature()
}
