package internal_bms

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

// sCabinetBms 实现 c_base.IBmsBasic 接口
type sCabinetBms struct {
	*c_base.SAlarmHandler
	alarmChannel chan *c_base.SAlarmDetail
	ctx          context.Context
	cabinetId    uint8 // 属于哪个柜子
	bms          c_device.IBms
}

func NewBms(ctx context.Context, cabinetId uint8, bms c_device.IBms) (c_device.IBmsBasic, chan<- *c_base.SAlarmDetail) {
	instance := &sCabinetBms{
		ctx:           context.WithValue(ctx, "DeviceName", "CabinetBms_"+string(cabinetId)),
		cabinetId:     cabinetId,
		bms:           bms,
		SAlarmHandler: &c_base.SAlarmHandler{},
		alarmChannel:  make(chan *c_base.SAlarmDetail),
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(instance.alarmChannel)
				return
			case alarm := <-instance.alarmChannel:
				g.Log().Infof(ctx, "BMS收到：Alarm: %v", alarm)
			}
		}
	}()
	fmt.Printf("NewBms: %v\n", instance)

	return instance, instance.alarmChannel
}

func (c *sCabinetBms) GetAlarmNotifyChannel() chan<- *c_base.SAlarmDetail {
	return c.alarmChannel
}

func (c *sCabinetBms) GetRatedPower() (float64, error) {
	return c.bms.GetRatedPower()
}

func (c *sCabinetBms) GetMaxInputPower() (float64, error) {
	return c.bms.GetMaxInputPower()
}

func (c *sCabinetBms) GetMaxOutputPower() (float64, error) {
	return c.bms.GetMaxOutputPower()
}

func (c *sCabinetBms) GetDcPower() (float64, error) {
	return c.bms.GetDcPower()
}

func (c *sCabinetBms) GetDcVoltage() (float64, error) {
	return c.bms.GetDcVoltage()
}

func (c *sCabinetBms) GetDcCurrent() (float64, error) {
	return c.bms.GetDcCurrent()
}

func (c *sCabinetBms) GetTodayIncomingQuantity() (float64, error) {
	return c.bms.GetTodayIncomingQuantity()
}

func (c *sCabinetBms) GetHistoryIncomingQuantity() (float64, error) {
	return c.bms.GetHistoryIncomingQuantity()
}

func (c *sCabinetBms) GetTodayOutgoingQuantity() (float64, error) {
	return c.bms.GetTodayOutgoingQuantity()
}

func (c *sCabinetBms) GetHistoryOutgoingQuantity() (float64, error) {
	return c.bms.GetHistoryOutgoingQuantity()
}

func (c *sCabinetBms) SetReset() error {
	return c.bms.SetReset()
}

func (c *sCabinetBms) SetBmsStatus(status c_device.EBmsStatus) error {
	return c.bms.SetBmsStatus(status)
}

func (c *sCabinetBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	return c.bms.GetBmsStatus()
}

func (c *sCabinetBms) GetSoc() (float32, error) {
	return c.bms.GetSoc()
}

func (c *sCabinetBms) GetSoh() (float32, error) {
	return c.bms.GetSoh()
}

func (c *sCabinetBms) GetCellTemp() (float32, float32, float32, error) {
	return c.bms.GetCellTemp()
}

func (c *sCabinetBms) GetCellVoltage() (float32, float32, float32, error) {
	return c.bms.GetCellVoltage()
}

func (c *sCabinetBms) GetCapacity() (uint16, error) {
	return c.bms.GetCapacity()
}

func (c *sCabinetBms) GetCycleCount() (uint, error) {
	return c.bms.GetCycleCount()
}
