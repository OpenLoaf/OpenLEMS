package c_default

import (
	"common/c_base"
)

type SDefaultPcs struct {
}

func (S *SDefaultPcs) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultPcs) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
