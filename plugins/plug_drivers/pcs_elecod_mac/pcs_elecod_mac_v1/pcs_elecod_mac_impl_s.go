package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
	"common/c_log"
	"common/c_timer"
	"context"
	"pcs_elecod/pcs_elecod_mac_v1/elecod_mac_defined"
	"time"
)

type sPcsElecodBasic struct {
	p_canbus.ICanbusProtocol
	ctx          context.Context
	deviceConfig *c_base.SDeviceConfig
	*c_base.SDriverDescription
}

func (s *sPcsElecodBasic) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.ICanbusProtocol = protocol.(p_canbus.ICanbusProtocol)

	for _, task := range elecod_mac_defined.AnalogAllTasks {
		s.RegisterRead(task)
	}
	for _, task := range elecod_mac_defined.ConfigAllTasks {
		s.RegisterRead(task)
	}

	// 使用自研定时器，监听 ctx
	c_timer.SetInterval(s.ctx, time.Second, func(ctx context.Context) {
		c_log.Debugf(s.ctx, "测试发送数据")
		// _ = s.SendMessage(sandBy, nil)
	})

	c_log.Info(s.ctx, "测试结束！！！！")
}

func (s *sPcsElecodBasic) Shutdown() {
	c_log.Info(s.ctx, "Shutdown")
}

func (s *sPcsElecodBasic) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsElecodBasic) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
