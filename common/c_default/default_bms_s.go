package c_default

import (
	"common/c_device"
)

type SDefaultBms struct {
}

func (S *SDefaultBms) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellMinTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellMaxTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellAvgTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellMinVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellMaxVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCellAvgVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetSoc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetSoh() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCapacity() (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetCycleCount() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (S *SDefaultBms) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
