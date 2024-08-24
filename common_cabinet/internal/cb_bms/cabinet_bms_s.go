package cb_bms

import (
	"context"
	"ems-plan/c_device"
)

// CabinetBms 实现 c_device.IBmsBasic 接口
type CabinetBms struct {
	ctx context.Context
	bms c_device.IBms
}

func (c *CabinetBms) GetRatedPower() (float64, error) {
	return c.bms.GetRatedPower()
}

func (c *CabinetBms) GetMaxInputPower() (float64, error) {
	return c.bms.GetMaxInputPower()
}

func (c *CabinetBms) GetMaxOutputPower() (float64, error) {
	return c.bms.GetMaxOutputPower()
}

func (c *CabinetBms) GetDcPower() (float64, error) {
	return c.bms.GetDcPower()
}

func (c *CabinetBms) GetDcVoltage() (float64, error) {
	return c.bms.GetDcVoltage()
}

func (c *CabinetBms) GetDcCurrent() (float64, error) {
	return c.bms.GetDcCurrent()
}

func (c *CabinetBms) GetTodayIncomingQuantity() (float64, error) {
	return c.bms.GetTodayIncomingQuantity()
}

func (c *CabinetBms) GetHistoryIncomingQuantity() (float64, error) {
	return c.bms.GetHistoryIncomingQuantity()
}

func (c *CabinetBms) GetTodayOutgoingQuantity() (float64, error) {
	return c.bms.GetTodayOutgoingQuantity()
}

func (c *CabinetBms) GetHistoryOutgoingQuantity() (float64, error) {
	return c.bms.GetHistoryOutgoingQuantity()
}

func (c *CabinetBms) SetReset() error {
	return c.bms.SetReset()
}

func (c *CabinetBms) SetBmsStatus(status c_device.EBmsStatus) error {
	return c.bms.SetBmsStatus(status)
}

func (c *CabinetBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	return c.bms.GetBmsStatus()
}

func (c *CabinetBms) GetSoc() (float32, error) {
	return c.bms.GetSoc()
}

func (c *CabinetBms) GetSoh() (float32, error) {
	return c.bms.GetSoh()
}

func (c *CabinetBms) GetCellTemp() (float32, float32, float32, error) {
	return c.bms.GetCellTemp()
}

func (c *CabinetBms) GetCellVoltage() (float32, float32, float32, error) {
	return c.bms.GetCellVoltage()
}

func (c *CabinetBms) GetCapacity() (uint16, error) {
	return c.bms.GetCapacity()
}

func (c *CabinetBms) GetCycleCount() (uint, error) {
	return c.bms.GetCycleCount()
}
