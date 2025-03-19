package pcs_enjoy_basic_v1

import (
	"common/c_base"
	"common/c_error"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"modbus/p_modbus"
)

type sPcsEnjoyBasic struct {
	p_modbus.IModbusProtocol
	ctx                 context.Context
	log                 *glog.Logger
	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率
	deviceConfig        *c_base.SDriverConfig
	*c_base.SDescription
}

func (s *sPcsEnjoyBasic) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.IModbusProtocol = protocol.(p_modbus.IModbusProtocol)
	s.deviceConfig = deviceConfig

	// 注册
	s.RegisterRead(s.ctx,
		GroupAcInfo,
		GroupPowerInfo,
		GroupBasicInfo,
		GroupSetting,
	)

	g.Log().Noticef(s.ctx, "sPcsEnjoyBasic Init")
}

func (s *sPcsEnjoyBasic) Destroy() {
	_ = s.SetPower(0)
	_ = s.SetStatus(c_base.EPcsStatusOff)
	g.Log().Noticef(s.ctx, "[%s]%s销毁成功,设置PCS状态为Off!", s.deviceConfig.Id, s.deviceConfig.Name)
}

func (s *sPcsEnjoyBasic) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsEnjoyBasic) SetReset() error {
	return c_error.NonSupportError
}

func (s *sPcsEnjoyBasic) SetStatus(status c_base.EEnergyStoreStatus) error {
	// TODO 先使用简捷的EMS去开机
	return c_error.NonSupportError
}

func (s *sPcsEnjoyBasic) SetGridMode(mode c_base.EGridMode) error {
	return c_error.NonSupportError
}

func (s *sPcsEnjoyBasic) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return c_base.EPcsStatusUnknown, c_error.NonSupportError
}

func (s *sPcsEnjoyBasic) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) SetPower(power int32) error {
	return s.WriteSingleRegister(Set_Ap, power)
}

func (s *sPcsEnjoyBasic) SetReactivePower(power int32) error {
	return s.WriteSingleRegister(Set_Qp, power)
}

func (s *sPcsEnjoyBasic) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetTargetPower() int32 {
	value, err := s.GetInt32Value(Set_Ap)
	if err != nil {
		return 0
	}
	return value
}

func (s *sPcsEnjoyBasic) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetPower() (float64, error) {
	return s.GetFloat64Value(Pt)
}

func (s *sPcsEnjoyBasic) GetApparentPower() (float64, error) {
	return s.GetFloat64Value(St)
}

func (s *sPcsEnjoyBasic) GetReactivePower() (float64, error) {
	return s.GetFloat64Value(Qt)
}

func (s *sPcsEnjoyBasic) GetRatedPower() int32 {
	// TODO 以后从配置中读取
	return 100
}

func (s *sPcsEnjoyBasic) GetMaxInputPower() (float32, error) {
	return float32(s.GetRatedPower()), nil
}

func (s *sPcsEnjoyBasic) GetMaxOutputPower() (float32, error) {
	return float32(s.GetRatedPower()), nil
}

func (s *sPcsEnjoyBasic) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsEnjoyBasic) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
