package internal_pcs

import (
	"context"
	"ems-plan/c_device"
)

// sCabinetPcs 实现了 c_base.IPcsBasic 接口
type sCabinetPcs struct {
	//info      *plugin.DriverInfo
	cabinetId uint8
	ctx       context.Context
	rootPcs   c_base.IPcs
	pcsList   []c_base.IPcs
}

func NewPcs(ctx context.Context, cabinetId uint8, master c_base.IPcs, pcsList []c_base.IPcs) c_base.IPcsBasic {
	if len(pcsList) == 0 && master == nil {
		panic("master and pcsList cannot be empty at the same time")
	}

	/*info := newCabinetInfo(cabinetId, config.Pcs, _pcs,
	"sCabinetPcs", "{#canbintPcsName}",
	map[string]*plugin.PointInfo{
		"targetPower":           {I18nName: "targetPower", Unit: "kW"},
		"power":                 {I18nName: "power", Unit: "kW"},
		"apparentPower":         {I18nName: "apparentPower", Unit: "kVA"},
		"reactivePower":         {I18nName: "reactivePower", Unit: "kVar"},
		"todayIncomingQuantity": {I18nName: "todayIncomingQuantity", Unit: "kWh"},
		"todayOutgoingQuantity": {I18nName: "todayOutgoingQuantity", Unit: "kWh"},
	})*/

	instance := &sCabinetPcs{
		ctx:       context.WithValue(ctx, "DeviceName", "CabinetPcs_"+string(cabinetId)),
		cabinetId: cabinetId,
		//info:      info,
		rootPcs: master,
		pcsList: pcsList,
	}

	//device.RegisterInstance(instance)
	return instance
}

func (p *sCabinetPcs) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sCabinetPcs) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}
