package ammeter_acrel_10r_v1

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"plug_protocol_modbus/p_modbus"
)

type sAmmeterAcrel10r struct {
	ctx context.Context
	p_modbus.IModbusProtocol
	*c_base.SDescription
}

func (s *sAmmeterAcrel10r) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.IModbusProtocol = protocol.(p_modbus.IModbusProtocol)

	// 注册
	s.IModbusProtocol.RegisterRead(s.ctx, GDatetime, GRealtimeInfo, GTotal, GSwitch)
}

func (s *sAmmeterAcrel10r) Destroy() {
	g.Log().Noticef(s.ctx, "[%s]%s销毁成功", s.GetDeviceConfig().Id, s.GetDeviceConfig().Name)
}

func (s *sAmmeterAcrel10r) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceAmmeter
}

func (s *sAmmeterAcrel10r) GetUa() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetUb() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetUc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetIa() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetIb() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetIc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPa() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPb() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPtCt() (float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetFrequency() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPowerFactor() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}
