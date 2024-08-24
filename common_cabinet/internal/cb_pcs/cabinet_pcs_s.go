package cb_pcs

import (
	"context"
	"ems-plan/c_device"
)

// CabinetPcs 实现了 c_device.IPcsBasic 接口
type CabinetPcs struct {
	//info      *plugin.DriverInfo
	cabinetId uint8
	ctx       context.Context
	rootPcs   c_device.IPcs
	pcsList   []c_device.IPcs
}

func (p *CabinetPcs) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) SetStatus(status c_device.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) SetGridMode(mode c_device.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetStatus() (c_device.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *CabinetPcs) GetGridMode() (c_device.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}
