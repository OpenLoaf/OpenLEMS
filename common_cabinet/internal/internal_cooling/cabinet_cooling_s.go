package internal_cooling

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
)

// sCabinetCooling 实现 c_base.ICoolingBasic 接口
type sCabinetCooling struct {
	ctx       context.Context
	cabinetId uint8
	cooling   c_base.ICoolingBasic
}

func NewCooling(ctx context.Context, cabinetId uint8, cooling c_base.ICoolingBasic) c_base.ICoolingBasic {
	instance := &sCabinetCooling{
		ctx:       context.WithValue(ctx, "DeviceName", "CabinetBms_"+string(cabinetId)),
		cabinetId: cabinetId,
		cooling:   cooling,
	}
	return instance

}

func (c *sCabinetCooling) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

func (c *sCabinetCooling) SetTemperature(temperature float32) error {
	return c.cooling.SetTemperature(temperature)
}

func (c *sCabinetCooling) GetTemperature() (float32, error) {
	return c.cooling.GetTemperature()
}

func (c *sCabinetCooling) GetTargetTemperature() (float32, error) {
	return c.cooling.GetTargetTemperature()
}
