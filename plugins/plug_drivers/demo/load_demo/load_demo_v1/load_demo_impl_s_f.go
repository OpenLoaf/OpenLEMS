package load_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
)

type sLoadDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sLoadDemo) Init() error {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(ReadTask)
		return nil
	})
	return nil
}

func (s *sLoadDemo) Shutdown() {
}

func (s *sLoadDemo) SetPower(power float64) error {
	return c_base.NotSupport
}

func (s *sLoadDemo) SetReactivePower(power float64) error {
	return c_base.NotSupport
}

func (s *sLoadDemo) SetPowerFactor(factor float32) error {
	return c_base.NotSupport
}

func (s *sLoadDemo) GetTargetPower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetTargetReactivePower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetTargetPowerFactor() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetPower() (*float64, error) {
	return s.GetFromPointFloat64(Power)
}

func (s *sLoadDemo) GetApparentPower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetReactivePower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetTodayIncomingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetHistoryIncomingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(Energy)
}

func (s *sLoadDemo) GetTodayOutgoingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetHistoryOutgoingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sLoadDemo) GetRatedPower() (*uint32, error) {
	return s.GetFromPointUint32(MaxLoad)
}

func (s *sLoadDemo) GetMaxInputPower() (*float64, error) {
	return s.GetFromPointFloat64(MaxLoad)
}

func (s *sLoadDemo) GetMaxOutputPower() (*float64, error) {
	return nil, nil
}

// 实现新的IDevice接口方法
func (s *sLoadDemo) GetDevicePoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 遥测点位
		telemetryPowerPoint,
		telemetryEnergyPoint,
		telemetryMaxLoadPoint,
		// 协议点位
		Status,
		Power,
		Energy,
		MaxLoad,
		PowerFactor,
		LoadRate,
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sLoadDemo) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		telemetryPowerPoint,  // 功率 - 核心运行参数
		telemetryEnergyPoint, // 电量 - 累计统计
	}
}
