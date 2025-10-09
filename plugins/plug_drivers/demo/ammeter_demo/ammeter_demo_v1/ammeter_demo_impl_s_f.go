package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
)

type sAmmeterDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sAmmeterDemo) Init() error {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(ReadTask)
		return nil
	})
	return nil
}

func (s *sAmmeterDemo) Shutdown() {
}

func (s *sAmmeterDemo) GetUa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAVoltage)
}

func (s *sAmmeterDemo) GetUb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBVoltage)
}

func (s *sAmmeterDemo) GetUc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCVoltage)
}

func (s *sAmmeterDemo) GetIa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseACurrent)
}

func (s *sAmmeterDemo) GetIb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBCurrent)
}

func (s *sAmmeterDemo) GetIc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCCurrent)
}

func (s *sAmmeterDemo) GetPa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAActivePower)
}

func (s *sAmmeterDemo) GetPb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBActivePower)
}

func (s *sAmmeterDemo) GetPc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCActivePower)
}

func (s *sAmmeterDemo) GetPTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalActivePower)
}

func (s *sAmmeterDemo) GetQa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAReactivePower)
}

func (s *sAmmeterDemo) GetQb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBReactivePower)
}

func (s *sAmmeterDemo) GetQc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCReactivePower)
}

func (s *sAmmeterDemo) GetQTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalReactivePower)
}

func (s *sAmmeterDemo) GetSa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAApparentPower)
}

func (s *sAmmeterDemo) GetSb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBApparentPower)
}

func (s *sAmmeterDemo) GetSc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCApparentPower)
}

func (s *sAmmeterDemo) GetSTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalApparentPower)
}

func (s *sAmmeterDemo) GetPfa() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfb() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfc() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfTotal() (*float32, error) {
	return s.GetFromPointFloat32(PowerFactor)
}

func (s *sAmmeterDemo) GetPtCt() (*float32, *float32, error) {
	return nil, nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetFrequency() (*float32, error) {
	return s.GetFromPointFloat32(Frequency)
}

func (s *sAmmeterDemo) GetTodayIncomingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetHistoryIncomingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ForwardActiveEnergy)
}

func (s *sAmmeterDemo) GetTodayOutgoingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetHistoryOutgoingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ReverseActiveEnergy)
}

// 实现新的IDevice接口方法
func (s *sAmmeterDemo) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		telemetryPTotalPoint,
		telemetryFrequencyPoint,
		telemetryPfTotalPoint,
		telemetryHistoryIncomingQuantityPoint,
		telemetryHistoryOutgoingQuantityPoint,
	}
}

func (s *sAmmeterDemo) GetProtocolPoints() []c_base.IPoint {
	return []c_base.IPoint{
		Status,
		PhaseAVoltage,
		PhaseBVoltage,
		PhaseCVoltage,
		PhaseACurrent,
		PhaseBCurrent,
		PhaseCCurrent,
		PhaseAActivePower,
		PhaseBActivePower,
		PhaseCActivePower,
		PhaseAReactivePower,
		PhaseBReactivePower,
		PhaseCReactivePower,
		PhaseAApparentPower,
		PhaseBApparentPower,
		PhaseCApparentPower,
		TotalActivePower,
		TotalReactivePower,
		TotalApparentPower,
		ForwardActiveEnergy,
		ReverseActiveEnergy,
		Frequency,
		PowerFactor,
		RatedLineVoltage,
		RatedFrequency,
	}
}

func (s *sAmmeterDemo) GetConfigPoints() []*c_base.SConfigPoint {
	// 电表驱动没有配置点位，返回空列表
	return []*c_base.SConfigPoint{}
}
