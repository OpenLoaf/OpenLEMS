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
	s.IModbusProtocol.RegisterRead(s.ctx, GRealtimeInfo, GTotal)
}

func (s *sAmmeterAcrel10r) Destroy() {
	g.Log().Noticef(s.ctx, "[%s]%s销毁成功", s.GetDeviceConfig().Id, s.GetDeviceConfig().Name)
}

func (s *sAmmeterAcrel10r) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceAmmeter
}

func (s *sAmmeterAcrel10r) GetUa() (float32, error) {
	return s.GetFloat32Value(Ua)
}

func (s *sAmmeterAcrel10r) GetUb() (float32, error) {
	return s.GetFloat32Value(Ub)
}

func (s *sAmmeterAcrel10r) GetUc() (float32, error) {
	return s.GetFloat32Value(Uc)
}

func (s *sAmmeterAcrel10r) GetIa() (float32, error) {
	return s.GetFloat32Value(Ia)
}

func (s *sAmmeterAcrel10r) GetIb() (float32, error) {
	return s.GetFloat32Value(Ib)
}

func (s *sAmmeterAcrel10r) GetIc() (float32, error) {
	return s.GetFloat32Value(Ic)
}

func (s *sAmmeterAcrel10r) GetPa() (float32, error) {
	return s.GetFloat32Value(Pa)
}

func (s *sAmmeterAcrel10r) GetPb() (float32, error) {
	return s.GetFloat32Value(Pb)
}

func (s *sAmmeterAcrel10r) GetPc() (float32, error) {
	return s.GetFloat32Value(Pc)
}

func (s *sAmmeterAcrel10r) GetPtCt() (float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetPTotal() (float32, error) {
	return s.GetFloat32Value(Pt)
}

func (s *sAmmeterAcrel10r) GetQa() (float32, error) {
	return s.GetFloat32Value(Qa)
}

func (s *sAmmeterAcrel10r) GetQb() (float32, error) {
	return s.GetFloat32Value(Qb)
}

func (s *sAmmeterAcrel10r) GetQc() (float32, error) {
	return s.GetFloat32Value(Qc)
}

func (s *sAmmeterAcrel10r) GetQTotal() (float32, error) {
	return s.GetFloat32Value(Qt)
}

func (s *sAmmeterAcrel10r) GetSa() (float32, error) {
	return s.GetFloat32Value(Sa)
}

func (s *sAmmeterAcrel10r) GetSb() (float32, error) {
	return s.GetFloat32Value(Sb)
}

func (s *sAmmeterAcrel10r) GetSc() (float32, error) {
	return s.GetFloat32Value(Sc)
}

func (s *sAmmeterAcrel10r) GetSTotal() (float32, error) {
	return s.GetFloat32Value(St)
}

func (s *sAmmeterAcrel10r) GetPfa() (float32, error) {
	return s.GetFloat32Value(Pfa)
}

func (s *sAmmeterAcrel10r) GetPfb() (float32, error) {
	return s.GetFloat32Value(Pfb)
}

func (s *sAmmeterAcrel10r) GetPfc() (float32, error) {
	return s.GetFloat32Value(Pfc)
}

func (s *sAmmeterAcrel10r) GetPfTotal() (float32, error) {
	return s.GetFloat32Value(Pft)
}

func (s *sAmmeterAcrel10r) GetFrequency() (float32, error) {
	return s.GetFloat32Value(F)
}

func (s *sAmmeterAcrel10r) GetPowerFactor() (float32, error) {
	return s.GetFloat32Value(Pft)
}

func (s *sAmmeterAcrel10r) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetHistoryIncomingQuantity() (float64, error) {
	return s.GetFloat64Value(Epi)
}

func (s *sAmmeterAcrel10r) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sAmmeterAcrel10r) GetHistoryOutgoingQuantity() (float64, error) {
	return s.GetFloat64Value(Eqe)
}
