package internal_cooling

import (
	"context"
	"ems-plan/c_device"
)

// sCabinetCooling 实现 c_base.ICoolingBasic 接口
type sCabinetCooling struct {
	ctx       context.Context
	cabinetId uint8
	cooling   c_device.ICoolingBasic
}

func NewCooling(ctx context.Context, cabinetId uint8, cooling c_device.ICoolingBasic) c_device.ICoolingBasic {
	instance := &sCabinetCooling{
		ctx:       context.WithValue(ctx, "DeviceName", "CabinetBms_"+string(cabinetId)),
		cabinetId: cabinetId,
		cooling:   cooling,
	}
	return instance

}

func (s *sCabinetCooling) SetTemperature(temperature float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetCooling) GetTemperature() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetCooling) GetTargetTemperature() (float32, error) {
	//TODO implement me
	panic("implement me")
}
