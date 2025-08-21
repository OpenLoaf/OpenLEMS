package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
	"common/c_log"
	"common/c_timer"
	"context"
	"fmt"
	"pcs_elecod/pcs_elecod_mac_v1/elecod_mac_defined"
	"time"
)

type sPcsElecodMac struct {
	p_canbus.ICanbusProtocol
	ctx          context.Context
	deviceConfig *c_base.SDeviceConfig
	pcsConfig    *sPcsElecodMacConfig
	*c_base.SDriverDescription
}

func (s *sPcsElecodMac) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.ICanbusProtocol = protocol.(p_canbus.ICanbusProtocol)

	s.pcsConfig = &sPcsElecodMacConfig{}
	err := deviceConfig.ScanParams(s.pcsConfig)
	if err != nil {
		panic(fmt.Errorf("PcsElecodMac配置解析失败：内容:%v 原因: %s", deviceConfig.Params, err.Error()))
	}

	if s.pcsConfig.MacAddress == nil || s.pcsConfig.SelfAddress == nil {
		panic(fmt.Errorf("PcsElecodMac配置解析失败：缺少配置项！当前配置：%v", deviceConfig.Params))
	}

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

func (s *sPcsElecodMac) Shutdown() {
	c_log.Info(s.ctx, "Shutdown")
}

func (s *sPcsElecodMac) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsElecodMac) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodMac) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
